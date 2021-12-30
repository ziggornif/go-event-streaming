# Go events streaming example with JetStream

## Install

```sh
go install
```

## Run

```sh
go run main.go
```

## Deploy dependencies

```sh
docker compose up -d
```

## How to use it ?

Open http://localhost:8080/listener/ URL in your browser.

Then call http://localhost:8080/tweet endpoint with a request like the following example :

```
POST http://localhost:8080/tweet

{
    "message": "Hello world !",
    "author": "ziggornif"
}
```

Once the entity is created, an event will be sent in Jetstream then retrieved by the event listener and displayed in the page opened earlier.  