package services

import (
	"fmt"
)

// ErrMissingNamespace is raised when config namespace is missing.
type ErrMissingNamespace struct {
	Namespace string
}

// Error implementation for ErrMissingNamespace.
func (e ErrMissingNamespace) Error() string {
	return fmt.Sprintf("missing namespace %s", e.Namespace)
}

// ErrEmptyNamespace is raised when config namespace is empty.
type ErrEmptyNamespace struct{}

// Error implementation for ErrEmptyNamespace.
func (e ErrEmptyNamespace) Error() string {
	return "empty namespace"
}

// ErrMissingKey is raised when a config key is missing.
type ErrMissingKey struct {
	Key string
}

// Error implementation for ErrMissingKey.
func (e ErrMissingKey) Error() string {
	return fmt.Sprintf("missing key %s", e.Key)
}

// ErrInvalidType is raised when config value is wrong type.
type ErrInvalidType struct {
	Key    string
	Expect string
	Value  interface{}
}

// Error implementation for ErrInvalidType.
func (e ErrInvalidType) Error() string {
	return fmt.Sprintf("key %s invalid. expect %s got %v", e.Key, e.Expect, e.Value)
}
