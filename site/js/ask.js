document.addEventListener('DOMContentLoaded', () => {
    const form = document.querySelector('.prompt-form');
    const responseDiv = document.getElementById('ask-response');

    form.addEventListener('submit', async (e) => {
        e.preventDefault();
        const q = form.querySelector('.prompt-input').value.trim();
        if (!q) return;

        responseDiv.innerHTML = '<p>Thinking...</p>';

        const resp = await fetch(`/ask?q=${encodeURIComponent(q)}`);
        const text = await resp.text();
        responseDiv.innerHTML = text;
    });
});
