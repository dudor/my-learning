import _superagent from 'superagent'

const API_ROOT = 'https://conduit.productionready.io/api'
const responseBody = res => {
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
    },
    delete: url => {
        return _superagent.delete(`${API_ROOT}${url}`).use(tokenPlugin).withCredentials().then(responseBody)
    }
}

const limit = (count, p) => `limit=${count}&offset=${p ? p * count : 0}`
const encode = encodeURIComponent
const Articles = {
    all: page => requests.get(`/articles?${limit(10, page)}`),
    delete: slug => requests.delete(`/articles/${slug}`),
    get: slug => requests.get(`/articles/${slug}`),
    byAuthor: (author, page) => requests.get(`/articles?author=${encode(author)}&${limit(10, page)}`),
    favoritedBy: (author, page) => requests.get(`/articles?favorited=${encode(author)}&${limit(10, page)}`),
    feed: page => requests.get(`/articles/feed?${limit(10, page)}`),
}
const Comments = {
    create: (slug, comment) => requests.post(`/articles/${slug}/comments`, { comment }),
    delete: (slug, commentID) => requests.delete(`/articles/${slug}/comments/${commentID}`),
    forArticle: slug => requests.get(`/articles/${slug}/comments`)
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
const Profile = {
    follow: username => requests.post(`/profiles/${username}/follow`),
    unfollow: username => requests.delete(`/frofiles/${username}/follow`),
    get: username => requests.get(`/profiles/${username}`),
}
const Tags = {
    getAll: () => requests.get(`/tags`)
}

export default {
    Articles,
    Auth,
    Comments,
    Profile,
    Tags,
    SetToken: _token => { token = _token }
};


