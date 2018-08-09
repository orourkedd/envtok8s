package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/orourkedd/envtok8s/pkg/envtok8s"
)

func main() {
	filePtr := flag.String("file", ".env", "an env file")
	namespacePtr := flag.String("namespace", "default", "a K8S namespace")
	secretNamePtr := flag.String("secret-name", "", "a K8S secret name")

	flag.Parse()

	env, err := envtok8s.ReadEnv(*filePtr)
	if err != nil {
		panic(err)
	}

	if len(*secretNamePtr) == 0 {
		panic(errors.New("--secret-name is required."))
	}

	secret := envtok8s.CreateSecret(env, *secretNamePtr, *namespacePtr)

	fmt.Println(secret)
}
