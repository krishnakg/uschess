import React, { Component } from 'react';
import axios from 'axios';
import Configs from './configs.js'
import { NavLink } from 'react-router-dom'

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
      <div className="nav flex-column nav-pills">
        {this.state.sections.map(section =>
          <div key={section.id}>
            <NavLink className="nav-item nav-link" activeClassName='active' to={{ pathname: getAbsolutePathForSection(section.id) }}>{section.name}</NavLink>
          </div>
        )}
      </div>
    );
  }
}

export default SectionInfo;