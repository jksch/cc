# Coding-Challenge Schulungsanbieter (freieTechnologiewahl)
Ein Anbieter von Schulungen möchte seine Produkte online anbieten. Dazu benötigt er eine Web-Anwendung, auf der sich potentielle Kunden
über das Schulungsangebot informieren und Schulungen buchen können. Jede Schulung wird an verschiedenen Terminen angeboten. An jedem
Termin stehen 10 Teilnehmerplätze zur Verfügung, die von potentiellen Kunden einzeln gebucht werden können.
Software-architektonisch ist ein Frontend in Javascript angedacht, das über eine REST-API mit einem Backend kommuniziert.

## 1. REST-API-Design
Designe eine REST-API, mit der die folgenden Frontend-Funktionalitäten abgebildet werden können:
* Übersicht über alle angebotenen Schulungen (Name der Schulung, Beschreibung, Name des Dozenten, Preis, ...)
* Anzeige der Schulungen in einem bestimmten Zeitraum
* Anzeige der Termine zu einer bestimmten Schulung
* Anlegen/Verändern einer neuen Schulung
* Anlegen/Verändern eines neuen Termins für eine Schulung
* Buchung einer Schulung an einem bestimmten Termin

Diese Aufgabe ist reine Konzeption. Es geht nicht um das Schreiben von Quellcode, sondern um die Definition der API.

---
### 1.1 Übersicht über alle angebotenen Schulungen:
```
GET /api/courses?startDate=[RFC3339]&endDate=[RFC3339]
```
Success:
```
Code: 200
ContentType: 'application/json'
Body: [
    {
        "id": 1,
        "name": "Go for Beginners",
        "description": "Learn Go programming for beginners.",
        "instructor": "John Doe",
        "centPrice": 2000,
    }
]
```
Errors:
```
Code: 400
ContentType: 'text/plain'
Body: 'Please provide the start and end dates in RFC3339'
```
```
Code: 500
ContentType: 'text/plain'
Body: 'Could not load course list'
```

### 1.2 Anlegen/Verändern einer neuen Schulung:
```
POST /api/courses
ContentType: 'application/json'
Body:
{
    "id": 1,
    "name": "Go for Beginners",
    "description": "Learn Go programming for beginners.",
    "instructor": "John Doe",
    "price": 2000
}
```
Success:
```
Code: 201
ContentType: 'text/plain'
Body: '1' // Updated ID
```
Errors:
```
Code: 400
ContentType: 'text/plain'
Body: 'Could not parse JSON body'
```
```
Code: 400
ContentType: 'text/plain'
Body: 'Invalid req, course, name is required'
```
```
Code: 400
ContentType: 'text/plain'
Body: 'Invalid req, course, description is required'
```
```
Code: 400
ContentType: 'text/plain'
Body: 'Invalid req, course, instructor is required'
```
```
Code: 400
ContentType: 'text/plain'
Body: 'Invalid req, course, price is required'
```
```
Code: 500
ContentType: 'text/plain'
Body: 'Could not save course'
```

### 1.3 Anzeige der Termine zu einer bestimmten Schulung
```
GET /api/courses/sessions?id=[int]&startDate=[RFC3339]&endDate=[RFC3339]
```
Success:
```
Code: 200
ContentType: 'application/json'
Body: 
[
    {
        courseId: 1,
        time: 1,
        date: '2013-02-03 19:54:00 +0000 CET'
    },
    {
        courseId: 1,
        time: 2,
        date: '2013-03-03 19:54:00 +0000 CET'
    },
]
```
Errors:
```
Code: 400
ContentType: 'text/plain'
Body: 'Please provide the start and end dates in RFC3339'
```
```
Code: 400
ContentType: 'text/plain'
Body: 'Please provide a course ID'
```
```
Code: 500
ContentType: 'text/plain'
Body: 'Could not load course sessions'
```

### 1.4 Anlegen/Verändern eines neuen Termins für eine Schulung
```
POST /api/courses/sessions
ContentType: 'application/json'
Body:
{
    courseId: 1,
    time: 1,
    date: '2013-03-03 19:54:00 +0000 CET'
}
```
Success:
```
Code: 201
```
Errors
```
Code: 400
ContentType: 'text/plain'
Body: 'Could not parse JSON body'
```
```
Code: 400
ContentType: 'text/plain'
Body: 'Invalid req, no course ID provided'
```
```
Code: 400
ContentType: 'text/plain'
Body: 'Invalid req, invalid date privided'
```
```
Code: 500
ContentType: 'text/plain'
Body: 'Could not save course session'
```
### 1.5 Buchung einer Schulung an einem bestimmten Termin
```
POST /api/corses/bookings
ContentType: 'application/json'
Body:
{
    courseId: 1,
    time: 1,
    participantId: 20
}
```
Success:
```
Code: 201
```
Errors:
```
Code: 400
ContentType: 'text/plain'
Body: 'Could not parse JSON body'
```
```
Code: 400
ContentType: 'text/plain'
Body: 'Invalid req, no course ID provided'
```
```
Code: 400
ContentType: 'text/plain'
Body: 'Invalid req, no time provided'
```
```
Code: 400
ContentType: 'text/plain'
Body: 'Invalid req, no participent ID provided'
```
```
Code: 409
ContentType: 'text/plain'
Body: 'Course session is booked up'
```
```
Code: 500
ContentType: 'text/plain'
Body: 'Could not book course session'
```
---
## 2. Umsetzung Server-Seite
Programmiere die Server-Seite der API in einer Programmiersprache Deiner Wahl. Baue hierbei Tests für die API ein. Setze mindestens einen
lesenden und einen modifizierenden REST-Call um. Treffe eine sinnvolle Auswahl der oben genannten Funktionalitäten, um das Prinzip der
Umsetzung zu zeigen. Es ist nicht notwendig, alle oben genannte Funktionalität umzusetzen. Die Daten müssen nicht über das Beenden der
Anwendung hinaus persistiert werden.

## 3. Umsetzung Client-Seite
Erstelle in einer Frontend-Technologie Deiner Wahl einen Client, mit dem exemplarisch ein bis zwei REST-Calls abgesetzt werden können. Bei
dieser Aufgabe geht es nicht um eine grafisch schicke UI, sondern um einen PoC des Zusammenspiels zwischen Client und Server.

## 4. Security
Was würdest Du in dieser Anwendung zum Thema Security machen? Was muss warum abgesichert werden und welche Konzepte und
Technologien würdest Du dafür einsetzen?
Diese Aufgabe ist wieder reine Konzeption. Eine Umsetzung ist nicht erforderlich.

1. Es wird eine Session-Management mit Login und Passwort benötigt um die Benutzer zu identifizieren. Hierfür würde ich github.com/gorilla/sessions verwenden.

Beispiel für einen Session ckeck:
```
func (s *server) checkSession(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.store.Get(r, sessionName)
		if err != nil {
			http.Error(w, "Session now invalid", http.StatusForbidden)
			return
		}
        ...
	})
}
```
2. Nur die Kursleitung darf die Kurse bearbeiten oder anlegen. Da sonst die Teilnehmer beispielsweise den Preis ändern könnten.
3. Ein Teilnehmer kann nur sich selbst für einen Kurs ein- oder austragen.

Für Punkt 2 und 3 würde ich die Eigenschaften des Benutzers (Kursleitung/Teilnehmer) auf dem Server mit einen Middleware wie in 1 überprüfen und ggf. den Request abbrechen.

Hierfür sollte es jeweils einen eigenen Middleware geben da nur bestimmte Teile der API für Teilnehmer gesperrt sein sollen.

Um Beispielsweise den Benutzer bei zwei Middlewares nicht zweimal laden zu müssen kann der Benutzer durch den ersten Interseptor an den nächsten weitergereicht werden. Hier kann z. B. der http.Request Context verwendet werden.

