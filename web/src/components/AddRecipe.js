import { useState } from 'react'

const AddRecipe = ({ onAdd }) => {
  const [title, setTitle] = useState('')
  // const [ingredients, setIngredients] = useState([])
  
  const onSubmit = (e) => {
    e.preventDefault()

    if (!title) {
      alert('Please add a recipe')
      return
    }

    onAdd({ title })

    setTitle('')
  }

  return (
    <form className='add-form' onSubmit={onSubmit}>
      <div className='form-control'>
        <label>Recipe</label>
        <input type='text' placeholder='Add Recipe' value={title} onChange={(e) => setTitle(e.target.value)}/>
      </div>

      <div className='form-control'>
        <label>Ingredients</label>
      </div>

      <input type='submit' value='Save Recipe' className='btn btn-block' />
    </form>
  )
}

export default AddRecipe
