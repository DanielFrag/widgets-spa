# Widget-spa (Red Venture)
API escrita em Golang para fornecer recursos para o frontend contido em:

```
https://github.com/RedVentures/widgets-spa
```

# Requisitos

- Criar todo o serviço de backend para integrar com a aplicação acima;
- Autenticar a API via token;
- Usar a linguagem Go (Golang);
- Persistir os dados em alguma base de dados de sua preferência;

# Descrição da Solução

A aplicação foi organizada nos seguintes pacotes:
- infra - contém o código de acesso e manipulação da base de dados (MongoDB)
- repository - contém a definição das interfaces, usadas pela camada de negócio, para acesso a base de dados
- handler - contém a lógica de manipulação de dados e disponibilização dos recursos
- model - contém a definição das estruturas de dados a serem usadas na aplicação (objetos do domínio)
- router - contém a configuração das formas de disponibilização dos recursos
- utils - contém as ferramentas utilizadas em diversas partes da aplicação que não fazem parte da lógica do domínio

## Testes

Para rodar os testes do pacote 'handler' ou do pacote utils basta navegar até a pasta que contém o pacote e digitar o comando:
```
go test
```
Vale ressaltar que os testes do pacote 'handler' fazem uso de um mock que deve respeitar a interface definida no pacote 'repository'

Para rodar os testes do pacote de 'infra' são necessárias algumas observações.
Seus testes fazem acesso real ao Mongo por isso, o mesmo precisa estar rodando.
Os testes desse pacote incluem e manipulam dados em uma base de teste a ser excluídaquando o teste é concluído. 
Os testes 'TestUserMGO' e 'TestWidgetMGO' devem ser rodados separadamente, uma vez que cada um deles, ao finalizar, fecha a seção com o banco. Para rodá-los basta navegar até a pasta 'infra' e digitar o comando:

```
go test -run <nome_do_teste>
```

## Observações quanto aos Middlewares

Para isolar responsabilidades e organizar a hierarquia de execução das funções que atendem aos recursos solicitados foi utilizada uma abordagem das funções semelhante a organização de middlewares adotada pelo framework Express.js do Node.js.
Esta organização impõe que todo middleware receba como parâmetro um adapter 'http.HandlerFunc' e que retorne um adapter do mesmo tipo.
Assim as funções são organizadas e injetadas pelas instâncias da struct do tipo 'HandlerFuncInjector'.

## Variáveis de ambiente

As seguintes variáveis de ambiente possuem valores padrões que podem ser redefinidos para esta aplicação:
- MONGODB_URI contém a url de conexão com o banco (mongodb://localhost:27017)
- DB_NAME contém o nome da base de dados (widgets-spa-rv)
- PORT contém a porta a ser escutada pelo servidor (8080)
- TOKEN_SECRET contém a chave para assinatura do token (secret)
- TOKEN_EXPIRATION_TIME contém o número de minutos para expiração de um token (60)

## Bibliotecas de terceiros

As seguintes bibliotecas foram usadas neste projeto:
- github.com/gorilla/context
- github.com/gorilla/mux
- github.com/dgrijalva/jwt-go

## Rodar a API

Após instalar as dependências listadas acima, basta rodar no terminal o comando:

```
go build
```

Um novo executável será criado na raiz do projeto. Então, basta digitar no terminal o comando:

```
./<nome_do_executável>
```

## Teste das rotas

O arquivo 'widgets-spa-rv.postman_collection.json' pode ser importado pelo software 'Postman'.
Ao importar o arquivo, será criada uma nova coleção no Postman. Esta coleção contém os dados para requisitar os recursos da API, além de alguns resultados já gravados como exemplo.

## Rotas

/login
```
Método
	POST
Headers
	Content-Type: application/json
Body
	{
		"login": "admin",
		"password": "admin"
	}
```

/users
```
Método
	GET
Headers
	Authorization: <token>
```

/users/:id
```
Método
	GET
Headers
	Authorization: <token>
```

/widgets
```
Método
	GET
Headers
	Authorization: <token>
```

/widgets/:id
```
Método
	GET
Headers
	Authorization: <token>
```

/widgets
```
Método
	POST
Headers
	Authorization: <token>
	Content-Type: application/json
Body
	{
		"name": "morango dmfs",
		"color": "red",
		"price": "6.40",
		"inventory": 6,
		"melts": false
	}
```

/widgets/:id
```
Método
	PUT
Headers
	Authorization: <token>
	Content-Type: application/json
Body
	{
		"name": "morango DMFS",
		"color": "blue",
		"inventory": -1
	}
```

## Autenticação na API

Todas as rotas realizam checagem do token do usuário, com exceção da rota que fornece o token. O token será fornecido pela rota '/login' usando o médoto POST e passando um json que contenha os dados de login e senha do usuário.

## Inserir um usuário na base de dados

Para inserir o primeiro usuário no MongoDB é necessário seguir os seguintes passos:
- Iniciar o MongoDB (pular este passo caso o MongoDB já esteja sendo executado)
```
$ mongod
```
- Abrir shell do MongoDB
```
$ mongo
```
- Criar a database
```
> use widgets-spa-rv
```
- Criar a collection de usuário
```
> db.createCollection("user")
```
- Inserir o primeiro documento
```
> db.user.insert({login: "admin", password: "admin", session: "", gravatar: "/admin"})
```
- Encerre o mongo
```
exit
```