# Server in Dockercontainer laufen lassen

* Status: akzeptiert
* Datum: 2020-04-16

## Kontext und Problemstellung

Für das 2 Feature soll u.a. die Anwendung per Dockercontainer lauffähig sein. Soll nur die Anwendung lauffähig sein oder auch als Entwicklungsumgebung dienen? Für eine Entwicklungsumgebung müsste überprüft werden, inwieweit innerhalb des Containers die Anwendung bei Änderung des Quellcodes neu kompiliert würde, ohne jedes mal den Image neu zu bauen.

## Betrachtete Optionen

* Dockerfile
* Docker Compose 

## Ergebnis der Entscheidung

Es wurde eine Dockerfile im Backend angelegt die dann vom Docker Compose genutzt wird. Zudem ist der Container in der Lage die Server bei Änderung des Quellcodes neu zu compilieren und zu starten.

## Links

* [Hotreload in Docker Container für Server] [https://levelup.gitconnected.com/docker-for-go-development-a27141f36ba9]
* [Github githubneom/CompileDaemon] [https://github.com/githubnemo/CompileDaemon]