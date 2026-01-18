package main

import (
	"flag"
	"log"

	"github.com/Leonid-Yakovlev63/password-generator/internal/app"
)

func main() {

	count := flag.Int("count", 1, "number of passwords to generate")
	flag.IntVar(count, "c", 1, "number of passwords to generate (shorthand)")

	mode := flag.String("mode", "random", "generation mode: 'random' or 'seed'")
	flag.StringVar(mode, "m", "random", "generation mode: 'random' or 'seed' (shorthand)")

	flag.Parse()

	if *count <= 0 {
		log.Fatal("Error: count must be positive")
	}

	if *mode != "random" && *mode != "seed" {
		log.Fatal("Error: mode must be 'random' or 'seed'")
	}

	app.Run(*count, *mode)
}
