import json
import os
import random
import re
import time
from datetime import datetime

import vlc


def play(path):
    if os.path.exists('./resources/{}.wav'.format(path)):
        vlc.MediaPlayer('./resources/{}.wav'.format(path)).play()
        time.sleep(1)
    else:
        print('The path \'{}\' does not exist.'.format(path))


def main():
    now = datetime.now()

    with open('data.json', 'r') as json_file:
        irregular_verbs = json.load(json_file)['irregular_verbs']

    if 23 > now.hour < 12:
        play(random.choice(['good_morning', 'morning']))
    elif 18 > now.hour > 11:
        play('good_afternoon')
    else:
        play('good_evening')

    while True:
        irregular_verb = random.choice(irregular_verbs)
        irregular_verb, tenses = irregular_verb.values()

        print(irregular_verb)
        play(irregular_verb)

        for letter in re.split(r'(\w)', irregular_verb):
            if letter.isalpha():
                play(letter)

        for tense in tenses:
            play(tense)

        while input('type \'{}\': '.format(irregular_verb)).lower() != irregular_verb:
            play(random.choice(['try_again', 'try_it_one_more_time']))

        play(random.choice(['approved', 'go_ahead', 'nice', 'ok', 'okeydokey', 'very_well']))


if __name__ == '__main__':
    main()
