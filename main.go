package main

import (
	"go-cli/app"
	"log"
	"os"
)

func main() {
	goCLI := app.Generate()

	if error := goCLI.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
