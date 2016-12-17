#!/bin/bash
spark-submit src/word_count_spark.py $1 src/InvertedIndexSpark.json
