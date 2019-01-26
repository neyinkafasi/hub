package cli

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/common"

	"github.com/ironman0x7b2/sentinel-sdk/apps/vpn"
)

func AddGenesisAccountCmd(ctx *server.Context, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-genesis-account [address_or_key_name] [coin][,[coin]]",
		Short: "Add genesis account to genesis.json",
		Args:  cobra.ExactArgs(2),
		RunE: func(_ *cobra.Command, args []string) error {
			config := ctx.Config
			config.SetRoot(viper.GetString(cli.HomeFlag))

			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				kb, err := keys.GetKeyBaseFromDir(viper.GetString(flagClientHome))
				if err != nil {
					return err
				}
				info, err := kb.Get(args[0])
				if err != nil {
					return err
				}
				addr = info.GetAddress()
			}
			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}
			coins.Sort()

			genFile := config.GenesisFile()
			if !common.FileExists(genFile) {
				return fmt.Errorf("%s does not exist, run `vpnd init` first", genFile)
			}
			genDoc, err := loadGenesisDoc(cdc, genFile)
			if err != nil {
				return err
			}

			var appState vpn.GenesisState
			if err = cdc.UnmarshalJSON(genDoc.AppState, &appState); err != nil {
				return err
			}

			appStateJSON, err := addGenesisAccount(cdc, appState, addr, coins)
			if err != nil {
				return err
			}

			return ExportGenesisFile(genFile, genDoc.ChainID, nil, appStateJSON)
		},
	}

	cmd.Flags().String(cli.HomeFlag, vpn.DefaultNodeHome, "node's home directory")
	cmd.Flags().String(flagClientHome, vpn.DefaultCLIHome, "client's home directory")
	return cmd
}

func addGenesisAccount(cdc *codec.Codec, appState vpn.GenesisState, addr sdk.AccAddress, coins sdk.Coins) (json.RawMessage, error) {
	for _, stateAcc := range appState.Accounts {
		if stateAcc.Address.Equals(addr) {
			return nil, fmt.Errorf("the application state already contains account %v", addr)
		}
	}

	acc := auth.NewBaseAccountWithAddress(addr)
	acc.Coins = coins
	appState.Accounts = append(appState.Accounts, vpn.NewGenesisAccount(&acc))
	return cdc.MarshalJSON(appState)
}
