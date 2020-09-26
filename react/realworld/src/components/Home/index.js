import React from 'react'
import { connect } from 'react-redux'
import Banner from './Banner'
import MainView from './MainView'
import agent from '../../agent'
import Tags from './Tags'

class Home extends React.Component {
    componentWillMount() {
        const tab = this.props.token ? 'feed' : 'all'
        const articlesPromise = this.props.token ? agent.Articles.feed() : agent.Articles.all()
        this.props.onLoad(tab, Promise.all([agent.Tags.getAll(), articlesPromise]))

    }
    componentWillUnmount() {
        this.props.onUnload()
    }
    render() {
        return (
            <div className="home-page">
                <Banner appName={this.props.appName} token={this.props.token} />
                <div className="container page">
                    <div className="row">
                        <MainView />

                        <div className="col-md-3">
                            <div className="sidebar">
                                <p>popular tags</p>
                                <Tags
                                    tags={this.props.tags}
                                    onClickTag={this.props.onClickTag}
                                />
                            </div>
                        </div>

                    </div>

                </div>

            </div>
        )
    }
}

const mapStateToProps = state => {
    return {
        ...state.home,
        appName: state.common.appName,
        token: state.common.token,
    }
}
const mapDispatch = dispatch => ({
    onLoad: (tab, payload) => {
        dispatch({ type: 'HOME_PAGE_LOADED', tab, payload })
    },
    onUnload: () => {
        dispatch({ type: 'HOME_PAGE_UNLOADED' })
    },
    onClickTag: (tag, payload) => {
        dispatch({ type: 'APPLY_TAG_FILTER', tag, payload })
    }
})

export default connect(mapStateToProps, mapDispatch)(Home)