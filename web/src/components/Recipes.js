// Import react components
import { useState } from 'react'

// Import custom components
import Recipe from './Recipe'

const Recipes = ({ onDelete }) => {
  const [recipes, setRecipes] = useState([
    {
      title: 'Fried Chicken',
      ingredients: 'Chicken',
      instructions: '',
      open: false
    },
    {
      title: 'Spaghetti',
      ingredients: 'Ground Turkey',
      instructions: '',
      open: false
    },
    {
      title: 'Baked Wings',
      ingredients: '1/2 lb of wings',
      instructions: 'Bake chicken for 45 minutes',
      open: false
    }
  ]);

  const toggleRecipes = (index) => {
    setRecipes(recipes.map((recipe, idx) => {
      if (idx === index) {
        recipe.open = !recipe.open
      } else {
        recipe.open = false;
      }

      return recipe;
    }))
  }

  return (
    <>
      <div className="recipes">
        {recipes.map((recipe, idx) => (
          <Recipe recipe={recipe} key={idx} index={idx} toggleRecipe={toggleRecipes} onDelete={onDelete} />
        ))}
      </div>
    </>
  )
}

export default Recipes
