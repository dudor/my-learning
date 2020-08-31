import {createStore,applyMiddleware} from 'redux'
import {promiseMiddleware} from './middleware'
const defaultState = {
    appName: "conduit",
    articles: null
  }
  const reducer = (state = defaultState, action) => {
    switch (action.type) {
      case 'HOME_PAGE_LOADED':
        return { ...state, articles: action.payload.articles }
        break;
    }
    return state;
  }
  const store = createStore(reducer, applyMiddleware(promiseMiddleware));

  export default store