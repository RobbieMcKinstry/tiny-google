#!/usr/bin/env python

import sys

# doc_name = sys.stdin.readline().strip()
# print ('doc_name - %s') % doc_name
doc_name = "/mounted/docs/doc.txt"


for line in sys.stdin:
    line = line.strip()

    #input is - document name::word\tfrequency
    doc_word, count = line.split('\t', 1)
    # print('doc_word --> %s') % (doc_word)
    doc_name, word = doc_word.split('::')

    print ('%s\t%s:%s') % (word, doc_name, count)
