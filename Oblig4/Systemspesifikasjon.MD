Systemspesifikasjon og systemarkitektur

Oppgavebeskrivelse: Beskrive samfunnets og/eller brukeres nytteverdi av applikasjonen. I tillegg skal dere utrede, velge og dokumentere systemets arkitektur.  Systemarkitekturen viser til alle nodene, som er involvert i systemet og beskriver deres funksjon.  

Systemarkitektur
Nodene i denne systemarkitekturen består av klient, database server og web server. Applikasjonen åpner web server på port ":8080" og presenterer dataen der basert på hvilken sti som blir aksessert. Klienten sender også en forespørsel til Akamai Technologies sine database servere, brukt av earthquake.usgs.gov, for å hente data om jordskjelv. Denne dataen blir brukt i applikasjonen.

Ignore:
http://www.fdsn.org/webservices/FDSN-WS-Specifications-1.0.pdf
https://en.wikipedia.org/wiki/Client_(computing)
https://en.wikipedia.org/wiki/Client%E2%80%93server_model
Datacenter: Akamai Technologies
Server IP: 172.229.233.152
Location: United States of America
City: Cambridge

