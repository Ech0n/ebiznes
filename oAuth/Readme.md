**ZADANIE 8.** OAuth

:white_check_mark: 3.0 [Logowanie przez aplikację serwerową (bez Oauth2)](https://github.com/Ech0n/ebiznes/blob/main/oAuth/oAuth-server/controllers/auth.go)
:white_check_mark: 3.5 [Rejestracja przez aplikację serwerową (bez Oauth2)](https://github.com/Ech0n/ebiznes/blob/main/oAuth/oAuth-server/controllers/auth.go)
:white_check_mark: 4.0 [Logowanie via Google OAuth2](https://github.com/Ech0n/ebiznes/blob/main/oAuth/oAuth-server/controllers/google.go)
:white_check_mark: 4.5 [Logowanie via Facebook lub Github OAuth2](https://github.com/Ech0n/ebiznes/blob/main/oAuth/oAuth-server/controllers/github.go)
:white_check_mark: 5.0 [Zapisywanie danych logowania OAuth2 po stronie serwera](https://github.com/Ech0n/ebiznes/blob/main/oAuth/oAuth-server/db/db.go#L23)

Aby aplikacja działała poprawnie należy dodać do zmiennych środowiskowych lub pliku /oAuth/oAuth-server/.env zmienne: GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET, GITHUB_CLIENT_ID, GITHUB_CLIENT_SECRET

Kod: [/oAuth](https://github.com/Ech0n/ebiznes/tree/main/oauth)  

## Jak uruchomić:  

1. Przed uruchomieniem upewnij się, że wszystkie zmienne środowiskowe potrzebne dla GoogleOauth i GithubOauth są uzupełnione.
2. `docker compose up` - Uruchamia oba serwer frontendowy i backendowy. Po komunikacie, że serwer nasłuchuje z service'u client-1 i go-service-1 możemy otworzyć stronę pod adresem 'http://localhost:5173/'