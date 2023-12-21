// Copyright (c) 2022-2023 Contributors to the Eclipse Foundation
//
// See the NOTICE file(s) distributed with this work for additional
// information regarding copyright ownership.
//
// This program and the accompanying materials are made available under the
// terms of the Apache License, Version 2.0 which is available at
// https://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
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

