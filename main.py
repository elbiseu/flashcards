import json
import os
import re
import time
from random import randint

import vlc
from ibm_cloud_sdk_core.authenticators import IAMAuthenticator
from ibm_watson import TextToSpeechV1


def generate_speech(path, text, text_to_speech):
    if not exists(path):
        with open(file=path, mode='wb') as audio_file:
            audio_file.write(
                text_to_speech.synthesize(
                    text=text,
                    voice='en-US_AllisonV3Voice',
                    accept='audio/wav'
                ).get_result().content)


def exists(path):
    return os.path.exists('./{}'.format(path))


def main():
    text_to_speech_iam_apikey = os.getenv('TEXT_TO_SPEECH_IAM_APIKEY')
    text_to_speech_url = os.getenv('TEXT_TO_SPEECH_URL')

    authenticator = IAMAuthenticator(text_to_speech_iam_apikey)
    text_to_speech = TextToSpeechV1(
        authenticator=authenticator
    )

    text_to_speech.set_service_url(text_to_speech_url)

    with open(file='data.json', mode='r') as json_file:
        irregular_verbs = json.load(json_file)['irregular_verbs']

    generate_speech(path='good.wav',
                    text='good!',
                    text_to_speech=text_to_speech)

    generate_speech(path='try_again.wav',
                    text='try again!',
                    text_to_speech=text_to_speech)

    for irregular_verb in irregular_verbs:
        irregular_verb, tenses = irregular_verb.values()
        generate_speech(path='{}.{}'.format(irregular_verb, 'wav'),
                        text='{}, {}'.format(irregular_verb, ', and '.join(tenses)),
                        text_to_speech=text_to_speech)
        generate_speech(path='{}_spell.{}'.format(irregular_verb, 'wav'),
                        text=', '.join([letter for letter in re.split(r'(\w)', irregular_verb) if letter.isalpha()]),
                        text_to_speech=text_to_speech)

    while True:
        irregular_verb = irregular_verbs[randint(0, len(irregular_verbs) - 1)]
        irregular_verb, tenses = irregular_verb.values()
        print('\n{}\n{}\n'.format(irregular_verb, ', and '.join(tenses)))
        vlc.MediaPlayer('./{}.{}'.format(irregular_verb, 'wav')).play()
        time.sleep(3)
        vlc.MediaPlayer('./{}_spell.{}'.format(irregular_verb, 'wav')).play()
        print(' '.join([letter for letter in re.split(r'(\w)', irregular_verb) if letter.isalpha()]))
        while input('type \'{}\': '.format(irregular_verb)).lower() != irregular_verb:
            vlc.MediaPlayer('./try_again.wav').play()
            time.sleep(1)
        vlc.MediaPlayer('./good.wav').play()
        time.sleep(1)


if __name__ == '__main__':
    main()
