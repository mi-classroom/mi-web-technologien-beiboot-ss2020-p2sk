# Dinghy Project - Web Technologies - SS2020

[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0%20adopted-ff69b4.svg)](docs/CODE_OF_CONDUCT.md) 
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE.md)

The "dinghy-project" at TH Köln - University of Applied Sciences in the course of [studies information technology](https://www.medieninformatik.th-koeln.de/study/master/schwerpunkte/weaving-the-web/) - is to train and improve the working and development of [code tasks](https://github.com/mi-classroom/mi-master-wt-beiboot-2020/issues).

This project is split up in two parts. One part describes a backend server writen in golang as the main project. The another part is a basic frontend development environment. where others can develope there ideas based on the images provided through the backend. This README will give you an overview and a basic unterstanding of the project and how to use it in a whole. For a deeper understanding for one of the parts, check out the [backend README](backend/README.md) and/or [frontend README](frontend/README.md).

While its possible to run the services on your local machine, it is highly recommend to use docker-compose to spin up the services. But we will go in depth in the next sections.


## Clone or fork

To start you have to clone this repository.

```
git clone https://github.com/mi-classroom/mi-web-technologien-beiboot-ss2020-p2sk
cd <pathTo>/mi-web-technologien-beiboot-ss2020-p2sk
```

Or fork this repo into your account and clone it from there.

## Docker

The whole project uses docker to simplify the development and deployment process. The `docker-compose.yml` holds two services, the `frontend` and the `backend`. Also every service distinguish between a `development` - which is the default - and a `production` environment. This section describes the basic commands to start the services. 


<!--Its possible to run the backend with docker. Then you don't need to install go. In the follow sections are the steps to run and reproduce the functions described above. The container is able to automatically recompile the backend while running. Therefore you don't need to restart the container when developing.-->

### Start container

```
docker-compose up -d
```

If you start the container for the first time, docker will download the images and build the container.

### Observe container state

To check if the containers are running use the follow command.

```
docker-compose ps
```

You will most likly see somthing like this.

```
              Name                            Command               State           Ports         
 -------------------------------------------------------------------------------------------------
 mi-wt-beiboot-ss2020_backend_1    /bin/sh -c CompileDaemon - ...   Up      0.0.0.0:8080->8080/tcp
 mi-wt-beiboot-ss2020_frontend_1   docker-entrypoint.sh npm r ...   Up      0.0.0.0:8081->8081/tcp
```

### Open in browser

The backend is reachable over [localhost:8080](http://localhost:8080) and comes with its own little user interface.
For the frontend open [localhost:8081](http://localhost:8081).

### Monitor the container output

Especially in `development` mode you will receive valuable feedback from the different watchers, running in the `frontend` container background. See [frontend README](frontend/README.md#watchers). To observe these use this command. The `-f` option will continue the output. To abort use `Ctrl+c`.

```
docker-compose logs -f [frontend|backend]
```

### Stop container

To stop the container use the following command.

```
docker-compose down [--remove-orphans]
```

## Example PWA implementation

An example is accessible via the [github.io page](https://p2sk.github.io/mi-web-technologien-beiboot-ss2020-Dominikdeimel/Frontend/devPage/public/). Descisions made for the PWA can be viewed in the [Decision.md](https://github.com/p2sk/mi-web-technologien-beiboot-ss2020-Dominikdeimel/blob/master/Frontend/Desicion.md). The implementation is based on Dominik Deimel Repos. The [fork is here](https://github.com/p2sk/mi-web-technologien-beiboot-ss2020-Dominikdeimel).

## How to contribute

For contributing to this project please review the [Contributing](docs/CONTRIBUTING.md) guide. If you encounter any issue or can't fix it yourself, don't hesitate to open an issue.

## License

This project is licensed under the terms of the MIT license. See [License](LICENSE.md).

### Third party software

This project uses third party software.

#### Frontend

* https://github.com/nodejs/node License: see https://raw.githubusercontent.com/nodejs/node/master/LICENSE
* https://github.com/remy/nodemon License: MIT
* https://github.com/cloudhead/node-static License: MIT
* https://github.com/eslint/eslint License: MIT
* https://github.com/microsoft/TypeScript License: Apache-2.0
* https://github.com/sass/sass License: MIT

#### Backend

* https://github.com/gin-gonic/gin License: MIT
* https://github.com/disintegration/imaging License: MIT
* https://github.com/RobCherry/vibrant License: MIT
* https://github.com/githubnemo/CompileDaemon License: BSD-2
