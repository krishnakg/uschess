import React, { Component } from 'react';
import axios from 'axios';
import Configs from './configs.js'

class PlayerEvents extends Component {
  constructor(props) {
    super(props)
    this.state = {
      events: []
    }
  }

  componentDidMount() {
    this.fetchPlayerEvents(this.props.playerId)
  } 

  componentWillReceiveProps(nextProps) {
    if (nextProps.playerId !== this.props.playerId) {
      this.fetchPlayerEvents(nextProps.playerId)
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
      {this.state.events.map(event =>
        // Key is a combination of event id and and rating type which should be unique
        <div className="row" key={event.sectionId}>
            <div className="col-10 mb-1">{event.name}</div>
            <div className="col-1 mb-1">{event.postRating}</div>
            <ProgressArrow change={event.postRating-event.preRating}/>
        </div>
      )}
    </div>
  );
}
}

class ProgressArrow extends Component {
  render() {
    return <div className="col-1 mb-1">
      {this.props.change > 0 ? <GreenArrow/> : <RedArrow/>}
      </div>;
  }
}

function RedArrow(props) {
  return <div style={{color:"#FF0000"}}>&#8595;</div>;
}

function GreenArrow(props) {
  return <div style={{color:"#00FF00"}}>&#8593;</div>;
}

export default PlayerEvents;