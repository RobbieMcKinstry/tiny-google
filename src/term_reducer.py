#!/usr/bin/env python

import sys
from operator import itemgetter

#input - word\tdocument:frequency
#ouput - word: document1\tfrequency, document2\tfrequency, etc.

current_word = None
value = None
word = None
doc_freq_table = []

for line in sys.stdin:
    #Strip \n off the right side
    #For some reason .strip() would take away the 'word\t'
    line = line.rstrip('\n')

    #Get the word and document:frequency
    word, doc_freq = line.split('\t')
    doc_name, frequency = doc_freq.split(':')

    #Cast is necessary here in order for the .sort() to work later
    frequency = int(frequency)

    #Same idea as the word count reducer
    if current_word == word:
        doc_freq_table.append([doc_name, frequency])
    else:
        if current_word:
            #Sort the list by frequency in descending order
            doc_freq_table.sort(key=itemgetter(1), reverse=True)
            
            print('%s\n{\n') % (word),
            #Iterate through the sorted list
            for value in doc_freq_table:
                print('Document:\t%s\nFrequency:\t%s\n') % (value[0], value[1]),
            print('}\n')
            #Clear the list
            doc_freq_table = []

        current_word = word
        doc_freq_table.append([doc_name, frequency])
if current_word == word:
    #Sort the list by frequency in descending order
    doc_freq_table.sort(key=itemgetter(1), reverse=True)
    
    print('%s\n{\n') % (word),
    #Iterate through the sorted list
    for value in doc_freq_table:
        print('\n{\nDocument:\t%s\nFrequency:\t%s\n') % (value[0], value[1]),
    print('}\n')
    doc_freq_table = []