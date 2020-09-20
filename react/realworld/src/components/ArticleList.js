import React from 'react'
import ArticlePreview from './ArticlePreview'
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
    return (<div>
        {
            props.articles.map(article => {
                return (
                    <ArticlePreview article={article} key={article.slug} />
                )
            })
        }
    </div>
    )
}
export default ArticleList