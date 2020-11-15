let irregularVerbs = infinitive.concat(pastSimple);
irregularVerbs = irregularVerbs.concat(pastParticiple);
let identifiers = [];
console.log(identifiers.length);

while (identifiers.length !== irregularVerbs.length) {
  const random = Math.random() * irregularVerbs.length;
  const floor = Math.floor(random);
  if (!identifiers.includes(floor)) {
    identifiers.push(floor);
  }
}

for (let identifier of identifiers) {
  let irregularVerb = irregularVerbs[identifier];
  const flashcard = document.createElement("div");
  flashcard.id = identifier;
  flashcard.className = "flashcard";
  flashcard.draggable = true;
  flashcard.innerText = irregularVerb;
  flashcard.classifications = [];
  flashcard.classifications.push("flashcard");
  if (infinitive.includes(irregularVerb)) {
    flashcard.classifications.push("infinitive");
  }
  if (pastSimple.includes(irregularVerb)) {
    flashcard.classifications.push("past_simple");
  }
  if (pastParticiple.includes(irregularVerb)) {
    flashcard.classifications.push("past_participle");
  }
  flashcard.onclick = function () {
    const src = "./resources/" + irregularVerb + ".wav";
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

for (let classification of ["flashcard", "infinitive", "past_simple", "past_participle"]) {
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
    if (target.className !== "flashcard") {
      if (element.classifications.includes(classification)) {
        element.style.backgroundColor = "white";
        element.style.borderColor = "black";
        element.style.color = "black";
      } else {
        element.style.backgroundColor = "red";
        element.style.borderColor = "red";
        element.style.color = "white";
      }
      target.appendChild(element);
      container.style.removeProperty("background-color");
    }
  };
}
