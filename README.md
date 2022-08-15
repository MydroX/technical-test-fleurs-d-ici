# Technical test fleurs d'ici 

## Install project

```bash
git clone https://github.com/MydroX/technical-test-fleurs-d-ici
```

Create a `.env` and `.env.db` for production, or `.env.local` and `.env.local.db` for dev,  following the examples 

- Example .env/.env.local file

```env
PORT=3000
JWT_SECRET="secret"
ENV=dev # dev or prod
```

- Example .env.db/.env.local.db file

```
POSTGRES_USER=user
POSTGRES_PASSWORD=root
POSTGRES_DB=fleurs-d-ici
POSTGRES_PORT=5432
POSTGRES_HOST=fleursdici-postgres

PGDATA=/data/postgres

PGADMIN_DEFAULT_EMAIL=john@doe.com
PGADMIN_DEFAULT_PASSWORD=root
PGADMIN_LISTEN_PORT=5050
```

Note : If you modify the ports in the files you must modify the ports also in the docker-compose.yml files.

## Run project

Run dev mode
```bash
docker-compose -f deploy/docker-compose.yml -f deploy/dev/docker-compose.yml up
```

Run prod mode
```bash
docker-compose -f deploy/docker-compose.yml -f deploy/prod/docker-compose.yml up --build 
```

## API Documentation

### Server

Server is running at http://localhost:3000 , unless you change the port in the config file.

### Login

**URL** : `/login`

**Method** : `POST`

**Data constraints**

```json
{
    "username": "[valid email address]",
    "password": "[password in plain text]"
}
```

**Data example**

```json
{
    "username": "iloveauth@example.com",
    "password": "abcd1234"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "token": "93144b288eb1fdccbe46d6fc0f241a51766ecd3d"
}
```

### Register

**URL** : `/register`

**Method** : `POST`

**Data constraints**

```json
{
    "username": "[valid email address]",
    "password": "[password in plain text]",
    "password_confirmation": "[password in plain text]"
}
```

**Data example**

```json
{
    "username": "iloveauth@example.com",
    "password": "abcd1234",
    "password_confirmation": "abcd1234"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "message": "user successfully created"
}
```