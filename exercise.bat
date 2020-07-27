@ECHO OFF

SET /p COUNTER=<template\counter_exercise
SET COUNTER=%COUNTER: =%
SET /a COUNTER = COUNTER + 1

SET PADDED=000%COUNTER%
SET PADDED=%PADDED:~-3%

MKDIR -p _exercises\%PADDED%_%1
COPY template\main.go _exercises\%PADDED%_%1\ >NUL

ECHO %COUNTER% > template\counter_exercise

start /b code _exercises\%PADDED%_%1\main.go

watch.bat _exercises\%PADDED%_%1