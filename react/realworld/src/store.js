import { createStore, applyMiddleware, combineReducers } from 'redux'
import { promiseMiddleware, localStorageMiddleware } from './middleware'
import auth from './reducers/auth'
import common from './reducers/common'
import home from './reducers/home'



const reducer = combineReducers(
  {auth, common, home}
)

const store = createStore(reducer, applyMiddleware(promiseMiddleware, localStorageMiddleware));

export default store