from ibm_cloud_sdk_core.authenticators import IAMAuthenticator
from ibm_watson import TextToSpeechV1
from googletrans import Translator
from random import randint
import time
import enum
import os
import vlc


class Voice(enum.Enum):
    Allison = 'en-US_AllisonV3Voice'


class Accept(enum.Enum):
    WAV = 'wav'


class Text:
    def __init__(self, text: str, voice: str, accept: str, text_to_speech: TextToSpeechV1):
        self.text = text
        self.voice = voice
        self.accept = accept
        self.audio_file = '%s.%s' % (self.text, self.accept)
        self.text_to_speech = text_to_speech

    def play(self):
        if not (self.__exists()):
            self.__generate_speech()
        vlc.MediaPlayer(self.audio_file).play()

    def __exists(self):
        return os.path.exists('./%s' % self.audio_file)

    def __generate_speech(self):
        with open(file=self.audio_file, mode='wb') as audio_file:
            audio_file.write(
                self.text_to_speech.synthesize(
                    text=self.text,
                    voice=self.voice,
                    accept='audio/%s' % self.accept
                ).get_result().content)


def main():
    apikey = os.getenv('TEXT_TO_SPEECH_IAM_APIKEY')
    service_url = os.getenv('TEXT_TO_SPEECH_URL')

    authenticator = IAMAuthenticator(apikey)
    text_to_speech = TextToSpeechV1(
        authenticator=authenticator
    )

    text_to_speech.set_service_url(service_url)

    translator = Translator()
    language = 'en'
    file = open('words.txt', 'r')

    list_of_words = []
    translator_words = []

    missing_words = {'stir': 'revolver', 'carve': 'trinchar', 'break': 'romper', 'beat': 'batir', }

    # translator_words.append(translator.translate(line, src=language, dest='es').text)

    text = Text('been', Voice.Allison.value, Accept.WAV.value, text_to_speech)
    text.play()

    time.sleep(1)


if __name__ == "__main__":
    main()
