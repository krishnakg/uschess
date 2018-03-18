import React, { Component } from 'react';
import axios from 'axios';

import Configs from './configs.js'
import {uscfPlayerURL} from './Utils.js'

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
        <h6 className="col-4">USCF Id</h6><h6 className="col-8">
          <a href={uscfPlayerURL(this.state.player.id)} target="_blank">{this.state.player.id}</a>
        </h6>
        <h6 className="col-4">State</h6><h6 className="col-8">{this.state.player.state}</h6>
    </div>
  );
}
}

export default PlayerInfo;