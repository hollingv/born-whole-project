(function () {
    const today = new Date();
    today.setHours(0, 0, 0, 0);

    document.querySelectorAll('[data-date]').forEach(function (el) {
        const parts = el.dataset.date.split('-');
        if (parts.length !== 3) return;

        // Construct date from parts to avoid timezone shifts from Date.parse.
        const eventDate = new Date(
            parseInt(parts[0], 10),
            parseInt(parts[1], 10) - 1,
            parseInt(parts[2], 10)
        );

        if (eventDate < today) {
            el.classList.add('event--archived');
            const dateEl = el.querySelector('.home-event-date, .events-date');
            if (dateEl) {
                const label = document.createElement('span');
                label.className = 'archived-label';
                label.textContent = ' (archived)';
                dateEl.appendChild(label);
            }
        }
    });
}());
