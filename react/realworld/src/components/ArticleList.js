import React from 'react'

const ArticleList = props => {
    if (!props.articles) {

        return (
            <div className="article-preview">
                loading
            </div>
        )
    }
    if (props.articles.length === 0) {
        return (
            <div className="article-preview">
                no articles yet
            </div>
        )
    }

    return (
        <div>
            {
                props.articles.map(article => {
                    return (
                        <h2>{article.title}</h2>
                    )
                })
            }
        </div>
    )
}
export default ArticleList