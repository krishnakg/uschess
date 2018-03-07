import React, { Component } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom'
import Configs from './configs.js'

class SectionResult extends Component {
  constructor(props) {
    super(props)
    this.state = {
      // Info for each player in the section.
      results : [],
      // Info on all games in the section.
      games : []
    }
  }
  componentDidMount() {
    this.fetchSectionResults(this.props.sectionId)
    this.fetchSectionGames(this.props.sectionId)
  }  
  componentWillReceiveProps(nextProps) {
    if (nextProps.sectionId !== this.props.sectionId) {
      this.fetchSectionResults(nextProps.sectionId)
      this.fetchSectionGames(nextProps.sectionId)
    }
  }

  fetchSectionResults(sectionId) {
    axios.get(Configs.sectionResultUrl + sectionId)
    .then(res => {
      const results = res.data;
      this.setState({ results:results });
    });
  }

  fetchSectionGames(sectionId) {
    axios.get(Configs.sectionGamesUrl + sectionId)
    .then(res => {
      const games = res.data;
      this.setState({ games:games });
    });
  }

  getGames(playerId) {
    if (this.state.games.length===0) {
      return [];
    }
    var games =  this.state.games;
    if (games[playerId] == null || games[playerId].length === 0) {
      return [];
    }
    // Convert data on each player from map to an array.
    var array = Object.keys(games[playerId]).map( key => games[playerId][key]);
    if (array == null || array.length === 0) {
      return [];
    }

    // Convert data on all rounds, which is a map into an array for react to render it.
    return Object.keys(array).map(key => array[key]);
  }
  render() {
    return (
      <div>
        {this.state.results.map((result, index) =>
        // TODO: Using index as a key for now as there are cases where all values are same for both rows.
          <SectionResultRow result={result} games={this.getGames(result.playerId)} position={index + 1} key={result.playerId + result.score + index}/>          
        )}
      </div>

    );
  }
}

function SectionResultRow(props) {
    return (
      <div className="row">
        <div className="col-1 mb-1">{props.position}</div>
        <div className="col-7 mb-1">
        <Link to={{ pathname: "/players/" + props.result.playerId }}>{props.result.playerName}</Link>
        </div>
        <div className="col-2 mb-1">
          <a href={"#collapse"+ props.result.playerId} role="button" data-toggle="collapse">{props.result.score}</a>
        </div>
        <div className="col-2 mb-1">{props.result.postRating}</div>
        <SectionPairings playerId={props.result.playerId} games={props.games}/>          
      </div>
    );
}

class SectionPairings extends Component {
  render() {
    return (
      <div className="col-12 collapse" id={"collapse"+this.props.playerId}>
      <div className="card card-body">
      {this.props.games.map((game, index) =>
        // TODO: Using index as a key is a bad idea. Need to fix this.
        <div className="row" key={"pairings" + this.props.playerId + game.playerId + index}>
          <div className="col-2 mb-1"></div>
          <div className="col-1 mb-1">{index+1}</div>
          <div className="col-7 mb-1">{game.playerName}</div>
          <div className="col-2 mb-1">{game.result}</div>
        </div>
      )}
      </div>
      </div>
    );
  }
}
export default SectionResult;