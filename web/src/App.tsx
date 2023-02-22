// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0

import React from 'react';
import logo from './Catena-X_Logo_ohne_Zusatz_2021.svg';
import './App.css';
import Faucet from './Faucet';

import SuggestKeplrChain from "./KeplrChainSuggestion";

function App() {
  // Register the keplr chain suggestion to the App component
  React.useEffect( () => {
    SuggestKeplrChain()
  },[])
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <Faucet/>
      </header>
    </div>
  );
}

export default App;
