package internal

type RepositoryInterface[T any] interface {
	GetByID(id int) (T, error)
	Create(value T) (T, error)
	Update(value T) (T, error)
	Delete(id int) error
}
