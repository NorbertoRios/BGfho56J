package unitofwork

import "geometris-go/repository"

//UnitOfWork ...
type UnitOfWork struct {
	mysqlRepository  repository.IRepository
	rabbitRepository repository.IRepository
}
