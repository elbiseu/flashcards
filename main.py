from ibm_cloud_sdk_core.authenticators import IAMAuthenticator
from ibm_watson import TextToSpeechV1
import time
import enum
import os
import vlc
import json


class Voice(enum.Enum):
    Allison = 'en-US_AllisonV3Voice'


class Accept(enum.Enum):
    WAV = 'wav'


class Tense(enum.Enum):
    Infinitive = 'infinitive'
    PastSimple = 'past simple'
    PastParticiple = 'past participle'


class Speech:
    def __init__(self, text: str, voice: str, accept: str, text_to_speech: TextToSpeechV1):
        self.text = text
        self.voice = voice
        self.accept = accept
        self.audio_file = '{}.{}'.format(text, accept)
        self.text_to_speech = text_to_speech

    @property
    def text(self):
        return self.__text

    @text.getter
    def text(self):
        return self.text

    @text.setter
    def text(self, text: str):
        self.__text = text

    def play(self):
        if not (self.__exists()):
            self.__generate_speech()
        vlc.MediaPlayer(self.audio_file).play()

    def __exists(self):
        return os.path.exists('./{}'.format(self.audio_file))

    def __generate_speech(self):
        with open(file=self.audio_file, mode='wb') as audio_file:
            audio_file.write(
                self.text_to_speech.synthesize(
                    text=self.text,
                    voice=self.voice,
                    accept='audio/{}'.format(self.accept)
                ).get_result().content)


class Verb(Speech):
    def __init__(self, text: str, voice: str, accept: str, text_to_speech: TextToSpeechV1, type: str, tenses: list):
        super().__init__(text, voice, accept, text_to_speech)
        self.type = type
        self.tenses = tenses

    @property
    def type(self):
        return self._type

    @type.getter
    def type(self):
        return self.type

    @type.setter
    def type(self, value):
        self._type = value

    @property
    def tenses(self):
        return self._tenses

    @tenses.getter
    def tenses(self):
        return self.tenses

    @tenses.setter
    def tenses(self, value):
        self._tenses = value


def main():
    apikey = os.getenv('TEXT_TO_SPEECH_IAM_APIKEY')
    service_url = os.getenv('TEXT_TO_SPEECH_URL')

    authenticator = IAMAuthenticator(apikey)
    text_to_speech = TextToSpeechV1(
        authenticator=authenticator
    )

    text_to_speech.set_service_url(service_url)

    speeches = []

    with open(file='data.json', mode='r') as json_file:
        data = json.load(json_file)
        for verb in data['verbs']:
            text, properties = verb.values()
            speech = Speech(text=text,
                            voice=Voice.Allison.value,
                            accept=Accept.WAV.value,
                            text_to_speech=text_to_speech)
            speeches.append(speech)

    while True:
        for speech in speeches:
            print(speech)
            speech.play()
            time.sleep(1)


if __name__ == '__main__':
    main()
