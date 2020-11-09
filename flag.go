package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog/log"
)

var (
	modeFlag = flag.String("mode", "dev", "runing mode")
)

func parseFlag() {
	flag.Parse()
	log.Info().
		Str("version", SemVer).
		Str("gitCommit", GitCommit).
		Msg("APP starting ...")

	if *modeFlag == "dev" {
		log.Info().
			Int("pid", os.Getpid()).
			Msg("pid for dlv debug attach")
	}
	log.Info().
		Str("mode", *modeFlag).
		Msg("APP arguments")
}
