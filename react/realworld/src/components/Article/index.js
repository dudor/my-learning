import React from 'react'
import agent from '../../agent'
import marked from 'marked'
import ArticleMeta from './ArticleMeta'
const mapStateToProps = state => {
    return {
        ...state.Article,
        currentUser: state.common.currentUser
    }
}
const mapDispatchToProps = dispatch => {
    return {
        onLoad: payload => dispatch({ type: 'ARTICLE_PAGE_LOADED', payload }),
        onUnload: () => dispatch({ type: 'ARTICLE_PAGE_UNLOADED' })
    }
}

class Article extends React.Component {
    componentWillMount() {
        this.props.onLoad(Promise.all([
            agent.Articles.get(this.props.params.id),
            agent.Comments.forArticle(this.props.params.id)
        ]))
    }
    componentWillUnmount() {
        this.props.onUnload()
    }
    render() {
        if (!this.props.article) {
            return null;
        }
        const markup = { __html: marked(this.props.article.body) }
        const canModify = this.props.currentUser &&
            this.props.currentUser.username === this.props.author.username;

        return (
            <div className='article-page'>
                <div className='banner'>
                    <div className='container'>
                        <h1>{this.props.article.title}</h1>
                        <ArticleMeta
                            article={this.props.article}
                            canModify={canModify} />
                    </div>
                </div>
            </div>
        )

    }
}