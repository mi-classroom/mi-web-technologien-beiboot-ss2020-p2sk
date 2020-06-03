# Scaffolding with picsum.photo

* Status: akzeptiert
* Workload: 5h
* Datum: 3.6.2020

## Kontext und Problemstellung

Nutzer sollen die Möglichkeit erhalten, exemplarisch Bilder zur Verfügung gestellt zu bekommen, um die Photobox direkt zu testen.

## Betrachtete Optionen

* flickr
* pixabay
* unsplash
* picsum.photos
* manuell Bilder im Web Suchen und in den Uploadfolder commiten

## Ergebnis der Entscheidung

Die Option manuell Bilder in das Repo zu comitten ist weniger elegant und widerspricht der Möglichkeit eines automatisierten Verfahrens. Die größeren Imagesites bieten allesamt eigene APIs an. Für die Nutzung wird aber zwingend ein API Key benötigt. Daher wird sich für die Lösung per automatisiertem download von picsum.photos entschieden, wenngleich der Zugriff irgendwann nicht mehr funktionieren könnte.

## Links

* [picsum.photos](https://picsum.photos/)