package scaffold

import (
	"fmt"
	"github.com/otiai10/copy"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// CopyProjectSkeleton copies a new project into the $GOPATH/src/repoName e.g: $GOPATH/src/github.com/githubuser/newproject
func CopyProjectSkeleton(targetPath string) {
	GOPATH := GetGOPATH()
	PKGPATH := GetPKGPATH()
	templateProjectPath := fmt.Sprintf("%s/src/%s", GOPATH, PKGPATH)

	_, err := os.Stat(templateProjectPath)
	if os.IsNotExist(err) {
		log.Fatalf("retsgo not found under %s.\n Try running go get %s", templateProjectPath, PKGPATH)
	}

	// Check if targetPath exist
	_, err = os.Stat(targetPath)
	if os.IsNotExist(err) {
		log.Fatalf("%s, directory not found", targetPath)
	}

	// Begin Copy Project Contents from templateProjectPath into targetPath
	copy.Copy(templateProjectPath, targetPath)

	// Remove unnecessary files & folders from target path (generated project)
	removeFiles := []string{
		"/main.go",
		"/app/core/commands",
		"/.git",
		"/docker-compose.local.yml",
		"/Dockerfile.dev",
	}
	for _, file := range removeFiles {
		err = os.Remove(fmt.Sprintf("%s%s", targetPath, file))

		if err != nil { // Probably because directory content is not empty
			ClearDir(fmt.Sprintf("%s%s", targetPath, file))
			os.Remove(fmt.Sprintf("%s%s", targetPath, file))
		}
	}
}

type ProjectFile struct {
	Path     string
	FileInfo os.FileInfo
}

func ChangeImportName(targetPath string, repo string) {
	// Get All files in project
	fileList := []ProjectFile{}
	err := filepath.Walk(targetPath, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, ProjectFile{
			Path:     path,
			FileInfo: f,
		})
		return nil
	})
	if err != nil {
		log.Fatalln("Project creation error")
	}

	// Replace def package name e.g: github.com/fernandochristyanto/retsgo with repo name
	PKGNAME := GetPKGPATH()
	for _, projectfile := range fileList {
		fmt.Printf("Generating %s\n", strings.Replace(projectfile.Path+"/", targetPath, "", -1))
		ReplaceStrInFile(projectfile.Path, projectfile.FileInfo, PKGNAME, repo)
	}
}
