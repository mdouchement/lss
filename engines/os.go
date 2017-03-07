package engines

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mdouchement/lss/errors"
)

// An OS structure delegates Engine's methods on Golang os package.
// It implements the Engine interface.
type OS struct {
	Workspace string
}

// IsPathValid implements the Engine interface.
func (o *OS) IsPathValid(path string) bool {
	cPath := filepath.Clean(path)
	return strings.HasPrefix(cPath, o.Workspace)
}

// Exist implements the Engine interface.
func (o *OS) Exist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true // ignoring error
}

// Metadata implements the Engine interface.
func (o *OS) Metadata(path string) M {
	m := M{}
	f, err := os.Stat(path)

	if err != nil {
		m["error"] = err.Error()
		return m
	}

	m["size"] = f.Size()
	m["directory"] = f.IsDir()
	m["updated_at"] = f.ModTime()

	return m
}

// Reader implements the Engine interface.
func (o *OS) Reader(path string) (io.ReadCloser, error) {
	rc, err := os.Open(path)
	if err != nil {
		return rc, errors.NewEnginesError("open", errors.M{
			"reason": err.Error(),
		})
	}
	return rc, err
}

// Writer implements the Engine interface.
func (o *OS) Writer(path string) (io.WriteCloser, error) {
	wc, err := os.Create(path)
	if err != nil {
		return wc, errors.NewEnginesError("create", errors.M{
			"reason": err.Error(),
		})
	}
	return wc, err
}

// ListFiles implements the Engine interface.
func (o *OS) ListFiles(path string, depth int) M {
	m := M{}
	fmt.Println(path)
	err := filepath.Walk(path, func(p string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if depth > 0 && o.pathDepth(path, p) > depth {
			return nil
		}

		relativePath := ""
		if p == o.Workspace {
			relativePath = "/"
		} else {
			relativePath = strings.Replace(p, o.Workspace, "", -1)
		}

		m[relativePath] = M{
			"size":       f.Size(),
			"directory":  f.IsDir(),
			"updated_at": f.ModTime(),
		}
		return nil
	})

	if err != nil {
		m["error"] = err.Error()
	}

	return m
}

// MkdirAllWithFilename implements the Engine interface.
func (o *OS) MkdirAllWithFilename(path string) {
	o.MkdirAll(filepath.Dir(path))
}

// MkdirAll implements the Engine interface.
func (o *OS) MkdirAll(path string) {
	if !o.Exist(path) {
		os.MkdirAll(path, 0755)
	}
}

// Remove implements the Engine interface.
func (o *OS) Remove(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return errors.NewEnginesError("delete", errors.M{
			"reason": err.Error(),
		})
	}
	return nil
}

func (o *OS) pathDepth(base, path string) int {
	rp := strings.Replace(path, base, "", -1)
	return len(strings.Split(rp, "/")) - 1
}
