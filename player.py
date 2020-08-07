import os
import time

import vlc


def play(file: str):
    if os.path.exists('./resources/{}.wav'.format(file)):
        vlc.MediaPlayer('./resources/{}.wav'.format(file)).play()
        time.sleep(1)
