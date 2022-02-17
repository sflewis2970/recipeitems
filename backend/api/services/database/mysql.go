package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sflewis2970/recipes/api/services/recipes"
)

func open(sqlDriverName string) (*sql.DB, error) {
	fmt.Println("Opening MySQL database")

	// Open database connection
	db, err := sql.Open(sqlDriverName, "root:devStation@tcp(127.0.0.1:3306)/maindb")

	if err != nil {
		return nil, err
	}

	return db, nil
}

// Get all records from Recipes table
func GetRecipes() ([]recipes.Recipe, error) {
	db, err := open("mysql")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	fmt.Println("Getting all the recipes from the database")
	results, err := db.Query("SELECT * FROM recipes")
	if err != nil {
		return nil, err
	}

	var recipeList []recipes.Recipe

	for results.Next() {
		var recipe recipes.Recipe

		err = results.Scan(&recipe.Recipe_ID, &recipe.Name, &recipe.Ingredients, &recipe.Instructions, &recipe.Opened)
		if err != nil {
			return nil, err
		}

		// Build list if recipe items
		recipeList = append(recipeList, recipe)
	}

	return recipeList, nil
}

// Get a single record from Recipes table
func GetRecipe(recipeID string) (recipes.Recipe, error) {
	db, err := open("mysql")
	if err != nil {
		return recipes.Recipe{}, err
	}
	defer db.Close()

	var recipe recipes.Recipe

	fmt.Println("Getting a single recipe from the database")
	err = db.QueryRow("SELECT * FROM recipes WHERE recipe_id = ?", recipeID).Scan(&recipe.Recipe_ID, &recipe.Name, &recipe.Ingredients, &recipe.Instructions, &recipe.Opened)
	if err != nil {
		return recipes.Recipe{}, err
	}

	return recipe, nil
}

func AddRecipe(recipe recipes.Recipe) *sql.Row {
	db, err := open("mysql")
	if err != nil {
		return nil
	}
	defer db.Close()

	fmt.Println("Adding a new record to the database")
	return db.QueryRow("INSERT INTO recipes VALUES (?, ?, ?, ?, ?)", recipe.Recipe_ID, recipe.Name, recipe.Ingredients, recipe.Instructions, recipe.Opened)
}

// Update a single record in Recipes table
func UpdateRecipe(recipe recipes.Recipe) error {
	db, err := open("mysql")
	if err != nil {
		return err
	}
	defer db.Close()

	fmt.Println("Updating a single recipe in the database")
	db.QueryRow("UPDATE recipes SET name = ?, ingredients = ?, instructions = ?, opened = ? WHERE recipe_id = ?", recipe.Name, recipe.Ingredients, recipe.Instructions, recipe.Opened, recipe.Recipe_ID)

	return nil
}

// Delete a single record from Recipes table
func DeleteRecipe(recipeID string) error {
	db, err := open("mysql")
	if err != nil {
		return err
	}
	defer db.Close()

	fmt.Println("deleting a single recipe from the database")
	db.QueryRow("DELETE FROM recipes WHERE recipe_id = ?", recipeID)

	return nil
}
