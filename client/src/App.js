import React, { Component } from 'react';
import Header from './components/Header'
import Directory from './components/Directory'
import DirectoryView from './components/DirectoryView'
import './App.scss';

class App extends Component {
  render() {
    return (
      <div className="doc-wrapper">
        <div className="hdr-wrapper">
          <Header />
        </div>
        <div className="dir-wrapper">
          <Directory />
        </div>
        <div className="view-wrapper">
          <DirectoryView />
        </div>
      </div>
    );
  }
}

export default App;
