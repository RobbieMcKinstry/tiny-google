#!/bin/bash
hadoop jar src/hadoop-streaming-2.7.3.jar -mapper src/word_count_mapper.py -reducer src/word_count_reducer.py -input $2 -output ./output
hadoop jar src/hadoop-streaming-2.7.3.jar -mapper src/term_mapper.py -reducer src/term_reducer.py -input output/part-00000 -output ./output2
rm -rf ./output
rm -rf ./output2
