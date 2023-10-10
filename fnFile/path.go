package fnFile

import (
	"os"
	"path/filepath"
	"strings"
)

type Path string

func (x Path) Path() (path string, err error) {
	var pwd string
	var isAbsolutePath = strings.HasPrefix(x.String(), "/") || strings.HasPrefix(x.String(), "\\")
	if !isAbsolutePath {
		if pwd, err = os.Getwd(); err != nil {
			return
		}
	}

	var ls = filepath.SplitList(pwd)
	ls = append(ls, filepath.SplitList(x.String())...)
	path = filepath.Join(ls...)
	return
}

func (x Path) String() string {
	return string(x)
}
