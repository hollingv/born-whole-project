const AI_MODEL = '@cf/meta/llama-3.1-8b-instruct';

const SYSTEM_PROMPT = `You are a helpful assistant answering questions about circumcision,
bodily autonomy, and children's rights. Answer based only on the provided context.
Be concise, factual, and compassionate. If the context does not contain enough
information to answer, say so.`;

export async function onRequestGet(context) {
    const q = new URL(context.request.url).searchParams.get('q') || '';
    const query = q.toLowerCase();

    // Load KB manifest
    const manifestResp = await fetch(new URL('/kb/manifest.json', context.request.url));
    if (!manifestResp.ok) {
        return respond('<p>Knowledge base not available.</p>');
    }
    const filenames = await manifestResp.json();

    // Keyword search for relevant chunks
    const chunks = [];
    for (const file of filenames) {
        const resp = await fetch(new URL(`/kb/${file}`, context.request.url));
        if (!resp.ok) continue;
        const text = await resp.text();
        const source = fileToSource(file);
        for (const line of text.split('\n')) {
            const trimmed = line.trim();
            if (trimmed && trimmed.toLowerCase().includes(query)) {
                chunks.push({ text: trimmed, source });
            }
            if (chunks.length >= 10) break;
        }
        if (chunks.length >= 10) break;
    }

    if (chunks.length === 0) {
        return respond(`<p>No information found for: <strong>${q}</strong></p>`);
    }

    // Collect unique sources
    const sources = [...new Set(chunks.map(c => c.source))];
    const contextText = chunks.map(c => c.text).join('\n\n');

    // Call Workers AI
    try {
        const aiResponse = await context.env.AI.run(AI_MODEL, {
            messages: [
                { role: 'system', content: SYSTEM_PROMPT },
                { role: 'user', content: `Context:\n${contextText}\n\nQuestion: ${q}` }
            ]
        });
        return respond(`<p>${aiResponse.response}</p><p class="ask-sources">Sources: ${sources.join(', ')}</p>`);
    } catch (err) {
        let html = '<p>AI unavailable. Here are relevant excerpts:</p><ul>';
        for (const c of chunks) html += `<li>${c.text}</li>`;
        html += `</ul><p class="ask-sources">Sources: ${sources.join(', ')}</p>`;
        return respond(html);
    }
}

function respond(html) {
    return new Response(html, {
        headers: { 'Content-Type': 'text/html; charset=utf-8' },
    });
}
