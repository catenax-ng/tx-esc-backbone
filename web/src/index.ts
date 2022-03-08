import catenax_suggestion from './catenax-testnet-1-suggestion.json';
import { Window as KeplrWindow } from '@keplr-wallet/types';

declare global {
  // eslint-disable-next-line @typescript-eslint/no-empty-interface
  interface Window extends KeplrWindow {}
}
// eslint-disable-next-line prefer-arrow/prefer-arrow-functions
async function onWindowLoad() {
  // This code is based on https://github.com/chainapsis/keplr-example/blob/master/src/main.js
  if (!window.getOfflineSigner || !window.keplr) {
    alert('Keplr extension not found. Please install it.');
  } else {
    if (window.keplr.experimentalSuggestChain) {
      try {
        await window.keplr.experimentalSuggestChain(catenax_suggestion);
      } catch (e: any) {
        alert(`Add the chain ${catenax_suggestion.chainName} failed: ${String(e)}`);
      }
    } else {
      alert('Please use the recent version of Keplr extension');
    }
  }
}

// eslint-disable-next-line @typescript-eslint/no-misused-promises
window.addEventListener('load', onWindowLoad);
