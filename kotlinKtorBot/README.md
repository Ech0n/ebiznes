# ZADANIE 3. Ktor Discord bot

:white_check_mark: 3.0 Należy stworzyć aplikację kliencką w Kotlinie we frameworku Ktor, która pozwala na przesyłanie wiadomości na platformę Discord  
:white_check_mark: 3.5 Aplikacja jest w stanie odbierać wiadomości użytkowników z platformy Discord skierowane do aplikacji (bota)  
:white_check_mark: 4.0 Zwróci listę kategorii na określone żądanie użytkownika  
:white_check_mark: 4.5 Zwróci listę produktów wg żądanej kategorii  
:x: 5.0 Aplikacja obsłuży dodatkowo jedną z platform: Slack, Messenger,  
Webex

### Działanie bota:
1. Bot wysyła wiadomość powitalną do użytkownika określonego w pliku .env
2. Bot za pomocą gateway api nasluchuje komend od uzytkownika  
    Obsługiwane komendy:
    - `list categories` wysyla uzytkownikowi dostepne kategorie produktow
    - `list <category>` wysyla uzytkownikowi produkty danej kategorii


### Uruchamianie:

1. Należy utworzyć plik .env i wypelnić go zmiennymi:
- `DISCORDUSERID` - id użytkownika do ktorego bot wysle wiadomosc powitalną
- `DISCORDBOTID` - id bota na discordzie
- `DISCORDBOTTOKEN` - token bota
2. Windows:
    `./gradlew.bat run`  
   Linux:
    `./gradlew run`
