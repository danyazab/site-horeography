// Анімація меню для мобільних пристроїв
const menuToggle = document.getElementById('menu-toggle');
const menuItems = document.querySelector('nav .menu-items');
const menu = document.getElementById('menu');
const menuButton = document.getElementById('menu-button');

if (menuToggle) {
    menuToggle.addEventListener('click', () => {
        menuItems.classList.toggle('active');
        menuToggle.classList.toggle('active');
    });
}

// Плавна прокрутка до секцій
const menuLinks = document.querySelectorAll('nav .menu-items a, .header-buttons .btn');

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
            if (menuToggle) menuToggle.classList.remove('active');
        }
    });
});

// Приховування меню при прокрутці та показ кнопки
let lastScrollTop = 0;
window.addEventListener('scroll', () => {
    let scrollTop = window.pageYOffset || document.documentElement.scrollTop;
    if (scrollTop > lastScrollTop) {
        // Прокрутка вниз
        menu.classList.add('hidden');
        menuButton.classList.add('visible');
    } else {
        // Прокрутка вгору
        menu.classList.remove('hidden');
        menuButton.classList.remove('visible');
    }
    lastScrollTop = scrollTop;
});

// Клік по кнопці для відкриття меню
menuButton.addEventListener('click', () => {
    menu.classList.remove('hidden');
    menuButton.classList.remove('visible');
});

// Анімація елементів при прокрутці з використанням Intersection Observer
function revealOnScroll() {
    const reveals = document.querySelectorAll('section h2, section p, .gallery-container, #contact-form');

    const options = {
        threshold: 0.1
    };

    const observer = new IntersectionObserver((entries, observer) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.classList.add('active');
                observer.unobserve(entry.target);
            }
        });
    }, options);

    reveals.forEach(reveal => {
        observer.observe(reveal);
    });
}

// Викликаємо функцію при завантаженні сторінки
document.addEventListener('DOMContentLoaded', () => {
    revealOnScroll();
});

// Слайдер галереї
const galleryWrapper = document.querySelector('.gallery-wrapper');
const images = document.querySelectorAll('.gallery-wrapper img');
const prevBtn = document.querySelector('.gallery-nav .prev');
const nextBtn = document.querySelector('.gallery-nav .next');

let currentIndex = 0;
let imageWidth = images[0].offsetWidth + 20; // ширина зображення + відступи

// Функція для оновлення ширини зображення при зміні розміру вікна
function updateImageWidth() {
    imageWidth = images[0].offsetWidth + 20;
    maxIndex = images.length - visibleImages();
    moveGallery();
}

// Визначаємо кількість видимих зображень залежно від ширини екрану
function visibleImages() {
    return window.innerWidth <= 768 ? 1 : 3;
}

let maxIndex = images.length - visibleImages(); // кількість зображень мінус видимі

nextBtn.addEventListener('click', () => {
    if (currentIndex < maxIndex) {
        currentIndex++;
        moveGallery();
    }
});

prevBtn.addEventListener('click', () => {
    if (currentIndex > 0) {
        currentIndex--;
        moveGallery();
    }
});

function moveGallery() {
    const translateX = -currentIndex * imageWidth;
    galleryWrapper.style.webkitTransform = `translateX(${translateX}px)`;
    galleryWrapper.style.transform = `translateX(${translateX}px)`;
}

// Анімація переходу між слайдами
galleryWrapper.style.transition = 'transform 0.5s ease-in-out';
galleryWrapper.style.webkitTransition = 'transform 0.5s ease-in-out';

// Додати можливість свайпів на мобільних пристроях
let touchStartX = 0;
let touchEndX = 0;

galleryWrapper.addEventListener('touchstart', (e) => {
    touchStartX = e.changedTouches[0].clientX;
}, false);

galleryWrapper.addEventListener('touchend', (e) => {
    touchEndX = e.changedTouches[0].clientX;
    handleGesture();
}, false);

function handleGesture() {
    if (touchEndX < touchStartX - 30) {
        // Свайп вліво
        if (currentIndex < maxIndex) {
            currentIndex++;
            moveGallery();
        }
    }
    if (touchEndX > touchStartX + 30) {
        // Свайп вправо
        if (currentIndex > 0) {
            currentIndex--;
            moveGallery();
        }
    }
}

// Додано обробку події touchmove для запобігання стандартній поведінці браузера
galleryWrapper.addEventListener('touchmove', (e) => {
    e.preventDefault();
}, false);

// Оновлення ширини зображень при зміні розміру вікна
window.addEventListener('resize', updateImageWidth);
