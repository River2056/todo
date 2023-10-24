package assets

import (
	"os"
	"path/filepath"
	"regexp"
)

func GetAllTemplates() []string {
	fileList := make([]string, 0)
	filepath.Walk("./templates", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList
}

func GetSpecificHtml(allHtmls []string, name string) string {
	r := regexp.MustCompile(name)
	for _, htmlPath := range allHtmls {
		if r.Match([]byte(htmlPath)) {
			return htmlPath
		}
	}
	return ""
}
