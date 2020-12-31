import React from 'react'
import ReactDOM from 'react-dom'
import { Button } from 'antd'
import Popup from './popup'
import './index.scss'

const App: React.FC = () => {
    return (
        <div className="App">
            <Popup/>
        </div>
    )
}

ReactDOM.render(<App />, document.getElementById('root'))