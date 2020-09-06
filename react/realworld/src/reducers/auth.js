export default (state = {}, action) => {
    console.log(action)
    switch (action.type) {
        case 'LOGIN':
            return {
                ...state,
                inProgress: false,
                errors: action.error ? action.payload.errors : null
            }
        case 'ASYNC_START':
            return {
                ...state,
            }
        case 'UPDATE_FIELD_AUTH':
            return {
                ...state,
                [action.key]: action.value
            }
    }

    return state;
}