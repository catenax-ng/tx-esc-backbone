import React from 'react';
import {Window as KeplrWindow} from '@keplr-wallet/types';
// import catenax_suggestion from './catenax-testnet-1-suggestion.json';

declare global {
    // eslint-disable-next-line @typescript-eslint/no-empty-interface
    interface Window extends KeplrWindow {
    }
}


async function SuggestKeplrChain() {
    if (!window.getOfflineSigner || !window.keplr) {
        alert('Keplr extension not found. Please install it.');
    } else {
        if (window.keplr.experimentalSuggestChain) {
            let suggestion = await createChainSuggestion()
            try {
                await window.keplr.experimentalSuggestChain(suggestion)
            } catch (e){
                console.log("ERROR", e);
                console.log('Chain %s added.',suggestion.chainName);
                alert(`Add the chain ${suggestion.chainName} failed: ${String(e)}`)
            }
        } else {
            alert('Please use the recent version of Keplr extension');
        }
    }
}

async function createChainSuggestion() {
    const catenax_suggestion=await fetch("/chain/catenax-testnet-1-suggestion.json")
    return catenax_suggestion.json();
}

export default SuggestKeplrChain;

