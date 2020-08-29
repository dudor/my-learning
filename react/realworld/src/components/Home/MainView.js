import React from 'react'
import { connect } from 'react-redux'
import ArticleList from '../ArticleList'

const MainView = props => {
    return (
        <div className="col-md-9">
            <div className="feed-toggle">
                <ul className="nav nav-pills outline-active">
                    <li className="nav-item">
                        <a href="/" className="nav-link active">
                            Global feeds
                        </a>
                    </li>
                </ul>
            </div>
            <ArticleList articles={props.articles} />
        </div>
    )
}

const mapState = state => {
    return {
        articles: state.articles
    }
}
export default connect(mapState)(MainView)