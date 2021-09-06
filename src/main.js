import {infinitive, pastSimple, pastParticiple} from './data.js';

const htmlParagraphElementText = document.getElementById('text');
const htmlButtonElementPlay = document.getElementById('play');
const htmlButtonElementInfinitive = document.getElementById('infinitive');
const htmlButtonElementPastSimple = document.getElementById('past_simple');
const htmlButtonElementPastParticiple = document.getElementById('past_participle');

let irregularVerbs = infinitive.concat(pastSimple)
    .concat(pastParticiple);

htmlButtonElementInfinitive.addEventListener('click', function () {
    checkAnswer(infinitive);
});

htmlButtonElementPastSimple.addEventListener('click', function () {
    checkAnswer(pastSimple);
});

htmlButtonElementPastParticiple.addEventListener('click', function () {
    checkAnswer(pastParticiple);
});

htmlButtonElementPlay.addEventListener('click', function () {
    const htmlAudioElement = new Audio();

    htmlAudioElement.src = './assets/audios/' + htmlParagraphElementText.innerText + '.wav';

    htmlAudioElement.play().then(r => console.log(r));
});

function pickUpIrregularVerb() {
    const index = Math.floor(Math.random() * irregularVerbs.length);

    return irregularVerbs.splice(index, 1)[0]
}

function updateFlashcard() {
    htmlParagraphElementText.className = 'flashcard flashcard--plain';
    htmlParagraphElementText.innerHTML = pickUpIrregularVerb();
}

function checkAnswer(elements) {
    if (elements.includes(htmlParagraphElementText.innerText)) {
        updateFlashcard();
    } else {
        htmlParagraphElementText.className = 'flashcard flashcard--mistake';
    }
}

updateFlashcard();
