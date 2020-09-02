import React from 'react'
import {connect} from 'react-redux'
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
                            <form>
                                <fieldset>
                                    <fieldset className='form-group'>
                                        <input className='form-control form-control-lg' type='email' placeholder='Email'>
                                        </input>
                                    </fieldset>
                                    <fieldset className='form-group'>
                                        <input className='form-control form-control-lg' type='password' placeholder='Password'>
                                        </input>
                                    </fieldset>
                                    <button className='btn btn-lg btn-primary pull-xs-right' type='submit'>
                                        SIGN IN
                                    </button>
                                </fieldset>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}
export default connect(() => ({}), () => ({}))(Login);