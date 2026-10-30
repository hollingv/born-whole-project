// Shared utilities for Cloudflare Pages Functions.

const STOP_WORDS = new Set([
    'a','an','the','is','are','was','were','be','been','being','have','has','had',
    'do','does','did','will','would','could','should','may','might','can',
    'what','why','how','when','where','who','which','that','this','these','those',
    'i','me','my','we','our','you','your','it','its','they','them','their',
    'of','in','to','for','on','with','at','by','from','about','and','or',
    'not','no','so','if','than','very',
]);

export function extractKeywords(query) {
    return query.toLowerCase()
        .split(/[^a-z]+/)
        .filter(w => w.length > 2 && !STOP_WORDS.has(w));
}

export function sourceFromContent(text, fallback) {
    for (const line of text.split('\n').slice(0, 5)) {
        const trimmed = line.trim();
        if (trimmed.startsWith('# Source:')) {
            return trimmed.replace('# Source:', '').trim();
        }
    }
    return fallback;
}

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

// searchBundle fetches a pre-built bundle (array of file-content strings) from
// bundleUrl and returns the top 10 scored chunks matching the query.
// One subrequest total instead of one per KB file.
export async function searchBundle(q, bundleUrl) {
    const resp = await fetch(bundleUrl);
    if (!resp.ok) return [];
    const fileContents = await resp.json();

    const keywords = extractKeywords(q);
    const searchTerms = keywords.length > 0 ? keywords : [q.toLowerCase()];
    const scored = [];

    for (const content of fileContents) {
        const source = sourceFromContent(content, '');
        for (const para of content.split('\n\n')) {
            const trimmed = para.trim();
            if (!trimmed || trimmed.startsWith('#')) continue;
            const paraLower = trimmed.toLowerCase();
            const score = searchTerms.filter(kw => paraLower.includes(kw)).length;
            if (score > 0) scored.push({ text: trimmed, source, score });
        }
    }

    scored.sort((a, b) => b.score - a.score);
    return scored.slice(0, 10);
}
