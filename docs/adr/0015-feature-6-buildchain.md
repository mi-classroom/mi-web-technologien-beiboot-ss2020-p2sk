# Buildchain für Dev und Prod

* Status: akzeptiert
* Workload: 6h
* Datum: 12.08.2020

## Kontext und Problemstellung

Für das Issue #6 soll ein Frontend Dev Container bereitgestellt werden, welcher eine geradlinige Buildchain aufweist. Wie soll die Buildchain aussehen, wenn diese SASS kompilieren, JS/TS linten und transspilen sowie minifizieren soll?

## Betrachtete Optionen

* Sass Watch
* Typescript Watch
* eslint Watch
* nodemon
* node-minify

## Ergebnis der Entscheidung

Für den Entwicklungsmodus werden Watcher parallel laufend eingesetzt, die bei Änderung des jeweiligen Assets arbeiten. 

* SCSS Dateien in `styles/` werden nach `public/css` kompiliert
* Javascript / Typescript Dateien in `src/` werden gemäß der `tsconfig.json` nach `public/js` transspilt
* Javascript / Typescript Dateien werden per `esw` gelintet (Konfiguration siehe `.eslintrc.json`)

Ändert sich zudem die `server.js`, wird per nodemon der Frontend Server neugestartet.

---

Im Production Mode führt die Buildchain sequenziell den Bau der Styles, der Skripte, das Minification und den Start des Servers aus.