// import react components
import { useState, useEffect } from 'react'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'

// import custom components
import Header from './components/Header'
import Footer from './components/Footer'
import Recipes from './components/Recipes'
import AddRecipe from './components/AddRecipe'
import About from './components/About'

function App() {
  const [showAddRecipe, setShowAddRecipe] = useState(false)
  const [recipes, setRecipes] = useState([])
  const beServerURL = 'http://localhost:8080'

  useEffect(() => {
    const getRecipes = async () => {
      const tasksFromServer = await getAllRecipes()
      setRecipes(tasksFromServer)
    }

    getRecipes()
  }, [])

  // Fetch All Recipes
  const getAllRecipes = async () => {
    const resp = await fetch(`${beServerURL}/recipes`, {
      method: 'GET',
    })

    const data = await resp.json()

    return data
  }

  // Add Recipe
  const addRecipe = async (recipe) => {
    const resp = await fetch(`${beServerURL}/recipe`, {
      method: 'POST',
      headers: {
        'Content-type': 'application/json',
      },
      body: JSON.stringify(recipe)
    })

    const data = await resp.json()

    setRecipes([...recipes, data])
  }

  const updateRecipe = async (recipe) => {
    console.log("updateRecipe Function...")

    const resp = await fetch(`${beServerURL}/recipe`, {
      method: 'PUT',
      headers: {
        'Content-type': 'application/json',
      },
      body: JSON.stringify(recipe)
    })

    const data = await resp.json()
    return data
  }

  // Delete Recipe
  const deleteRecipe = async (id) => {
    const resp = await fetch(`${beServerURL}/recipe/${id}`, {
      method: 'DELETE',
    })
    
    // We should control the response status to decide if we will change the state or not.
    resp.status === 200
      ? setRecipes(recipes.filter((recipe) => recipe.id !== id))
      : alert('Error Deleting This Recipe')
  }

  const toggleRecipes = async (id) => {
    console.log("toggleRecipes Function...")

    setRecipes(recipes.map((recipe) => { 
      if (recipe.recipe_id === id) {
        recipe.opened = !recipe.opened

        updateRecipe(recipe)
      } else {
        recipe.opened = false
      }
      
      return recipe;
    }))
  }

  return (
    <Router>
      <div className='container'>
        <Header onAdd={() => setShowAddRecipe(!showAddRecipe)} showAdd={showAddRecipe} title='Recipe Manager'/>

        <Routes>
          <Route path='/' element={
                            <>
                              {showAddRecipe && <AddRecipe onAdd={addRecipe} />}
                              {recipes.length > 0 ? 
                                (<Recipes recipes={recipes} onToggle={toggleRecipes} onDelete={deleteRecipe} />) : 
                                ('No Recipes To Show')}
                            </>} />
          <Route path='/about' element={<About />} />
        </Routes>

        <Footer />
      </div>
    </Router>
  )
}

export default App
