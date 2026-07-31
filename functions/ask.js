export async function onRequestGet(context) {
    const q = new URL(context.request.url).searchParams.get('q') || '';
    const html = `<p>You asked: <strong>${q}</strong></p><p>Stub response: more information coming soon.</p>`;
    return new Response(html, {
        headers: { 'Content-Type': 'text/html; charset=utf-8' },
    });
}
