import React, { Component } from 'react';
import {XYPlot, XAxis, YAxis, HorizontalGridLines, LineSeries} from 'react-vis';
import axios from 'axios';

import Configs from './configs.js'
import PlayerInfo from './PlayerInfo'
import PlayerEvents from './PlayerEvents'

import '../node_modules/react-vis/dist/style.css'; 

class Player extends Component {
  constructor(props) {
    super(props)
    this.state = {
      events: []
    }
  }

  componentDidMount() {
    this.fetchPlayerEvents(this.props.match.params.id)
  } 

  componentWillReceiveProps(nextProps) {
    if (nextProps.match.params.id !== this.props.match.params.id) {
      this.fetchPlayerEvents(nextProps.match.params.id)
    }
  }

  fetchPlayerEvents(playerId) {
    axios.get(Configs.playerEventsUrl + playerId)
    .then(res => {
      const events = res.data;
      this.setState({ events });
    });
  }

  render() {
    return (
      <div className="container">

        <div className="row">
          <div className="col-6">
            <PlayerInfo playerId={this.props.match.params.id}/> 
          </div>
          <div className="col-6">
            <RatingPlot events={this.state.events}/> 
          </div>
        </div>

        <div className="row">
          <div className="col-8">
            <h4 className="mb-4">Recent Events</h4>
            <PlayerEvents playerId={this.props.match.params.id} events={this.state.events}/> 
          </div>
          <div className="col-4">              
            <div></div>
          </div>         
        </div>
      
      </div>
    );
  }
}

class RatingPlot extends Component {
  getDataFromEvents(events) {
    // TODO: It might be interesting to use the prerating for the first available tournament 
    // as the starting point so that the graph looks reasonable for players with 1 tournament.
    return events.slice(0).reverse().map((tournament, index) => {
      return {x:index, y:tournament.postRating};
    });
  }

  render() {
    
    var data = this.getDataFromEvents(this.props.events);
    const {yMin, yMax} = data.reduce((acc, row) => ({
        yMax: Math.max(acc.yMax, row.y),
        yMin: Math.min(acc.yMin, row.y)
        }), {yMin: Infinity, yMax: -Infinity})  

    // Y-axis is rating. By specifying a reasonable rating tick values, we are avoiding 
    // the crazy/empty tickValues that react-vis produces.
    var yTickValues = [];
    for (var i  = yMin; i <= yMax; i += 50) {
      yTickValues.push(i);
    }

    // X-axis is number of tournaments.
    var xTickValues = [];    
    for (i  = 0; i <= data.length; i += 5) {
      xTickValues.push(i);
    }

    return (
      <XYPlot width={400} height={300}>
      <HorizontalGridLines />
      <LineSeries
        data={data}
        curve={'curveMonotoneX'}
      />
      <XAxis title="Tournaments" tickValues={xTickValues}/>
      <YAxis title="Rating" tickPadding={1} tickValues={yTickValues}/>
      </XYPlot>
    )
  }
}
export default Player;