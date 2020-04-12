## Gedanken Feature 1 - Image Upload

### Aufgabe
Entwickeln Sie eine Client und eine Server Komponente, die einen Bildupload via http ermöglicht. Die Komponente muss folgende Funktionen bereitstellen:

- GUI zum Upload

Entscheidung gegen JS-Framework und CSS-Framework -> Vanilla HTML5/JS/CSS
Grund: in diesem frühen Stadium so flexibel wie möglich zu bleiben

- Backend zur Verarbeitung des Uploads

Golang Server (neue Technologie ausprobieren) + Gin Framework auf Basis von net/http

- Persistieren der Bilder auf der Serverseite

Middleware zur Persitierung, jeder Upload in eigenen Ordner, Ordnername Datum Uhrzeit

- Skalierung der Bilder auf eine konfigurierbare Größe

Middleware zur Skalierung, Golang Image Bib [Imaging][7], konfigurierbar per HTML (was bedeutet konfigurierbar) -> Ansatz per %-Faktor

- Skalierung der Bilder auf drei Größen zur Nutzung für verschiedene Geräteklassen

Was sind geeignete Größen? Wie sollten Bilder im Web verwendet werden? 
[Web Dev][1] gibt einen guten Einblick, inwieweit Bilder in Webseiten eingebunden werden sollten.
Interessant ist das picture-Element für Art Directions. Unter [statcounter][4] finden sich die aktuellen 
Bildschirmauflösungen. Geräteklassen differenzieren sich in Desktop [1366x768 (\~23%), 1920x1080 (\~21%)], Tablet [768x1024 (\~51%), 1280x800 (\~7%)] und Mobile [360x640 (\~19%), 375x667 (\~7%)].

- Erzeugung eines quadratischen Derivats des Bildes

Je nach Seitenverhältnis auszuwählen (Ausschnitt)

- Anzeigen aller hochgeladenen Bilder auf einer Übersichtsseite




Skalierung auf CLient oder Server vornehmen?
Was ist mit konfigurierbar gemeint? Prozentangabe? Auch Hochskalierung möglich?

Skalierung auf drei definierte Größen für verschiedene Geräteklassen siehe [Web Dev][1] für geeignete Betrachtungen

Quadratisches Derivat erzeugen: welche Größe? Ausschnitt aussreichend?

Übersichtsseite mit Uploads, Pagination? Liefert Server fertiges Html oder soll Frontend Bilder einzeln Fetchen?

### Client-Komponente
GUI für Upload
- Format
- Dateigröße
- AJAX-Post
- Frontend Framework suchen
- Vorhande Componenten suchen


### Server-Komponente
- http (node oder go)
- Nimmt Bilder entgegen (POST application-form?)
- Speichert dieses in geeigneter Form und Ort
- 


[1]: https://developers.google.com/web/fundamentals/design-and-ux/responsive/images?hl=de
[2]: http://responsiveimages.org/demos/
[3]: https://cloudfour.com/thinks/a-framework-for-discussing-responsive-images-solutions/
[4]: https://gs.statcounter.com/screen-resolution-stats/mobile/germany
[5]: https://www.html5rocks.com/en/mobile/high-dpi/#toc-tech-overview
[6]: https://de.wikipedia.org/wiki/Skalierung_(Computergrafik)
[7]: https://github.com/disintegration/imaging
[8]: https://github.com/imgproxy/imgproxy