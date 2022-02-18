// Import react components

// Import custom components
import Recipe from './Recipe'

const Recipes = ({ recipes, onToggle, onDelete }) => {
  return (
    <>
      <div className="recipes">
        {recipes.map((recipe, idx) => (
          <Recipe key={idx} recipe={recipe} onToggle={onToggle} onDelete={onDelete} />
        ))}
      </div>
    </>
  )
}

export default Recipes
