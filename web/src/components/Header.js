import PropTypes from 'prop-types'
import { useLocation } from 'react-router-dom'
import Button from './Button'

const Header = ({ title, onAdd, showAdd }) => {
  const location = useLocation()

  return (
    <header className='header'>
      <h1>{title}</h1>

      {location.pathname === '/' && (
        <Button color={showAdd ? 'red' : 'green'}
                text={showAdd ? 'Close' : 'Add'}
                onClick={onAdd}/>
      )}
    </header>
  )
}

// CSS in JS
// const HeaderStyle = {
//   color: 'red',
//   backgroundColor: 'black',
// }

Header.defaultProps = {
  title: 'App Name',
}

Header.propTypes = {
  title: PropTypes.string,
}

export default Header
