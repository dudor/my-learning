import React from 'react'
import { connect } from 'react-redux'
import Banner from './Banner'
import MainView from './MainView'

class Home extends React.Component {

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
    return { appName: state.appName }
}

export default connect(mapStateToProps)(Home)