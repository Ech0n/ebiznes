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
