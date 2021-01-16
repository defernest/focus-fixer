package cmd

import (
	"time"

	"github.com/rs/zerolog"
	zl "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	timeout   time.Duration
	botConfig string
)

var botCmd = &cobra.Command{
	Use:     "bot",
	Short:   "is a best way to start autofocus process periodically",
	Long:    `Start autofocus cameras from your device list`,
	Example: "focusfix bot -t 3 -c config.csv",
	Run: func(cmd *cobra.Command, args []string) {
		zerolog.DurationFieldUnit = time.Minute
		for {
			cameras, err := GetCameras(botConfig)
			if err != nil {
				zl.Fatal().Str("File", botConfig).Err(err).Msg("Failed to get cameras array")
			}
			zl.Info().Str("config", botConfig).Dur("Timeout(minute)", timeout.Truncate(time.Minute)).Msgf("Start autofocus iteration in %d cameras!", len(cameras))
			success := 0
			for _, cam := range cameras {
				err := cam.Autofocus()
				if err != nil {
					zl.Error().Str("Camera", cam.IP.String()).Str("Login", cam.Login).Str("Password", cam.Password).Err(err).Msg("Failed to set focus on camera")
				}
			}
			zl.Info().Str("config", botConfig).Msgf("Successful setup focus in %d/%d cameras!", success, len(cameras))
			time.Sleep(timeout)
		}
	},
}

func init() {
	botCmd.Flags().DurationVarP(&timeout, "timeout", "t", 1, "timeout between iterations of setting autofocus for the camera list")
	botCmd.MarkFlagRequired("timeout")
	botCmd.Flags().StringVarP(&botConfig, "config", "c", "config.csv", "cameras config file")
	botCmd.MarkFlagRequired("config")
}
