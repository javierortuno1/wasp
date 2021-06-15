package chain

import (
	"strconv"

	"github.com/iotaledger/wasp/packages/hashing"

	"github.com/iotaledger/wasp/packages/coretypes/chainid"

	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/publisher"
)

// LogStateTransition also used in testing
func LogStateTransition(msg *ChainTransitionEventData, reqids []coretypes.RequestID, log *logger.Logger) {
	if msg.ChainOutput.GetStateIndex() > 0 {
		log.Infof("STATE TRANSITION TO #%d. Chain output: %s, block size: %d",
			msg.VirtualState.BlockIndex(), coretypes.OID(msg.ChainOutput.ID()), len(reqids))
		log.Debugf("STATE TRANSITION. State hash: %s",
			msg.VirtualState.Hash().String())
	} else {
		log.Infof("ORIGIN STATE SAVED. State output id: %s", coretypes.OID(msg.ChainOutput.ID()))
		log.Debugf("ORIGIN STATE SAVED. state hash: %s",
			msg.VirtualState.Hash().String())
	}
}

// LogGovernanceTransition
func LogGovernanceTransition(msg *ChainTransitionEventData, log *logger.Logger) {
	stateHash, _ := hashing.HashValueFromBytes(msg.ChainOutput.GetStateData())
	log.Infof("GOVERNANCE TRANSITION state index #%d, anchor output: %s, state hash: %s",
		msg.VirtualState.BlockIndex(), coretypes.OID(msg.ChainOutput.ID()), stateHash.String())
}

func LogSyncedEvent(outputID ledgerstate.OutputID, blockIndex uint32, log *logger.Logger) {
	log.Infof("EVENT: state was synced to block index #%d, approving output: %s", blockIndex, coretypes.OID(outputID))
}

func PublishStateTransition(stateOutput *ledgerstate.AliasOutput, reqids []coretypes.RequestID) {
	stateHash, _ := hashing.HashValueFromBytes(stateOutput.GetStateData())
	chainID := chainid.NewChainID(stateOutput.GetAliasAddress())

	publisher.Publish("state",
		chainID.String(),
		strconv.Itoa(int(stateOutput.GetStateIndex())),
		strconv.Itoa(len(reqids)),
		coretypes.OID(stateOutput.ID()),
		stateHash.String(),
	)
	for _, reqid := range reqids {
		publisher.Publish("request_out",
			chainID.String(),
			reqid.String(),
			strconv.Itoa(int(stateOutput.GetStateIndex())),
			strconv.Itoa(len(reqids)),
		)
	}
}

func PublishGovernanceTransition(stateOutput *ledgerstate.AliasOutput) {
	stateHash, _ := hashing.HashValueFromBytes(stateOutput.GetStateData())
	chainID := chainid.NewChainID(stateOutput.GetAliasAddress())

	publisher.Publish("rotate",
		chainID.String(),
		strconv.Itoa(int(stateOutput.GetStateIndex())),
		coretypes.OID(stateOutput.ID()),
		stateHash.String(),
	)
}
