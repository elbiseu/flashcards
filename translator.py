#!/usr/bin/python

import json
import os
from urllib import request


def translate(word: str):
    key_spanish = os.getenv('KEY_SPANISH')
    request_url = 'https://www.dictionaryapi.com/api/v3/references/spanish/json/{}?key={}'.format(word, key_spanish)
    response = request.urlopen(request_url).read().decode('utf-8')
    for definitions in json.loads(response):
        for definition in definitions['shortdef']:
            print('* {}'.format(definition))
