const version = document.querySelector('footer').dataset.version;
const navVersion = document.getElementById('nav-version');
if (navVersion) navVersion.textContent = `v: ${version}`;
