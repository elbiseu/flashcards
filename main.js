let listOfIrregularVerbs = [
  {
    id: 1,
    infinitive: "be",
    past_simple: "was/were",
    past_participle: "been"
  },
  {
    id: 2,
    infinitive: "beat",
    past_simple: "beat",
    past_participle: "beaten"
  },
  {
    id: 3,
    infinitive: "become",
    past_simple: "became",
    past_participle: "become"
  },
  {
    id: 4,
    infinitive: "begin",
    past_simple: "began",
    past_participle: "begun"
  },
];

for (let irregularVerb of listOfIrregularVerbs) {
  for (let key of Object.keys(irregularVerb)) {
    if (key !== "id") {
      const flashcard = document.createElement("div");
      flashcard.id = irregularVerb["id"] + key;
      flashcard.className = "flashcard";
      flashcard.draggable = true;
      flashcard.innerText = irregularVerb[key];
      flashcard.onclick = function () {
        const src = "./resources/" + irregularVerb[key] + ".wav";
        const pronunciation = new Audio(src);
        const promise = pronunciation.play();
        if (promise !== null) {
          promise.then(function () {
            console.log("Automatic playback started!");
          }).catch(function (error) {
            console.log("Automatic playback failed!");
            console.log(error);
          });
        }
      };
      flashcard.ondragstart = function (event) {
        event.dataTransfer.setData("text/plain", this.id);
      };
      const element = document.getElementById("flashcard_container");
      element.appendChild(flashcard);
    }
  }
}

for (let classification of ["infinitive", "past_simple", "past_participle", "flashcard"]) {
  const id = classification + "_container";
  const container = document.getElementById(id);
  container.ondragend = function () {
    container.style.removeProperty("background-color");
  };
  container.ondragleave = function () {
    container.style.removeProperty("background-color");
  };
  container.ondragover = function (event) {
    event.preventDefault();
    container.style.backgroundColor = "lightgreen";
  };
  container.ondrop = function (event) {
    const data = event.dataTransfer.getData("text/plain");
    const element = document.getElementById(data);
    const target = event.target;
    target.appendChild(element);
    container.style.removeProperty("background-color");
  };
}
