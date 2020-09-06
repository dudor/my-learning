import _superagent from 'superagent'

const API_ROOT = 'https://conduit.productionready.io/api'
const responseBody = res => {
    console.log(res);
    return res.body;
}

let token = null
const tokenPlugin = req => {
    if (token) {
        req.set('authorization', `Token ${token}`)
    }
}

const requests = {
    get: url => {
        return _superagent.get(`${API_ROOT}${url}`)
            .use(tokenPlugin).withCredentials()
            .then(responseBody)
    },
    post: (url, body) => {
        return _superagent.post(`${API_ROOT}${url}`, body)
            .use(tokenPlugin).withCredentials()
            .then(req => { console.log(req); return req.body; })
            .catch(err => { console.log(err); return err; })
    }
}

const Articles = {
    all: page => requests.get('/articles?limit=10')
}
const Auth = {
    current: () => requests.get('/user'),
    login: (email, password) => {
        return requests.post('/users/login', { user: { email, password } })
    }
}

export default {
    Articles,
    Auth,
    SetToken: _token => { token = _token }
};


