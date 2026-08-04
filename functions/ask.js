const AI_MODEL = '@cf/meta/llama-3.1-8b-instruct-fast';

const STOP_WORDS = new Set([
    'a','an','the','is','are','was','were','be','been','being','have','has','had',
    'do','does','did','will','would','could','should','may','might','can',
    'what','why','how','when','where','who','which','that','this','these','those',
    'i','me','my','we','our','you','your','it','its','they','them','their',
    'of','in','to','for','on','with','at','by','from','about','and','or',
    'not','no','so','if','than','very',
]);

function extractKeywords(query) {
    return query.toLowerCase()
        .split(/[^a-z]+/)
        .filter(w => w.length > 2 && !STOP_WORDS.has(w));
}

const SYSTEM_PROMPT = `You are a helpful assistant answering questions about circumcision,
bodily autonomy, and children's rights. Answer based only on the provided context.
Be concise, factual, and compassionate. If the context does not contain enough
information to answer, say so.`;

// Sanitise an error message to remove any account-specific information.
function sanitiseError(err) {
    let msg = err?.message || String(err);
    msg = msg.replace(/accounts\/[a-f0-9]+/gi, 'accounts/[redacted]');
    msg = msg.replace(/Bearer\s+\S+/gi, 'Bearer [redacted]');
    return msg;
}

export async function onRequestGet(context) {
    try {
        return await handleRequest(context);
    } catch (err) {
        const msg = sanitiseError(err);
        return respond(`<p>An error occurred: <code>${msg}</code></p>`);
    }
}

async function handleRequest(context) {
    const q = new URL(context.request.url).searchParams.get('q') || '';
    const query = q.toLowerCase();

    // Load KB manifest
    const manifestResp = await fetch(new URL('/kb/manifest.json', context.request.url));
    if (!manifestResp.ok) {
        return respond('<p>Knowledge base not available.</p>');
    }
    const filenames = await manifestResp.json();

    // Extract keywords and search for relevant chunks
    const keywords = extractKeywords(q);
    const searchTerms = keywords.length > 0 ? keywords : [query];
    const scored = [];
    for (const file of filenames) {
        const resp = await fetch(new URL(`/kb/${file}`, context.request.url));
        if (!resp.ok) continue;
        const text = await resp.text();
        const source = fileToSource(file);
        for (const line of text.split('\n')) {
            const trimmed = line.trim();
            if (!trimmed) continue;
            const lineLower = trimmed.toLowerCase();
            const score = searchTerms.filter(kw => lineLower.includes(kw)).length;
            if (score > 0) scored.push({ text: trimmed, source, score });
        }
    }
    scored.sort((a, b) => b.score - a.score);
    const chunks = scored.slice(0, 10);

    if (chunks.length === 0) {
        return respond(`<p>No information found for: <strong>${q}</strong></p>`);
    }

    const sources = [...new Set(chunks.map(c => c.source))];
    const contextText = chunks.map(c => c.text).join('\n\n');

    // Call Workers AI
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
        return respond(`<p>${aiResponse.response}</p><p class="ask-sources">Sources: ${sources.join(', ')}</p>`);
    } catch (err) {
        const msg = sanitiseError(err);
        let html = `<p>AI unavailable: <code>${msg}</code></p>`;
        html += '<p>Here are relevant excerpts:</p><ul>';
        for (const c of chunks) html += `<li>${c.text}</li>`;
        html += `</ul><p class="ask-sources">Sources: ${sources.join(', ')}</p>`;
        return respond(html);
    }
}

function fileToSource(filename) {
    let name = filename.replace(/\.txt$/, '');
    const idx = name.lastIndexOf('-org');
    if (idx !== -1) name = name.slice(0, idx) + '.org';
    return name.replace(/-/g, '.');
}

function respond(html) {
    return new Response(html, {
        headers: { 'Content-Type': 'text/html; charset=utf-8' },
    });
}
