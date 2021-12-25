package mysql

type Storage interface {
	Get(stmt string) ([]interface{}, error)
	Create(stmt string) (bool, error)
}
