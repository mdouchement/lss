package engines

import "io"

// M is a metadata structure.
type M map[string]interface{}

// Engine is the interface that wraps the basic FileSytem operations.
type Engine interface {
	// IsPathValid validates that the given path is in the workspace.
	IsPathValid(path string) bool

	// Exist checks whether the path exists.
	Exist(path string) bool
	// Metadata returns informations about the given file.
	Metadata(path string) M
	// Reader returns a ReadCloser of the file.
	Reader(path string) (io.ReadCloser, error)
	// Reader returns a WriteCloser of the file.
	Writer(path string) (io.WriteCloser, error)

	// ListFiles lists all files for a given path.
	// It returns pathes relative to the workspace.
	ListFiles(path string) M
	// MkdirAll creates the given directory and its parents.
	MkdirAll(path string)
	// MkdirAllWithFilename creates all parent directories of the given filepath.
	MkdirAllWithFilename(path string)

	// Remove deletes the given path (file or whole directory).
	Remove(path string) error
}
