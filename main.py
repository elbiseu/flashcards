import enum
import json
import os
import random
import re
import time
from datetime import datetime

import vlc


class Hello(enum.Enum):
    GoodMorning = 'good_morning'
    GoodAfternoon = 'good_afternoon'
    GoodEvening = 'good_evening'
    Morning = 'morning'


class Tense(enum.Enum):
    Infinitive = 'infinitive'
    PastSimple = 'past_simple'
    PastParticiple = 'past_participle'


def play(path):
    if os.path.exists('./{}.wav'.format(path)):
        vlc.MediaPlayer('./{}.wav'.format(path)).play()
    else:
        print('The path \'{}\' does not exist.'.format(path))


def main():
    now = datetime.now()

    with open(file='data.json', mode='r') as json_file:
        irregular_verbs = json.load(json_file)['irregular_verbs']

    if 23 > now.hour < 12:
        if random.choice([True, False]):
            play(Hello.GoodMorning.value)
        else:
            play(Hello.Morning.value)
    elif 18 > now.hour > 11:
        play(Hello.GoodAfternoon.value)
    else:
        play(Hello.GoodEvening.value)

    time.sleep(1)

    while True:
        irregular_verb = random.choice(irregular_verbs)
        irregular_verb, tenses = irregular_verb.values()

        print(irregular_verb)
        play(irregular_verb)
        time.sleep(1)

        for letter in re.split(r'(\w)', irregular_verb):
            if letter.isalpha():
                print(letter, end=' ')
                play(letter)
                time.sleep(1)

        print('\n{}'.format(', and '.join(tenses)))

        if 'infinitive' in tenses:
            play(Tense.Infinitive.value)
            time.sleep(1)
        if 'past simple' in tenses:
            play(Tense.PastSimple.value)
            time.sleep(1)
        if 'past participle' in tenses:
            play(Tense.PastParticiple.value)
            time.sleep(1)

        while input('type \'{}\': '.format(irregular_verb)).lower() != irregular_verb:
            play(random.choice(['try_again', 'try_it_one_more_time']))
            time.sleep(1)

        play(random.choice(['approved', 'go_ahead', 'nice', 'ok', 'okeydokey', 'very_well']))
        time.sleep(1)


if __name__ == '__main__':
    main()
