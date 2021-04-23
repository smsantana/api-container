# API Construida em camadas |  Dockerizada

# Realizando build usando o cli do go
o comando abaixo alem de ja subir um conteiner de todas as dependencias da aplicação(Redis, Postgres) tambem realiza o build da imagem atravez da chamada ao Dockerfile

    #go get
	#go build -o apifood
    #./apifood


# Realizando build via Docker
o comando abaixo alem de ja subir um conteiner de todas as dependencias da aplicação(Redis, Postgres) tambem realiza o build da imagem atravez da chamada ao Dockerfile

	#docker-compose build
    #docker-compose up -d

# Sugestão de comida

Essa api foi construida usando as melhores praticas de desenvolvimento e usando o conceito de camadas


# Funcionalidades

 1. Manter usuario
 2. Manter comidas associado ao usuario


Lista de endpoints

POST   /usuarios

    {	"id": 1,
    	"first_name": "saul",
    	"last_name": "sobrenome",
    	"email": "meuemail@hotmail.com",
    	"password": "minhasenha"
    }

GET    /usuarios
GET    /usuarios/:usuar_id
POST   /comida

    {
        "id": 9,
        "user_id": 1,
        "title": "Hambunger Vegetariano",
        "description": "Hamburguer Vegetariano Com Queijo Branco",
        "food_image": "",
        "created_at": "2021-04-23T03:05:57.270487-03:00",
        "updated_at": "2021-04-23T03:05:57.270487-03:00",
        "deleted_at": null
        }


PUT    /comida/:com_id

    ...

GET    /comida/:com_id

    ...

DELETE /comida/:com_id

    ...

GET    /comida

    [
        {
            "id": 9,
            "user_id": 1,
            "title": "comida6",
            "description": "desc comida6",
            "food_image": "",
            "created_at": "2021-04-23T06:05:57.270487Z",
            "updated_at": "2021-04-23T06:05:57.270487Z",
            "deleted_at": null
        },
        {
            "id": 2,
            "user_id": 1,
            "title": "comida2",
            "description": "desc comida2",
            "food_image": "",
            "created_at": "2021-04-23T04:28:45.225982Z",
            "updated_at": "2021-04-23T04:28:45.225982Z",
            "deleted_at": null
        },
        {
            "id": 1,
            "user_id": 1,
            "title": "comida1",
            "description": "desc comida1",
            "food_image": "",
            "created_at": "2021-04-23T04:27:43.545339Z",
            "updated_at": "2021-04-23T04:27:43.545339Z",
            "deleted_at": null
        }
    ]
    ...

POST   /login

    ...

POST   /logout

    ...

# Referencias

* https://golang.org/doc/
* https://pt.stackoverflow.com/
* https://dev.to/(Steven Victor)
* https://docs.docker.com/
* https://www.fiap.com.br/
