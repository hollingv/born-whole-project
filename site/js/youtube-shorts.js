/**
 * youtube-shorts.js
 * Handles transcript expand/collapse on the YouTube Shorts page.
 */
(function () {
    'use strict';

    function closeAllTranscripts() {
        document.querySelectorAll('.resource-transcript').forEach(function (t) {
            if (!t.hidden) {
                t.hidden = true;
                var b = t.closest('.resource-list-item').querySelector('.expand-btn');
                if (b) { b.setAttribute('aria-expanded', 'false'); b.textContent = '+'; }
            }
        });
    }

    document.addEventListener('click', function (e) {
        var btn = e.target.closest('.expand-btn');

        if (!btn) {
            if (!e.target.closest('.resource-transcript')) {
                closeAllTranscripts();
            }
            return;
        }

        var item       = btn.closest('.resource-list-item');
        var transcript = item && item.querySelector('.resource-transcript');
        if (!transcript) return;

        var wasExpanded = btn.getAttribute('aria-expanded') === 'true';
        closeAllTranscripts();
        if (!wasExpanded) {
            transcript.hidden = false;
            btn.setAttribute('aria-expanded', 'true');
            btn.textContent = '−';
        }
    });

}());
