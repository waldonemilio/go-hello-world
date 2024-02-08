// Inside public/script.js
const menuItems = document.querySelectorAll('nav li');

menuItems.forEach(item => {
    item.addEventListener('click', () => {
        alert('You clicked me! Add your navigation logic here.');
    });
});
