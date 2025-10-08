(() => {
  const dataEl = document.getElementById("flashstudy-data");
  const CARDS = dataEl ? JSON.parse(dataEl.textContent || "[]") : [];

  const inner   = document.getElementById("card-inner");
  const frontT  = document.getElementById("front-text");
  const backT   = document.getElementById("back-text");
  const counter = document.getElementById("counter");
  const cardEl  = document.getElementById("flashcard");
  const prevBtn = document.getElementById("prev");
  const nextBtn = document.getElementById("next");

  if (!inner || !frontT || !backT || !counter || !cardEl) return;

  let i = 0;
  let flipped = false;

  // write content; CSS handles centering/wrapping; JS doesn’t touch layout
  const render = () => {
    const c = CARDS[i] || {};
    frontT.textContent = c.Term ?? c.term ?? "";
    backT.textContent  = c.Definition ?? c.definition ?? c.def ?? "";
    counter.textContent = `${CARDS.length ? i + 1 : 0}/${CARDS.length || 0}`;
    inner.style.transform = flipped ? "rotateX(180deg)" : "rotateX(0deg)";
  };

  // carousel + flip
  const next = () => { if (!CARDS.length) return; i = (i + 1) % CARDS.length; flipped = false; render(); };
  const prev = () => { if (!CARDS.length) return; i = (i - 1 + CARDS.length) % CARDS.length; flipped = false; render(); };
  const flip = () => { if (!CARDS.length) return; flipped = !flipped; render(); };

  // interactions
  cardEl.addEventListener("click", (e) => {
    if (e.target.closest(".nav")) return; // clicking arrows shouldn't flip
    flip();
  });
  prevBtn?.addEventListener("click", (e) => { e.stopPropagation(); prev(); });
  nextBtn?.addEventListener("click", (e) => { e.stopPropagation(); next(); });

  document.addEventListener("keydown", (e) => {
    if (e.code === "ArrowRight") { e.preventDefault(); next(); }
    else if (e.code === "ArrowLeft") { e.preventDefault(); prev(); }
    else if (e.code === "Space") { e.preventDefault(); flip(); }
  });

  // init
  if (!Array.isArray(CARDS) || CARDS.length === 0) {
    frontT.textContent = "No cards yet.";
    counter.textContent = "0/0";
  }
  render();
})();
