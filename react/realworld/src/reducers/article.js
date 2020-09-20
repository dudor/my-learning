export default (state = {}, action) => {
    switch (action.type) {
        case 'ARTICLE_PAGE_LOADED':
            return {
                ...state,
                article: action.payload[0].article,
                comments: action.payload[1].commments
            }
        case 'ARTICLE_PAGE_UNLOADED':
            return {};
        case 'ADD_COMMENT':
            return {
                ...state,
                comments: action.error ? null : (state.comments || []).concat([action.payload.comment]),
                commentErrors: action.error ? action.payload.errors : null
            }
        case 'DELETE_COMMENT':
            const commentID = action.commentID
            return {
                ...state,
                comments: state.comments.filter(c => c.id === commentID)
            }
    }
    return state;

}