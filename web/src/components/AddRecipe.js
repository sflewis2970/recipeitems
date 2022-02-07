import { useState } from 'react'

const AddRecipe = ({ onAdd }) => {
  const [title, setTitle] = useState('')
  const [ingredients, setIngredients] = useState('')
  const [instructions, setInstructions] = useState('')
  const [opened, setOpened] = useState(false)
  
  const onSubmit = (e) => {
    e.preventDefault()

    if (!title) {
      alert('Please add a recipe')
      return
    }

    // Send add request up the chain
    onAdd({ title, ingredients, instructions, opened })

    setTitle('')
    setIngredients('')
    setInstructions('')
    setOpened(false)
  }

  return (
    <form className='add-form' onSubmit={onSubmit}>
      <div className='form-control'>
        <label>Recipe</label>

        <input type='text' placeholder='Add Recipe' value={title} onChange={(e) => setTitle(e.target.value)}/>
      </div>

      <div className='form-control'>
        <label>Ingredients</label>

        <textarea placeholder='Add Ingredients' value={ingredients} onChange={(e) => setIngredients(e.target.value)}/>
      </div>

      <div className='form-control'>
        <label>Instructions</label>

        <textarea placeholder='Add Instructions' value={instructions} onChange={(e) => setInstructions(e.target.value)}/>
      </div>

      <input type='submit' value='Save Recipe' className='btn btn-block' />
    </form>
  )
}

export default AddRecipe
