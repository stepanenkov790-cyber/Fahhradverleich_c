# Projektregeln

## Ziel
Einfacher, verständlicher Prototyp für Fahrradverleih.

## Stack
- Backend: Go, Structs und Methoden
- Datenhaltung: thread-sicherer In-Memory-Store
- Frontend: HTML/CSS/JavaScript
- Keine zusätzlichen Dependencies ohne Nachfrage

## Architektur
- `main.go`: HTTP-Server und Routing
- `internal/domain/`: Datenmodelle
- `internal/store/`: Geschäftslogik und In-Memory-Daten
- `internal/handler/`: HTTP-Handler
- `web/`: Frontend

## Konventionen
- Methoden an Structs, wo sinnvoll
- Fehler nicht verschlucken
- Eingaben validieren
- Keine Secrets im Code
- Frontend kommuniziert ausschließlich über HTTP-API
- Einfach und klar statt unnötig komplex
