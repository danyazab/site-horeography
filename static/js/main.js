window.onload = () => {
    // Якщо є прелоадер, ховаємо його
    const preloader = document.querySelector('.loader');
    if (preloader) {
        preloader.style.display = 'none';
    }
};

const header = document.querySelector('.main-header');
const scrollMenuBtn = document.querySelector('.scroll-menu-btn');
const navList = document.querySelector('.menu');
const toggleButton = document.querySelector('.mobile-menu-toggle');
const body = document.body; // Для блокування скролу
let lastScrollTop = 0;

/* ============ ЛОГІКА ПРИ СКРОЛІ (ХОВАТИ/ПОКАЗУВАТИ ШАПКУ) ============ */
window.addEventListener('scroll', () => {
    const st = window.pageYOffset || document.documentElement.scrollTop;
    if (st > lastScrollTop && st > 110) {
        // Скролимо вниз
        header.classList.add('hide-header');
        scrollMenuBtn.style.display = 'flex';
    } else {
        // Скролимо вгору
        header.classList.remove('hide-header');
        scrollMenuBtn.style.display = 'none';
    }
    lastScrollTop = st <= 0 ? 0 : st;

    // Можна додати ефект масштабування слайдера, якщо він є
    const sliderContainer = document.querySelector('.slider-container');
    if (sliderContainer) {
        const scaleVal = 1 + window.scrollY / 2000;
        sliderContainer.style.transform = `scale(${Math.min(scaleVal, 1.1)})`;
    }
}, false);

/* ============ ВІДКРИТТЯ/ЗАКРИТТЯ МЕНЮ (МОБІЛЬНЕ) ============ */
toggleButton.addEventListener('click', () => {
    const isMenuOpen = navList.classList.toggle('show');

    // Ховаємо шапку та блокуємо скрол, якщо меню відкрите
    if (isMenuOpen) {
        header.classList.add('hide-header');
        body.style.overflow = 'hidden'; // Блокування скролу
        scrollMenuBtn.style.display = 'flex';

    } else {
        header.classList.remove('hide-header');
        scrollMenuBtn.style.display = 'none';
        body.style.overflow = ''; // Розблокування скролу
    }
});

/* Якщо хочете, щоб ще й ця кнопка відкривала/закривала меню */
scrollMenuBtn.addEventListener('click', () => {

    if (window.innerWidth >= 1124) {
        navList.classList.toggle('show');
        if (window.innerWidth >= 768) {
            header.classList.remove('hide-header');
            scrollMenuBtn.style.display = 'none';

        }
    } else {
        const isMenuOpen = navList.classList.toggle('show');

        if (isMenuOpen) {
            header.classList.add('hide-header');
            if (window.innerWidth <= 1124) {
                body.style.overflow = 'hidden'; // Блокування скролу
            }
        } else {
            header.classList.remove('hide-header');
            scrollMenuBtn.style.display = 'none';
            body.style.overflow = ''; // Розблокування скролу


        }
    }
        // Ховаємо шапку та блокуємо скрол, якщо меню відкрите

});

/* ============ ВІДКРИТТЯ/ЗАКРИТТЯ ПІДМЕНЮ НА МОБІЛЬНИХ (АКОРДЕОН) ============ */
document.querySelectorAll('.item.has-submenu > .link').forEach(link => {
    link.addEventListener('click', (e) => {
        // Спрацьовує тільки якщо ширина <= 768px
        if (window.innerWidth <= 1124) {
            e.preventDefault();
            const parent = link.closest('.item.has-submenu');
            const submenu = parent.querySelector('.submenu');

            // Перемикаємо клас .open
            if (submenu.classList.contains('open')) {
                submenu.classList.remove('open');
                submenu.style.maxHeight = '0';

            } else {
                submenu.classList.add('open');
                submenu.style.maxHeight = submenu.scrollHeight + 'px';
            }
        }
    });
});

/* ============ АНІМАЦІЯ ПРИ СКРОЛІ (Intersection Observer) ============ */
const animatedElements = document.querySelectorAll('.animate-on-scroll, .adv-item');
const observer = new IntersectionObserver(entries => {
    entries.forEach(entry => {
        if (entry.isIntersecting) {
            entry.target.classList.add('visible');
        } else {
            entry.target.classList.remove('visible');
        }
    });
}, {threshold: 0.1});

animatedElements.forEach(el => observer.observe(el));

/* ============ ВІДПРАВКА ФОРМИ (ОПЦІЙНО) ============ */
const form = document.getElementById('contact-form');
const successMsg = document.getElementById('form-success-message');
if (form) {
    form.addEventListener('submit', (e) => {
        e.preventDefault();
        if (!form.checkValidity()) {
            form.reportValidity();
            return;
        }
        const formData = new FormData(form);
        fetch('/contact', {
            method: 'POST',
            body: formData
        })
            .then(res => res.json())
            .then(data => {
                if (data.status === 'success') {
                    successMsg.querySelector('p').textContent = data.message;
                    successMsg.style.display = 'block';
                    form.reset();
                    setTimeout(() => {
                        successMsg.style.display = 'none';
                    }, 4000);
                }
            })
            .catch(err => console.log(err));
    });
}
