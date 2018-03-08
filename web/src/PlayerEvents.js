import React, { Component } from 'react';
import axios from 'axios';
import Configs from './configs.js'
import ProgressArrow from './ProgressArrow.js'
import { Link } from 'react-router-dom'
import {getAbsolutePathForSection} from './Utils.js'

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
        {this.state.events.map((event, index) =>
          // REVIEW: Look for better way to organize keys for various components.
          // TODO: Using index as a key for now as there are cases where all values are same for both rows.
          <EventRow key={event.sectionId + index} event={event}/>
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
          <Link to={{ pathname: getAbsolutePathForSection(event.sectionId) }}>{event.name}</Link>
        </div>
        <div className="col-1 mb-1">{event.postRating}</div>
        <ProgressArrow change={event.postRating-event.preRating}/>
      </div>
    );
  }
}

export default PlayerEvents;