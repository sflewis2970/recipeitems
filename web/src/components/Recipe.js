import { FaTimes } from 'react-icons/fa'

const Recipe = ({ recipe, onDelete }) => {
  return (
    <div className={`recipe`}>
      <h3>{recipe.title}{' '} <FaTimes style={{ color: 'red', cursor: 'pointer' }} onClick={() => onDelete(recipe.id)}/></h3>
    </div>
  )
}

export default Recipe
