import React from 'react';
import { connect } from 'react-redux'
import './App.css';
import Home from './Home';

class App extends React.Component {

  render() {
    return (
      <div>
        <Home appName={this.props.appName} />
      </div>
    )
  }
}

const mapStateToProps = state => ({
  appName: state.appName
});
export default connect(mapStateToProps, () => ({}))(App);
