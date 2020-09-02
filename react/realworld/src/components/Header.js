import React from 'react'
import { Link } from 'react-router-dom'

const LogoutView = props => {
    return (<div>
        <ul className='nav navbar-nav pull-xs-right'>
            <li className='nav-item'>
                <Link to='/' className='nav-link'>Home</Link>
            </li>
            <li className='nav-item'>
                <Link to='/login' className='nav-link'>Login</Link>
            </li>
        </ul>
    </div>)
}

class Header extends React.Component {
    render() {
        return (
            <nav className="navbar navbar-light">
                <div className="container">
                    <Link to='/' className='navbar-brand'>
                        {this.props.appName}
                    </Link>
                    <LogoutView></LogoutView>
                </div>
            </nav>

        )
    }
}

export default Header