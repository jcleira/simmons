package simmons

// Storer interface define all the CRUD operations to create data driven
// applications.
type Storer interface {
	Create(interface{}) error
	Update(interface{}) error
	Delete(interface{}) error
	All(interface{}) error
	Find(interface{}, int) error
	Query(interface{}, string) error
}

// Create persists models for first time.
//
// - model: The model to persists.
//
// Returns an error if any.
func (c *Client) Create(model interface{}) error {
	return c.DB.Create(model).Error
}

// Update modify a previously persisted model.
//
// - model: The model to update.
//
// Returns an error if any.
func (c *Client) Update(model interface{}) error {
	return c.DB.Save(model).Error
}

// Delete removes a previously persisted model.
//
// - model: The model to delete.
//
// Returns an error if any.
func (c *Client) Delete(model interface{}) error {
	return c.DB.Delete(model).Error
}

// All fetch and returns all the persisted records for a given model.
//
// - models: The model to fetch.
//
// Returns an error if any.
func (c *Client) All(models interface{}) error {
	return c.DB.Find(models).Error
}

// Find fetch a model by it's given ID
//
// - model: Struct to fill with the record found.
// - ID: The ID of the model to fetch.
//
// Rerturns an error if any.
func (c *Client) Find(model interface{}, ID int) error {
	return c.DB.First(model, ID).Error
}

// Query fetch multiple models by querying by a SQL statement.
//
// - models: Struct array to fill with the records found.
// - where; Where SQL statement to query.
//
// Rerturns an error if any.
func (c *Client) Query(models interface{}, where string) error {
	return c.DB.Where(where).Find(models).Error
}
