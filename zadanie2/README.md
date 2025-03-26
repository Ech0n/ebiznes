# ZADANIE 2. Scala & Play framework

:white_check_mark: 3.0 Należy stworzyć kontroler do Produktów  
:white_check_mark: 3.5 Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane pobierane z listy  
:white_check_mark: 4.0 Należy stworzyć kontrolery do Kategorii oraz Koszyka + endpointy zgodnie z CRUD  
:white_check_mark: 4.5 Należy aplikację uruchomić na dockerze (stworzyć obraz) oraz dodać skrypt uruchamiający aplikację via ngrok  
:white_check_mark: 5.0 Należy dodać konfigurację CORS dla dwóch hostów dla metod CRUD   

### Wymagania:
- docker
- ngrok (+ skonfigurowany z authtoken)

### Uruchamianie:

`./run.sh`:
1. Sprawdza czy istnieje obraz Docker.
2. Jezeli obraz nie isntnieje tworzy go z pliku Dockerfile
3. W trakcie tworzenia obrazu kompiluje program
4. Uruchamia utworzony obraz
5. W kontenerze uruchamiana jest aplikacja w scala na porcie 9000
6. uruchamiany jest ngrok

`./corsTest.sh`: 
1. Wysyła zapytanie do serwera z domeny akceptowanej przez cors
2. Wysyła zapytanie do serwera z domeny nie akceptowanej przez cors

Demo docker + ngrok + endpointy:

https://github.com/user-attachments/assets/d5c86b8d-a118-4d5f-92f5-a9606eef2bae

Demo cors:

https://github.com/user-attachments/assets/50c114f2-624b-4366-ad6d-547a8d10d35b

