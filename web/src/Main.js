import React, { Component } from 'react';
import { Switch, Route, Redirect } from 'react-router-dom'

import Home from './Home'
import Player from './Player'
import Tournament from './Tournament'
import {getAbsolutePathForTournament} from './Utils.js'

class Main extends Component {
  render() {
    return (
      <main>
        <Switch>
          {/* REVIEW: All url paths to make sure we are consistent. */}
          <Route exact path='/' component={Home}/>
          <Route path='/players/:id' component={Player}/>
          {/* TODO: We should make this a utility function soon. */}
          <Route exact path='/tournaments/:id' render={({match}) => (<Redirect to={{pathname: match.params.id + '/1'}} />)}/>
          <Route path='/tournaments/:id/:section' component={Tournament}/>
          {/* TODO: May be also support directly going to the tournament page without the section information. */}
        </Switch>
      </main>
    );
  }
}

export default Main;