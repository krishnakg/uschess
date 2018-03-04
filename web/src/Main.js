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
          <Route exact path='/' component={Home}/>
          <Route path='/players/:id' component={Player}/>
          <Route path='/tournaments/:id' component={Tournament}/>
        </Switch>
      </main>
    );
  }
}

export default Main;