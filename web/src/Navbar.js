import React, { Component } from 'react';

import PlayerSuggest from './PlayerSuggest';
import {getAbsolutePathForPlayer} from './Utils.js'
import SuggestionRenderer from './SuggestionRenderer'


class Navbar extends Component {
  // This is how suggestions would be rendered in the drop down. This has to be a pure function.
  renderSuggestion = suggestion => (
    <SuggestionRenderer pathName={getAbsolutePathForPlayer(suggestion.id)} suggestion={suggestion} />
  );

render() {
  return (
    <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
      {/* Add the icon here. */}
      <a className="navbar-brand" href="/">Chess Matrix</a> 
      {/* We have a menu bar followed by the search bar.      */}
      <div className="collapse navbar-collapse" id="navbarCollapse">
        <ul className="navbar-nav mr-auto"></ul>
      </div>
      <PlayerSuggest renderSuggestion={this.renderSuggestion}/>
    </nav>
  );
}
}

export default Navbar;