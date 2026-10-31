document.addEventListener('DOMContentLoaded', () => {
    const form = document.querySelector('.prompt-form');
    const responseDiv = document.getElementById('ask-response');
    const input = form.querySelector('.prompt-input');

    input.value = 'Is genital cutting illegal for both girls and boys?';

    // Handle form submission
    form.addEventListener('submit', async (e) => {
        e.preventDefault();
        const q = input.value.trim();
        if (!q) return;

        responseDiv.classList.add('has-content');
        responseDiv.innerHTML = `<p class="ask-question">You asked: <strong>${q}</strong></p><p>Thinking...</p>`;

        const resp = await fetch(`/ask-attorney-clopper?q=${encodeURIComponent(q)}`);
        const text = await resp.text();
        responseDiv.innerHTML = `<p class="ask-question">You asked: <strong>${q}</strong></p>${text}`;
        const meta = document.getElementById('ask-response-meta');
        if (meta) {
            meta.innerHTML = '';
            ['ask-sources', 'ask-disclaimer'].forEach(cls => {
                const el = responseDiv.querySelector('.' + cls);
                if (el) meta.appendChild(el);
            });
        }
    });
});
