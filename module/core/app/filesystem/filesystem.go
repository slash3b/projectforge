package filesystem

import (
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"

	"github.com/pkg/errors"
)

var DefaultMode = os.FileMode(0o755)

type FileSystem struct {
	root     string
	logger   *zap.SugaredLogger
	children []FileLoader
}

var _ FileLoader = (*FileSystem)(nil)

func NewFileSystem(root string, logger *zap.SugaredLogger) *FileSystem {
	return &FileSystem{root: root, logger: logger.With(zap.String("service", "filesystem"))}
}

func (f *FileSystem) GetChildren() []FileLoader {
	return append([]FileLoader{}, f.children...)
}

func (f *FileSystem) AddChildren(fls ...FileLoader) {
	f.children = append(f.children, fls...)
}

func (f *FileSystem) getPath(ss ...string) string {
	s := filepath.Join(ss...)
	if strings.HasPrefix(s, f.root) {
		return s
	}
	return filepath.Join(f.root, s)
}

func (f *FileSystem) Root() string {
	return f.root
}

func (f *FileSystem) Clone() FileLoader {
	return NewFileSystem(f.root, f.logger)
}

func (f *FileSystem) Stat(path string) (os.FileInfo, error) {
	p := f.getPath(path)
	s, err := os.Stat(p)
	if err == nil {
		return s, nil
	}
	for _, c := range f.children {
		if x, e := c.Stat(path); x != nil && e == nil {
			return x, nil
		}
	}
	return nil, err
}

func (f *FileSystem) Exists(path string) bool {
	x, _ := f.Stat(path)
	for _, c := range f.children {
		if c.Exists(path) {
			return true
		}
	}
	return x != nil
}

func (f *FileSystem) IsDir(path string) bool {
	s, err := f.Stat(path)
	if s == nil || err != nil {
		return false
	}
	return s.IsDir()
}

func (f *FileSystem) CreateDirectory(path string) error {
	p := f.getPath(path)
	if err := os.MkdirAll(p, DefaultMode); err != nil {
		return errors.Wrapf(err, "unable to create data directory [%s]", p)
	}
	return nil
}
