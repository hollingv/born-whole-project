/**
 * youtube-shorts.js
 * Handles the grid/list view toggle and per-row transcript expand/collapse
 * on the YouTube Shorts page.
 */
(function () {
    'use strict';

    // --- View toggle ---

    var btnGrid  = document.getElementById('btn-grid');
    var btnList  = document.getElementById('btn-list');
    var viewGrid = document.getElementById('view-grid');
    var viewList = document.getElementById('view-list');

    if (btnGrid && btnList && viewGrid && viewList) {
        btnGrid.addEventListener('click', function () {
            showView('grid');
        });
        btnList.addEventListener('click', function () {
            showView('list');
        });
    }

    function showView(which) {
        var isGrid = which === 'grid';

        viewGrid.hidden = !isGrid;
        viewList.hidden = isGrid;

        btnGrid.classList.toggle('view-btn--active', isGrid);
        btnList.classList.toggle('view-btn--active', !isGrid);

        btnGrid.setAttribute('aria-pressed', String(isGrid));
        btnList.setAttribute('aria-pressed', String(!isGrid));
    }

    // --- Transcript expand / collapse ---

    document.addEventListener('click', function (e) {
        var btn = e.target.closest('.resource-expand-btn');
        if (!btn || btn.classList.contains('resource-expand-btn--none')) return;

        var row        = btn.closest('.resource-list-item');
        var transcript = row && row.querySelector('.resource-transcript');
        if (!transcript) return;

        var expanded = btn.getAttribute('aria-expanded') === 'true';
        expanded = !expanded;

        transcript.hidden = !expanded;
        btn.setAttribute('aria-expanded', String(expanded));
        btn.textContent   = expanded ? '−' : '+';
        btn.setAttribute('aria-label',
            (expanded ? 'Hide' : 'Show') + ' transcript for ' +
            (row.querySelector('.resource-list-title') || {}).textContent
        );
    });

}());
