package cmd

import (
	"fmt"

	"github.com/cheggaaa/pb"
	"github.com/rs/zerolog"
	zl "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var fixConfig string

var fixCmd = &cobra.Command{
	Use:     "fix",
	Short:   "is a way to manually start autofocus process",
	Long:    `Start autofocus cameras from your device list`,
	Example: "focusfix fix -c config.csv",
	Run: func(cmd *cobra.Command, args []string) {
		cameras, err := GetCameras(fixConfig)
		if err != nil {
			zl.Fatal().Str("File", fixConfig).Err(err).Msg("Failed to get cameras array")
		}
		var fails []*zerolog.Event
		success := 0
		bar := pb.StartNew(len(cameras))
		for _, cam := range cameras {
			bar.Prefix(fmt.Sprintf("[%s]", cam.IP.String()))
			err := cam.Autofocus()
			if err != nil {
				f := zl.Error().Str("Camera", cam.IP.String()).Str("Login", cam.Login).Str("Password", cam.Password).Err(err)
				fails = append(fails, f)
			} else {
				success++
			}
			bar.Increment()
		}
		bar.Postfix(" DONE")
		bar.Finish()
		zl.Info().Str("config", fixConfig).Msgf("Successful setup focus in %d/%d cameras!", success, len(cameras))
		if len(fails) != 0 {
			zl.Error().Str("config", fixConfig).Msgf("Failed setup focus in cameras:")
			for _, f := range fails {
				f.Msg("Cam: ")
			}
		}
	},
}

func init() {
	fixCmd.Flags().StringVarP(&fixConfig, "config", "c", "config.csv", "cameras config file")
	fixCmd.MarkFlagRequired("config")
}
