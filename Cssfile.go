package cssminify

import (
	"os"
	"path/filepath"
	"strings"
)

// returns all .css files in the passed directory
func CssFiles(dir string) []string {
	files := make([]string, 0)
	filepath.Walk(dir,
		func(root string, info os.FileInfo, err error) error {
			if strings.HasSuffix(root, ".css") && !strings.HasSuffix(root, "min.css") {
				files = append(files, root)
			}
			return err
		})
	return files
}

// returns all CSS files in the current directory
// might in the future also return HTML files to allow inline style minification
func Files() []string {
	return CssFiles(".")
}
