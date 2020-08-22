# Dinghy Project - Backend

This doc describes mainly the relevant features of the backend server. 

The backend is developed in the [golang](https://golang.org) programming language. The docker container provides the necessery go binaries and dependencies to develope and run the backend.

## Dev environment

The container is able to automatically recompile the backend while running. Therefore you don't need to restart the container while developing.

### Start container

Please read the [main README](../README.md) to start the backend service or use the following command from the **main directory**.

```
docker-compose up -d backend
```

### Open browser

```
[localhost:8080](localhost:8080)
```

### Stop container

```
docker-compose down [--remove-orphans]
```

## Testing

To test the app with the docker environment use this shell command.

```
docker-compose exec backend go test gallery/
```

## Benchmark

If you Benchmark linl golang testing.

```
docker-compose exec backend go test -bench gallery/
```

## Scaffolding

If you run the backend for the first time, there will be no images to show. To scaffold some example images you can run the scaffolding cli. This command will download images directly into the `uploads/` directory. The API is provied by [picsum.photos](https://picsum.photos/).

```
docker-compose exec backend go run scaffolding/scaffolding.go -h
```

Use the `-h` parameter to see all the possible flags. For example use `-c <int>` to specify the image count to download and `-d` to clear the `uploads/` dir.

## REST API v1

The Backend provides a REST API for accessing the existing image collections. The API is accessible via the URI `<host>/rest/v1/collections`. The collections can be retrieved via GET requests.

The collections can also be restricted and sorted using query parameters. To determine the number of collections, use the parameter `count`. With `from` you can set the startposition. The values `alpha`, `color`, `date`, and `random` are available for the sorting parameter `sort`. The defaults are `count=10`, `from=0` and `sort=alpha`.

For example:

```
GET http://localhost:8080/rest/v1/collections?count=10&from=0&sort=alpha <- this would also be the default

GET http://localhost:8080/rest/v1/collections?sort=random
```

The API delivers a JSON in the following format.

```
[{
    id: <int>,
    images: [{
        name: <string>,
        uri: <string>,
        width: <int>,
        height: <int>,
        orientation: <string>
    }],
    colors: [
         {
            "rgba": {
                "R": <int>,
                "G": <int>,
                "B": <int>,
                "A": <int>
            },
            "hsl": {
                "H": <float>,
                "S": <float>,
                "L": <float>,
                "A": <int>
            },
            "quantity": <int>,
            "vibrant": <string>
        },
    ]
    },
... 
]
```

The `uri` will be relativ to the host `uploads/<id>/<name>`. So you have to add the hostname by yourself.

## Decoupled frontend dev environment with docker

If your not interested in the backend, this project provides a decoupled frontend dev environment to develope an own frontend and at the same time using the backend REST API. Open the frontend [README.md](../frontend/README.md) to see what features are there and how to use them.

## License

This project is licensed under the terms of the MIT license. See [License](LICENSE.md).