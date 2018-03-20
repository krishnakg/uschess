import React, { Component } from 'react';
import axios from 'axios';

import Configs from './configs.js'
import UscfLink from './UscfLink.js'
import {uscfPlayerURL, chessDbPlayerURL} from './Utils.js'

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
          <div className="row col-12"><h3 className="col-12">{this.state.player.name}&nbsp;
            <UscfLink url={uscfPlayerURL(this.state.player.id)} />
          </h3></div>
          <UscfInfo id={this.state.player.id} />
          {this.state.player.fideId &&
            <FideInfo id={this.state.player.fideId} />
          }
          {this.state.player.state &&
            <StateInfo state={this.state.player.state} />
          }
      </div>
    );
  }
}

function UscfInfo(props) {
  return (
    <div className="row col-12">
      <h6 className="col-4">USCF Id</h6>
      <h6 className="col-8">{props.id}</h6>
    </div>
  )
}

function FideInfo(props) {
  return (
    <div className="row col-12">
      <h6 className="col-4">FIDE Id</h6>
      <h6 className="col-3">{props.id}</h6>
      <h6 className="col-5">
        <a href={chessDbPlayerURL(props.id)}>Games</a>
      </h6>
    </div>
  )
}

function StateInfo(props) {
  return (
    <div className="row col-12">
      <h6 className="col-4">State</h6>
      <h6 className="col-8">{props.state}</h6>
    </div>
  )
}


export default PlayerInfo;