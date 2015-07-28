package which

import (
	"os"
	"path/filepath"
)

func One(cmd string) string {
	return OneWithPath(cmd, os.Getenv("PATH"))
}

func All(cmd string) []string {
	return AllWithPath(cmd, os.Getenv("PATH"))
}

func OneWithPath(cmd, pathenv string) string {
	paths := which(cmd, pathenv, true)

	if len(paths) > 0 {
		return paths[0]
	}
	return ""
}

func AllWithPath(cmd, pathenv string) []string {
	return which(cmd, pathenv, false)
}

func isExecutable(f os.FileInfo) bool {
	return !f.IsDir() && f.Mode()&0111 != 0
}

func which(cmd, pathenv string, onlyOne bool) (paths []string) {

	for _, dir := range filepath.SplitList(pathenv) {
		path := filepath.Join(dir, cmd)
		if fi, err := os.Stat(path); err != nil || !isExecutable(fi) {
			continue
		}

		paths = append(paths, path)

		if onlyOne {
			break
		}
	}

	return paths
}
