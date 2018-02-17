package simmons

import "sync"

type FakeClient struct {
	CreateFunc  func(interface{}) error
	CreateCalls int
	UpdateFunc  func(interface{}) error
	UpdateCalls int
	DeleteFunc  func(interface{}) error
	DeleteCalls int
	AllFunc     func(interface{}) error
	AllCalls    int
	FindFunc    func(interface{}, int) error
	FindCalls   int
	QueryFunc   func(interface{}, string) error
	QueryCalls  int

	sync.Mutex
}

func (fc *FakeClient) Create(model interface{}) error {
	fc.Lock()
	defer fc.Unlock()
	fc.CreateCalls++
	return fc.CreateFunc(model)
}

func (fc *FakeClient) Update(model interface{}) error {
	fc.Lock()
	defer fc.Unlock()
	fc.UpdateCalls++
	return fc.UpdateFunc(model)
}

func (fc *FakeClient) Delete(model interface{}) error {
	fc.Lock()
	defer fc.Unlock()
	fc.DeleteCalls++
	return fc.DeleteFunc(model)
}

func (fc *FakeClient) All(model interface{}) error {
	fc.Lock()
	defer fc.Unlock()
	fc.AllCalls++
	return fc.AllFunc(model)
}

func (fc *FakeClient) Find(model interface{}, ID int) error {
	fc.Lock()
	defer fc.Unlock()
	fc.FindCalls++
	return fc.FindFunc(model, ID)
}

func (fc *FakeClient) Query(model interface{}, where string) error {
	fc.Lock()
	defer fc.Unlock()
	fc.QueryCalls++
	return fc.QueryFunc(model, where)
}
