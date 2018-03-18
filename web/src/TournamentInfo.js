import React, { Component } from 'react';
import axios from 'axios';

import Configs from './configs.js'
import UscfLink from './UscfLink.js'
import {uscfTournamentURL} from './Utils.js'

class TournamentInfo extends Component {
  constructor(props) {
    super(props)
    this.state = {
      tournament : []
    }
  }
  componentDidMount() {
    this.fetchTournamentInfo(this.props.tournamentId)

  }  
  componentWillReceiveProps(nextProps) {
    if (nextProps.tournamentId !== this.props.tournamentId) {
      this.fetchTournamentInfo(nextProps.tournamentId)
    }
  }

  fetchTournamentInfo(tournamentId) {
    axios.get(Configs.tournamentInfoUrl + tournamentId)
    .then(res => {
      const tournament = res.data;
      this.setState({ tournament });
    });
  }

  render() {
    return (
        <div className="row">
          <div className="col-12">
            <h4>{this.state.tournament.name}&nbsp;
              <UscfLink url={uscfTournamentURL(this.state.tournament.id)}/>
            </h4>
            <h5>{this.state.tournament.city}, {this.state.tournament.state}</h5>
            <h6>{this.state.tournament.sections} sections, {this.state.tournament.players} players</h6>
          </div>
        </div>
    );
  }
}

export default TournamentInfo;