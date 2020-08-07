import json
import os
from urllib import request


def definition(word: str):
    key_school_dictionary = os.getenv('KEY_SCHOOL_DICTIONARY')
    request_url = 'https://www.dictionaryapi.com/api/v3/references/sd4/json/{}?key={}'.format(word, key_school_dictionary)
    response = request.urlopen(request_url).read().decode('utf-8')
    for definitions in json.loads(response):
        for definition in definitions['shortdef']:
            print('* {}'.format(definition.replace(' : ', ' ')))
