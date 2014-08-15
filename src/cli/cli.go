package cli

import (
	"log"
	"bufio"
	"os"
	"fmt"
)

type cliApp struct {
	name string
}

func Create() (cliApp){
	var cl cliApp
	cl.name = "Console interface"
	return cl
}

func run() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(" > ");
		text, _ := reader.ReadString(byte('\n'))
		fmt.Printf("Got command [%s]", text)
		if text == "quit\n"{
			break
		}
	}
	log.Printf("Done")
}

func (cl *cliApp) RunningLoop() {
	log.Printf("Running application")
	run()
}
