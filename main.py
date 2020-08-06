import json
import os
import time

import vlc
from ibm_cloud_sdk_core.authenticators import IAMAuthenticator
from ibm_watson import TextToSpeechV1


def generate_speech(path, text, text_to_speech):
    if exists(path):
        print(text)
    else:
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

    for irregular_verb in irregular_verbs:
        irregular_verb, tenses = irregular_verb.values()
        generate_speech(path='{}.{}'.format(irregular_verb, 'wav'),
                        text='{}, {}'.format(irregular_verb, ', and '.join(tenses)),
                        text_to_speech=text_to_speech)

    for irregular_verb in irregular_verbs:
        irregular_verb, tenses = irregular_verb.values()
        path = '{}.{}'.format(irregular_verb, 'wav')
        if exists(path):
            print('./{}'.format(path))
            vlc.MediaPlayer('./{}'.format(path)).play()
            time.sleep(5)


if __name__ == '__main__':
    main()
