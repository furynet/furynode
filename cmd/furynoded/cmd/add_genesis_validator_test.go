package cmd

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/Furynet/furynode/app"
	//furynodedcmd "github.com/Furynet/furynode/cmd/furynoded/cmd"
	"github.com/Furynet/furynode/x/oracle"
)

func TestAddGenesisValidatorCmd(t *testing.T) {
	homeDir, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer func(path string) {
		err := os.RemoveAll(path)
		require.NoError(t, err)
	}(homeDir)
	initCmd, _ := NewRootCmd()
	initBuf := new(bytes.Buffer)
	initCmd.SetOut(initBuf)
	initCmd.SetErr(initBuf)

	initCmd.SetArgs([]string{"init", "test", "--home=" + homeDir})
	app.SetConfig(false)
	expectedValidatorBech32 := "furyvaloper1rwqp4q88ue83ag3kgnmxxypq0td59df4782tjn"
	expectedValidator, err := sdk.ValAddressFromBech32(expectedValidatorBech32)
	require.NoError(t, err)
	addValCmd, _ := NewRootCmd()
	addValBuf := new(bytes.Buffer)
	addValCmd.SetOut(addValBuf)
	addValCmd.SetErr(addValBuf)
	addValCmd.SetArgs([]string{"add-genesis-validators", expectedValidatorBech32, "--home=" + homeDir})
	// Run init
	err = svrcmd.Execute(initCmd, homeDir)
	require.NoError(t, err)
	// Run add-genesis-validators

	err = svrcmd.Execute(addValCmd, homeDir)
	require.NoError(t, err)
	// Load genesis state from temp home dir and parse JSON
	serverCtx := server.GetServerContextFromCmd(addValCmd)
	genFile := serverCtx.Config.GenesisFile()
	appState, _, err := genutiltypes.GenesisStateFromGenFile(genFile)
	require.NoError(t, err)
	// Setup app to get oracle keeper and ctx.
	furyapp := app.Setup(false)
	ctx := furyapp.BaseApp.NewContext(false, tmproto.Header{})
	// Run loaded genesis through InitGenesis on oracle module
	mm := module.NewManager(
		oracle.NewAppModule(furyapp.OracleKeeper),
	)
	_ = mm.InitGenesis(ctx, furyapp.AppCodec(), appState)
	// Assert validator
	validators := furyapp.OracleKeeper.GetOracleWhiteList(ctx)
	require.Equal(t, []sdk.ValAddress{expectedValidator}, validators)
}
