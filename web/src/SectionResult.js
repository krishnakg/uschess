import React, { Component } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom'
import Configs from './configs.js'

class SectionResult extends Component {
  constructor(props) {
    super(props)
    this.state = {
      results : []
    }
  }
  componentDidMount() {
    this.fetchSectionResults(this.props.sectionId)

  }  
  componentWillReceiveProps(nextProps) {
    if (nextProps.sectionId !== this.props.sectionId) {
      this.fetchSectionResults(nextProps.sectionId)
    }
  }

  fetchSectionResults(sectionId) {
    axios.get(Configs.sectionResultUrl + sectionId)
    .then(res => {
      const results = res.data;
      this.setState({ results });
    });
  }

  render() {
    return (
      <div>
        {this.state.results.map((result, index) =>
        // TODO: Using index as a key for now as there are cases where all values are same for both rows.
          <SectionResultRow result={result} position={index + 1} key={result.playerId + result.score + index}/>          
        )}
      </div>

    );
  }
}

function SectionResultRow(props) {
    return (
      <div className="row">
        <div className="col-1 mb-1">{props.position}</div>
        <div className="col-7 mb-1">
        <Link to={{ pathname: "/players/" + props.result.playerId }}>{props.result.playerName}</Link>
        </div>
        <div className="col-2 mb-1">{props.result.score}</div>
        <div className="col-2 mb-1">{props.result.postRating}</div>
      </div>
    );
}

export default SectionResult;