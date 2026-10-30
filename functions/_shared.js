// formatParagraphs splits plain AI response text into <p> elements by grouping
// every two sentences, regardless of how the model chose to format its output.
export function formatParagraphs(text) {
    const sentences = text.match(/[^.!?]*[.!?]+\s*/g) || [text];
    const paras = [];
    for (let i = 0; i < sentences.length; i += 2) {
        const para = sentences.slice(i, i + 2).join('').trim();
        if (para) paras.push(`<p>${para}</p>`);
    }
    return paras.length ? paras.join('\n') : `<p>${text}</p>`;
}
