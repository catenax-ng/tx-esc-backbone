import React from 'react';
import logo from './logo.svg';
import './App.css';

import SuggestKeplrChain from "./KeplrChainSuggestion";
function App() {
  // Register the keplr chain suggestion to window.onload
  React.useEffect( () => {
    window.addEventListener('load', SuggestKeplrChain);
  },[])
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          TODO make it catena-x branding
        </p>
      </header>
    </div>
  );
}

export default App;
