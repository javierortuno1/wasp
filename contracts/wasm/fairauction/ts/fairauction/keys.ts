// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib"
import * as sc from "./index";

export const IdxParamColor          = 0;
export const IdxParamDescription    = 1;
export const IdxParamDuration       = 2;
export const IdxParamMinimumBid     = 3;
export const IdxParamOwnerMargin    = 4;
export const IdxResultBidders       = 5;
export const IdxResultColor         = 6;
export const IdxResultCreator       = 7;
export const IdxResultDeposit       = 8;
export const IdxResultDescription   = 9;
export const IdxResultDuration      = 10;
export const IdxResultHighestBid    = 11;
export const IdxResultHighestBidder = 12;
export const IdxResultMinimumBid    = 13;
export const IdxResultNumTokens     = 14;
export const IdxResultOwnerMargin   = 15;
export const IdxResultWhenStarted   = 16;
export const IdxStateAuctions       = 17;
export const IdxStateBidderList     = 18;
export const IdxStateBids           = 19;
export const IdxStateOwnerMargin    = 20;

export let keyMap: string[] = [
    sc.ParamColor,
    sc.ParamDescription,
    sc.ParamDuration,
    sc.ParamMinimumBid,
    sc.ParamOwnerMargin,
    sc.ResultBidders,
    sc.ResultColor,
    sc.ResultCreator,
    sc.ResultDeposit,
    sc.ResultDescription,
    sc.ResultDuration,
    sc.ResultHighestBid,
    sc.ResultHighestBidder,
    sc.ResultMinimumBid,
    sc.ResultNumTokens,
    sc.ResultOwnerMargin,
    sc.ResultWhenStarted,
    sc.StateAuctions,
    sc.StateBidderList,
    sc.StateBids,
    sc.StateOwnerMargin,
];

export let idxMap: wasmlib.Key32[] = new Array(keyMap.length);
