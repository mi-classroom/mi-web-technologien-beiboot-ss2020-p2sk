# Quantisierung mit Paket go-quantize

* Status: ersetzt durch [ADR-0010](0010-feature-2-vibrant-as-new-color-quantizer.md)
* Workload: 8h
* Datum: 2020-04-16

## Kontext und Problemstellung

Es soll für die hochgeladenen Bilder eine Farbpalette erzeugt werden, die die wsesentlichsten Farben des Bildes representiert und als JSON Datei abgespeichert werden. Wie ist eine Farbquantisierung in golang vorzunehmen? 

## Betrachtete Optionen

* ericpauley/qo-quantize
* esimov/colorquant

## Ergebnis der Entscheidung

Für die Quantisierung wurde sich für ericpauley/qo-quantize entschieden, da dieses Paket die Möglichkeit hat, nur die quantisierte Farbpalette zu erhalten. 

## Links

* [Github ericpauley/qo-quantize] [https://github.com/ericpauley/go-quantize]
* [Github esimov/colorquant] [https://github.com/esimov/colorquant]