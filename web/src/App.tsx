import React from 'react';
import logo from './logo.svg';
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
        <p>
          TODO make it catena-x branding
        </p>
        <Faucet/>
      </header>
    </div>
  );
}

export default App;
