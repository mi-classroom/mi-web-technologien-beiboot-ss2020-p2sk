# Golang Imageing Modul für Skalierung und Zuschnitt

* Status: akzeptiert
* Datum: 2020-04-06

## Kontext und Problemstellung

Die hochgeladenen Bilder sind in verschiedene Größen zu skalieren und ein qudaratisches Derivat zu erzeugen. Wie können Bilder in Golang skaliert und zugeschnitten werden?

## Betrachtete Optionen

* Paket disintegration/imaging
* Paket nfnt/resize

## Ergebnis der Entscheidung

Das Paket disintegration/imaging bietet umfangreichere Bildverarbeitungsfunktionen als nfnt/resize und wurde deshalb implementiert.

## Links

* [Github disintegration/imaging] [https://github.com/disintegration/imaging]
* [Github nfnt/resize] [https://github.com/nfnt/resize]