# Fahrradverleih

Eigenständiger Prototyp für einen Fahrradverleih mit Go-Backend und HTML/CSS/JavaScript-Frontend.

## Start
Voraussetzung: Go 1.22 oder neuer.

```bash
go run .
```

Browser: http://localhost:8081

## Tests und Formatierung

```bash
go test ./...
gofmt -w .
```

## Funktionen
- Dashboard mit verfügbaren und ausgeliehenen Fahrrädern
- Fahrräder mit individuellem Tagespreis hinzufügen und löschen
- Kunden registrieren und löschen
- Fahrradverleih mit Mietdauer und automatischer Preisberechnung
- Reservierungen erstellen und anzeigen
- In-Memory-Speicherung mit thread-sicherem Store

Die Daten werden beim Neustart zurückgesetzt. Das ist für diesen Prototyp vorgesehen.
