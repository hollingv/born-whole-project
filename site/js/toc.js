// Automatically builds a table of contents from .resource-section h2 elements.
// Adding a new section to the template will appear in the TOC without any code changes.
document.addEventListener('DOMContentLoaded', () => {
    const toc = document.getElementById('toc');
    if (!toc) return;

    const sections = document.querySelectorAll('.resource-section');
    if (sections.length === 0) return;

    const heading = document.createElement('p');
    heading.className = 'toc-heading';
    heading.textContent = 'On this page';
    toc.appendChild(heading);

    const list = document.createElement('ul');
    sections.forEach(section => {
        const h2 = section.querySelector('h2');
        if (!h2) return;

        // Generate a slug ID from the heading text and assign it to the section
        const id = h2.textContent.trim().toLowerCase().replace(/[^a-z0-9]+/g, '-');
        section.id = id;

        const li = document.createElement('li');
        const a = document.createElement('a');
        a.href = `#${id}`;
        a.textContent = h2.textContent.trim();
        li.appendChild(a);
        list.appendChild(li);
    });

    toc.appendChild(list);
});
