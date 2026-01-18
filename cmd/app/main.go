package main

import (
	"flag"
	"log"

	"github.com/Leonid-Yakovlev63/password-generator/internal/app"
)

func main() {

	count := flag.Int("count", 1, "number of passwords to generate")
	flag.IntVar(count, "c", 1, "number of passwords to generate (shorthand)")
	flag.Parse()

	if *count <= 0 {
		log.Fatal("Error: count must be positive")
	}

	app.Run(*count)
}
