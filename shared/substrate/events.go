// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	events "github.com/automata-network/chainbridge-substrate-events"
	"github.com/automata-network/go-substrate-rpc-client/v3/types"
)

type EventExampleRemark struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

type Events struct {
	types.EventRecords
	events.ChainBridgeEvents
	events.BridgeTransferEvents
	events.GameEvents
	Example_Remark []EventExampleRemark //nolint:stylecheck,golint
}
