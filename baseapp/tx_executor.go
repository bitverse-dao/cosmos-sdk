package baseapp

import (
	"context"
	"cosmossdk.io/store/types"
	abci "github.com/cometbft/cometbft/abci/types"
)

type TxExecutor func(
	ctx context.Context,
	blockSize int,
	cms types.MultiStore,
	deliverTxWithMultiStore func(int, types.MultiStore) *abci.ExecTxResult,
) ([]*abci.ExecTxResult, error)
