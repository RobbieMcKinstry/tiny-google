from __future__ import print_function

import sys
import json
from operator import add 
from operator import itemgetter
import unicodedata

from pyspark.sql import SparkSession
# from pyspark import SparkContext

word_dict = {}
inner_dict = {}
if __name__ == "__main__":
    spark = SparkSession\
        .builder\
        .appName("PythonWordCount")\
        .getOrCreate()

    # sc = SparkContext()

    inverted_index = None
    input_path, inverted_index_path = sys.argv[1], sys.argv[2]
    with open(inverted_index_path) as f:
        inverted_index = json.load(f)['Links']

    # with open(input_path) as f:
    #     text_blob = input_path.read_lines()

    #Get the document name from the path
    doc_name = input_path.split('/')[-1]    
    doc_path = "/" + doc_name

    filename = spark.read.text(input_path).rdd.map(lambda r: r[0])
    counts = filename.flatMap(lambda x: x.split(' ')) \
                  .map(lambda x: (x.encode('ascii', 'ignore'), 1)) \
                  .reduceByKey(add).collect()

    if len(counts) == 0:
        raise Exception("Uh oh counts is 0 length")
    dictionary_list = {}
    for word_tuple in counts:
        word, count = word_tuple[0], word_tuple[1]
        dictionary_list[word] = {
            'Frequency': count,
            'DocumentPath': doc_path,
            'DocumentName': doc_name,
        }

    for word, word_data in dictionary_list.iteritems():
        if word in inverted_index:
           for doc in inverted_index.get(word):
                if doc['DocumentName'] == word_data['DocumentName']:
                    continue
                else:
                    inverted_index[word].append(word_data)
        else:
            inverted_index[word] = [ word_data ] 

    with open(inverted_index_path, 'r+') as fp:
        json.dump( { 'Links': inverted_index }, fp)
    spark.stop() 


"""
    doc_map = doc_map.collect()

    #Create a list to that will hold (document, count)
    doc_freq_list = []
    current_word = None
    word = None

    #Iterate through the map output
    for (word, doc_count) in doc_map:
        doc, count = doc_count.split(":")
        count = int(count)

        if current_word == word:
            inner_dict["Frequency"] = count
            inner_dict["DocumentPath"] = doc_path
            inner_dict["DocumentName"] = doc_name
            doc_freq_list.append(inner_dict)
        else:
            #Check if current_word is defined
            if current_word:
                #Sort the list of (document, count) tuples for the current_word
                #Update the dictionary with Key(word)/Value(list->doc, count)
                word_dict[current_word] = [ doc_freq_list ]
                #Clear the list
                doc_freq_list = []

            inner_dict = {}
            current_word = word    
            doc_freq_list.append([doc, count])
            current_word = word    
            inner_dict["Frequency"] = count
            inner_dict["DocumentPath"] = doc_path
            inner_dict["DocumentName"] = doc_name
            doc_freq_list.append(inner_dict)

    if current_word == word:
        #Sort the list of (document, count) tuples for the current_word
        #Update the dictionary with Key(word)/Value(list->doc, count)
        word_dict[current_word] = [ doc_freq_list ]
        #Clear the list

    #Print values in dictionary
    # for word, doc_freq in word_dict.iteritems():
    #     print('\n%s\n{' % (word))
    #     for value in doc_freq:
    #         print('Document:\t%s\nFrequency:\t%s\n}' % (value['Document'], value['Path']))
    
    #Iterate through the dictionary for this book
    #If the json that was loaded in also contains a word in word_dict
    # append the word_dict value to the json
    for word, value in word_dict.iteritems():
        if inverted_index.get(word):
            
            # then, we need to see if the array in the index
            # already contains the document.
            

        #else create a new key/value pair in the dictionary
        else:
            inverted_index[word] = [ value ]
"""
