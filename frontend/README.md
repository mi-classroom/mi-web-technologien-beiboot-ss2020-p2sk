# Dinghy Project - Frontend

The container supportes some functions to facilitate a build chain. That meens you 

The docker container provides all necessary functions to develope your own frontend. For the ease of use it holds all functions. This means that the container will serve the site (nodejs server), watches and compiles the sass/scss & javascript/typescript files and gives you feedback, if your scripts are written correct with a linter. Based on the choosen mode ´develompment´ or ´production´ the buildchain will be slietly different. E.g. in production mode the buildchain will compile, transpil and minify the files once into the container and will serve your site without the ability to modify it.

Please be aware that you have two different uris to access the backend service. To access the backend from within your js scripts (see for example ´frontend/src/scripts.ts´) you are actually accessing it from outside the frontend container so that you need to use `http://localhost:8080`. From within your frontend container you need to change the hostname to `http://backend:8080` to access the backend container.

## Features

* Sass/Scss
* Typescript
* Eslint
* nodemon
* Minification

https://github.com/cloudhead/node-static

## Directory structure

The server.js will serve the `public` folder. That means all importand files have to go here.


## Development Mode

If in developement mode, this command will give you feedback from the sass, typescript and eslint watcher.

```
docker-compose logs -f frontend
```

### Wachter

The watchers are running in parallel.

#### Sass/Scass

The sass watcher will compile all style files from `styles/` to `public/css/`.

#### Typescript

The typescript/javascript files will also be watched and if changes happen gets compiled from `src/` to `public/js/`. The typescript is configured in `tsconfig.json` and will compile `*.js` files too. More infos about the `tsconfig.json` can be found here https://www.typescriptlang.org/tsconfig. General infos about typescript can be found in the official docs https://www.typescriptlang.org/docs.

#### Eslint

todo


## Production Mode

In production mode the files will be compiled and minified.

## Whats missing

To be small and lean the container provides only basic functions. If you need to add more modules, please install them directly into the container:

```
docker-compose exec -u root frontend npm install [--save-dev] [your-package]
```

Finally, if you encounter any issues and could't fix them yourself, please open an issue.