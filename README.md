# Primeiro
## Dentro da pasta do seu package execute go mod init github.com/seu-usuario/nome-da-pasta-do-seu-package

# Segundo
## Crie o repositório no GitHub e nome do repositório e o mesmo da pasta.
echo "# package_cachorro" >> README.md
git init
git add .
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/seu-usuario/nome-da-pasta-do-seu-package.git
git push -u origin main

# Terceiro
## Adicione a tag com git tag "v1.0.0" por exemplo e depois git push --tags