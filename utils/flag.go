package utils

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	modeFlag = flag.String("mode", "dev", "dev/test/production")
)

// ParseFlag ...
func ParseFlag() {
	flag.Parse()

	switch *modeFlag {
	case "dev":
		log.Info().
			Int("pid", os.Getpid()).
			Msg("pid for dlv debug attach")
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "prod":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	log.Info().
		Str("mode", *modeFlag).
		Msg("APP arguments")
}
