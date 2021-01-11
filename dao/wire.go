// +build wireinject


package dao

import "github.com/google/wire"

func initializeDao() (MongoBaseDao, error) {
	wire.Build(NewBaseDao,NewProjectCollection)
	return MongoBaseDao{}, nil
}
