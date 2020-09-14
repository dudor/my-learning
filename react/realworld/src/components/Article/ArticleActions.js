import { Link } from 'react-router-dom'
import React from 'react'
import agent from '../../agent'
import { connect } from 'react-redux'

const mapDispatchToProps = dispatch => {
    return {
        onClickDelete: payload => dispatch({ type: 'DELETE_ARTICLE', payload })
    }
}

const ArticleActions = props => {
    const article = props.article
    const del = () => {
        props.onClickDelete(agent.Articles.delete(article.slug))
    }
    if (props.canModify) {
        return <span>
            <Link to={`/editor/${article.slug}`}
                className="btn btn-outline-secondary btn-sm">
                <i className="ion-edit"></i>EDIT ARTICLE
            </Link>
            <button className="btn btn-outline-danger btn-sm" onClick={del}>
                <i className="ion-trash-a"></i> DELETE ARTICLE
            </button>
        </span>
    }
    return null
}

export default connect(null, mapDispatchToProps)(ArticleActions)