Systemspesifikasjon og systemarkitektur

Oppgavebeskrivelse: Beskrive samfunnets og/eller brukeres nytteverdi av applikasjonen. I tillegg skal dere utrede, velge og dokumentere systemets arkitektur.  Systemarkitekturen viser til alle nodene, som er involvert i systemet og beskriver deres funksjon.  

Systemspesifikasjon
Formålet med applikasjonen er å gi brukeren en enkel og intuitiv måte å finne informasjon om nylige jordskjelv. Jordskjelv forutsetning er en egen gren innenfor seismologi. Informasjon gitt av applikasjonen som kraft og lokasjon er informasjon som blir brukt for å forutsi jordskjelv. Selv om applikasjonen ikke lagrer informasjonen på lang sikt, kan applikasjonen ha nytteverdi innenfor ikke bare jordskjelv forutsetning, men hele fagfeltet seismologi. Dette kan videre gi en nytteverdi til samfunnet gjennom for eksempel forvarsler av jordskjelv. Applikasjonen kan også være nyttig for brukere som bare har en generell interesse i jordskjelv.


Systemarkitektur
Applikasjonen har en "single-page" web applikasjon arkitektur. Den lagrer eller oppdaterer ikke data automatisk, men henter alt gjennom en eneste innlasting. Det åpnes en lokal web server som lytter til port ":8080". Avhengig av hvilken sti som blir aksessert vil klienten hente JSON-basert data. JSON dataen blir hentet gjennom et API kall til en rekke av earthquake.usgs.gov sine API'er. Aplikasjonen oversetter deretter dataen til et leselig format og presenterer den i web serveren. 

Ignore:
http://www.fdsn.org/webservices/FDSN-WS-Specifications-1.0.pdf
https://en.wikipedia.org/wiki/Client_(computing)
https://en.wikipedia.org/wiki/Client%E2%80%93server_model
https://www.scnsoft.com/blog/web-application-architecture
https://no.wikipedia.org/wiki/Webtjeneste
https://stackify.com/web-application-architecture/
Datacenter: Akamai Technologies
Server IP: 172.229.233.152
Location: United States of America
City: Cambridge

