import json
import random

from dictionary import definition
from player import play


def main():
    condolences = ['try again', 'try it one more time']
    congratulations = ['approved', 'go ahead', 'nice', 'ok', 'okeydokey', 'very well']

    with open('data.json', 'r') as json_file:
        irregular_verbs = json.load(json_file)['irregular_verbs']

    while True:
        irregular_verb, tenses = random.choice(irregular_verbs).values()

        print(irregular_verb)
        play(irregular_verb)

        if input('Do you want to know the meaning of this word?\nIf yes, type \'yes\': ') == 'yes':
            definition(irregular_verb)

        for letter in irregular_verb:
            play(letter)

        for tense in tenses:
            print(tense)
            play(tense)

        while not input('Type the word\'{}\': '.format(irregular_verb)) == irregular_verb:
            play(random.choice(condolences))

        play(random.choice(congratulations))


if __name__ == '__main__':
    main()
