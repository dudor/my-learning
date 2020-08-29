import _superagent from 'superagent'

const API_ROOT = 'https://conduit.productionready.io/api'
const responseBody = res => res.body

const requests = {
    get: url => {
        return _superagent.get(`${API_ROOT}${url}`).withCredentials().then(responseBody)
    }
}

const Articles = {
    all: page => requests.get('/articles?limit=10')
}

export default {
    Articles
  };