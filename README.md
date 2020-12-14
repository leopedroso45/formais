# Manual Instalação Go
# Linux:
link para download: https://golang.org/dl/
1. Abra o Terminal
2. Certifique-se de estar no diretório home
3. Execute o comando: curl -O https://dl.google.com/go/go1.12.1.linux-amd64.tar.gz
4. Logo após, execute: sha256sum go1.12.1.linux-amd64.tar.gz
5. Obs: O hash que é exibido a partir da execução do comando acima deve corresponder ao hash que estava
na página de downloads. Se não, então este não é um arquivo válido e você deve baixar o arquivo novamente.
6. Em seguida, extraia o arquivo baixado e instale-o no local desejado no sistema: sudo tar -xvf go1.12.1.linux-amd64.tar.gz -C /usr/local
7. Você terá agora um diretório chamado go no diretório /usr/local. Em seguida, 
altere recursivamente o proprietário e o grupo deste diretório para root: sudo chown -R root:root /usr/local/go

### 8. Outra alternativa (minha favorita) é a instalação pelo snap através do comando: sudo snap install go --classic
### verifique se o snap esta instalado, caso não: sudo apt install snapd 

9. Abra o gitHub para obter o código fonte, abra o arquivo main na sua IDE ou por comando no linux: go run main.go
esse comando deve ser utilizado no terminal da IDE também, se for o caso.

# Windows:
1. Abra o link: https://golang.org/dl/
3. Selecione a opção de dowload para o windows
4. após ter tudo instalado, abra sua IDE  e o arquivo e execute o comando: go run main.go no terminal
