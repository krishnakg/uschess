import React, { Component } from 'react';
import axios from 'axios';

import Configs from './configs.js'

class PlayerInfo extends Component {
  constructor(props) {
    super(props)
    this.state = {
      player: []
    }
  }
  componentDidMount() {
    this.fetchPlayerInfo(this.props.playerId)

  }  
  componentWillReceiveProps(nextProps) {
    if (nextProps.playerId !== this.props.playerId) {
      this.fetchPlayerInfo(nextProps.playerId)
    }
  }

  fetchPlayerInfo(playerId) {
    axios.get(Configs.playerInfoUrl + playerId)
    .then(res => {
      const player = res.data;
      this.setState({ player });
    });
  }
render() {
  return (
    <div className="row" key={this.state.player.id}>
        <h3 className="col-12">{this.state.player.name}</h3>
        <h6 className="col-3">USCF Id</h6><h6 className="col-9">{this.state.player.id}</h6>
        <h6 className="col-3">State</h6><h6 className="col-9">{this.state.player.state}</h6>
    </div>
  );
}
}

export default PlayerInfo;