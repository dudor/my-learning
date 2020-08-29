

const promiseMiddleware = store => next => action => {
    if (isPromise(action.payload)) {
        action.payload.then(
            res => {
                action.payload = res;
                console.log(action)
                store.dispatch(action)
            },
            error => {
                action.type='faild'
                action.error = true; 
                //console.log(error);
                action.payload = error.response; 
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

export {
    promiseMiddleware
};