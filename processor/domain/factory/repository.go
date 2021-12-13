package factory

import "github.com/igor-gregori/imersao5-gateway/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
