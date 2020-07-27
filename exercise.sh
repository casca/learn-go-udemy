#!/bin/sh
counter=$(<template/counter_exercise)
((counter++))
padded=$(printf "%03d" $counter)
mkdir _exercises/${padded}_$1
cp template/main.go _exercises/${padded}_$1/
echo $counter > template/counter_exercise
code _exercises/${padded}_$1/main.go
watch _exercises/${padded}_$1