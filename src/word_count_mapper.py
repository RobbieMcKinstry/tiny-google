import sys
import os

#Get the filename
filename = os.environ[mapreduce_map_input_file]

#Print filename, i.e. pass it as input to the reducer
print filename

for line in sys.stdin:
	line = line.strip()
	words = line.split()

	for word in words:
		#Write words to stdout
		#i.e. input to the WordCount reducer
		#
		#tuple with each word accompained by a 1
		print '%s\t%s' % (word, 1)