import json
import os
import random
import time

import vlc


def play(path):
    if os.path.exists('./resources/{}.wav'.format(path)):
        vlc.MediaPlayer('./resources/{}.wav'.format(path)).play()
        time.sleep(1)
    else:
        print('The path \'{}\' does not exist.'.format(path))


def main():
    condolences = ['try again', 'try it one more time']
    congratulations = ['approved', 'go ahead', 'nice', 'ok', 'okeydokey', 'very well']

    with open('data.json', 'r') as json_file:
        irregular_verbs = json.load(json_file)['irregular_verbs']

    while True:
        irregular_verb = random.choice(irregular_verbs)
        irregular_verb, tenses = irregular_verb.values()

        print(irregular_verb)
        print('https://www.merriam-webster.com/dictionary/{}'.format(irregular_verb))
        play(irregular_verb)

        for letter in irregular_verb:
            play(letter)

        for tense in tenses:
            print(tense)
            play(tense)

        while not input('type \'{}\': '.format(irregular_verb)).lower() == irregular_verb:
            play(random.choice(condolences))

        play(random.choice(congratulations))


if __name__ == '__main__':
    main()
