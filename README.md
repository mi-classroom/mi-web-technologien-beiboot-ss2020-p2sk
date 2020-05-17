# Beibootprojekt - PictureBox

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0%20adopted-ff69b4.svg)](docs/CODE_OF_CONDUCT.md) 
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE.md)

The "dinghy-project" at TH Köln - University of Applied Sciences in the course of [studies information technology](https://www.medieninformatik.th-koeln.de/study/master/schwerpunkte/weaving-the-web/) is to train and improve the working and development of [code tasks](https://github.com/mi-classroom/mi-master-wt-beiboot-2020/issues). The focus of this project is the development of a PictureBox.

## How to contribute

For contributing to this project please review the [Contributing](docs/CONTRIBUTING.md) guide.

## Setting up the dev environment

This project is developed in the [golang](https://golang.org) programming language and should be installed beforehand to start the backend server (https://golang.org/doc/install).

### Clone repo

```
git clone https://github.com/mi-classroom/mi-web-technologien-beiboot-ss2020-p2sk
cd PathTo/mi-web-technologien-beiboot-ss2020-p2sk
```

### Get Go dependencies

```
go get
```

Go tries to download the dependencies in $GOPATH. Please see [setting GOPATH](https://github.com/golang/go/wiki/SettingGOPATH) and [Go Environment variables](https://golang.org/cmd/go/#hdr-Environment_variables). For more informationen check [understanding the GOPATH](https://www.digitalocean.com/community/tutorials/understanding-the-gopath).

### Build

```
go build
```

### Start the server

```
./backend[.exe]
```

### Docker

#### Start container

If your container doens't exists, docker will build it for the first time.

```
docker-compose up -d
```

#### Open in browser

```
localhost:8080
```

#### Stop container

```
docker-compose down [--remove-orphans]
```

## Testing

```
cd backend/gallery/
go test
```

## Benchmark

```
cd backend/gallery/
go test -bench .
```

## License

This project is licensed under the terms of the MIT license. See [License](LICENSE.md).

### Third party software

This project uses third party software.

* https://github.com/gin-gonic/gin License: MIT
* https://github.com/disintegration/imaging License: MIT
* https://github.com/ericpauley/go-quantize License: BSD-3
