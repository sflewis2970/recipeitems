// import react components
import { useState, useEffect } from 'react'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'

// import custom components
import Header from './components/Header'
import Footer from './components/Footer'
// import Tasks from './components/Tasks'
import Recipes from './components/Recipes'
// import AddTask from './components/AddTask'
import AddRecipe from './components/AddRecipe'
import About from './components/About'

function App() {
  // const [showAddTask, setShowAddTask] = useState(false)
  // const [tasks, setTasks] = useState([])
  const [showAddRecipe, setShowAddRecipe] = useState(false)
  const [recipes, setRecipes] = useState([])

  useEffect(() => {
    // const getTasks = async () => {
    //   const tasksFromServer = await fetchTasks()
    //   setTasks(tasksFromServer)
    // }

    // getTasks()

    const getRecipes = async () => {
      const tasksFromServer = await fetchRecipes()
      setRecipes(tasksFromServer)
    }

    getRecipes()
  }, [])

  // Fetch Tasks
  const fetchTasks = async () => {
    const resp = await fetch('http://localhost:5000/tasks')
    const data = await resp.json()

    return data
  }

  // Fetch Recipes
  const fetchRecipes = async () => {
    const resp = await fetch('http://localhost:5500/recipes')
    const data = await resp.json()

    return data
  }

  // Fetch Task
  // const fetchTask = async (id) => {
  //   const res = await fetch(`http://localhost:5000/tasks/${id}`)
  //   const data = await res.json()

  //   return data
  // }

  // // Add Task
  // const addTask = async (task) => {
  //   const res = await fetch('http://localhost:5000/tasks', {
  //     method: 'POST',
  //     headers: {
  //       'Content-type': 'application/json',
  //     },
  //     body: JSON.stringify(task),
  //   })

  //   const data = await res.json()

  //   setTasks([...tasks, data])

  // const id = Math.floor(Math.random() * 10000) + 1
  // const newTask = { id, ...task }
  // setTasks([...tasks, newTask])
  // }

  // Add Recipe
  const addRecipe = async (recipe) => {
    const resp = await fetch('http://localhost:6000/recipes', {
      method: 'POST',
      headers: {
        'Content-type': 'application/json',
      },
      body: JSON.stringify(recipe),
    })

    const data = await resp.json()

    setRecipes([...recipes, data])

    // const id = Math.floor(Math.random() * 10000) + 1
    // const newRecipe = { id, ...recipe }
    // setTasks([...recipes, newRecipe])
  }

  // Delete Task
  // const deleteTask = async (id) => {
  //   const res = await fetch(`http://localhost:5000/tasks/${id}`, {
  //     method: 'DELETE',
  //   })
    
  //   //We should control the response status to decide if we will change the state or not.
  //   res.status === 200
  //     ? setTasks(tasks.filter((task) => task.id !== id))
  //     : alert('Error Deleting This Task')
  // }

  // Delete Recipe
  const deleteRecipe = async (id) => {
    const resp = await fetch(`http://localhost:6000/recipes/${id}`, {
      method: 'DELETE',
    })
    
    //We should control the response status to decide if we will change the state or not.
    resp.status === 200
      ? setRecipes(recipes.filter((recipe) => recipe.id !== id))
      : alert('Error Deleting This Recipe')
  }

  // Toggle Reminder
  // const toggleReminder = async (id) => {
  //   const taskToToggle = await fetchTask(id)
  //   const updatedTask = { ...taskToToggle, reminder: !taskToToggle.reminder }

  //   const resp = await fetch(`http://localhost:5000/tasks/${id}`, {
  //     method: 'PUT',
  //     headers: {
  //       'Content-type': 'application/json',
  //     },
  //     body: JSON.stringify(updatedTask),
  //   })

  //   const data = await resp.json()

  //   setTasks(
  //     tasks.map((task) => task.id === id ? { ...task, reminder: data.reminder } : task)
  //   )
  // }

  return (
    <Router>
      <div className='container'>
        {/* <Header onAdd={() => setShowAddTask(!showAddTask)} showAdd={showAddTask}/> */}
        <Header onAdd={() => setShowAddRecipe(!showAddRecipe)} showAdd={showAddRecipe} title='Recipe Manager'/>

        <Routes>
          {/* <Route path='/' element={
                            <>
                              {showAddTask && <AddTask onAdd={addTask} />}
                              {tasks.length > 0 ? 
                                (<Tasks tasks={tasks} onDelete={deleteTask} onToggle={toggleReminder} />) : 
                                ('No Tasks To Show')}
                            </>} /> */}

          <Route path='/' element={
                            <>
                              {showAddRecipe && <AddRecipe onAdd={addRecipe} />}
                              {recipes.length > 0 ? 
                                (<Recipes recipes={recipes} onDelete={deleteRecipe} />) : 
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
