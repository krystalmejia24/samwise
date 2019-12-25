package test

import "errors"

// FakeRedis mocks Redis
type FakeRedis struct {
	GetCalledWith, SetCalledWith, DeleteCalledWith string
	ExpectError                                    bool
	DeleteReturns                                  int64
	GetReturns, SetObject                          interface{}
}

// Reset for use in tests
func (r *FakeRedis) Reset() {
	r.GetCalledWith = ""
	r.SetCalledWith = ""
	r.DeleteCalledWith = ""
	r.DeleteReturns = 0
	r.ExpectError = false
	r.GetReturns = nil
	r.SetObject = nil
}

// Get is a fake implementation of redis
func (r *FakeRedis) Get(key string) (interface{}, error) {
	r.GetCalledWith = key

	return r.GetReturns, r.returnError()
}

// Set is a fake implementation of redis
func (r *FakeRedis) Set(key string, value interface{}) error {
	r.SetCalledWith = key
	r.SetObject = value

	return r.returnError()
}

// Delete is a fake implementation of redis
func (r *FakeRedis) Delete(key string) (int64, error) {
	r.DeleteCalledWith = key

	return r.DeleteReturns, r.returnError()
}

func (r *FakeRedis) returnError() error {
	var err error
	if r.ExpectError {
		err = errors.New("forced by test")
	}

	return err
}
