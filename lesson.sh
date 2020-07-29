#!/bin/sh
counter=$(<template/counter_lesson)
((counter++))
padded=$(printf "%03d" $counter)
mkdir _lessons/${padded}_$1
cp template/main.go _lessons/${padded}_$1/
echo $counter > template/counter_lesson
code _lessons/${padded}_$1/main.go
./watch.sh _lessons/${padded}_$1