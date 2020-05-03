# Gin Gonic als HTTP Backend Web-Framework

* Status: akzeptiert
* Workload: 8h
* Datum: 2020-04-06

## Kontext und Problemstellung

Das Golang net/http und html/template Paket unterstützen alle notwendigen Funktionen um einen Server zu betreiben. Um Boilerplate zu vermeiden und wiederkehrende http Aufgaben zu abstrahieren sollte ein Blick in vorhande Golang Webframeworks geworfen werden.Welche Webframeworks sind vorhanden und welches Framework wäre ggf. geeignet?

## Betrachtete Optionen

* Martini
* Gin Gonic
* Golang net/http

## Ergebnis der Entscheidung

Es wurde sich für Gin Gonic entschieden. Gin hat eine 40x bessere Performance gegenüber Martini, unterstützt einen httprouter, Middlewares und die native Kompatibilität zum net/http sowie html/template Pakets.

## Links

* [Top Golang Web-Frameworks 2019-2020](https://www.mindinventory.com/blog/top-web-frameworks-for-development-golang/)
* [Gin Gonic](https://github.com/gin-gonic)
* [Golang net/http](https://golang.org/pkg/net/http/)
* [Golang html/template](https://golang.org/pkg/html/template/)

