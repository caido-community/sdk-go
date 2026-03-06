package caido

import "fmt"

// Error represents a Caido API error.
type Error struct {
	Op      string // Operation that failed
	Message string // Human-readable message
	Err     error  // Underlying error
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("caido: %s: %s: %v", e.Op, e.Message, e.Err)
	}
	return fmt.Sprintf("caido: %s: %s", e.Op, e.Message)
}

func (e *Error) Unwrap() error { return e.Err }

// GraphQLError represents a GraphQL user error returned by the API.
type GraphQLError struct {
	Typename string
}

func (e *GraphQLError) Error() string {
	return fmt.Sprintf("caido: graphql error: %s", e.Typename)
}

// NotFoundError indicates the requested resource was not found.
type NotFoundError struct {
	Resource string
	ID       string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("caido: %s not found: %s", e.Resource, e.ID)
}

// NotReadyError indicates the Caido instance is not ready.
type NotReadyError struct{}

func (e *NotReadyError) Error() string {
	return "caido: instance not ready"
}

func opErr(op, msg string, err error) error {
	return &Error{Op: op, Message: msg, Err: err}
}
