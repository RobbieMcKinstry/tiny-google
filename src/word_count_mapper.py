import sys

for line in sys.stdin:
	line = line.strip()
	words = line.split()

	for word in words:
		#Write words to stdout
		#i.e. input to the WordCount reducer
		#
		#tuple with each word accompained by a 1
		print '%s\t%s' % (word, 1)