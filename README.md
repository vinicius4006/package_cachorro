# Como Criar e Publicar Packages

Divido em duas partes

## Instalação

Crie a pasta que será seu package

```bash
mkdir package_exemplo
cd ./package_exemplo
```
Depois coloque o arquivo go.mod
## Adicionando go.mod

```go
go mod init github.com/seu-usuario/o-nome-da-pasta-que-tem-o-pkg
```

```go
go mod init github.com/teste_user/package_exemplo
```

## Depois de criar o seu repositório no GitHub
O código abaixo é dentro da pasta package_exemplo...
```git
git init
git add .
git commit -m "first commit"
git branch -M main
git remote add origin github.com/teste_user/package_exemplo.git
git push -u origin main
```

Coloque flags
```git
git tag "v1.0.0"
git push --tags
```

Agora crie outra pasta
```bash
mkdir testando_test
cd ./testando_test
```
Execute
```go
go mod init testando_test
go get -u github.com/teste_user/package_exemplo
```
Ele irá fazer o download e adicionará no arquivo .mod

Espero ter ajudado, forte abraço!
