#!/usr/bin/env python

#/usr/local/Cellar/hadoop/2.7.3/libexec/share/hadoop/tools/lib/hadoop-streaming-2.7.3.jar 

import sys
import os
import string

#Get the filename
try:
    doc_path = os.environ['mapreduce_map_input_file']
except KeyError:
    doc_path = os.environ['map_input_file']

#Split based on '/' and get the last item in the split list
doc_name = doc_path.split('/')[-1]  

#For testing only
# doc_name = "/mounted/docs/doc.txt"

#Print document name, i.e. pass it as input to the reducer
# print doc_name

for line in sys.stdin:
    line = line.strip()
    words = line.split()


    for word in words:
        #Strip the punctuation from each word
        word = word.translate(None, string.punctuation)
        word = ''.join(e for e in word if e.isalnum())

        key = ':'.join([doc_name, word])

        #Write words to stdout
        #i.e. input to the WordCount reducer
        #
        #tuple with each word accompained by a 1
        print '%s\t%s' % (key, 1)
