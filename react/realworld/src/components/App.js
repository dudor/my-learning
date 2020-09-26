import React from 'react';
import { connect } from 'react-redux'
import './App.css';
import Home from './Home';
import Header from './Header';
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";
import Login from './Login';
import agent from '../agent';
import Register from './Register';
import Settings from './Settings';
import Article from './Article';
import Profile from './Profile';
import ProfileFavorites from './ProfileFavorites';
import Editor from './Editor'

const mapStateToProps = state => {
  return {
    appName: state.common.appName,
    currentUser: state.common.currentUser,
    redirectTo: state.common.redirectTo,
  }
}
const mapDispatchToProps = dispatch => ({
  onLoad: (payload, token) => dispatch({ type: 'APP_LOAD', payload, token }),
  onRedirect: () => dispatch({ type: 'REDIRECT' })
})

class App extends React.Component {
  componentWillMount() {
    const token = window.localStorage.getItem('jwt')
    if (token) {
      agent.SetToken(token)
    }
    this.props.onLoad(token ? agent.Auth.current() : null, token);
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.redirectTo) {
      this.props.onRedirect();
    }
  }

  render() {
    return (
      <div>
        <Header appName={this.props.appName} currentUser={this.props.currentUser} />
        <Switch>
          <Route path='/' exact component={Home} />
          <Route path='/login' component={Login} />
          <Route path='/register' component={Register} />
          <Route path='/settings' component={Settings} />
          <Route path="/article/:id" component={Article} />
          <Route path='/settings' component={Settings} />
          <Route path="/@:username" exact component={Profile} />
          <Route path="/@:username/favorites" component={ProfileFavorites} />
          <Route path="/editor" component={Editor} />
          <Route path="/editor/:slug" component={Editor} />
        </Switch>
      </div>
    )
  }
}


export default connect(mapStateToProps, mapDispatchToProps)(App);
