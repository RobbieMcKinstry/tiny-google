#!/usr/bin/env python

import sys


for line in sys.stdin:
    line = line.strip()

    #input is - document name::word\tfrequency
    doc_word, count = line.split('\t', 1)
    #Get the document and word from previous .split
    doc_name, word = doc_word.split(':')
    #Join the document and frequency
    value = ":".join([doc_name, count])

    #Output as word\tdocument:frequency
    print ('%s\t%s') % (word, value)
