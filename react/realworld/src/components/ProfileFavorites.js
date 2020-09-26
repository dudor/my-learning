import { Profile, mapStateToProps } from './Profile';
import React from 'react';
import { Link } from 'react-router-dom';
import agent from '../agent';
import { connect } from 'react-redux';

const mapDispatchToProps = dispatch => ({
  onFollow: username => dispatch({
    type: 'FOLLOW_USER',
    payload: agent.Profile.follow(username)
  }),
  onLoad: (payload) =>
    dispatch({ type: 'PROFILE_FAVORITES_PAGE_LOADED', payload }),
  onUnfollow: username => dispatch({
    type: 'UNFOLLOW_USER',
    payload: agent.Profile.unfollow(username)
  }),
  onUnload: () =>
    dispatch({ type: 'PROFILE_FAVORITES_PAGE_UNLOADED' }),
  onSetPage: (page, payload) => dispatch({ type: 'SET_PAGE', page, payload })
});

class ProfileFavorites extends Profile {
  componentWillMount() {
    this.props.onLoad(Promise.all([
      agent.Profile.get(this.props.match.params.username),
      agent.Articles.favoritedBy(this.props.match.params.username)
    ]));
  }

  componentWillUnmount() {
    this.props.onUnload();
  }
  onSetPage(page) {
    const promise = agent.Articles.favoritedBy(this.props.profile.username, page)
    this.props.onSetPage(page, promise)
  }

  renderTabs() {
    console.log(this.props)
    return (
      <ul className="nav nav-pills outline-active">
        <li className="nav-item">
          <Link
            className="nav-link"
            to={`/@${this.props.profile.username}`}>
            My Articles
          </Link>
        </li>

        <li className="nav-item">
          <Link
            className="nav-link active"
            to={`/@${this.props.profile.username}/favorites`}>
            Favorited Articles
          </Link>
        </li>
      </ul>
    );
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(ProfileFavorites);
