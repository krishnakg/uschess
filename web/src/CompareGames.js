import React, { Component } from 'react';
import { Link } from 'react-router-dom'

import {getAbsolutePathForSection, tournamentIdToDateString} from './Utils.js'

import './Compare.css';

class CompareGames extends Component {
  render() {
    return (
      <div>
        {this.props.games.map((game, index) =>
          <GameRow key={game.gameId} game={game}/>
        )}
      </div>
    );
  }
}

function Result(props) {
  var resultStr = "";
  if (props.result === "L") { resultStr = "0 - 1"; }
  if (props.result === "W") { resultStr = "1 - 0"; }
  if (props.result === "D") { resultStr = "\u00bd - \u00bd"; }
  return <div>{resultStr}</div>;
}

function PlayerColor(props) {
  var color = "";
  if (props.color === 2) { color = "\u25cf"; }
  if (props.color === 1) { color = "\u25cb"; }  
  return <div>{color}</div>;
}

class GameRow extends Component {

  render() {
    var game = this.props.game;
    return (
      // Key is a combination of event id and and rating type which should be unique
      <div>
      <div className="row top-buffer">
        <div className="col-2 mb-1">{tournamentIdToDateString(game.sectionId)}</div>
        <div className="col-8 mb-1">
          <Link to={{ pathname: getAbsolutePathForSection(game.sectionId) }}>{game.eventName}</Link>
        </div>
      </div>
      <div className="row">
        <div className="col-1 mb-1"></div>
        <div className="col-1 mb-1">{game.round}</div>
        <div className="col-3 mb-1">{game.player1Name}</div>
        <div className="col-1 mb-1"><PlayerColor color={game.player1Color}/></div>
        <div className="col-1 mb-1"><Result result={game.result}/></div>
        <div className="col-1 mb-1"><PlayerColor color={game.player2Color}/></div>
        <div className="col-3 mb-1">{game.player2Name}</div>
      </div>
      </div>
    );
  }
}

export default CompareGames;