# Helper

**How to Debug Dockerfile**

- run docker run -it {img_name} /bin/bash

## Generate Mock

**Generate Mock Cli**

```
mockery --name={interface_name} --dir={directory_name} --output={directory_output_name} --filename={filename}
```

**Generate Mock Example**

```
mockery --name=BaseRepositoryInterface --dir=./app/repository  --output=./app/mocks --filename=base_repository_mock.go
```

## Run Test

```
go test -v -count=1 ./...
```

## Generate Dependency Injection

```
wire gen ./app/module
```

## Swagger UI

```
http://localhost/api/docs/
```
