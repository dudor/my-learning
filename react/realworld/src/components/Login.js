import React from 'react'
import { connect } from 'react-redux'
import agent from '../agent'
import ListErrors from './ListErrors'
import { Link } from 'react-router-dom'

const mapStateToProps = state => {
    return { ...state.auth }
}

const mapDispatchToProps = dispatch => ({
    onChangeEmail: value => dispatch({ type: 'UPDATE_FIELD_AUTH', key: 'email', value }),
    onChangePassword: value => dispatch({ type: 'UPDATE_FIELD_AUTH', key: 'password', value }),
    onSubmit: (email, password) => dispatch({ type: 'LOGIN', payload: agent.Auth.login(email, password) })
})

class Login extends React.Component {
    constructor() {
        super();
        this.changeEmail = ev => this.props.onChangeEmail(ev.target.value);
        this.changePassword = ev => this.props.onChangePassword(ev.target.value);
        this.submitForm = (email, password) => ev => {
            ev.preventDefault();
            this.props.onSubmit(email, password);
        };
    }

    render() {
        const email = this.props.email || '';
        const password = this.props.password || '';
        return (
            <div className='auth-page'>
                <div className='container page'>
                    <div className='row'>
                        <div className='col-md-6 offset-md-3 col-xs-12'>
                            <h1 className='text-xs-center'>SIGN IN</h1>
                            <p className='text-xs-center'>
                                <Link to='/register'>
                                    NEED AN ACCOUNT?
                                    </Link>
                            </p>


                            <form onSubmit={this.submitForm(email, password)}>
                                <fieldset>
                                    <fieldset className='form-group'>
                                        <input className='form-control form-control-lg'
                                            type='email' placeholder='Email'
                                            value={email}
                                            onChange={this.changeEmail}>
                                        </input>
                                    </fieldset>
                                    <fieldset className='form-group'>
                                        <input className='form-control form-control-lg'
                                            type='password' placeholder='Password'
                                            value={password}
                                            onChange={this.changePassword}>
                                        </input>
                                    </fieldset>
                                    <button className='btn btn-lg btn-primary pull-xs-right'
                                        disabled={this.props.inProgress}
                                        type='submit'>
                                        SIGN IN
                                    </button>
                                </fieldset>
                            </form>
                            <ListErrors errors={this.props.errors} />

                        </div>
                    </div>
                </div>
            </div>
        )
    }
}
export default connect(mapStateToProps, mapDispatchToProps)(Login);