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

/*
function getRandomIrregularVerb() {
  const random = Math.random() * listOfIrregularVerbs.length;
  const floor = Math.floor(random);
  return listOfIrregularVerbs[floor]
}
*/

for (let irregularVerb of listOfIrregularVerbs) {
  for (let key of Object.keys(irregularVerb)) {
    if (key !== "id") {
      let flashcard = document.createElement("div");
      flashcard.id = irregularVerb["id"];
      flashcard.className = "flashcard";
      flashcard.draggable = true;
      flashcard.innerText = irregularVerb[key];
      document.getElementById("flashcard_container").appendChild(flashcard);
    }
  }
}
