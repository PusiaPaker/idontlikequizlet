document.addEventListener("DOMContentLoaded", () => {
  document.querySelectorAll(".delete-btn").forEach(btn => {
    btn.addEventListener("click", async () => {
      const cardID = btn.dataset.cardId;
      const deckID = btn.dataset.deckId;
      if (!cardID || !deckID) return;

      // confirm deletion
      if (!confirm("Delete this card?")) return;

      try {
        const res = await fetch(`/edit/${deckID}/delete/${cardID}`, {
          method: "DELETE",
        });

        if (!res.ok) throw new Error(`HTTP ${res.status}`);

        // Remove card from the page visually
        const cardElement = btn.closest(".edit-card");
        if (cardElement) cardElement.remove();

        console.log("Card deleted successfully");
      } catch (err) {
        console.error("Delete failed:", err);
        alert("Failed to delete card");
      }
    });
  });
});
