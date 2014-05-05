/*
apto is a command line tool automatically generated by ´gobi´. Happy hacking!
*/
package main

import (
	"flag"
	"log"
	"os"

	"github.com/fern4lvarez/apto/apto"
)

func main() {
	_ = flag.Bool("f", false, "force command")
	flag.Parse()

	if l := len(os.Args); l == 1 {
		log.Println("Hello apto")
		return
	} else {
		switch first := os.Args[1]; first {
		case "install":
			err := apto.Install(flag.Args())
			if err != nil {
				log.Println(err)
				return
			}
		case "uninstall":
			err := apto.Uninstall(flag.Args())
			if err != nil {
				log.Println(err)
				return
			}
		case "file":
			apto.File(flag.Args())
		default:
			log.Println("Hello", first)
		}
	}
}
