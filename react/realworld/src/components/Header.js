import React from 'react'

class Header extends React.Component{

    render(){
        return (
            <nav>
                <div>
                    <a>
                        {this.props.appName}
                    </a>
                </div>
            </nav>

        )
    }
}

export default Header