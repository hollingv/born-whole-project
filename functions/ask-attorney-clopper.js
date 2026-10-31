// Cloudflare Pages Function handling the /ask-attorney-clopper endpoint.
// Mirrors the pattern in functions/ask.js — keep both in sync when making changes.
import { formatParagraphs, searchBundle } from './_shared.js';

const AI_MODEL = '@cf/meta/llama-3.1-8b-instruct-fast';

const SYSTEM_PROMPT = `You are attorney Clopper and speak in the first person.
Answer questions based exclusively on the transcript context provided below.
Do not use any outside knowledge. Be concise, factual, and direct.
Never say "according to the text", "according to the provided context", "based on the context", or any similar phrase. Speak naturally in the first person as if from your own knowledge.
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
    const bundleUrl = new URL('/harvested/attyclopper-kb-bundle.json', context.request.url).href;
    const chunks = await searchBundle(q, bundleUrl);

    if (chunks.length === 0) {
        return respond(`<p>No information found for: <strong>${q}</strong></p>`);
    }

    return await askAI(q, chunks, context);
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
            max_tokens: 1024,
            messages: [
                { role: 'system', content: SYSTEM_PROMPT },
                { role: 'user', content: `Context:\n${contextText}\n\nQuestion: ${q}` }
            ]
        });
        return respond(
            formatParagraphs(aiResponse.response) +
            `<p class="ask-sources">Sources:<br>${sources.join('<br>')}</p>` +
            `<p class="ask-disclaimer">Answers are drawn from Attorney Clopper's YouTube Shorts transcripts. Always verify with the source videos.</p>`
        );
    } catch (err) {
        let html = `<p>AI unavailable: <code>${sanitiseError(err)}</code></p>`;
        html += '<p>Here are relevant excerpts:</p><ul>';
        for (const c of chunks) html += `<li>${c.text}</li>`;
        html += `</ul><p class="ask-sources">Sources:<br>${sources.join('<br>')}</p>`;
        return respond(html);
    }
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
