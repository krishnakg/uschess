import React, { Component } from 'react';

class ProgressArrow extends Component {
  render() {
    return <div className="col-1 mb-1">
      {this.props.change > 0 ? <GreenArrow/> : <RedArrow/>}
      </div>;
  }
}

function RedArrow(props) {
  return <div style={{color:"#FF0000"}}>&#8595;</div>;
}

function GreenArrow(props) {
  return <div style={{color:"#00FF00"}}>&#8593;</div>;
}

export default ProgressArrow;