# Convert .env files to Kubernetes secrets

### Input (.env)
```
SECRET=SECRETVALUE
WITHEQUALS=ABC=123=DEF=456
ENV=test
```

### Output
```
apiVersion: v1
kind: Secret
metadata:
  name: supersecret
  namespace: default
type: Opaque
data:
  SECRET: U0VDUkVUVkFMVUU=
  WITHEQUALS: QUJDPTEyMz1ERUY9NDU2
  ENV: dGVzdA==
```

### Usage
```
envtok8s --file=.env --namespace=default --secret-name=my-service-secrets
```


### Usage with defaults
```
envtok8s --secret-name=my-service-secrets
```