package executer

import (
	"migrator/configs"
	"migrator/models/schemas"
	"migrator/repository/tables"

	"github.com/jmoiron/sqlx"
)

type Executer struct {
	Repos []tables.TableRepositoryInterface[any]
}

func (e *Executer) Run() error {
	for _, repo := range e.Repos {
		if err := repo.ReadData(); err != nil {
			return err
		}
		if err := repo.WriteData(); err != nil {
			return err
		}
	}
	return nil
}

func NewExecuter(db *sqlx.DB, config *configs.Config) (*Executer, error) {

	repos := []tables.TableRepositoryInterface[any]{}

	// Transaction
	transaction, err := tables.NewTableRepository[schemas.TransactionSchema](
		config.TransactionTableName,
		config.TransactionTableSchemaName,
		db,
		config,
	)
	if err != nil {
		return nil, err
	}
	repos = append(repos, transaction)

	// TransactionProduct
	transactionProduct, err := tables.NewTableRepository[schemas.TransactionProductSchema](
		config.TransactionProductTableName,
		config.TransactionProductTableSchemaName,
		db,
		config,
	)
	if err != nil {
		return nil, err
	}
	repos = append(repos, transactionProduct)

	// Customer
	customer, err := tables.NewTableRepository[schemas.CustomerSchema](
		config.CustomerTableName,
		config.CustomerTableSchemaName,
		db,
		config,
	)
	if err != nil {
		return nil, err
	}
	repos = append(repos, customer)

	// Product
	product, err := tables.NewTableRepository[schemas.ProductSchema](
		config.ProductTableName,
		config.ProductTableSchemaName,
		db,
		config,
	)
	if err != nil {
		return nil, err
	}
	repos = append(repos, product)

	executer := &Executer{
		Repos: repos,
	}

	return executer, nil
}
