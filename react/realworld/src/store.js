import { createStore, applyMiddleware, combineReducers } from 'redux'
import { promiseMiddleware, localStorageMiddleware } from './middleware'
import auth from './reducers/auth'
import common from './reducers/common'
import home from './reducers/home'
import settings from './reducers/setttings'
import article from './reducers/article'

const reducer = combineReducers(
  { auth, common, home, settings, article }
)

const store = createStore(reducer, applyMiddleware(promiseMiddleware, localStorageMiddleware));

export default store