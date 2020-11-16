// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package dao

// Injectors from wire.go:

func newTestDao() (*dao, func(), error) {
	redis, cleanup, err := NewRedis()
	if err != nil {
		return nil, nil, err
	}
	memcache, cleanup2, err := NewMC()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	db, cleanup3, err := NewDB()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	client, cleanup4, err := NewMongo()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	daoDao, cleanup5, err := newDao(redis, memcache, db, client)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return daoDao, func() {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
