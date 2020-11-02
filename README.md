# JWT Implementation

## Commands

### Build

```console
docker-compose down && sudo rm -rf postgresql && go build main.go && docker-compose up -d --build
```

### Reset Database

```console
sudo rm -rf postgresql
```

## Sign Up

`POST` and located at [0.0.0.0:8000/signup](0.0.0.0:8000/signup) .

### Sign Up - cURL

```console
curl --location --request POST '0.0.0.0:8000/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "Adam",
    "password": "12345"
}'
```

### Sign Up - Note

After adding one...

![./docs/added%20new%20users.png](./docs/added%20new%20users.png)

You can check Adminer at [0.0.0.0:8080](0.0.0.0:8080) with this credential:

|Field|Value
|-|-|
|System|PostgreSQL|
|Server|go-auth-playground-postgresql|
|Username|gorm|
|Password|gorm|

There will be something as such:

![./docs/added%20new%20users%20database.png](./docs/added%20new%20users%20database.png)

## Sign In

`POST` and located at [0.0.0.0:8000/signin](0.0.0.0:8000/signin) .

### Sign In - cURL

```console
curl --location --request POST '0.0.0.0:8000/signin' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "Adam",
    "password": "12345"
}'
```

### Sign In - Note

![./docs/sign%20in.png](./docs/sign%20in.png)

## Accessing API of Which Needs User of Being Authorized First

`POST` and located at [0.0.0.0:8000/home](0.0.0.0:8000/home) .

### API - cURL

```console
curl --location --request POST '0.0.0.0:8000/home' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "Adam",
    "token": "<put-your-token-here>"
}'
```

### API - Note

#### With Random Token

![./docs/not%20authorized.png](././docs/not%20authorized.png)

#### With Valid Token

![./docs/authorized.png](././docs/authorized.png)

## To-Do and Improvements

* [ ] Unit tests duh...
* [ ] Lack of comments
* [ ] Config yang dari environment variables ganti pake config file
