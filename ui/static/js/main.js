const addLiveClassToCurrentNavLink = () => {
  const navLinks = document.querySelectorAll("nav a");
  for (let i = 0; i < navLinks.length; i++) {
    if (navLinks[i].getAttribute("href") === window.location.pathname) {
      navLinks[i].classList.add("live");
      break;
    }
  }
}

addLiveClassToCurrentNavLink();

document.addEventListener("htmx:afterRequest", function (ent) {
    addLiveClassToCurrentNavLink();
});
