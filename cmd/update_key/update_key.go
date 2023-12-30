package update_key

import (
	"log"

	"github.com/ViRb3/wgcf/cloudflare"
	. "github.com/ViRb3/wgcf/cmd/shared"
	"github.com/ViRb3/wgcf/config"
	"github.com/ViRb3/wgcf/util"
	"github.com/ViRb3/wgcf/wireguard"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deviceName string

var shortMsg = "Updates the current WireGuard key"

var Cmd = &cobra.Command{
	Use:   "update_key",
	Short: shortMsg,
	Long: FormatMessage(shortMsg, `
Update to new privateKey and publicKey of WireGuard.`),
	Run: func(cmd *cobra.Command, args []string) {
		if err := UpdateWireguardKey(); err != nil {
			log.Fatal(util.GetErrorMessage(err))
		}
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&deviceName, "name", "n", "", "Device name displayed under the 1.1.1.1 app")
}

func UpdateWireguardKey() error {
	if !IsConfigValidAccount() {
		return errors.New("no account detected")
	}

	privateKey, err := wireguard.NewPrivateKey()
	if err != nil {
		return err
	}
	viper.Set(config.PublicKey, privateKey.Public().String())
	viper.Set(config.PrivateKey, privateKey.String())
	log.Println(privateKey.String())
	if err := viper.WriteConfig(); err != nil {
		return err
	}

	ctx := CreateContext()
	_, err = cloudflare.UpdatePrivateKey(ctx)
	if err != nil {
		return err
	}
	log.Println("Successfully updated WireGuard key")
	return nil
}
