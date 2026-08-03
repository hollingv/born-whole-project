export async function onRequestGet(context) {
    const q = new URL(context.request.url).searchParams.get('q') || '';
    const query = q.toLowerCase();

    // Fetch KB files from the static site
    const kbFiles = await listKBFiles(context);
    const matches = [];

    for (const file of kbFiles) {
        const resp = await fetch(new URL(`/kb/${file}`, context.request.url));
        if (!resp.ok) continue;

        const text = await resp.text();
        for (const line of text.split('\n')) {
            const trimmed = line.trim();
            if (trimmed && trimmed.toLowerCase().includes(query)) {
                matches.push(trimmed);
            }
            if (matches.length >= 5) break;
        }
        if (matches.length >= 5) break;
    }

    let html;
    if (matches.length === 0) {
        html = `<p>No results found for: <strong>${q}</strong></p>`;
    } else {
        html = `<p>Results for: <strong>${q}</strong></p><ul>`;
        for (const m of matches) {
            html += `<li>${m}</li>`;
        }
        html += `</ul>`;
    }

    return new Response(html, {
        headers: { 'Content-Type': 'text/html; charset=utf-8' },
    });
}

async function listKBFiles(context) {
    const resp = await fetch(new URL('/kb/', context.request.url));
    if (!resp.ok) return [];
    const text = await resp.text();
    const matches = [...text.matchAll(/href="([^"]+\.txt)"/g)];
    return matches.map(m => m[1]);
}
