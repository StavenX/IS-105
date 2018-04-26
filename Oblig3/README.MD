# IS-105, Obligatorisk oppgave 3

Adrian Nesvik, Edvin Myhren, Endre Grindhaug og Tord Vetle Gjertsen


## Oppgave 1

----------------------------------------------------------------------------------------------------------

### a) Sett opp en webserver som lytter til port 8080.
Programmet ligger i 'IS-105_oppgaver\Oblig3\src\oppgave1' med navnet http1.go.

Programmet bruker listenAndServe til å lytte til porten 8080.

### b) Når klienter aksesserer webserveren på path lik "/" skal klienten få tilbake strengen "Hello, client". Strengen skal vises i nettleseren.
Programmet er en metode med navnet func hello. Denne metoden sender setningen 'Hello client'.
Etter det bruker vi http.HandleFunc og setter pathen "/" til å bruke func hello.


![in cmd](https://github.com/StavenX/IS-105_oppgaver/blob/master/Oblig3/images/http1.png)




## Oppgave 2

----------------------------------------------------------------------------------------------------------


### a) Presenter data på webserveren slik at den er leselig for mennesker(Ren tekst, f.eks. "Sted: Jorenhholem, dato: 13.04.2018", ikke i curly brackets.) når klienter aksesserer stiene /1, /2, /3, /4, /5.
Vi fant 5 forskjellige API'er med json filer og lagde klasser med structs til hver av dem.
Vi henter også filene fra nettet istedenfor å laste dem ned og "structer" dem individuelt i 'Oblig3\src\oppgave2\pages'. Deretter i server.go importer vi go klassene med json structs og lager print funksjoner til dem. Vi bruker også html styling til å gi server websiden et bedre utseende. Til slutt bruker vi handlefunc til å gi stiene sine designerte json structs. Man kan kjøre serveren med run.bat


### b) Samtlige stier i oppgave 2 skal rendres til klienter ved bruk av Go templates.
Alle stiene rendres til klienter med Go templates. De blir gjort klare med 'func loadTemplate' og senere ble de rendret med
'func useTemplate'




## Oppgave 3

----------------------------------------------------------------------------------------------------------

### a) Implementer et serverprogram i henhold til RFC 865. Serveren skal svare både på UDP og TCP.
