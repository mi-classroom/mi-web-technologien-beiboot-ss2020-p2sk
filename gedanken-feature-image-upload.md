## Gedanken Feature 1 - Image Upload

### Aufgabe
Entwickeln Sie eine Client und eine Server Komponente, die einen Bildupload via http ermöglicht. Die Komponente muss folgende Funktionen bereitstellen:

- GUI zum Upload

Entscheidung gegen JS-Framework und CSS-Framework -> Vanilla HTML5/JS/CSS (Eye-Candy Materalize CSS [Materialize][8], aber leicht austauschbar)
Grund: in diesem frühen Stadium so flexibel wie möglich bleiben

- Backend zur Verarbeitung des Uploads

[Golang][10] Server (neue Technologie ausprobieren) + [Gin Framework][9] auf Basis von net/http

- Persistieren der Bilder auf der Serverseite

Middleware zur Persistierung, jeder Upload in eigenen Ordner, Ordnername Datum Uhrzeit

- Skalierung der Bilder auf eine konfigurierbare Größe

Middleware zur Skalierung, Golang Image Bib [Imaging][7], konfigurierbar per HTML (was bedeutet konfigurierbar) -> Ansatz per %-Faktor.
Auch Hochskalierung möglich? [Skalierung Definition][6]

- Skalierung der Bilder auf drei Größen zur Nutzung für verschiedene Geräteklassen

Was sind geeignete Größen? Wie sollten Bilder im Web verwendet werden? 
[Web Dev][1] gibt einen guten Einblick, inwieweit Bilder in Webseiten eingebunden werden sollten. Weitere Ressourcen: [responsiveimages.org][2]

Unter [statcounter][4] finden sich aktuelleBildschirmauflösungen.
Geräteklassen differenzieren sich in Desktop [1366x768 (\~23%), 1920x1080 (\~21%)], Tablet [768x1024 (\~51%), 1280x800 (\~7%)] und Mobile [360x640 (\~19%), 375x667 (\~7%)]. [HighDPI][5]

- Erzeugung eines quadratischen Derivats des Bildes

Je nach Seitenverhältnis auszuwählen (Ausschnitt). Ausschnitt ausreichend? [Imaging CropCenter][11]

- Anzeigen aller hochgeladenen Bilder auf einer Übersichtsseite

Übersichtsseite mit Uploads, Pagination? Liefert Server fertiges Html oder soll Frontend Bilder einzeln Fetchen?



[1]: https://developers.google.com/web/fundamentals/design-and-ux/responsive/images?hl=de
[2]: http://responsiveimages.org/demos/
[3]: https://cloudfour.com/thinks/a-framework-for-discussing-responsive-images-solutions/
[4]: https://gs.statcounter.com/screen-resolution-stats/mobile/germany
[5]: https://www.html5rocks.com/en/mobile/high-dpi/#toc-tech-overview
[6]: https://de.wikipedia.org/wiki/Skalierung_(Computergrafik)
[7]: https://github.com/disintegration/imaging
[8]: https://materializecss.com/
[9]: https://github.com/gin-gonic/gin
[10]: https://golang.org/
[11]: https://godoc.org/github.com/disintegration/imaging#CropCenter