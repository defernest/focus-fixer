package main

import (
	"fmt"
	cmd "focus-fixer/cmd"
	"time"

	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("Start focusfixer service!")
	color := colorable.NewColorableStdout()
	consoleWriter := zerolog.NewConsoleWriter(
		func(w *zerolog.ConsoleWriter) {
			w.TimeFormat = time.RFC822
		},
	)
	consoleWriter.Out = color
	log.Logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
	cmd.Execute()
}
