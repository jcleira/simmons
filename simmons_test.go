package simmons

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Id int
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	t.Run("when Create fails on the DB", func(t *testing.T) {
		expectedError := fmt.Errorf("error!")

		fc := &FakeClient{
			CreateFunc: func(interface{}) error { return expectedError },
		}

		client = fc
		err := Create(&testStruct{})

		assert.Equal(expectedError, err)
		assert.Equal(fc.CreateCalls, 1)
	})
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)

	t.Run("when Update fails on the DB", func(t *testing.T) {
		expectedError := fmt.Errorf("error!")

		fc := &FakeClient{
			UpdateFunc: func(interface{}) error { return expectedError },
		}

		client = fc
		err := Update(&testStruct{})

		assert.Equal(expectedError, err)
		assert.Equal(fc.UpdateCalls, 1)
	})
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)

	t.Run("when Delete fails on the DB", func(t *testing.T) {
		expectedError := fmt.Errorf("error!")

		fc := &FakeClient{
			DeleteFunc: func(interface{}) error { return expectedError },
		}

		client = fc
		err := Delete(&testStruct{})

		assert.Equal(expectedError, err)
		assert.Equal(fc.DeleteCalls, 1)
	})
}

func TestAll(t *testing.T) {
	assert := assert.New(t)

	t.Run("when All fails on the DB", func(t *testing.T) {
		expectedError := fmt.Errorf("error!")

		fc := &FakeClient{
			AllFunc: func(interface{}) error { return expectedError },
		}

		client = fc
		err := All([]*testStruct{})

		assert.Equal(expectedError, err)
		assert.Equal(fc.AllCalls, 1)
	})
}

func TestFind(t *testing.T) {
	assert := assert.New(t)

	t.Run("when Find fails on the DB", func(t *testing.T) {
		expectedError := fmt.Errorf("error!")

		fc := &FakeClient{
			FindFunc: func(interface{}, int) error { return expectedError },
		}

		client = fc
		err := Find(&testStruct{}, 1)

		assert.Equal(expectedError, err)
		assert.Equal(fc.FindCalls, 1)
	})
}

func TestQuery(t *testing.T) {
	assert := assert.New(t)

	t.Run("when Query fails on the DB", func(t *testing.T) {
		expectedError := fmt.Errorf("error!")

		fc := &FakeClient{
			QueryFunc: func(interface{}, string) error { return expectedError },
		}

		client = fc
		err := Query(&testStruct{}, "foo == 'bar'")

		assert.Equal(expectedError, err)
		assert.Equal(fc.QueryCalls, 1)
	})
}
