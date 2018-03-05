import React, { Component } from 'react';
import { Switch, Route } from 'react-router-dom'

import Home from './Home'
import Player from './Player'
import Tournament from './Tournament'

class Main extends Component {
  render() {
    return (
      <main>
        <Switch>
          {/* REVIEW: All url paths to make sure we are consistent. */}
          <Route exact path='/' component={Home}/>
          <Route path='/players/:id' component={Player}/>
          <Route path='/tournaments/:id/:section' component={Tournament}/>
          {/* TODO: May be also support directly going to the tournament page without the section information. */}
        </Switch>
      </main>
    );
  }
}

export default Main;