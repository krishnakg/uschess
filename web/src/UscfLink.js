import React, { Component } from 'react';

import uscfIcon from './uschess.ico'

class UscfLink extends Component {
  render() {
    return (
      <a href={this.props.url} target="_blank"><img src={uscfIcon} alt="US Chess Icon"/></a>
    );
  }
}

export default UscfLink;