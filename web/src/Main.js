import React, { Component } from 'react';
import { Switch, Route, Redirect } from 'react-router-dom'

import Home from './Home'
import Player from './Player'
import Compare from './Compare'
import Tournament from './Tournament'

class Main extends Component {
  render() {
    return (
      <main>
        <Switch>
          {/* REVIEW: All url paths to make sure we are consistent. */}
          <Route exact path='/' component={Home}/>
          <Route exact path='/players/:id' component={Player}/>
          <Route exact path='/players/:player1/vs/:player2' component={Compare}/>
          {/* Is someone visits the tournament page directly, redirect them to the first section. */}
          <Route exact path='/tournaments/:id' render={({match}) => (<Redirect to={{pathname: match.params.id + '/1'}} />)}/>
          <Route exact path='/tournaments/:id/:section' component={Tournament}/>
        </Switch>
      </main>
    );
  }
}

export default Main;