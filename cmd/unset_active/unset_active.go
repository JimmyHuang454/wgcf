package unset_active

import (
	"log"

	"github.com/ViRb3/wgcf/cloudflare"
	. "github.com/ViRb3/wgcf/cmd/shared"
	"github.com/ViRb3/wgcf/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var deviceName string

var shortMsg = "Updates the current WireGuard key"

var Cmd = &cobra.Command{
	Use:   "unset_active",
	Short: shortMsg,
	Long: FormatMessage(shortMsg, `
Update to new privateKey and publicKey of WireGuard.`),
	Run: func(cmd *cobra.Command, args []string) {
		if err := disableDevice(); err != nil {
			log.Fatal(util.GetErrorMessage(err))
		}
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&deviceName, "name", "n", "", "Device name displayed under the 1.1.1.1 app")
}

func disableDevice() error {
	if !IsConfigValidAccount() {
		return errors.New("no account detected")
	}

	ctx := CreateContext()
	_, err := cloudflare.GetSourceBoundDevice(ctx)
	if err != nil {
		return err
	}

	_, err = cloudflare.UpdateSourceBoundDeviceActive(ctx, false)
	return nil
}
