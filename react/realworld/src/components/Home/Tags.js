import React from 'react'
import agent from '../../agent'

const Tags = props => {
    const tags = props.tags
    if (tags) {
        return <div>
            {
                tags.map(tag => {
                    const handleClick = ev => {
                        ev.preventDefault()
                        props.onClickTag(tag, agent.Articles.byTag(tag))
                    }

                    return <a href='' className='tag-default tag-pill' key={tag} onClick={handleClick}>
                        {tag}
                    </a>
                })


            }

        </div>
    }
    else {
        return <div>
            LOADING TAGS...
        </div>
    }
}

export default Tags