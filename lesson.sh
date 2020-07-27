#!/bin/sh
counter=$(<template/counter_lesson)
((counter++))
padded=$(printf "%03d" $counter)
mkdir ${padded}_$1
cp template/main.go ${padded}_$1/
echo $counter > template/counter_lesson
code ${padded}_$1/main.go
./watch.sh ${padded}_$1