package client

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

type InputImageReference interface {
	Name() string
	Layout() bool
	FullName() (string, error)
}

type defaultInputImageReference struct {
	name string
}

type layoutInputImageReference struct {
	name string
}

func ParseInputImageReference(input string) InputImageReference {
	if strings.HasPrefix(input, "oci:") {
		imageNameParsed := strings.SplitN(input, ":", 2)
		return &layoutInputImageReference{
			name: imageNameParsed[1],
		}
	}
	return &defaultInputImageReference{
		name: input,
	}
}

func (d *defaultInputImageReference) Name() string {
	return d.name
}

func (d *defaultInputImageReference) Layout() bool {
	return false
}

func (d *defaultInputImageReference) FullName() (string, error) {
	return d.name, nil
}

func (l *layoutInputImageReference) Name() string {
	return l.name
}

func (l *layoutInputImageReference) Layout() bool {
	return true
}

func (l *layoutInputImageReference) FullName() (string, error) {
	var (
		fullImagePath string
		err           error
	)

	path := l.name

	// path/to/save/image:tag was provided
	if strings.Contains(path, ":") {
		split := strings.SplitN(path, ":", 2)
		// do not include the tag in the path
		path = split[0]
	}

	if fullImagePath, err = filepath.EvalSymlinks(path); err != nil {
		if !os.IsNotExist(err) {
			return "", errors.Wrap(err, "evaluate symlink")
		} else {
			fullImagePath = path
		}
	}

	if fullImagePath, err = filepath.Abs(fullImagePath); err != nil {
		return "", errors.Wrap(err, "resolve absolute path")
	}

	return fullImagePath, nil
}
