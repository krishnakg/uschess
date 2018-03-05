import React, { Component } from 'react';
import axios from 'axios';
import Configs from './configs.js'
import { Link } from 'react-router-dom'
import {getAbsolutePathForSection} from './Utils.js'

class SectionInfo extends Component {
  constructor(props) {
    super(props)
    this.state = {
      sections : []
    }
  }
  componentDidMount() {
    this.fetchSectionInfo(this.props.tournamentId)

  }  
  componentWillReceiveProps(nextProps) {
    if (nextProps.tournamentId !== this.props.tournamentId) {
      this.fetchSectionInfo(nextProps.tournamentId)
    }
  }

  fetchSectionInfo(tournamentId) {
    axios.get(Configs.tournamentInfoUrl + tournamentId + "/sections")
    .then(res => {
      const sections = res.data;
      this.setState({ sections });
    });
  }

  render() {
    return (
      <div>
        {this.state.sections.map(section =>
          <Link to={{ pathname: getAbsolutePathForSection(section.id) }} key={section.id}><div>{section.name}</div></Link>
        )}
      </div>
    );
  }
}

export default SectionInfo;