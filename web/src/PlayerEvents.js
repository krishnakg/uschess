import React, { Component } from 'react';
import axios from 'axios';
import Configs from './configs.js'
import { Link } from 'react-router-dom'

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
          <EventRow event={event}/>
        )}
      </div>
    );
  }
}

class EventRow extends Component {
  // The event Id is of the form YYYYMMDD. So we convert it to the form YYYY-MM-DD.
  eventIdToDateString(eventId) {
    return eventId.slice(0, 4) + '-' + eventId.slice(4,6) + '-' + eventId.slice(6,8);
  }

  render() {
    var event = this.props.event;
    return (
      // Key is a combination of event id and and rating type which should be unique
      <div className="row" key={event.sectionId}>
        <div className="col-2 mb-1">{this.eventIdToDateString(event.id)}</div>
        <div className="col-8 mb-1">
          <Link to={{ pathname: '/tournaments/'+event.id }}>{event.name}</Link>
        </div>
        <div className="col-1 mb-1">{event.postRating}</div>
        <ProgressArrow change={event.postRating-event.preRating}/>
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