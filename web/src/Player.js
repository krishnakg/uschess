import React, { Component } from 'react';
import PlayerInfo from './PlayerInfo'
import PlayerEvents from './PlayerEvents'
import axios from 'axios';
import Configs from './configs.js'
import {XYPlot, XAxis, YAxis, HorizontalGridLines, LineSeries} from 'react-vis';
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
        <div>
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
      </div>
    );
  }
}

class RatingPlot extends Component {
  getDataFromEvents(events) {
    return events.slice(0).reverse().map((tournament, index) => {
      return {x:index, y:tournament.postRating};
    });
  }

  render() {
    return (
    <XYPlot width={400} height={300}>
    <HorizontalGridLines />
    <LineSeries
      data={this.getDataFromEvents(this.props.events)}
      curve={'curveMonotoneX'}
    />
    <XAxis />
    <YAxis tickPadding={1}/>
    </XYPlot>
    )
  }
}
export default Player;