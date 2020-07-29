@ECHO OFF

SET /p COUNTER=<template\counter_lesson
SET COUNTER=%COUNTER: =%
SET /a COUNTER = COUNTER + 1

SET PADDED=000%COUNTER%
SET PADDED=%PADDED:~-3%

MKDIR  _lessons\%PADDED%_%1
COPY template\main.go  _lessons\%PADDED%_%1\ >NUL

ECHO %COUNTER% > template\counter_lesson

start /b code  _lessons\%PADDED%_%1\main.go

watch.bat  _lessons\%PADDED%_%1