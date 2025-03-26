# E-Biznes

## Zadanie 1. (Docker):

- lokalizacja: dockerExercise/
- obraz: [matid90/ebiznesdocker:latest](https://hub.docker.com/r/matid90/ebiznesdocker)

### wykonane zadania:

- użyto ubuntu 24.04
- zainstalowano python 3.10
- zainstalowano jave 8, kotlina i gradle przez sdkman
- utworzono prosty program helloworld w java z gradle
- dodano paczke jdbc do dependencji w gradle
- buildowanie i uruchamianie aplikacji przez "CMD gradle run"
- dodano prostą konfiguracje docker compose

---

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
