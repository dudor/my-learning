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
            .then(responseBody)
    },
    put: (url, body) => {
        return _superagent.put(`${API_ROOT}${url}`, body).use(tokenPlugin).withCredentials().then(responseBody)
    }
}

const Articles = {
    all: page => requests.get('/articles?limit=10')
}
const Auth = {
    current: () => requests.get('/user'),
    login: (email, password) => {
        return requests.post('/users/login', { user: { email, password } })
    },
    register: (username, email, password) => {
        return requests.post('/users', { user: { username, email, password } })
    },
    save: user => {
        return requests.put('/user', user)
    }
}

export default {
    Articles,
    Auth,
    SetToken: _token => { token = _token }
};


