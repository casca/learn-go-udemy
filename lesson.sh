#!/bin/sh
counter=$(<template/counter_lesson)
((counter++))
padded=$(printf "%03d" $counter)
mkdir $padded_$1
cp template/main.go $padded_$1/
echo $counter > template/counter_lesson
code $padded_$1/main.go