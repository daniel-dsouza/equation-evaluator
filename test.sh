#!/bin/bash

# count files in test directory
NUM_TESTS=$(expr $(ls tests -1 | wc -l) / 2)
echo running $NUM_TESTS tests

# loop over tests
NUM_CORRECT=$NUM_TESTS
for i in `seq 1 $NUM_TESTS`; do
    go run main.go tests/test${i}.txt > tmp.txt
    diff -q -s tmp.txt tests/output${i}.txt
    NUM_CORRECT=$(expr $NUM_CORRECT - $?)
done

echo PASSED: $NUM_CORRECT
echo FAILED: $(expr $NUM_TESTS - $NUM_CORRECT)

rm tmp.txt