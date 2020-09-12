import React from 'react'
import agent from '../agent'
import { Link } from 'react-router-dom'
import ListErrors from './ListErrors'
import { connect } from 'react-redux'

const mapStateToProps = state => {
    return {
        ...state.auth
    }
}
const mapDispatchToProps = dispatch => {
    return {
        onChangeEmail: value => dispatch({ type: 'UPDATE_FIELD_AUTH', key: 'email', value }),
        onChangePassword: value => dispatch({ type: 'UPDATE_FIELD_AUTH', key: 'password', value }),
        onChangeUsername: value => dispatch({ type: 'UPDATE_FIELD_AUTH', key: 'username', value }),
        onSubmit: (username, email, password) => dispatch({ type: 'REGISTER', payload: agent.Auth.register(username, email, password) }),
        onUnload: () => dispatch({ type: 'REGISTER_PAGE_UNLOADED' })
    }
}

class Register extends React.Component {
    constructor() {
        super();
        this.ChangeEmail = ev => this.props.onChangeEmail(ev.target.value)
        this.ChangePassword = ev => this.props.onChangePassword(ev.target.value)
        this.ChangeUsername = ev => this.props.onChangeUsername(ev.target.value)
        this.submitForm = (username, email, password) => ev => {
            ev.preventDefault()
            this.props.onSubmit(username, email, password)
        }
    }
    componentWillUnmount() {
        this.props.onUnload()
    }

    render() {
        const email = this.props.email;
        const password = this.props.password;
        const username = this.props.username;

        return <div className='auth-page'>
            <div className='container page'>
                <div className='row'>
                    <div className='col-md-6 offset-md-3 col-xs-12'>
                        <h1 className='text-xs-center'>SIGN UP</h1>
                        <p className='text-xs-center'>
                            <Link to='/login'>HAVE A ACCOUNT?</Link>
                        </p>


                        <form onSubmit={this.submitForm(username, email, password)}>
                            <fieldset>
                                <fieldset className='form-group'>
                                    <input
                                        className='form-control form-control-lg'
                                        type='text'
                                        placeholder='Username'
                                        value={this.props.username}
                                        onChange={this.ChangeUsername}></input>
                                </fieldset>
                                <fieldset className='form-group'>
                                    <input
                                        className='form-control form-control-lg'
                                        type='email'
                                        placeholder='Email'
                                        value={this.props.email}
                                        onChange={this.ChangeEmail}></input>
                                </fieldset>
                                <fieldset className='form-group'>
                                    <input
                                        className='form-control form-control-lg'
                                        type='password'
                                        placeholder='Password'
                                        value={this.props.password}
                                        onChange={this.ChangePassword}></input>
                                </fieldset>
                                <button className='btn btn-lg btn-primary pull-xs-right'
                                    type='submit'
                                    disabled={this.props.inProgress}>
                                    SIGN UP
                                </button>
                            </fieldset>
                        </form>
                        <ListErrors errors={this.props.errors} />

                    </div>
                </div>

            </div>
        </div>

    }
}

export default connect(mapStateToProps, mapDispatchToProps)(Register)