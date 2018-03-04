import React, { Component } from 'react';
import PlayerSuggest from './PlayerSuggest';

class Navbar extends Component {
render() {
  return (
    <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
      {/* Add the icon here. */}
      <a className="navbar-brand" href="/">Ratings</a> 
      {/* We have a menu bar followed by the search bar.      */}
      <div className="collapse navbar-collapse" id="navbarCollapse">
        <ul className="navbar-nav mr-auto"></ul>
      </div>
      <PlayerSuggest />
    </nav>
  );
}
}

export default Navbar;