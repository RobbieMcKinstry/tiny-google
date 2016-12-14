#!/usr/bin/env bash

/usr/local/bin/hadoop jar /usr/local/Cellar/hadoop/2.7.3/libexec/share/hadoop/tools/lib/hadoop-streaming-2.7.3.jar -mapper "python ./word_count_mapper.py" -reducer "python ./word_count_reducer.py" -input "../books/floop.txt" -output "./output"