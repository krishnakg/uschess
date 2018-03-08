import React, { Component } from 'react';
import TournamentInfo from './TournamentInfo'
import SectionInfo from './SectionInfo';
import SectionResult from './SectionResult';

class Tournament extends Component {
  render() {
    var tournamentId = this.props.match.params.id;
    var sectionId = this.props.match.params.id + "." + this.props.match.params.section;
    return (
      <div className="container">
        <TournamentInfo tournamentId={tournamentId}/>
        <div className="row">
          <div className="col-3 border rounded">
            <h4 className="mb-4">Sections</h4>
            <SectionInfo tournamentId={tournamentId}/>
          </div>
          <div className="col-9 border rounded">
            {/* TODO: Make this a margin or padding instead. */}
            <h4>&nbsp;</h4> 
            <SectionResult sectionId={sectionId}/>
          </div>         
        </div>
      </div>
    );
  }
}

export default Tournament;