#!/usr/bin/env python

import sys


for line in sys.stdin:
    line = line.strip()

    #input is - document name::word\tfrequency
    doc_word, count = line.split('\t', 1)
    # print('doc_word --> %s') % (doc_word)
    doc_name, word = doc_word.split(':')

    print ('%s\t%s:%s') % (word, doc_name, count)
