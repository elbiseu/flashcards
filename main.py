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

    while True:
        irregular_verb, tenses = random.choice(irregular_verbs).values()

        print('Irregular verb: {}'.format(irregular_verb))
        play(irregular_verb)

        print('Spelling: ', end='')

        for letter in irregular_verb:
            print(letter, end=' ')
            play(letter)

        print('\nTenses: {}'.format(', and '.join(tenses)))

        for tense in tenses:
            play(tense)

        play('type')

        while not input('Type the word \'{}\': '.format(irregular_verb)) == irregular_verb:
            play(random.choice(condolences))

        play(random.choice(congratulations))
        play('definition')

        if input('Do you want to know the meaning of this word?\nIf yes, type \'y\': ') == 'y':
            play('definitions')
            definition(irregular_verb)

        play('translate')

        if input('Do you want to translate this word?\nIf yes, type \'y\': ') == 'y':
            play('translations')
            translate(irregular_verb)

        play('continue')

        if not input('Do you want to continue?\nIf yes, type \'y\': ') == 'y':
            break


if __name__ == '__main__':
    main()
