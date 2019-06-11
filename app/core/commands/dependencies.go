package commands

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os/exec"
)

type Dependency struct {
	Name      string
	GithubURL string
}

func Dependencies() *[]Dependency {
	dependencies := []Dependency{
		Dependency{Name: "justinas/alice", GithubURL: "github.com/justinas/alice"},
		Dependency{Name: "gorilla/mux", GithubURL: "github.com/gorilla/mux"},	
		Dependency{Name: "json-iterator", GithubURL: "github.com/json-iterator/go"},
		Dependency{Name: "urfave/cli", GithubURL: "github.com/urfave/cli"},
		Dependency{Name: "jinzhu/gorm", GithubURL: "github.com/jinzhu/gorm"},
		Dependency{Name: "google/wire", GithubURL: "github.com/google/wire"},
		Dependency{Name: "sirupsen/logrus", GithubURL: "github.com/sirupsen/logrus"},
		Dependency{Name: "otiai10/copy", GithubURL: "github.com/otiai10/copy"},
		Dependency{Name: "dgrijalva/jwt-go", GithubURL: "github.com/dgrijalva/jwt-go"},
	}
	return &dependencies
}

func PrintDependencies(c *cli.Context) error {
	fmt.Println("External Dependencies:")
	dependencies := Dependencies()
	for _, dependency := range *dependencies {
		fmt.Printf("- %s %s\n", dependency.Name, dependency.GithubURL)
	}
	return nil
}

func InstallDependencies(c *cli.Context) error {
	dependencies := Dependencies()
	for _, dependency := range *dependencies {
		fmt.Printf("Installing %s\n", dependency.Name)
		cmd := exec.Command("go", "get", dependency.GithubURL)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println()
	fmt.Println("All external dependencies installed successfully!")
	return nil
}
