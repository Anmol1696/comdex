package keeper_test

import (
	"testing"

	"github.com/comdex-official/comdex/app"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/comdex-official/comdex/x/collector/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	comdexApp := app.Setup(false)
	ctx := comdexApp.BaseApp.NewContext(false, tmproto.Header{})

	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	comdexApp.CollectorKeeper.SetParams(ctx, params)

	response, err := comdexApp.CollectorKeeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
