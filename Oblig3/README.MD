# IS-105, Obligatorisk oppgave 3

Adrian Nesvik, Edvin Myhren, Endre Grindhaug og Tord Vetle Gjertsen


### Oppgave 1

----------------------------------------------------------------------------------------------------------

#### a) **Sett opp en webserver som lytter til port 8080.**
Programmet ligger i 'IS-105_oppgaver\Oblig3\src\oppgave1' med navnet http1.go.

Programmet bruker listenAndServe til å lytte til porten 8080.

#### b) **Når klienter aksesserer webserveren på path lik "/" skal klienten få tilbake strengen "Hello, client". Strengen skal vises i nettleseren.**
Programmet er en metode med navnet func hello. Denne metoden sender setningen 'Hello client'.
Etter det bruker vi http.HandleFunc og setter pathen "/" til å bruke func hello.


![in cmd](https://github.com/StavenX/IS-105_oppgaver/blob/master/Oblig3/images/http1.png)

### Oppgave 2

----------------------------------------------------------------------------------------------------------


#### a)

### Oppgave 3

----------------------------------------------------------------------------------------------------------

#### a)
