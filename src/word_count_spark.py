from __future__ import print_function

import sys
import json
from operator import add 
from operator import itemgetter
import unicodedata

from pyspark.sql import SparkSession
# from pyspark import SparkContext

word_dict = {}
if __name__ == "__main__":
    spark = SparkSession\
        .builder\
        .appName("PythonWordCount")\
        .getOrCreate()

    # sc = SparkContext()

    inverted_index = None
    input_path, inverted_index_path = sys.argv[1], sys.argv[2]
    with open(inverted_index_path) as f:
        inverted_index = json.load(f)

    # with open(input_path) as f:
    #     text_blob = input_path.read_lines()

    #Get the document name from the path
    doc_name = input_path.split('/')[-1]    

    filename = spark.read.text(input_path).rdd.map(lambda r: r[0])
    counts = filename.flatMap(lambda x: x.split(' ')) \
                  .map(lambda x: (x.encode('ascii', 'ignore'), 1)) \
                  .reduceByKey(add)
    output_word_count = counts.collect()
    
    #Map (word, count) => (word, document:count)
    doc_map = counts.map(lambda x: (x[0], ':'.join([doc_name, str(x[1])])))
    doc_map = doc_map.collect()

    #Create a list to that will hold (document, count)
    doc_freq_list = []
    current_word = None
    

    #Iterate through the map output
    for (word, doc_count) in doc_map:
        doc, count = doc_count.split(":")
        count = int(count)

        if current_word == word:
            doc_freq_list.append([doc, count])
        else:
            #Check if current_word is defined
            if current_word:
                #Sort the list of (document, count) tuples for the current_word
                doc_freq_list.sort(key=itemgetter(1), reverse=True)
                #Update the dictionary with Key(word)/Value(list->doc, count)
                word_dict[current_word] = doc_freq_list
                #Clear the list
                doc_freq_list = []
            current_word = word    
            doc_freq_list.append([doc, count])

    if current_word == word:
        #Sort the list of (document, count) tuples for the current_word
        doc_freq_list.sort(key=itemgetter(1), reverse=True)
        #Update the dictionary with Key(word)/Value(list->doc, count)
        word_dict[current_word] = doc_freq_list
        #Clear the list
        doc_freq_list = []

    #Print values in dictionary
    for word, doc_freq in word_dict.iteritems():
        print('\n%s\n{' % (word))
        for value in doc_freq:
            print('Document:\t%s\nFrequency:\t%s\n}' % (value[0], value[1]))
    spark.stop()        
