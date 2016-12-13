#!/usr/bin/env python

import sys
import os

#Get the filename
# doc_name = os.environ[mapreduce_map_input_file]
doc_name = "/mounted/docs/doc.txt"

#Print document name, i.e. pass it as input to the reducer
print doc_name

for line in sys.stdin:
	line = line.strip()
	words = line.split()

	for word in words:
		#Write words to stdout
		#i.e. input to the WordCount reducer
		#
		#tuple with each word accompained by a 1
		print '%s\t%s' % (word, 1)
