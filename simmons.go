package simmons

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Client is simmons main struct it contains the ORM configuration and it
// implemnts the Storer interface for real ORM / Database operations.
type Client struct {
	DB *gorm.DB
}

// client is simmons package var to simplify the simmons API, with it simmons
// can be used as:
//
// simmons.Init(DB)
// simmons.Create(user)
//
// that would be better than a more common approach:
//
// client := simmons.Init()
// client.Create(user)
var client Storer

// Init initializes simmons using the given Database configuration to prepare
// the ORM client with it.
//
// - DB: A properly configured gorm ORM DB.
//
// Returns an ORM configuration option func.
func Init(conn string) error {
	DB, err := gorm.Open(conn)
	if err != nil {
		return fmt.Errorf("error intializing simmons Database, Err: %v", err)
	}

	client = &Client{
		DB: DB,
	}

	return nil
}

// Init test initializes simmons with a custom storer, usually intended to be
// initialized with a simmons.FakeClient with customized funcs for mocking.
//
// - storer: A simmons.Storer interface implementer.
//
// Returns nothing.
func InitTest(storer Storer) {
	client = storer
}

// Create persists models for first time.
//
// - model: The model to persists.
//
// Returns an error if any.
func Create(model interface{}) error {
	return client.Create(model)
}

// Update modify a previously persisted model.
//
// - model: The model to update.
//
// Returns an error if any.
func Update(model interface{}) error {
	return client.Update(model)
}

// Delete removes a previously persisted model.
//
// - model: The model to delete.
//
// Returns an error if any.
func Delete(model interface{}) error {
	return client.Delete(model)
}

// All fetch and returns all the persisted records for a given model.
//
// - models: The model to fetch.
//
// Returns an error if any.
func All(models interface{}) error {
	return client.All(models)
}

// Find fetch a model by it's given ID
//
// - model: Struct to fill with the record found.
// - ID: The ID of the model to fetch.
//
// Rerturns an error if any.
func Find(model interface{}, ID int) error {
	return client.Find(model, ID)
}

// Query fetch multiple models by querying by a SQL statement.
//
// - models: Struct array to fill with the records found.
// - where; Where SQL statement to query.
//
// Rerturns an error if any.
func Query(models interface{}, where string) error {
	return client.Query(models, where)
}
