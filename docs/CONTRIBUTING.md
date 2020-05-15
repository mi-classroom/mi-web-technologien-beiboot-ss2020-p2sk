# Contributing

## Code of Conduct



## Design decisions


## Setting up a dev environment

### clone repo

```
git clone https://github.com/mi-classroom/mi-web-technologien-beiboot-ss2020-p2sk
cd PathTo/mi-web-technologien-beiboot-ss2020-p2sk
```

### Without Docker

#### Get Go dependencies

```
go get
```

#### Build

```
go build
```

#### Start the server

```
./backend[.exe]
```

#### Benchmark

```
cd backend/gallery/
go test -bench .
```

### With Docker

#### Build image

```
docker-compose build
```

#### Start container

```
docker-compose up -d
```

#### Open in browser

```
localhost:8080
```

### Stop container

```
docker-compose down [--remove-orphans]
```

## Testing

```
cd backend/gallery/
go test
```


## Styleguides for Go Code

* https://golang.org/doc/effective_go.html
* https://github.com/golang/go/wiki/CodeReviewComments