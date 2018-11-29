package cmd

import (
	"github.com/spf13/cobra"
	"github.com/kaifei-bianjie/mock/conf"
	"github.com/spf13/viper"
	"github.com/kaifei-bianjie/mock/util/helper/account"
	"github.com/kaifei-bianjie/mock/util/constants"
	"fmt"
)

func FaucetInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "init faucet account",
		Long:  `init faucet account`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			seed := viper.GetString(FlagFaucetSeed)
			addr := viper.GetString(FlagFaucetAddress)

			address, err := account.CreateAccount(constants.MockFaucetName, constants.MockFaucetPassword, seed)

			if err != nil {
				return err
			}
			if address != addr {
				return fmt.Errorf("faucet address generated by seed not equal given")
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&conf.FaucetSeed, FlagFaucetSeed, "", "", "seed")
	cmd.Flags().StringVarP(&conf.FaucetAddress, FlagFaucetAddress, "", "", "bech32 address")

	cmd.MarkFlagRequired(FlagFaucetSeed)
	cmd.MarkFlagRequired(FlagFaucetAddress)

	return cmd
}