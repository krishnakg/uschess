import React, { Component } from 'react';
import { Link } from 'react-router-dom'
import Autosuggest from 'react-autosuggest';
import axios from 'axios';

import Configs from './configs.js'
import {getAbsolutePathForPlayer} from './Utils.js'
import './PlayerSuggest.css';

// When suggestion is clicked, Autosuggest needs to populate the input
// based on the clicked suggestion. Teach Autosuggest how to calculate the
// input value for every given suggestion.
const getSuggestionValue = suggestion => suggestion.name;

// This is how suggestions would be rendered in the drop down.
const renderSuggestion = suggestion => (
  // Refactor this into a simple component.
  <span>
    <Link to={{ pathname: getAbsolutePathForPlayer(suggestion.id) }} style={{ textDecoration: 'none' }}>
      <span style={{display:"block"}}>{suggestion.name}
        <span style={{float:"right"}}> {suggestion.state}</span>
      </span>
    </Link>
  </span>
);

class PlayerSuggest extends Component {
  constructor() {
    super();

    // Autosuggest is a controlled component.
    // This means that you need to provide an input value
    // and an onChange handler that updates this value (see below).
    // Suggestions also need to be provided to the Autosuggest,
    // and they are initially empty because the Autosuggest is closed.
    this.state = {
      value: '',
      suggestions: []
    };
  }

  onChange = (event, { newValue }) => {
    this.setState({
      value: newValue
    });
  };

  // Autosuggest will call this function every time you need to update suggestions.
  // You already implemented this logic above, so just use it.
  onSuggestionsFetchRequested = ({ value }) => {
    value = value.trim().toLowerCase();
    // Start suggestions only after there are 3 characters.
    if (value.length >= 3) {
    axios.get(Configs.playerSearchUrl + value)
    .then(res => {
      const players = res.data;
      this.setState({
        suggestions: (players != null && players.length > 0) ? players : [],
      });
    });
  }
  };

  // Autosuggest will call this function every time you need to clear suggestions.
  onSuggestionsClearRequested = () => {
    this.setState({
      suggestions: []
    });
  };

  render() {
    const { value, suggestions } = this.state;

    // Autosuggest will pass through all these props to the input.
    const inputProps = {
      placeholder: 'Player Name',
      value,
      onChange: this.onChange
    };

    // Finally, render it!
    return (
      <Autosuggest
        suggestions={suggestions}
        onSuggestionsFetchRequested={this.onSuggestionsFetchRequested}
        onSuggestionsClearRequested={this.onSuggestionsClearRequested}
        getSuggestionValue={getSuggestionValue}
        renderSuggestion={renderSuggestion}
        inputProps={inputProps}
      />
    );
  }
}

export default PlayerSuggest;