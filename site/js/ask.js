const exampleQuestions = [
    'What is genital cutting?',
    'What are the medical risks?',
    'What do doctors say about genital cutting?',
    'Is genital cutting legal?',
    'How common is genital cutting worldwide?',
    'What is bodily autonomy?',
    'What does the law say about genital cutting?',
];

document.addEventListener('DOMContentLoaded', () => {
    const form = document.querySelector('.prompt-form');
    const responseDiv = document.getElementById('ask-response');
    const input = form.querySelector('.prompt-input');

    // Cycle example questions as actual input value every 3 seconds
    // Stop cycling once the user starts typing
    let index = 0;
    let userTyped = false;
    input.value = exampleQuestions[0];
    input.addEventListener('input', () => { userTyped = true; });
    setInterval(() => {
        if (userTyped) return;
        index = (index + 1) % exampleQuestions.length;
        input.value = exampleQuestions[index];
    }, 3000);

    // Handle form submission
    form.addEventListener('submit', async (e) => {
        e.preventDefault();
        const q = input.value.trim();
        if (!q) return;

        responseDiv.innerHTML = `<p class="ask-question">You asked: <strong>${q}</strong></p><p>Thinking...</p>`;

        const resp = await fetch(`/ask?q=${encodeURIComponent(q)}`);
        const text = await resp.text();
        responseDiv.innerHTML = `<p class="ask-question">You asked: <strong>${q}</strong></p>${text}`;
    });
});
