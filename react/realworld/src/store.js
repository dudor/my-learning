import { createStore, applyMiddleware, combineReducers } from 'redux'
import { promiseMiddleware, localStorageMiddleware } from './middleware'
import auth from './reducers/auth'
import common from './reducers/common'
import home from './reducers/home'
import settings from './reducers/setttings'
import article from './reducers/article'
import articleList from './reducers/articleList';
import profile from './reducers/profile';

const reducer = combineReducers(
  { article, articleList, auth, common, home, settings, profile }
)

const store = createStore(reducer, applyMiddleware(promiseMiddleware, localStorageMiddleware));

export default store