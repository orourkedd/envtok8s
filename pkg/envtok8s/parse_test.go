package envtok8s

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadEnv(t *testing.T) {
	env, err := ReadEnv(".env.test")
	assert.Nil(t, err)

	assert.Equal(t, "SECRET", env[0].Key)
	assert.Equal(t, "SECRETVALUE", env[0].Value)

	assert.Equal(t, "WITHEQUALS", env[1].Key)
	assert.Equal(t, "ABC=123=DEF=456", env[1].Value)

	assert.Equal(t, "ENV", env[2].Key)
	assert.Equal(t, "test", env[2].Value)
}
func TestCreateSecret(t *testing.T) {
	env, err := ReadEnv(".env.test")
	assert.Nil(t, err)

	secretName := "secret-name"
	namespace := "test"

	expected := `apiVersion: v1
kind: Secret
metadata:
  name: secret-name
  namespace: test
type: Opaque
data:
  SECRET: U0VDUkVUVkFMVUU=
  WITHEQUALS: QUJDPTEyMz1ERUY9NDU2
  ENV: dGVzdA==
`

	actual := CreateSecret(env, secretName, namespace)
	assert.Equal(t, expected, actual)
}
