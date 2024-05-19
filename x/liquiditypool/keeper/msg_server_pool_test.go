package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/sunriselayer/sunrise/x/liquiditypool/types"
)

func TestPoolMsgServerCreate(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	sender := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreatePool(wctx, &types.MsgCreatePool{Sender: sender})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestPoolMsgServerUpdate(t *testing.T) {
	sender := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdatePool
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdatePool{Sender: sender},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdatePool{Sender: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdatePool{Sender: sender, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreatePool(wctx, &types.MsgCreatePool{Sender: sender})
			require.NoError(t, err)

			_, err = srv.UpdatePool(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestPoolMsgServerDelete(t *testing.T) {
	sender := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeletePool
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeletePool{Sender: sender},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeletePool{Sender: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeletePool{Sender: sender, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreatePool(wctx, &types.MsgCreatePool{Sender: sender})
			require.NoError(t, err)
			_, err = srv.DeletePool(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
