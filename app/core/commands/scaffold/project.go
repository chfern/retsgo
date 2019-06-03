package scaffold

import (
	"fmt"
	"github.com/otiai10/copy"
	"log"
	"os"
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
