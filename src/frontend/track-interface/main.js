import './fonts/neue-montreal/neue-montreal.css';
import './css/layout.css';
import './css/style.css';

function main() {
    const contentContainer = document.getElementById("contentContainer");
    const refreshButton = document.getElementById("BUTTON_refreshFeed");

    // Add refresh button listener
    refreshButton.addEventListener("click", loadContent)
    loadContent();
}

async function loadContent() {
    console.log("load content");
}


// Make sure that all the code only ever runs once.
if (!window.hasCodeRun) {
    main();

    // Set the flag to indicate that the code has run
    window.hasCodeRun = true;
}