import React, { Component } from 'react';
import { Link } from 'react-router-dom'

class SuggestionRenderer extends Component {
  render() {
    return (
      <span>
        <Link to={{ pathname: this.props.pathName }} style={{ textDecoration: 'none' }}>
          <span style={{display:"block"}}>{this.props.suggestion.name}
            <span style={{float:"right"}}>{this.props.suggestion.state}</span>
          </span>
        </Link>
      </span>
    );
  }
}

export default SuggestionRenderer;