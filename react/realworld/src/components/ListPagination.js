import React from 'react'
import agent from '../agent'

const ListPagination = props => {
    if (props.articlesCount <= 10)
        return null

    const range = []
    for (let i = 0; i < Math.ceil(props.articlesCount / 10); i++) {
        range.push(i)
    }

    const setPage = page => props.onSetPage(page)
    return <nav>
        <ul className='pagination'>
            {
                range.map(v => {
                    const isCurrent = v === props.currentPage
                    const onClick = ev => {
                        ev.preventDefault()
                        setPage(v)
                    }
                    return <li className={isCurrent ? 'page-item active' : 'page-item'} onClick={onclick} key={v.toString()}>
                        <a className='page-link' href=''>{v + 1}</a>
                    </li>
                })
            }

        </ul>
    </nav>
}

export default ListPagination