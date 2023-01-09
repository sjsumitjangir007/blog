let anchors = document.querySelectorAll('a');
let btns = document.querySelectorAll('button');
anchors.forEach((elm) => {
    // elm.href = "#";
    // elm.addEventListener('click', (e) => {
    //     e.preventDefault();
    // })
})
btns.forEach((btn) => {
    btn.addEventListener('click', (e) => {
        e.preventDefault();
    });
})