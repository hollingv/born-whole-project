export async function onRequestGet(context) {
    const q = new URL(context.request.url).searchParams.get('q') || '';
    const query = q.toLowerCase();

    const kbFiles = await listKBFiles(context);
    const matches = [];

    for (const file of kbFiles) {
        const resp = await fetch(new URL(`/kb/${file}`, context.request.url));
        if (!resp.ok) continue;

        const text = await resp.text();
        const source = fileToSource(file);
        for (const line of text.split('\n')) {
            const trimmed = line.trim();
            if (trimmed && trimmed.toLowerCase().includes(query)) {
                matches.push({ source, line: trimmed });
            }
            if (matches.length >= 10) break;
        }
        if (matches.length >= 10) break;
    }

    let html;
    if (matches.length === 0) {
        html = `<p>No results found for: <strong>${q}</strong></p>`;
    } else {
        html = `<p>Results for: <strong>${q}</strong></p>`;
        html += `<table class="ask-results"><thead><tr><th>Source</th><th>Content</th></tr></thead><tbody>`;
        for (const m of matches) {
            html += `<tr><td>${m.source}</td><td>${m.line}</td></tr>`;
        }
        html += `</tbody></table>`;
    }

    return new Response(html, {
        headers: { 'Content-Type': 'text/html; charset=utf-8' },
    });
}

function fileToSource(filename) {
    let name = filename.replace(/\.txt$/, '');
    const idx = name.lastIndexOf('-org');
    if (idx !== -1) name = name.slice(0, idx) + '.org';
    return name.replace(/-/g, '.');
}

async function listKBFiles(context) {
    const resp = await fetch(new URL('/kb/manifest.json', context.request.url));
    if (!resp.ok) return [];
    return await resp.json();
}
