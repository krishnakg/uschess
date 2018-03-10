import React, { Component } from 'react';
import { Link } from 'react-router-dom'

import {getAbsolutePathForTournament, tournamentIdToDateString} from './Utils.js'

class TournamentList extends Component {
  render() {
    var tournaments = this.props.tournaments || [];
    return ( 
      <div>
        {tournaments.map((tournament, index) =>
          // REVIEW: Look for better way to organize keys for various components.
          <TournamentRow key={tournament.id} tournament={tournament}/>
        )}
      </div>
    );
  }
}

class TournamentRow extends Component {

  render() {
    var tournament = this.props.tournament;
    return (
      // Key is a combination of event id and and rating type which should be unique
      <div className="row" key={tournament.id}>
        <div className="col-2 mb-1">{tournamentIdToDateString(tournament.id)}</div>
        <div className="col-7 mb-1">
          <Link to={{ pathname: getAbsolutePathForTournament(tournament.id) }}>{tournament.name}</Link>
        </div>
        <div className="col-3 mb-1">{tournament.city}, {tournament.state}</div>
      </div>
    );
  }
}

export default TournamentList;
