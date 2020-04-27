package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().
		Str("version", SemVer).
		Str("gitCommit", GitCommit).
		Msg("APP starting ...")
	fmt.Println("go project template")
}
