package envtok8s

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"
)

type ENVPairs struct {
	Key   string
	Value string
}

func ReadEnv(filepath string) ([]ENVPairs, error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return []ENVPairs{}, err
	}

	lines := strings.Split(string(content), "\n")
	trimmedLines := []string{}
	for _, line := range lines {
		if len(strings.TrimSpace(line)) > 0 {
			trimmedLines = append(trimmedLines, line)
		}
	}

	entries := []ENVPairs{}
	for _, line := range trimmedLines {
		parts := strings.Split(line, "=")

		if len(parts) < 2 {
			return []ENVPairs{}, fmt.Errorf("env must have key/value pairs: %v", parts)
		}

		if len(parts) > 2 {
			for i := 2; i < len(parts); i++ {
				parts[1] = fmt.Sprintf("%s=%s", parts[1], parts[i])
			}
		}

		parts = []string{parts[0], parts[1]}

		entries = append(entries, ENVPairs{
			Key:   strings.TrimSpace(parts[0]),
			Value: strings.TrimSpace(parts[1]),
		})
	}
	return entries, nil
}

func CreateSecret(env []ENVPairs, secretName string, namespace string) string {
	top := fmt.Sprintf(`apiVersion: v1
kind: Secret
metadata:
  name: %s
  namespace: %s
type: Opaque
data:`, secretName, namespace)

	lines := []string{}
	for _, entry := range env {
		encodedValue := b64.StdEncoding.EncodeToString([]byte(entry.Value))
		lines = append(lines, fmt.Sprintf("  %s: %s\n", entry.Key, encodedValue))
	}

	data := strings.Join(lines, "")

	return fmt.Sprintf("%s\n%s", top, data)
}
