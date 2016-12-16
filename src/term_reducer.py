#!/usr/bin/env python

import sys
from operator import itemgetter
import json

#input - word\tpath**document:frequency
#ouput - word: 
#              document1
#              path
#              frequency
#
#              document2
#              path
#              frequency, etc.

current_word = None
value = None
word = None
doc_freq_table = []
word_dict = {}

for line in sys.stdin:
    #Strip \n off the right side
    #For some reason .strip() would take away the 'word\t'
    line = line.rstrip('\n')


    #Get the word and document:frequency
    word, doc_freq = line.split('\t')
    #I got an error if I didn't do the split this way...
    splurg = doc_freq.split(':')
    doc_path_name = splurg[-2]
    frequency = splurg[-1]

    doc_path, doc_name = doc_path_name.split('**')

    #Cast is necessary here in order for the .sort() to work later
    frequency = int(frequency)

    #Same idea as the word count reducer
    if current_word == word:
        doc_freq_table.append([doc_name, doc_path, frequency])
    else:
        if current_word:
            #Sort the list by frequency in descending order
            doc_freq_table.sort(key=itemgetter(2), reverse=True)

            word_dict[current_word] = doc_freq_table
            
            # print('%s\n{\n') % (current_word),
            # #Iterate through the sorted list
            # for value in doc_freq_table:
            #     print('Document:\t%s\nPath:\t%s\nFrequency:\t%s\n') % (value[0], value[1], value[2]),
            # print('}\n')


            #Clear the list
            doc_freq_table = []

        current_word = word
        doc_freq_table.append([doc_name, doc_path, frequency])
if current_word == word:
    #Sort the list by frequency in descending order
    doc_freq_table.sort(key=itemgetter(2), reverse=True)

    word_dict[current_word] = doc_freq_table
    
    # print('%s\n{\n') % (current_word),
    # #Iterate through the sorted list
    # for value in doc_freq_table:
    #     print('Document:\t%s\nPath:\t%s\nFrequency:\t%s\n') % (value[0], value[1], value[2]),
    # print('}\n')

    doc_freq_table = []

#Dump the dictionary as a json
with open('result.json', 'w') as fp:
    json.dump(word_dict, fp)
