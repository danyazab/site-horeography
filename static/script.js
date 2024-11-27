// Анімація меню для мобільних пристроїв
const menuToggle = document.getElementById('menu-toggle');
const menuItems = document.querySelector('nav .menu-items');
const menu = document.getElementById('menu');

menuToggle.addEventListener('click', () => {
    menuItems.classList.toggle('active');
    menuToggle.classList.toggle('active');
});

// Плавна прокрутка до секцій
const menuLinks = document.querySelectorAll('nav .menu-items a');

menuLinks.forEach(link => {
    link.addEventListener('click', (e) => {
        e.preventDefault();
        const targetId = link.getAttribute('href').substring(1);
        const targetSection = document.getElementById(targetId);

        window.scrollTo({
            top: targetSection.offsetTop - 60,
            behavior: 'smooth'
        });

        // Закриваємо меню на мобільних після вибору
        if (menuItems.classList.contains('active')) {
            menuItems.classList.remove('active');
            menuToggle.classList.remove('active');
        }
    });
});

// Приховування меню при прокрутці
let lastScrollTop = 0;
window.addEventListener('scroll', () => {
    let scrollTop = window.pageYOffset || document.documentElement.scrollTop;
    if (scrollTop > lastScrollTop) {
        // Прокрутка вниз
        menu.classList.add('hidden');
    } else {
        // Прокрутка вгору
        menu.classList.remove('hidden');
    }
    lastScrollTop = scrollTop;
});

// Слайдер галереї
const slides = document.querySelector('.slides');
const images = document.querySelectorAll('.slides img');
const prevBtn = document.querySelector('.prev');
const nextBtn = document.querySelector('.next');

let counter = 0;
const size = images[0].clientWidth;

nextBtn.addEventListener('click', () => {
    if (counter >= images.length - 1) return;
    counter++;
    slides.style.transform = 'translateX(' + (-size * counter) + 'px)';
});

prevBtn.addEventListener('click', () => {
    if (counter <= 0) return;
    counter--;
    slides.style.transform = 'translateX(' + (-size * counter) + 'px)';
});

// Анімація назви при прокрутці
const mainTitle = document.querySelector('.main-title');
window.addEventListener('scroll', () => {
    const titlePosition = mainTitle.getBoundingClientRect().top;
    const screenPosition = window.innerHeight / 1.5;

    if (titlePosition < screenPosition) {
        mainTitle.classList.add('active');
    } else {
        mainTitle.classList.remove('active');
    }
});
