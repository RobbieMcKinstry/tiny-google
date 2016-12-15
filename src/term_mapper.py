#!/usr/bin/env python

import sys

for line in sys.stdin:
    line = line.strip()
    #input is - document name::word\tfrequency
    doc_word, count = line.split('\t', 1)
    #Get the document and word from previous .split
    #I got an error if I didn't do the split this way...
    splurg = doc_word.split(':')
    doc_name = splurg[-2]
    word = splurg[-1]

    doc_path, document = doc_name.split('**')
    #Join the document and frequency
    # value = ":".join([count, doc_name])

    #Output as word\tdocument:frequency
    print ('%s\t%s:%s') % (word, doc_name, count)
