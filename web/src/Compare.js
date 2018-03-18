import React, { Component } from 'react';
import axios from 'axios';

import Configs from './configs.js'
import PlayerInfo from './PlayerInfo'
import CompareGames from './CompareGames'

class Compare extends Component {
  constructor(props) {
    super(props)
    this.state = {
      games: []
    }
  }

  componentDidMount() {
    this.fetchPlayerCompareEvents(this.props.match.params.player1, this.props.match.params.player2)
  } 

  componentWillReceiveProps(nextProps) {
    if ( (nextProps.match.params.player1 !== this.props.match.params.player1) ||
  (nextProps.match.params.player2 !== this.props.match.params.player2) ) {
      this.fetchPlayerCompareEvents(nextProps.match.params.player1, nextProps.match.params.player2)
    }
  }

  fetchPlayerCompareEvents(player1Id, player2Id) {
    axios.get(Configs.playerCompareUrl + player1Id + '/' + player2Id)
    .then(res => {
      const games = res.data || [];
      this.setState({ games });
    });
  }

  render() {
    return (
      <div className="container">

        <div className="row">
          <div className="col-xs-12 col-md-6">
            <PlayerInfo playerId={this.props.match.params.player1}/> 
          </div>
          <div className=" col-xs-12 col-md-6">
            <PlayerInfo playerId={this.props.match.params.player2}/> 
          </div>
        </div>

        <div className="row top-buffer">
          <div className="col-12">
            <h4 className="mb-4">{this.state.games.length} Recent Games</h4>            
            <CompareGames games={this.state.games} />
          </div>
        </div>
      
      </div>
    );
  }
}
export default Compare;