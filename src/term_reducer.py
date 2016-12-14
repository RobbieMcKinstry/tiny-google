#!/usr/bin/env python

import sys

#input - word\tdocument:frequency
#ouput - word: document1\tfrequency, document2\tfrequency, etc.

current_word = None
value = None
word = None


for line in sys.stdin:
    #Strip \n off the right side
    #For some reason .strip() would take away the 'word\t'
    line = line.rstrip('\n')

    #Get the word and document:frequency
    word, doc_freq = line.split('\t')

    #Same idea as the word count reducer
    if current_word == word:
        value = ', '.join([doc_freq, value])
    else:
        if current_word:
            print('%s\t%s') % (current_word, value)
        current_word = word
        value = doc_freq
if current_word == word:
    print('%s\t%s') % (current_word, value)