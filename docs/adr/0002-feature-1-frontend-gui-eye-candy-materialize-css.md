# Fronted GUI minimal wie möglich Eye-Candy Materialize CSS

* Status: akzeptiert
* Workload: 3h
* Datum: 2020-04-06

## Kontext und Problemstellung

Der Bildupload soll über einen GUI erfolgen. Wie soll diese GUI aussehen? Soll eine SPA entwicklet oder das HTML per Backend ausgelifert werden?

## Betrachtete Optionen

* Single Page Application (JS-Framework: React.js, Vue.js oder Angular.js)
* Vanilla wie möglich (HTML vom Backend ausgeliefert)
* CSS-Framework

## Ergebnis der Entscheidung

Es wurde sich für "Vanilla wie möglich" und "CSS-Framework" (konkret Materialize-CSS) entschieden. Der Grund hierfür ist, dass zum jetzigen Zeitpunkt im Frontend Designentscheidungen offen gehalten werden soll. Damit das GUI trotzdem etwas schöner aussieht, wird das CSS-Framework Materialize verwendet. Dies macht das GUI ansprechend, aber ist gleichzeitig unproblematisch zu entfernen/ersetzen. Das HTML wird über das Backend ausgelifert und per go "html/templates" orchestriert. Sollte sich im Laufe des Beibootprojekts Features oder Anforderungen ergeben, die für eine SPA oder JS-Framework sprechen, soll dieses ADR abgelöst werden.
