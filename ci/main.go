package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	// create dagger client
	ctx := context.Background()
	client, err := dagger.Connect(
		ctx,
		dagger.WithLogOutput(os.Stderr),
		dagger.WithWorkdir("../api"),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// login to docker hub
	dhUsername := os.Getenv("DOCKERHUB_USERNAME")
	dhPassword := client.SetSecret("password", os.Getenv("DOCKERHUB_PASSWORD"))

	// set dependency files
	dependecyFiles := []string{"go.mod", "go.sum"}
	dependencies := client.Host().Directory(".", dagger.HostDirectoryOpts{
		Include: dependecyFiles,
	})

	// set host directory
	project := client.Host().Directory(".")

	// build base image
	builderImage := client.Container().
		From("golang:1.20").
		WithWorkdir("/src").
		WithDirectory(".", dependencies).
		WithExec([]string{"go", "mod", "download"}).
		WithDirectory(".", project).
		WithExec([]string{"go", "build", "-o", "main", "main.go"})

	// build alpine image
	deployImage := client.Container().
		From("alpine:latest").
		WithExec([]string{
			"apk",
			"--no-cache",
			"add",
			"ca-certificates",
			"libc6-compat",
		}).
		WithWorkdir("/bin").
		WithFile("main", builderImage.File("main")).
		// WithExposedPort(8080).
		WithEntrypoint([]string{"./main"})

	// publish image to registry
	address, err := deployImage.WithRegistryAuth(
		"docker.io",
		dhUsername,
		dhPassword,
	).
		Publish(ctx, fmt.Sprintf("%s/gin", dhUsername))
	if err != nil {
		panic(err)
	}

	// print image address
	fmt.Println("Image published at:", address)
}
