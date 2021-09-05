import {infinitive, pastSimple, pastParticiple} from './data.js';

let irregularVerbs = infinitive.concat(pastSimple)
    .concat(pastParticiple);

const flashcardContainer = document.getElementById('flashcard_container');
const infinitivePod = 'infinitive_pod';
const pastSimplePod = 'past_simple_pod';
const pastParticiplePod = 'past_participle_pod';

function pickUpIrregularVerb() {
    return irregularVerbs.pop()
}

function buildPod(elementId, irregularVerbs) {
    const pod = document.getElementById(elementId);

    pod.addEventListener('dragover', function (dragEvent) {
        if (pod.className.endsWith('pod--drag-over') === false) {
            pod.className = 'pod pod--drag-over';
        }
    });

    pod.addEventListener('dragleave', function (dragEvent) {
        if (pod.className.endsWith('pod--drag-over') !== false) {
            pod.className = 'pod';
        }
    });

    pod.addEventListener('dragover', function (dragEvent) {
        dragEvent.preventDefault();
    });

    pod.addEventListener('drop', function (dragEvent) {
        dragEvent.preventDefault();

        let flashcardId = dragEvent.dataTransfer.getData('text/plain');
        let flashcard = document.getElementById(flashcardId);

        pod.className = 'pod';

        if (irregularVerbs.includes(flashcard.innerText)) {
            flashcard.className = 'flashcard flashcard--fade-out flashcard--success';
            flashcard.draggable = false;
            flashcard.id = undefined;

            setTimeout(function () {
                flashcard.remove();
            }, 6000);

            buildFlashcard();
        } else {
            flashcard.className = 'flashcard flashcard--mistake';
        }

        pod.appendChild(flashcard);
    });
}

function buildFlashcard() {
    const irregularVerb = pickUpIrregularVerb();
    const flashcard = document.createElement('div');

    flashcard.draggable = true;
    flashcard.className = 'flashcard flashcard--fade-in flashcard--plain';
    flashcard.id = 'flashcard';
    flashcard.innerHTML = `<span>${irregularVerb}</span>`;

    flashcardContainer.appendChild(flashcard);

    flashcard.addEventListener('dragstart', function (dragEvent) {
        dragEvent.dataTransfer.setData('text/plain', flashcard.id);
    });

    flashcard.addEventListener('click', function (mouseEvent) {
        mouseEvent.preventDefault();

        const pronunciation = new Audio();
        pronunciation.src = './resources/' + irregularVerb + '.wav';
        const promise = pronunciation.play();

        promise.then(function () {
            console.log("Automatic playback started!");
        }).catch(function (error) {
            console.log("Automatic playback failed!");
            console.log(error);
        });
    });
}

buildFlashcard();
buildPod(infinitivePod, infinitive);
buildPod(pastSimplePod, pastSimple);
buildPod(pastParticiplePod, pastParticiple);
