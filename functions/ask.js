// Cloudflare Pages Function handling the /ask endpoint for the siteName
// The logic in this file mirrors src/server/server.go — keep both in sync when making changes.

const AI_MODEL = '@cf/meta/llama-3.1-8b-instruct-fast';

const STOP_WORDS = new Set([
    'a','an','the','is','are','was','were','be','been','being','have','has','had',
    'do','does','did','will','would','could','should','may','might','can',
    'what','why','how','when','where','who','which','that','this','these','those',
    'i','me','my','we','our','you','your','it','its','they','them','their',
    'of','in','to','for','on','with','at','by','from','about','and','or',
    'not','no','so','if','than','very',
]);

const SYSTEM_PROMPT = `You are a helpful assistant answering questions about genital cutting,
bodily autonomy, and children's rights. Only use the text provided below as context.
Do not use any outside knowledge. Be concise, factual, and compassionate.
If the provided context does not contain enough information to answer, say so explicitly.`;

export async function onRequestGet(context) {
    try {
        return await handleRequest(context);
    } catch (err) {
        return respond(`<p>An error occurred: <code>${sanitiseError(err)}</code></p>`);
    }
}

async function handleRequest(context) {
    const q = new URL(context.request.url).searchParams.get('q') || '';
    const chunks = await searchKB(q, context.request.url);

    if (chunks.length === 0) {
        return respond(`<p>No information found for: <strong>${q}</strong></p>`);
    }

    return await askAI(q, chunks, context);
}

async function searchKB(q, requestUrl) {
    const manifestResp = await fetch(new URL('/kb/manifest.json', requestUrl));
    if (!manifestResp.ok) return [];
    const filenames = await manifestResp.json();

    const keywords = extractKeywords(q);
    const searchTerms = keywords.length > 0 ? keywords : [q.toLowerCase()];
    const scored = [];

    for (const file of filenames) {
        const resp = await fetch(new URL(`/kb/${file}`, requestUrl));
        if (!resp.ok) continue;
        const text = await resp.text();
        const source = sourceFromContent(text, file);
        for (const para of text.split('\n\n')) {
            const trimmed = para.trim();
            if (!trimmed) continue;
            const paraLower = trimmed.toLowerCase();
            const score = searchTerms.filter(kw => paraLower.includes(kw)).length;
            if (score > 0) scored.push({ text: trimmed, source, score });
        }
    }

    scored.sort((a, b) => b.score - a.score);
    return scored.slice(0, 10);
}

async function askAI(q, chunks, context) {
    const sourceNames = [...new Set(chunks.map(c => c.source))];
    const sources = sourceNames.map(s => `<a href="${s}" target="_blank">${s}</a>`);
    const contextText = chunks.map(c => c.text).join('\n\n');

    try {
        if (!context.env.AI) {
            throw new Error('AI binding not configured — add the AI binding in the Cloudflare Pages dashboard under Settings → Functions → AI Bindings');
        }
        const aiResponse = await context.env.AI.run(AI_MODEL, {
            messages: [
                { role: 'system', content: SYSTEM_PROMPT },
                { role: 'user', content: `Context:\n${contextText}\n\nQuestion: ${q}` }
            ]
        });
        return respond(
            `<p>${aiResponse.response}</p>` +
            `<p class="ask-sources">Sources:<br>${sources.join('<br>')}</p>` +
            `<p class="ask-disclaimer">Answers are generated from curated sources. Always verify with the linked organisations.</p>`
        );
    } catch (err) {
        let html = `<p>AI unavailable: <code>${sanitiseError(err)}</code></p>`;
        html += '<p>Here are relevant excerpts:</p><ul>';
        for (const c of chunks) html += `<li>${c.text}</li>`;
        html += `</ul><p class="ask-sources">Sources:<br>${sources.join('<br>')}</p>`;
        return respond(html);
    }
}

function extractKeywords(query) {
    return query.toLowerCase()
        .split(/[^a-z]+/)
        .filter(w => w.length > 2 && !STOP_WORDS.has(w));
}

function sourceFromContent(text, filename) {
    for (const line of text.split('\n').slice(0, 5)) {
        const trimmed = line.trim();
        if (trimmed.startsWith('# Source:')) {
            return trimmed.replace('# Source:', '').trim();
        }
    }
    return filename.replace(/\.txt$/, '');
}

function sanitiseError(err) {
    let msg = err?.message || String(err);
    msg = msg.replace(/accounts\/[a-f0-9]+/gi, 'accounts/[redacted]');
    msg = msg.replace(/Bearer\s+\S+/gi, 'Bearer [redacted]');
    return msg;
}

function respond(html) {
    return new Response(html, {
        headers: { 'Content-Type': 'text/html; charset=utf-8' },
    });
}
