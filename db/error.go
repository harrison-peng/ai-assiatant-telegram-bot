package db

const (
	ErrorNotFound = "mongo: no documents in result"
)

// IsNotFoundError checks if the error is a not found error.
func IsNotFoundError(err error) bool {
	return err.Error() == ErrorNotFound
}
