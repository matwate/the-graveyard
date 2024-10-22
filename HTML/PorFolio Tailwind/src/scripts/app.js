const btn = document.querySelector(".languageChange");

const englishText = document.querySelectorAll(".en-text");
const spanishText = document.querySelectorAll(".es-text");

let isEnglish = true;
function changeLanguage() {
  console.log("The Function Activates");
  if (isEnglish) {
    for (var i = 0; i < englishText.length; i++) {
      englishText[i].classList.add("hidden");
    }
    for (var i = 0; i < spanishText.length; i++) {
        spanishText[i].classList.remove("hidden");
      }
    isEnglish = false;
  } else {
    for (var i = 0; i < englishText.length; i++) {
      englishText[i].classList.remove("hidden");
    }
    for (var i = 0; i < spanishText.length; i++) {
        spanishText[i].classList.add("hidden");
    }
    isEnglish = true;
  }
}

btn.addEventListener("click", changeLanguage);
