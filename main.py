#!/usr/bin/python

import json
import random

from dictionary import definition
from player import play
from translator import translate


def main():
    randomly = False
    infinitely = True
    congratulations = ['approved', 'go ahead', 'nice', 'ok', 'okeydokey', 'very well']

    with open('data.json', 'r') as json_file:
        irregular_verbs = json.load(json_file)['irregular_verbs']

    while True:
        irregular_verb = random.choice(irregular_verbs) if randomly else irregular_verbs.pop(0)
        word = irregular_verb["word"]
        # root = irregular_verb["root"]
        tenses = irregular_verb["tenses"]
        # translations = irregular_verb["translations"]

        print('Irregular verb: {}'.format(word))
        play(word)

        print('\nTenses: {}'.format(', and '.join(tenses)))

        for tense in tenses:
            play(tense)

        arguments = input('More information about \'{}\'?: '.format(word)).split()
        for argument in arguments:
            if '-d' == argument:
                print('Definition: ', end='')
                definition(word)
            elif '-t' == argument:
                print('Translation: ', end='')
                translate(word)
            elif '-s' == argument:
                print('Spelling: ', end='')
                for letter in word:
                    print(letter, end=' ')
                    play(letter)

        play(random.choice(congratulations))

        if not infinitely:
            if input('Do you want to continue?\nType \'n\' to end. Type any other letter to continue: ') == 'n':
                break


if __name__ == '__main__':
    main()
