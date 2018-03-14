import React, { Component } from 'react';

import SectionInfo from './SectionInfo';
import SectionResult from './SectionResult';
import TournamentInfo from './TournamentInfo'

import './Tournament.css';

class Tournament extends Component {
  render() {
    var tournamentId = this.props.match.params.id;
    var sectionId = this.props.match.params.id + "." + this.props.match.params.section;
    return (
      <div className="container">
        <TournamentInfo tournamentId={tournamentId}/>
        <div className="row">
          <div className="col-xs-6 col-md-3 top-buffer">
            <h4 className="mb-4">Sections</h4>
            <SectionInfo tournamentId={tournamentId}/>
          </div>
          <div className="col-xs-12 col-md-9 top-buffer">
            <h4 className="mb-4">Results</h4>
            <SectionResult sectionId={sectionId}/>
          </div>         
        </div>
      </div>
    );
  }
}

export default Tournament;