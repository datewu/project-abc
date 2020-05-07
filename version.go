package main

import "github.com/rs/zerolog/log"

// GitCommit SemVer the build ldflags slugs
var (
	GitCommit string
	SemVer    string
)

func init() {
	log.Info().
		Str("version", SemVer).
		Str("gitCommit", GitCommit).
		Msg("starting ...")
}
