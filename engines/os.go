package engines

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mdouchement/lss/errors"
)

// An OS structure delegates Engine's methods on Golang os package.
type OS struct {
	Workspace string
}

func (o *OS) IsPathValid(path string) bool {
	cPath := filepath.Clean(path)
	return strings.HasPrefix(cPath, o.Workspace)
}

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

func (o *OS) Reader(path string) (io.ReadCloser, error) {
	rc, err := os.Open(path)
	if err != nil {
		return rc, errors.NewEnginesError("open", errors.M{
			"reason": err.Error(),
		})
	}
	return rc, err
}

func (o *OS) Writer(path string) (io.WriteCloser, error) {
	wc, err := os.Create(path)
	if err != nil {
		return wc, errors.NewEnginesError("create", errors.M{
			"reason": err.Error(),
		})
	}
	return wc, err
}

func (o *OS) ListFiles(path string) M {
	m := M{}
	err := filepath.Walk(path, func(p string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		m[strings.Replace(p, o.Workspace, "", -1)] = M{
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

func (o *OS) MkdirAllWithFilename(path string) {
	o.MkdirAll(filepath.Dir(path))
}

func (o *OS) MkdirAll(path string) {
	if !o.Exist(path) {
		os.MkdirAll(path, 0755)
	}
}

func (o *OS) Remove(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return errors.NewEnginesError("delete", errors.M{
			"reason": err.Error(),
		})
	}
	return nil
}
