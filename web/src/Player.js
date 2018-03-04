import React, { Component } from 'react';
import PlayerInfo from './PlayerInfo'
import PlayerEvents from './PlayerEvents'

class Player extends Component {
  render() {
    return (
        <div>
        <div className="container">

          <div className="row">
            <div className="col-6 border rounded">
              <PlayerInfo playerId={this.props.match.params.id}/> 
            </div>
          </div>

          <div className="row">
            <div className="col-8 border rounded">
              <h4 className="mb-4">Recent Events</h4>
              <PlayerEvents playerId={this.props.match.params.id}/> 
            </div>
            <div className="col-4 border rounded">
              <h4 className="mb-4">Opponents</h4>
              <div></div>
            </div>         
          </div>
        
        </div>
      </div>
    );
  }
}

export default Player;