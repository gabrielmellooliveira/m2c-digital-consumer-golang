# M2C Digital Consumer - Golang

Sistema da M2C Digital responsável processar as mensagens das campanhas.

## Rodando o projeto

### Instalando as dependencias

Após baixar o projeto na sua máquina, rode o seguinte comando para instalar as dependencias do mesmo:

```
go mod tidy
```

### Docker Compose

Para criar a instância do ```MongoDB``` e do ```RabbitMQ``` com Docker Compose, deve ser utilizado o seguinte comando:

```
docker-compose up -d
```

### Variaveis de ambiente

No projeto, há um arquivo chamado ```cmd/.env-example``` em que as informações devem ser copiadas para um arquivo chamado ```cmd/.env```.

Caso necessário, poderá alterar as informações do .env para apontar para sua aplicação, banco de dados ou ferramenta.

### Inicializando o projeto

Para rodar o projeto, utilize o comando:

```
go run cmd/main.go
```