package scaffold

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"strings"
)

func BeginProjectScaffold(c *cli.Context) {
	if len(c.Args()) != 1 {
		log.Fatalln("Please provide the project repo path.")
	}
	repo := c.Args().First()
	GOPATH := GetGOPATH()
	targetPath := fmt.Sprintf("%s/src/%s", GOPATH, repo)
	slicedRepo := strings.Split(repo, "/")
	if len(slicedRepo) < 2 {
		log.Fatal("Invalid repo format.")
	}
	fmt.Printf("Creating new project in: %s\n\n", targetPath)
	CopyProjectSkeleton(targetPath)
}
