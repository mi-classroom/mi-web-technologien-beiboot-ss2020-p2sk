# Beibootprojekt - PictureBox

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0%20adopted-ff69b4.svg)](docs/CODE_OF_CONDUCT.md) 
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE.md)

The "dinghy-project" at TH Köln - University of Applied Sciences in the course of [studies information technology](https://www.medieninformatik.th-koeln.de/study/master/schwerpunkte/weaving-the-web/) is to train and improve the working and development of [code tasks](https://github.com/mi-classroom/mi-master-wt-beiboot-2020/issues). The focus of this project is the development of a PictureBox.

## How to contribute

For contributing to this project please review the [Contributing](docs/CONTRIBUTING.md) guide.

## Setting up a local dev environment

This project is developed in the [golang](https://golang.org) programming language and should be installed beforehand to start the backend server (https://golang.org/doc/install). If you don't want to install go, it is also possible to run and develope the backend in a docker container. Please jump to section [dev environment with docker](#dev-environment-with-docker).

### 1. Clone repo

```
git clone https://github.com/mi-classroom/mi-web-technologien-beiboot-ss2020-p2sk
cd PathTo/mi-web-technologien-beiboot-ss2020-p2sk
```

### 2. Get Go dependencies

```
go get
```

Go tries to download the dependencies in $GOPATH. Please see [setting GOPATH](https://github.com/golang/go/wiki/SettingGOPATH) and [Go Environment variables](https://golang.org/cmd/go/#hdr-Environment_variables). For more informationen check [understanding the GOPATH](https://www.digitalocean.com/community/tutorials/understanding-the-gopath).

### 3. Build

```
go build
```

### 4. Start the server

```
cd backend
./backend[.exe]
```

## Dev environment with docker

Its possible to run the backend with docker. Then you don't need to install go. In the follow sections are the steps to run and reproduce the functions described above. The container is able to automatically recompile the backend while running. Therefore you don't need to restart the container when developing.

### Start container

If your container doens't exists, docker will build it for the first time.

```
docker-compose up -d
```

### Open in browser

```
localhost:8080
```

### Stop container

```
docker-compose down [--remove-orphans]
```

## Testing

To test in the local environment please use the following shell commands.

```
cd backend/gallery/
go test
```

### With docker

To test the app with the docker environment use this shell command.

```
docker-compose exec backend-golang-server go test gallery/
```

## Benchmark

```
cd backend/gallery/
go test -bench .
```

### With docker

```
docker-compose exec backend-golang-server go test -bench gallery/
```

## Scaffolding

To scaffold some example images you can run the scaffolding cli. This command will download images directly into the `uploads/` directory. The API is provied by [picsum.photos](https://picsum.photos/).

```
cd backend/scaffolding
go run scaffolding.go
```

Use the `-h` parameter to see all the possible flags. For example use `-c <int>` to specify the image count to download and `-d` to clear the `uploads/` dir.

```
go run scaffolding.go -h
```

### With docker

```
docker-compose exec backend-golang-server go run scaffolding/scaffolding.go -h
```

## REST API v1

The Backend provides a REST API for accessing the existing image collections. The API is accessible via the URI `/rest/v1/collections`. The collections can be retrieved via GET requests.

The collections can also be restricted and sorted using query parameters. To determine the number of collections, use the parameter `count`. The values `alpha`, `color`, `date`, and `random` are available for sorting parameter `sort`. The defaults are `count=10` and `sort=alpha`.

For example:

```
GET http://localhost:8080/rest/v1/collections?count=10&sort=alpha <- this would also be the default

GET http://localhost:8080/rest/v1/collections?sort=random
```

The API delivers a JSON in the following format.

```
[{
    id: <id>,
    images: [{
        name: <name>,
        width: <width>,
        height: <height>,
    }],
    colors: [
        {
            R: <r>
            G: <g>
            B: <b>
            A: <a>
        }
    ]
    },
... 
]
```

## Decoupled fronted dev environment with docker

If your not interested in the backend, this project provides a decoupled frontend dev enrivonment to develope an own frontend and at the same time using the backend REST API. Open the frontend [README.md](/frontend/README.md) to see what features are there and how to use them.


## License

This project is licensed under the terms of the MIT license. See [License](LICENSE.md).

### Third party software

This project uses third party software.

* https://github.com/gin-gonic/gin License: MIT
* https://github.com/disintegration/imaging License: MIT
* https://github.com/RobCherry/vibrant License: MIT
* https://github.com/githubnemo/CompileDaemon License: BSD-2