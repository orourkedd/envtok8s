package main

import (
	"flag"
	"fmt"

	"github.com/orourkedd/envtok8s/pkg/envtok8s"
)

func main() {
	filePtr := flag.String("file", ".env", "an env file")
	namespacePtr := flag.String("namespace", "test", "a K8S namespace")
	secretNamePtr := flag.String("secret-name", "my-secret", "a K8S secret name")

	flag.Parse()

	env, err := envtok8s.ReadEnv(*filePtr)
	if err != nil {
		panic(err)
	}

	secret := envtok8s.CreateSecret(env, *secretNamePtr, *namespacePtr)

	fmt.Println(secret)
}
