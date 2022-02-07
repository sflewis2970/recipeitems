// Import react components

// Import custom components
import Recipe from './Recipe'

const Recipes = ({ recipes, toggleRecipes, onDelete }) => {

  return (
    <>
      <div className="recipes">
        {recipes.map((recipe, idx) => (
          <Recipe recipe={recipe} key={recipe.id} index={recipe.id} toggleRecipes={toggleRecipes} onDelete={onDelete} />
        ))}
      </div>
    </>
  )
}

export default Recipes
