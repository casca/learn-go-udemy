#!/bin/sh
nodemon --watch ../ -e go --exec go run $1/main.go
