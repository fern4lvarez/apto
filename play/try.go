package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("sudo", "apt-get", "update")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	rd := bufio.NewReader(stdout)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	go io.Copy(os.Stdout, stdout)

	for {
		str, err := rd.ReadString('\n')
		if err != nil {
			log.Fatal("Read Error:", err)
			return
		}
		fmt.Println(str)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
