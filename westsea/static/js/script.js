// static/js/script.js

document.addEventListener('DOMContentLoaded', function () {
    console.log("JavaScript loaded!");

    const header = document.querySelector('header');
    header.addEventListener('click', () => {
        alert('You clicked the header!');
    });
});
