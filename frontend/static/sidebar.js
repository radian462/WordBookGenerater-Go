fetch('./sidebar.html')
  .then(res => res.text())
  .then(html => {
    const container = document.getElementById('sidebar');
    container.innerHTML = html;

    const path = window.location.pathname;
    const filename = path.split('/').pop();
    const navLinks = container.querySelectorAll(".nav-link");
    navLinks.forEach(link => {
      const href = link.getAttribute("href");
      if (href.split('/').pop() === filename) {
        link.classList.add("active");
        link.classList.remove("text-white");
      } else {
        link.classList.remove("active");
        link.classList.add("text-white");
      }
    });
  });
