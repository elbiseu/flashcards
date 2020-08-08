#!/usr/bin/python

import json
import random

from dictionary import definition
from player import play
from translator import translate


def main():
    condolences = ['try again', 'try it one more time']
    congratulations = ['approved', 'go ahead', 'nice', 'ok', 'okeydokey', 'very well']

    with open('data.json', 'r') as json_file:
        irregular_verbs = json.load(json_file)['irregular_verbs']

    randomly = input('Do you want to learn the English irregular verbs randomly?\nType \'y\' to learn the English '
                     'irregular verbs randomly.\nType any other letter to continue: ')

    infinitely = input('Do you want to run the script infinitely?\nType \'y\' to run the script infinitely.\nType any '
                       'other letter to continue: ')

    while True:
        if randomly:
            irregular_verb, tenses = random.choice(irregular_verbs).values()
        else:
            irregular_verbs.pop()
            irregular_verb, tenses = irregular_verbs.values()

        print('Irregular verb: {}'.format(irregular_verb))
        play(irregular_verb)

        print('Spelling: ', end='')

        for letter in irregular_verb:
            print(letter, end=' ')
            play(letter)

        print('\nTenses: {}'.format(', and '.join(tenses)))

        for tense in tenses:
            play(tense)

        response = input('Type the word \'{}\': '.format(irregular_verb)).split()

        while irregular_verb not in response:
            play(random.choice(condolences))
            response[0] = input('Sorry! Type the word \'{}\' again: '.format(irregular_verb))

        if '-d' in response:
            definition(irregular_verb)

        if '-t' in response:
            translate(irregular_verb)

        play(random.choice(congratulations))

        if not infinitely:
            if input('Do you want to continue?\nType \'n\' to end. Type any other letter to continue: ') == 'n':
                break


if __name__ == '__main__':
    main()
