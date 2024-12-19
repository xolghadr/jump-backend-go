package workdir

import (
	"os"
	"path/filepath"
	"strings"
)

// you can use this library freely: "github.com/otiai10/copy"

type WorkDir struct {
	Files map[string]string
}

func InitEmptyWorkDir() *WorkDir {
	return &WorkDir{Files: make(map[string]string)}
}

func (w *WorkDir) CatFile(filePath string) (string, error) {
	content, exists := w.Files[filePath]
	if !exists {
		return "", os.ErrNotExist
	}
	return content, nil
}

func (w *WorkDir) CreateFile(filename string) error {
	if _, exists := w.Files[filename]; exists {
		return os.ErrExist
	}
	w.Files[filename] = ""
	return nil
}

func (w *WorkDir) CreateDir(path string) error {
	return nil
}

func (w *WorkDir) WriteToFile(filename string, content string) error {
	if _, exists := w.Files[filename]; !exists {
		return os.ErrNotExist
	}
	w.Files[filename] = content
	return nil
}

func (w *WorkDir) AppendToFile(filename string, content string) error {
	if _, exists := w.Files[filename]; !exists {
		return os.ErrNotExist
	}
	w.Files[filename] += content
	return nil
}

func (w *WorkDir) Clone() *WorkDir {
	clone := &WorkDir{Files: make(map[string]string)}
	for k, v := range w.Files {
		clone.Files[k] = v
	}
	return clone
}

func (w *WorkDir) ListFilesRoot() []string {
	var files []string
	for file := range w.Files {
		files = append(files, file)
	}
	return files
}

func (w *WorkDir) ListFilesIn(path string) ([]string, error) {
	var files []string
	for file := range w.Files {
		relPath, err := filepath.Rel(path, file)
		if err != nil {
			return nil, err
		}
		if !filepath.IsAbs(relPath) && !strings.HasPrefix(relPath, "..") {
			files = append(files, file)
		}
	}
	return files, nil
}
