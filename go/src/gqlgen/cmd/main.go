package main

import (
	"log"
	"os"

	"gqlgen"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("gqlgen ")

	gqlgen.NewAPI()
}
