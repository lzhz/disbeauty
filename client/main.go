package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/lzhz/disbeauty"
)

func main() {
	cli, err := disbeauty.NewClient("127.0.0.1:9090", 1024)
	if err != nil {
		log.Fatalln(err)
	}

	bio := bufio.NewReader(os.Stdin)

	for {
		line, _, err := bio.ReadLine()
		if err != nil {
			log.Println(err)
		}

		fmt.Println(cli.Display(string(line)))
	}
}
