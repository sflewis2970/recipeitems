import { FaTimes } from 'react-icons/fa'

const Recipe = ({ recipe, onToggle, onDelete }) => {
  let ingredientsFound = false
  if (recipe.ingredients.length > 0) {
    ingredientsFound = true
  }

  let instructionsFound = false
  if (recipe.instructions.length > 0) {
    instructionsFound = true
  }

  return (
    <div className={"recipe" + (recipe.opened ? ' opened' : '')} onClick={() => onToggle(recipe.recipe_id)}>
			<div className="recipe-title">
        {recipe.name}{' '} 
        <FaTimes style={{ color: 'red', cursor: 'pointer' }} onClick={() => onDelete(recipe.recipe_id)}/> 
			</div>

			<div className="recipe-ingredients">
        <label>Ingredients: </label>

        <p>
          {(ingredientsFound) ? recipe.ingredients : 'ingredients not found'}
        </p>
			</div>

			<div className="recipe-instructions">
        <label>Instructions: </label>

        <p>
          {(instructionsFound) ? recipe.instructions : 'instructions not found'}
        </p>
			</div>
    </div>
  )
}

export default Recipe
