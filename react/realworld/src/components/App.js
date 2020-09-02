import React from 'react';
import { connect } from 'react-redux'
import './App.css';
import Home from './Home';
import Header from './Header';
import { Switch, Route } from 'react-router-dom'
import Login from './Login';

class App extends React.Component {

  render() {
    return (
      <div>
        <Header appName={this.props.appName} />
        <Switch>
          <Route path='/' component={Home}/>
          <Route path='/login' component={Login}/>
        </Switch>
      </div>
    )
  }
}

const mapStateToProps = state => ({
  appName: state.appName
});
export default connect(mapStateToProps, () => ({}))(App);
