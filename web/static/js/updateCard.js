function debounce(fn, wait = 500) {
  let t;
  return (...args) => {
    clearTimeout(t);
    t = setTimeout(() => fn(...args), wait);
  };
}

async function patchCard(deckID, cardID, payload) {
  const res = await fetch(`/edit/${deckID}/update/${cardID}`, {
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
  if (!res.ok) throw new Error(`HTTP ${res.status}`);
  return res.json();
}

async function patchTitle(deckID, payload) {
  const res = await fetch(`/edit/${deckID}/update/title`, {
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
  if (!res.ok) throw new Error(`HTTP ${res.status}`);
  return res.json();
}

function showValidation(el, success) {
  el.classList.add(success ? "is-valid" : "is-invalid");
  setTimeout(() => el.classList.remove("is-valid", "is-invalid"), 1200);
}

document.addEventListener("DOMContentLoaded", () => {
  // save title
  const saveTitle = debounce(async (el) => {
    const deckID = el.dataset.deckId;
    const val = el.value.trim();
    if (!val) {
      showValidation(el, false);
      return;
    }

    try {
      await patchTitle(deckID, { name: val });
      showValidation(el, true);
    } catch {
      showValidation(el, false);
    }
  });

  // save term
  const saveTerm = debounce(async (el) => {
    const deckID = el.dataset.deckId;
    const cardID = el.dataset.cardId;
    const val = el.value.trim();
    if (!deckID || !cardID || !val) {
      showValidation(el, false);
      return;
    }

    try {
      await patchCard(deckID, cardID, { term: val });
      showValidation(el, true);
    } catch {
      showValidation(el, false);
    }
  });

  // save definition
  const saveDef = debounce(async (el) => {
    const deckID = el.dataset.deckId;
    const cardID = el.dataset.cardId;
    const val = el.value.trim();
    if (!deckID || !cardID || !val) {
      showValidation(el, false);
      return;
    }

    try {
      await patchCard(deckID, cardID, { definition: val });
      showValidation(el, true);
    } catch {
      showValidation(el, false);
    }
  });

  // listeners
  document.querySelectorAll(".deck-title").forEach((el) => {
    el.addEventListener("input", () => saveTitle(el));
    el.addEventListener("change", () => saveTitle(el));
  });

  document.querySelectorAll(".card-term").forEach((el) => {
    el.addEventListener("input", () => saveTerm(el));
    el.addEventListener("change", () => saveTerm(el));
  });

  document.querySelectorAll(".card-def").forEach((el) => {
    el.addEventListener("input", () => saveDef(el));
    el.addEventListener("change", () => saveDef(el));
  });
});
