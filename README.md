# E-Biznes

**Zadanie 1** Docker

:white_check_mark: 3.5 obraz ubuntu:24.02 z Javą w wersji 8 oraz Kotlinem  
:white_check_mark: 4.0 do powyższego należy dodać najnowszego Gradle’a oraz paczkę JDBC  
:white_check_mark: SQLite w ramach projektu na Gradle (build.gradle)  
:white_check_mark: 4.5 stworzyć przykład typu HelloWorld oraz uruchomienie aplikacji przez CMD oraz gradle  
:white_check_mark: 5.0 dodać konfigurację docker-compose  

obraz: [matid90/ebiznesdocker:latest](https://hub.docker.com/r/matid90/ebiznesdocker)  
Kod: [/dockerExerise](https://github.com/Ech0n/ebiznes/tree/main/dockerExercise)

**Zadanie 2** Scala & Play

:white_check_mark: 3.0 Należy stworzyć kontroler do Produktów  
:white_check_mark: 3.5 Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane pobierane z listy  
:white_check_mark: 4.0 Należy stworzyć kontrolery do Kategorii oraz Koszyka + endpointy zgodnie z CRUD  
:white_check_mark: 4.5 Należy aplikację uruchomić na dockerze (stworzyć obraz) oraz dodać skrypt uruchamiający aplikację via ngrok  
:white_check_mark: 5.0 Należy dodać konfigurację CORS dla dwóch hostów dla metod CRUD   

Kod: [/zadanie2](https://github.com/Ech0n/ebiznes/tree/main/zadanie2)  
Demo Videos: [/zadanie2](https://github.com/Ech0n/ebiznes/tree/main/demos/scala)

Demo docker + ngrok + endpointy:

https://github.com/user-attachments/assets/d5c86b8d-a118-4d5f-92f5-a9606eef2bae

Demo cors:

https://github.com/user-attachments/assets/50c114f2-624b-4366-ad6d-547a8d10d35b

**ZADANIE 3.** Ktor Discord bot

:white_check_mark: 3.0 Należy stworzyć aplikację kliencką w Kotlinie we frameworku Ktor, która pozwala na przesyłanie wiadomości na platformę Discord  
:white_check_mark: 3.5 Aplikacja jest w stanie odbierać wiadomości użytkowników z platformy Discord skierowane do aplikacji (bota)  
:white_check_mark: 4.0 Zwróci listę kategorii na określone żądanie użytkownika  
:white_check_mark: 4.5 Zwróci listę produktów wg żądanej kategorii  
:x: 5.0 Aplikacja obsłuży dodatkowo jedną z platform: Slack, Messenger,  
Webex

Kod: [/kotlinKtorBot](https://github.com/Ech0n/ebiznes/tree/main/kotlinKtorBot)  

Demo:

https://github.com/user-attachments/assets/c3c2c751-bad9-4db6-8078-bc7a23c816c3


**ZADANIE 4.** go + echo

:white_check_mark: 3.0 Należy stworzyć aplikację we frameworki echo w j. Go, która będzie miała kontroler Produktów zgodny z CRUD  
:white_check_mark: 3.5 Należy stworzyć model Produktów wykorzystując gorm oraz wykorzystać model do obsługi produktów (CRUD) w kontrolerze (zamiast listy)   
:white_check_mark: 4.0 Należy dodać model Koszyka oraz dodać odpowiedni endpoint  
:white_check_mark: 4.5 Należy stworzyć model kategorii i dodać relację między kategorią, a produktem  
:white_check_mark: 5.0 pogrupować zapytania w gorm’owe scope'y  

Kod: [/go](https://github.com/Ech0n/ebiznes/tree/main/go)  

**ZADANIE 5.** Frontend

:white_check_mark: 3.0 W ramach projektu należy stworzyć dwa komponenty: Produkty oraz Płatności; Płatności powinny wysyłać do aplikacji serwerowej dane, a w Produktach powinniśmy pobierać dane o produktach z aplikacji serwerowej;  
:white_check_mark: 3.5 Należy dodać Koszyk wraz z widokiem; należy wykorzystać routing
:white_check_mark: 4.0 Dane pomiędzy wszystkimi komponentami powinny być przesyłane za pomocą React hooks  
:white_check_mark: 4.5 Należy dodać skrypt uruchamiający aplikację serwerową oraz kliencką na dockerze via docker-compose  
:white_check_mark: 5.0 Należy wykorzystać axios’a oraz dodać nagłówki pod CORS  


Kod: [/react](https://github.com/Ech0n/ebiznes/tree/main/react)  


**ZADANIE 6.** Testy

:white_check_mark: 3.0 Należy stworzyć 20 przypadków testowych w CypressJS lub Selenium (Kotlin, Python, Java, JS, Go, Scala)  
:white_check_mark: 3.5 Testy muszą zawierać minimum 50 asercji funkcjonalnych  
:white_check_mark: 4.0 Należy dodać testy jednostkowe do wybranego wcześniej projektu z minimum 50 asercjami  
:x: 4.5 Należy dodać testy API pokrywające wszystkie endpointy, każdy z co najmniej jednym scenariuszem negatywnym  
:x: 5.0 Testy funkcjonalne muszą być uruchamiane na platformie Browserstack (konto z Github Education Pack)  

Kod: [commit z testami](https://github.com/Ech0n/blef/commit/cb650fdf4469316a1daf5b5d1d9617a0c9f1a3b6)


**ZADANIE 7.** Sonar

:white_check_mark: 3.0 Należy dodać litera (linter) do kodu aplikacji serwerowej poprzez hooki gita  
:white_check_mark: 3.5 Należy usunąć wszystkie bugi wykryte przez SonarCloud w kodzie aplikacji serwerowej  
:white_check_mark: 4.0 Należy usunąć wszystkie zapaszki (code smells) w kodzie aplikacji serwerowej  
:white_check_mark: 4.5 Należy usunąć wszystkie podatności i błędy bezpieczeństwa w kodzie aplikacji serwerowej  
:white_check_mark: 5.0 Należy usunąć wszystkie błędy oraz zapaszki w kodzie aplikacji klienckiej i dodać widżety Sonara w pliku README  

Kod: [/EbiznesReactApp](https://github.com/Ech0n/EbiznesReactApp)
