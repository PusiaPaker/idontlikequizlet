function debounce(fn, wait=500) {
  let t;
  return (...args) => { clearTimeout(t); t = setTimeout(() => fn(...args), wait); };
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

document.addEventListener("DOMContentLoaded", () => {
  const saveTerm = debounce(async (el) => {
    const deckID = el.dataset.deckId, cardID = el.dataset.cardId;
    try { await patchCard(deckID, cardID, { term: el.value }); el.classList.add("is-valid"); }
    catch { el.classList.add("is-invalid"); }
    finally { setTimeout(() => el.classList.remove("is-valid","is-invalid"), 1200); }
  });

  const saveDef = debounce(async (el) => {
    const deckID = el.dataset.deckId, cardID = el.dataset.cardId;
    try { await patchCard(deckID, cardID, { definition: el.value }); el.classList.add("is-valid"); }
    catch { el.classList.add("is-invalid"); }
    finally { setTimeout(() => el.classList.remove("is-valid","is-invalid"), 1200); }
  });

  document.querySelectorAll(".card-term").forEach(el => {
    el.addEventListener("input", () => saveTerm(el));
    el.addEventListener("change", () => saveTerm(el)); // fallback
  });

  document.querySelectorAll(".card-def").forEach(el => {
    el.addEventListener("input", () => saveDef(el));
    el.addEventListener("change", () => saveDef(el));
  });
});