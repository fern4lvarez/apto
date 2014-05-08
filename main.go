/*
apto is a command line tool automatically generated by ´gobi´. Happy hacking!
*/
package main

import (
	"flag"
	"log"

	"github.com/fern4lvarez/apto/apto"
)

func main() {
	force := flag.Bool("f", false, "force command (only available for uninstall)")
	flag.Parse()

	if l := len(flag.Args()); l == 0 {
		log.Println("Hello apto")
		return
	} else {
		switch first := flag.Args()[0]; first {
		case "install":
			err := apto.Install(flag.Args())
			if err != nil {
				log.Println(err)
				return
			}
		case "uninstall":
			args, force_ := apto.HandleFlag(flag.Args(), "-f")
			err := apto.Uninstall(args, *force || force_)
			if err != nil {
				log.Println(err)
				return
			}
		case "update":
			if len(flag.Args()) > 1 {
				log.Println("No parameters accepted for update command")
				return
			}
			err := apto.Update()
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
