// Automatically builds a table of contents from section headings.
// Configure via data attributes on the #toc element:
//   data-section  — CSS selector for each section (default: .resource-section)
//   data-heading  — heading tag within each section (default: h2)
// Adding a new section to the template will appear in the TOC without any code changes.
document.addEventListener('DOMContentLoaded', () => {
    const toc = document.getElementById('toc');
    if (!toc) return;

    const sectionSelector = toc.dataset.section || '.resource-section';
    const headingTag      = toc.dataset.heading  || 'h2';

    const sections = document.querySelectorAll(sectionSelector);
    if (sections.length === 0) return;

    const label = document.createElement('p');
    label.className = 'toc-heading';
    label.textContent = 'On this page';
    toc.appendChild(label);

    const list = document.createElement('ul');
    sections.forEach(section => {
        const heading = section.querySelector(headingTag);
        if (!heading) return;

        const id = heading.textContent.trim().toLowerCase().replace(/[^a-z0-9]+/g, '-');
        section.id = id;

        const li = document.createElement('li');
        const a  = document.createElement('a');
        a.href        = `#${id}`;
        a.textContent = heading.textContent.trim();
        li.appendChild(a);
        list.appendChild(li);
    });

    toc.appendChild(list);
});
