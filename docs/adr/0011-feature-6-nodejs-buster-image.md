# Nodejs Buster Image

* Status: akzeptiert
* Workload: 1h
* Datum: 10.08.2020

## Kontext und Problemstellung

Für das Issue #6 soll ein Frontend Dev Container bereitgestellt werden. Welches nodejs Image soll gewählt werden?

## Betrachtete Optionen

* nodejs:latest
* nodejs:buster
* nodejs:alpine

## Ergebnis der Entscheidung

Es wurde sich für das Image nodejs:buster entschieden. Das Buster Release stellt den aktuellen Entwicklungszyklus von Debian dar. Die Entwicklungsumgebung sollte so aktuell wie möglich sein. Daraus folgt auch Aktualität > Imagegröße.

## Links

* [Docker Hub nodejs](https://hub.docker.com/_/node/)