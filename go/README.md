**ZADANIE 4.** go + echo

:white_check_mark: 3.0 Należy stworzyć aplikację we frameworki echo w j. Go, która będzie miała kontroler Produktów zgodny z CRUD  
:white_check_mark: 3.5 Należy stworzyć model Produktów wykorzystując gorm oraz wykorzystać model do obsługi produktów (CRUD) w kontrolerze (zamiast listy)   
:white_check_mark: 4.0 Należy dodać model Koszyka oraz dodać odpowiedni endpoint  
:white_check_mark: 4.5 Należy stworzyć model kategorii i dodać relację między kategorią, a produktem  
:white_check_mark: 5.0 pogrupować zapytania w gorm’owe scope'y  

### Uruchamianie

linux:  
`./runDocker.sh` - jeżeli obraz dockerowy nie istnieje builduje go, a następnie uruchamia. Entrypointem obrazu jest skrypt uruchamiający serwer go. Serwer słucha na porcie 8080.
windows:  
1. `docker buildx build -t <własna_nazwa_obrazu> .` - buduje obraz
2. `docker run -it -p 8080:8080 --rm -v $(pwd)/app:/app <własna_nazwa_obrazu>` - uruchamia obraz docker wraz z serwerem go

### Testowanie
Serwer na [http://localhost:8080](http://localhost:8080)  publikuje widok który umożliwia testowanie zaimplementowanych endpointów