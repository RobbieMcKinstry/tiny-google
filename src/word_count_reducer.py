#!/usr/bin/env python

import sys


current_word = None
overall_count = 0
word = None

filename = sys.stdin.readline()
print filename

#get input from Stdin
for line in sys.stdin:
	line = line.strip()

	#get the tuple with the word and 1
	word, count = line.split('\t', 1)

	#convert count from string to int
	try:
		count = int(count)
	except ValueError:
		continue

	#Note: Hadoop will sort the output of mapper based on the key (word)
	#	therefore once we don't see a word anymore from Stdin, it will not appear again
    #check if the current_word = word
    if current_word == word:
    	overall_count += count
    else:
    	if current_word:
    		print('%s\t%s') % (current_word, overall_count)
    	current_word = word
    	overall_count = count

#print out last word
if current_word == word:
	print('%s\t%s') % (current_word, overall_count)