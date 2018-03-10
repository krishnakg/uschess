import React, { Component } from 'react';
import ProgressArrow from './ProgressArrow.js'
import { Link } from 'react-router-dom'
import {getAbsolutePathForSection, tournamentIdToDateString} from './Utils.js'

class PlayerEvents extends Component {
  render() {
    return (
      <div>
        {this.props.events.map((event, index) =>
          // REVIEW: Look for better way to organize keys for various components.
          // TODO: Using index as a key for now as there are cases where all values are same for both rows.
          <EventRow key={event.sectionId + index} event={event}/>
        )}
      </div>
    );
  }
}

class EventRow extends Component {

  render() {
    var event = this.props.event;
    return (
      // Key is a combination of event id and and rating type which should be unique
      <div className="row" key={event.sectionId}>
        <div className="col-2 mb-1">{tournamentIdToDateString(event.id)}</div>
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