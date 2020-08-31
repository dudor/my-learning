import React from 'react'

class Login extends React.Component {
    render() {
        return (
            <div className='auth-page'>
                <div className='container page'>
                    <div className='row'>
                        <div className='col-md-6 offset-md-3 col-xs-12'>
                            <h1 className='text-xs-center'>SIGN IN</h1>
                            <p className='text-xs-center'>
                                <a className='text-xs-center'>
                                    NEED AN ACCOUNT?
                                    </a>
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default Login