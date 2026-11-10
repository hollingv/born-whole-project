/**
 * youtube-shorts.js
 * Grid/list view toggle for the YouTube Shorts page.
 * Each button has its own self-contained handler — no shared helper function.
 */
(function () {
    'use strict';

    var btnGrid  = document.getElementById('btn-grid');
    var btnList  = document.getElementById('btn-list');
    var viewGrid = document.getElementById('view-grid');
    var viewList = document.getElementById('view-list');

    if (!btnGrid || !btnList || !viewGrid || !viewList) return;

    btnGrid.addEventListener('click', function () {
        viewGrid.hidden = false;
        viewList.hidden = true;
        btnGrid.classList.add('view-btn--active');
        btnList.classList.remove('view-btn--active');
        btnGrid.setAttribute('aria-pressed', 'true');
        btnList.setAttribute('aria-pressed', 'false');
    });

    btnList.addEventListener('click', function () {
        viewList.hidden = false;
        viewGrid.hidden = true;
        btnList.classList.add('view-btn--active');
        btnGrid.classList.remove('view-btn--active');
        btnList.setAttribute('aria-pressed', 'true');
        btnGrid.setAttribute('aria-pressed', 'false');
    });

    // --- Transcript expand / collapse ---

    document.addEventListener('click', function (e) {
        var btn = e.target.closest('.expand-btn');
        if (!btn) return;

        var item       = btn.closest('.resource-list-item');
        var transcript = item && item.querySelector('.resource-transcript');
        if (!transcript) return;

        var expanded = btn.getAttribute('aria-expanded') === 'true';
        expanded = !expanded;
        transcript.hidden = !expanded;
        btn.setAttribute('aria-expanded', String(expanded));
        btn.textContent = expanded ? '−' : '+';
    });

}());
