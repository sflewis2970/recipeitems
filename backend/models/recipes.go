package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const dbName = "root:devStation@tcp(127.0.0.1:3306)/maindb"

type DBModel interface {
	open(driverName string) (*sql.DB, error)
	GetRecipes() ([]Recipe, error)
	AddRecipe(recipe Recipe) *sql.Row
	UpdateRecipe(recipe Recipe) error
	DeleteRecipe(recipeID string) error
}

type dbModel struct {
	driverName string
}

func (dbm *dbModel) open(sqlDriverName string) (*sql.DB, error) {
	log.Println("Opening MySQL database")

	// Open database connection
	db, err := sql.Open(sqlDriverName, dbName)

	if err != nil {
		return nil, err
	}

	return db, nil
}

// Get all records from Recipes table
func (dbm *dbModel) GetRecipes() ([]Recipe, error) {
	db, err := dbm.open(dbm.driverName)
	if err != nil {
		return []Recipe{}, err
	}
	defer db.Close()

	log.Println("Getting all the recipes from the database")
	results, err := db.Query("SELECT * FROM recipes")
	if err != nil {
		return []Recipe{}, err
	}

	var recipeList []Recipe

	for results.Next() {
		var recipe Recipe

		err = results.Scan(&recipe.Recipe_ID, &recipe.Name, &recipe.Ingredients, &recipe.Instructions, &recipe.Opened)
		if err != nil {
			return []Recipe{}, err
		}

		// Build list if recipe items
		recipeList = append(recipeList, recipe)
	}

	return recipeList, nil
}

// Get a single record from Recipes table
func (dbm *dbModel) GetRecipe(recipeID string) (Recipe, error) {
	db, err := dbm.open(dbm.driverName)
	if err != nil {
		return Recipe{}, err
	}
	defer db.Close()

	var recipe Recipe

	log.Println("Getting a single recipe from the database")
	err = db.QueryRow("SELECT * FROM recipes WHERE recipe_id = ?", recipeID).Scan(&recipe.Recipe_ID, &recipe.Name, &recipe.Ingredients, &recipe.Instructions, &recipe.Opened)
	if err != nil {
		return Recipe{}, err
	}

	return recipe, nil
}

func (dbm *dbModel) AddRecipe(recipe Recipe) *sql.Row {
	db, err := dbm.open(dbm.driverName)
	if err != nil {
		return nil
	}
	defer db.Close()

	log.Println("Adding a new record to the database")
	return db.QueryRow("INSERT INTO recipes VALUES (?, ?, ?, ?, ?)", recipe.Recipe_ID, recipe.Name, recipe.Ingredients, recipe.Instructions, recipe.Opened)
}

// Update a single record in Recipes table
func (dbm *dbModel) UpdateRecipe(recipe Recipe) error {
	db, err := dbm.open(dbm.driverName)
	if err != nil {
		return err
	}
	defer db.Close()

	log.Println("Updating a single recipe in the database")
	db.QueryRow("UPDATE recipes SET name = ?, ingredients = ?, instructions = ?, opened = ? WHERE recipe_id = ?", recipe.Name, recipe.Ingredients, recipe.Instructions, recipe.Opened, recipe.Recipe_ID)

	return nil
}

// Delete a single record from Recipes table
func (dbm *dbModel) DeleteRecipe(recipeID string) error {
	db, err := dbm.open(dbm.driverName)
	if err != nil {
		return err
	}
	defer db.Close()

	log.Println("deleting a single recipe from the database")
	db.QueryRow("DELETE FROM recipes WHERE recipe_id = ?", recipeID)

	return nil
}

func New(driverName string) *dbModel {
	pDBModel := new(dbModel)
	pDBModel.driverName = driverName

	return pDBModel
}
