const version = document.querySelector('footer').dataset.version;
const navVersion = document.getElementById('nav-version');
if (navVersion) navVersion.textContent = `v: ${version}`;

// Dropdown toggle: tap/click to expand, again to collapse
document.querySelectorAll('.nav-links .dropdown > a').forEach(trigger => {
    trigger.addEventListener('click', e => {
        e.preventDefault();
        const menu = trigger.nextElementSibling;
        const isOpen = menu.classList.toggle('mobile-open');
        // Close all other open menus
        document.querySelectorAll('.dropdown-menu.mobile-open').forEach(m => {
            if (m !== menu) m.classList.remove('mobile-open');
        });
        // On mobile, position sub-row just below the nav
        if (isOpen && window.innerWidth <= 480) {
            const nav = document.querySelector('nav.site-nav');
            menu.style.top = nav.getBoundingClientRect().bottom + 'px';
        }
    });
});
// Close when clicking outside
document.addEventListener('click', e => {
    if (!e.target.closest('.nav-links')) {
        document.querySelectorAll('.dropdown-menu.mobile-open')
            .forEach(m => m.classList.remove('mobile-open'));
    }
});
