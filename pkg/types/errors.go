package types

import "errors"

var (
	// ErrStorageMongoDB ...
	ErrStorageMongoDB = errors.New("error storage mongodb")
	// ErrLenPage ...
	ErrLenPage = errors.New("page length must not be zero")
	// ErrSortField ...
	ErrSortField = errors.New("sort field must be filled")
	// ErrSortOrder ...
	ErrSortOrder = errors.New("sort order must be set")
)
