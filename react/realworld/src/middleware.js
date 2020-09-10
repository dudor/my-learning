import agent from './agent'

const promiseMiddleware = store => next => action => {
    if (isPromise(action.payload)) {
        store.dispatch({ type: 'ASYNC_START', subtype: action.type })
        action.payload.then(
            res => {
                action.payload = res;
                console.log(action)
                store.dispatch(action)
            },
            error => {
                action.error = true;
                action.payload = error.response.body;
                console.log(action);
                store.dispatch(action)
            }
        )
        return;
    }
    next(action);
}


function isPromise(v) {
    return v && typeof v.then === 'function';
}

const localStorageMiddleware = store => next => action => {
    if (action.type === 'REGISTER' || action.type === 'LOGIN') {
        if (!action.error) {
            window.localStorage.setItem('jwt', action.payload.user.token);
            agent.SetToken(action.payload.user.token);
        }
        else if (action.type === 'LOGOUT') {
            Window.localStorage.setItem('jwt', '');
            agent.SetToken(null);
        }
    }
    next(action);
}

export {
    promiseMiddleware,
    localStorageMiddleware
};