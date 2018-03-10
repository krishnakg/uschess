import React, { Component } from 'react';
import TournamentList from './TournamentList'
import axios from 'axios';
import Configs from './configs.js'

class Home extends Component {
  constructor(props) {
    super(props)
    this.state = {
      tournaments: []
    }
  }

  componentDidMount() {
    axios.get(Configs.tournamentInfoUrl)
    .then(res => {
      const tournaments = res.data;
      this.setState({ tournaments });
    });  
  } 

  render() {
    return (
      <div className="container">

      <div className="row">
        <div className="col-12">
          <h4>Recent tournaments</h4>
          <TournamentList tournaments={this.state.tournaments}/> 
        </div>        
      </div>
    
    </div>
    );
  }
}

export default Home;