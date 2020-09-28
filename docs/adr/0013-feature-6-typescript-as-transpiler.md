# Typescript als Transpiler

* Status: akzeptiert
* Workload: 1h
* Datum: 10.08.2020

## Kontext und Problemstellung

Für das Issue #6 soll ein Frontend Dev Container bereitgestellt werden. Welcher Transpiler soll gewählt werden?

## Betrachtete Optionen

* Webpack
* Babel
* Typescript

## Ergebnis der Entscheidung

Es wurde sich für Typescript als JS Transpiler entschieden. Treiber der Entscheidung war in diesem Fall den Overhead (Konfigurationsaufwand und Dependencies) gering zu halten. Typescript kann neben den .ts files auch Javascript Dateien kompilieren und führt dieses sehr leichtgewichtig durch. Webpack ist als Packer vom Funktionsumfang dem Transpiling von Typescript überlegen, allerdings steigt hier der Konfigurationsaufwand. 

## Links

* [Bable JS](https://babeljs.io/)
* [Typescript](https://www.typescriptlang.org/)
* [Webpack](https://webpack.js.org/)