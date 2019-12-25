@echo off
REM I know, I know. "Batch?! Really!" :(
echo Building...
go build -ldflags "-H windowsgui -s -w" .
echo Compressing...
upx --lzma go-pong.exe > nul
echo Moving...
mkdir out
move go-pong.exe .\out > nul
echo Copying res
xcopy /s res out\res\ > nul