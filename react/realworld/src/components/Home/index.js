import React from 'react'
import { connect } from 'react-redux'
import Banner from './Banner'
import MainView from './MainView'
import agent from '../../agent'

class Home extends React.Component {
    componentWillMount() {
        this.props.onLoad(agent.Articles.all())
    }
    componentWillUnmount() {
        this.props.onUnload()
    }
    render() {
        return (
            <div className="home-page">
                <Banner appName={this.props.appName} />
                <div className="container page">
                    <div className="row">
                        <MainView />

                        <div className="col-md=3">
                            <div className="sidebar">
                                <p>popular tags</p>
                            </div>
                        </div>

                    </div>

                </div>

            </div>
        )
    }
}

const mapStateToProps = state => {
    return { appName: state.common.appName }
}
const mapDispatch = dispatch => ({
    onLoad: (payload) => {
        dispatch({ type: 'HOME_PAGE_LOADED', payload })
    },
    onUnload: () => {
        dispatch({ type: 'HOME_PAGE_UNLOADED' })
    }
})

export default connect(mapStateToProps, mapDispatch)(Home)