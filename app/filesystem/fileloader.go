// Code generated by Project Forge, see https://projectforge.dev for details.
package filesystem

import (
	"os"
)

type FileLoader interface {
	Root() string
	ReadFile(path string) ([]byte, error)
	CreateDirectory(path string) error
	WriteFile(path string, content []byte, mode os.FileMode, overwrite bool) error
	CopyFile(src string, tgt string) error
	CopyRecursive(src string, tgt string, ignore []string) error
	Move(src string, tgt string) error
	ListFiles(path string, ignore []string) []os.FileInfo
	ListFilesRecursive(path string, ignore []string) ([]string, error)
	ListJSON(path string) []string
	ListExtension(path string, ext string) []string
	ListDirectories(path string) []string
	Stat(path string) os.FileInfo
	Exists(path string) bool
	IsDir(path string) bool
	Remove(path string) error
	RemoveRecursive(pt string) error
}
