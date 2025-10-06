// --- helpers ---
function escHTML(s) {
  if (s == null) return "";           // handles undefined/null
  return String(s).replace(/[&<>"']/g, c => (
    ({ '&':'&amp;', '<':'&lt;', '>':'&gt;', '"':'&quot;', "'":'&#39;' })[c]
  ));
}

function debounce(fn, wait=500){ let t; return (...a)=>{ clearTimeout(t); t=setTimeout(()=>fn(...a),wait); }; }

async function postNewCard(deckID) {
  const res = await fetch(`/add/${deckID}`, { method: "POST" });
  if (!res.ok) throw new Error(`HTTP ${res.status}`);
  const data = await res.json();
  // sanity check the payload shape
  if (!data || typeof data !== "object") {
    throw new Error("Bad JSON from /add: not an object");
  }
  if (data.id == null) {
    // still allow it (we’ll handle empty string), but log for you to debug server
    console.warn("Server returned card without id. Payload:", data);
  }
  return data;
}

// Re-attach autosave for just-added fields (unchanged from before)
function attachAutosaveHandlersFor(el) {
  const termInput = el.querySelector(".card-term");
  const defInput  = el.querySelector(".card-def");

  const showValidation = (node, ok) => {
    node.classList.add(ok ? "is-valid" : "is-invalid");
    setTimeout(() => node.classList.remove("is-valid","is-invalid"), 1200);
  };

  const saveTerm = debounce(async (node) => {
    const deckID = node.dataset.deckId;
    const cardID = node.dataset.cardId;
    const val = node.value.trim();
    if (!deckID || !cardID || !val) { showValidation(node, false); return; }
    const res = await fetch(`/edit/${deckID}/update/${cardID}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ term: val }),
    });
    showValidation(node, res.ok);
  });

  const saveDef = debounce(async (node) => {
    const deckID = node.dataset.deckId;
    const cardID = node.dataset.cardId;
    const val = node.value.trim();
    if (!deckID || !cardID || !val) { showValidation(node, false); return; }
    const res = await fetch(`/edit/${deckID}/update/${cardID}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ definition: val }),
    });
    showValidation(node, res.ok);
  });

  if (termInput) {
    termInput.addEventListener("input", () => saveTerm(termInput));
    termInput.addEventListener("change", () => saveTerm(termInput));
  }
  if (defInput) {
    defInput.addEventListener("input", () => saveDef(defInput));
    defInput.addEventListener("change", () => saveDef(defInput));
  }
}

// Build a new card editor block (matches your partial’s structure/classes)
function buildCardHTML(index, card, deckID) {
  // default any missing fields
  const id   = escHTML(card?.id ?? "");
  const term = escHTML(card?.term ?? "");
  const def  = escHTML(card?.definition ?? "");
  const dID  = escHTML(deckID ?? "");

  return `
  <div class="edit-card border rounded-3 mb-4 p-4 shadow-sm">
    <div class="row g-4 align-items-start">
      <div class="col-md-7">
        <label for="term-${index}" class="form-label fw-semibold">Term</label>
        <input
          type="text"
          class="form-control mb-3 card-term"
          id="term-${index}"
          name="term[]"
          data-card-id="${id}"
          data-deck-id="${dID}"
          placeholder="Enter term..."
          value="${term}">
      </div>

      <div class="col-md-5">
        <label class="form-label fw-semibold d-flex justify-content-between">
          Image
          <span class="text-muted small">optional</span>
        </label>
        <button type="button"
          class="image-slot w-100 border border-2 border-secondary-subtle bg-light-subtle"
          title="Add image (coming soon)">
          <div class="image-slot-inner">
            <span class="image-plus">+</span>
            <span class="image-text">Add Image</span>
          </div>
        </button>
      </div>
    </div>

    <div class="mt-4">
      <label for="def-${index}" class="form-label fw-semibold">Definition</label>
      <textarea
        class="form-control card-def"
        id="def-${index}"
        name="definition[]"
        rows="5"
        data-card-id="${id}"
        data-deck-id="${dID}"
        placeholder="Enter definition..."></textarea>
    </div>

    <div class="d-flex justify-content-end mt-4">
      <button
        type="button"
        class="btn btn-outline-danger px-4 fw-semibold delete-btn"
        data-card-id="${id}"
        data-deck-id="${dID}">DELETE</button>
    </div>
  </div>`;
}

document.addEventListener("DOMContentLoaded", () => {
  const addBtn = document.getElementById("add-card-btn");
  const list = document.getElementById("edit-list");
  if (!addBtn || !list) return;

  addBtn.addEventListener("click", async () => {
    const deckID = addBtn.dataset.deckId;
    if (!deckID) {
      console.error("add-card-btn is missing data-deck-id");
      alert("Cannot add: missing deck id");
      return;
    }

    try {
      const card = await postNewCard(deckID);
      const index = list.querySelectorAll(".edit-card").length;
      const html = buildCardHTML(index, card, deckID);

      const wrapper = document.createElement("div");
      wrapper.innerHTML = html.trim();
      const block = wrapper.firstElementChild;
      list.appendChild(block);
      attachAutosaveHandlersFor(block);
      block.scrollIntoView({ behavior: "smooth", block: "center" });

      // If server didn’t give an id, warn (autosave will fail without cardId)
      if (!card.id) {
        console.warn("New card has no id; autosave for this card will not work until page reload.");
      }
    } catch (e) {
      console.error(e);
      alert("Failed to add card");
    }
  });

  // Delegated delete (works for dynamically-added cards too)
  document.addEventListener("click", async (ev) => {
    const btn = ev.target.closest(".delete-btn");
    if (!btn) return;
    const cardID = btn.dataset.cardId;
    const deckID = btn.dataset.deckId;
    if (!cardID || !deckID) return;

    if (!confirm("Delete this card?")) return;

    try {
      const res = await fetch(`/edit/${deckID}/delete/${cardID}`, { method: "DELETE" });
      if (!res.ok) throw new Error(`HTTP ${res.status}`);
      const cardEl = btn.closest(".edit-card");
      if (cardEl) cardEl.remove();
    } catch (e) {
      console.error(e);
      alert("Failed to delete card");
    }
  });
});
