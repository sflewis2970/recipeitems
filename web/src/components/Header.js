import PropTypes from 'prop-types'
import { useLocation } from 'react-router-dom'
import Button from './Button'

const Header = ({ title, showAdd, onAdd, onRegister, onLogin }) => {
  const location = useLocation()

  return (
    <header className="header">
      <div className='flex-container'>
        <div className='flex-nav-container-items flex-nav-container-items-left'>
          Recipes
        </div>

        <div className='flex-nav-container-items flex-nav-container-items-right'>
          {location.pathname === '/' && (<button>Register</button>)}
          {location.pathname === '/' && (<button></button>)}
          {location.pathname === '/' && (<button>Login</button>)}
        </div>
      </div>

      <div className='flex-container'>
        <div className='flex-title-container-items flex-title-container-items-left'>
          <h1>{title}</h1>
        </div>

        <div className='flex-title-container-items flex-title-container-items-right'>
          {location.pathname === '/' && (
            <Button color={showAdd ? 'red' : 'green'}
                    text={showAdd ? 'Close' : 'Add'}
                    onClick={onAdd}/>
          )}
        </div>
      </div>
    </header>
  )
}

// CSS in JS
// const HeaderStyle = {
//   color: 'red',
//   backgroundColor: 'black',
// }

Header.defaultProps = {
  title: 'Place App Name Here',
}

Header.propTypes = {
  title: PropTypes.string,
}

export default Header
