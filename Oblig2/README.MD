# IS-105, Obligatorisk oppgave 2

Adrian Nesvik, Edvin Myhren, Endre Grindhaug og Tord Vetle Gjertsen


### Oppgave 1

----------------------------------------------------------------------------------------------------------

Programmene ligger i 'IS-105_oppgaver/Oblig2/src/oppgave1', under navnene fileinfo.go og fileinfo.exe.

a) Skriv et Golang program med navn fileinfo.go​.
Programmet returnerer informasjon om en fil. Denne informasjonen skaffes gjennom bruk av flere funksjoner fra 'os' pakken. Programmet kan også returnere feilmeldinger gjennom if/else statements.

b) Lag en kjørbart program (build) av fileinfo.go som kan utføres fra kommandolinje. Programmet skal ta et argument, som skal være et filnavn.
Filen 'runFileinfoGo.bat' tar et argument og kjører programmet 'fileinfo.go'. På linje 2 pauser vi programmet slik at terminalen ikke lukker seg etter programmet er kjørt.

### Oppgave 2

----------------------------------------------------------------------------------------------------------

Programmet ligger i 'IS-105_oppgaver/Oblig2/src/oppgave2/', og heter fielcount.go.

a) Skriv programmer med navn filecount.go som skal lese en tekst-fil, skrive ut totalt antall linjer og de fem “runes” (bokstaver/tegn), som forekommer mest i filen samt hvor mange ganger de telles i den medfølgende filen text.txt. Som forrige oppgave skal filen bygges til et kjørbart program og kunne utføres i terminal med kommandoen: filecount  <filnavn>


### Oppgave 3

----------------------------------------------------------------------------------------------------------

Programmene ligger i 'IS-105_oppgaver/Oblig2/src/oppgave3/', under navnene addtofile.go, addup.go og sumfromfile.go.

a) Skriv et program med navn addup.go bestående av to go funksjoner og kommunisere med hverandre gjennom channels


Funksjon A:


Funksjon B:
 

b) Skriv to programmer (addtofile.go og sumfromfile.go) som kommuniserer med hverandre gjennom en fil.


c) Beskriv og implementer feilhåndtering på all I/O i oppgave a) og b). Se https://blog.golang.org/error-handling-and-go (Lenker til en ekstern side.)Lenker til en ekstern side. for mer informasjon om feilhåndtering i Go.


d) Implementer håndtering av SIGINT i programmene i a) og b); programmene skal skrive ut en avslutningsmelding dersom de mottar SIGINT før de fullfører naturlig.


e) Bygg filene som i oppgave 1 og 2 til kjørbare filer på ditt operativsystem og legg dem ved i /bin mappen.

----------------------------------------------------------------------------------------------------------
