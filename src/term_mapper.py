#!/usr/bin/env python

import sys

for line in sys.stdin:
    line = line.strip()
    #input is - document name:word\tfrequency
    doc_word, count = line.split('\t', 1)
    #Get the document and word from previous .split
    #I got an error if I didn't do the split this way...
    splurg = doc_word.split(':')
    doc_name = splurg[-2]
    doc_name = doc_name.split('/')[-1]
    doc_name = '/' + doc_name
    word = splurg[-1]

    #Output as word\tdocument:frequency
    print ('%s\t%s:%s') % (word, doc_name, count)
