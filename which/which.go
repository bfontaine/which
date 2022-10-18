package which

import (
	"os"
	"path/filepath"
)

// One returns the first executable path matching the given command. It returns
// an empty string if there's no such path.
func One(cmd string) string {
	return OneWithPath(cmd, os.Getenv("PATH"))
}

// All returns all executable paths. matching the given command. It returns nil
// if there's no such path.
func All(cmd string) []string {
	return AllWithPath(cmd, os.Getenv("PATH"))
}

// OneWithPath is like One, but it takes the PATH to check as a second argument
// instead of using the PATH environment variable.
func OneWithPath(cmd, pathenv string) string {
	paths := which(cmd, pathenv, true)

	if len(paths) > 0 {
		return paths[0]
	}
	return ""
}

// AllWithPath is like All, but it takes the PATH to check as a second argument
// instead of using the PATH environment variable.
func AllWithPath(cmd, pathEnv string) []string {
	return which(cmd, pathEnv, false)
}

func isExecutable(filepath string) bool {
	f, err := os.Stat(filepath)
	return err == nil && !f.IsDir() && f.Mode()&0111 != 0
}

func which(cmd, pathenv string, onlyOne bool) (paths []string) {
	if filepath.IsAbs(cmd) && isExecutable(cmd) {
		paths = append(paths, cmd)

		if onlyOne {
			return
		}
	}

	for _, dir := range filepath.SplitList(pathenv) {
		path := filepath.Join(dir, cmd)
		if !isExecutable(path) {
			continue
		}

		paths = append(paths, path)

		if onlyOne {
			break
		}
	}

	return
}
