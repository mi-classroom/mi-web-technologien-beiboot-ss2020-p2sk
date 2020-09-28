# Development und Production Environment

* Status: akzeptiert
* Workload: 10h
* Datum: 12.08.2020

## Kontext und Problemstellung

Für das Issue #6 soll ein Frontend Dev Container bereitgestellt werden, welcher einen Entwicklungs und Produktions Modus bereit stellt. Wie sind diese sinnvoll und schlank zu integrieren?

## Betrachtete Optionen

* Docker Environment Variablen
* node Environment Variablen

## Ergebnis der Entscheidung

Die gewählte Lösung stellt eine Option beider Möglichkeiten dar und ist in seiner Komplexität nicht klar in einem ADR zu formulieren. Grund dafür ist, dass die Entwicklungsumgebung vollständig im Docker Kontext lauffähig sein soll, was wiederum die Frage nach den node_modules eröffnet. Das heißt, dass - nachdem der Container gestartet wurde - keine weiteren Abhängigkeiten auf dem Host installiert werden sollen, sondern diese im Container selbst vorliegen. Hieraus folgt wiederum, dass die Binaries nur aus dem Container aufgerufen werden dürfen, um Seiteneffekte (z.B. Probleme mit Dateirechten) zu vermeiden.

In der docker-compose Datei kann die NODE_ENV gesetzt werden. Das Image führt zum Start das Default Command unter der in der package.json definierten _start_ Skript aus. Diese deligiert - abhängig der definierten NODE_ENV Variable - entweder zum dev oder zum prod Skript und führt die dort beschriebene Buildchain aus.

Siehe zu Buildchain [ADR #0015 - Buildchain](0015-feature-6-buildchain.md)

## Links

* [ADR #0015 - Buildchain](0015-feature-6-buildchain.md)