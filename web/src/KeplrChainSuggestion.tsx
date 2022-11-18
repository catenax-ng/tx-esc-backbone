// Copyright (c) 2022 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0

import React from 'react';
import {ChainInfo, Window as KeplrWindow} from '@keplr-wallet/types';

declare global {
    // eslint-disable-next-line @typescript-eslint/no-empty-interface
    interface Window extends KeplrWindow {
    }
}


function SuggestKeplrChain():void {
    if (!window.getOfflineSigner || !window.keplr) {
        alert('Keplr extension not found. Please install it.');
    } else {
        const keplr = window.keplr
        createChainSuggestion().then(
            suggestion => {
                if (keplr.experimentalSuggestChain) {
                    keplr.experimentalSuggestChain(suggestion)
                        .then(() => {
                            console.log('Chain %s added.', suggestion.chainName)
                        })
                        .catch(e => {
                            console.log("ERROR", e);
                            alert(`Add the chain ${suggestion.chainName} failed: ${String(e)}`)
                        })
                } else {
                    alert('Please use the recent version of Keplr extension');
                }
            }
        )
    }
}

async function createChainSuggestion(): Promise<ChainInfo> {
    const catenax_suggestion = await fetch("/chain/catenax-testnet-1-suggestion.json")
    return await catenax_suggestion.json() as ChainInfo;
}

export default SuggestKeplrChain;

