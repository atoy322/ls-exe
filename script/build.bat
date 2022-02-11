@echo off
cd ..\cmd\ls
go build
cd ..\..\script
move ..\cmd\ls\ls.exe ..\dist
cd ..\dist
.\ls.exe
