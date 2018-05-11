## is-105 oblig 4
# Systemspesifikasjon og systemarkitektur

Adrian Nesvik, Edvin Myhren, Endre Grindhaug og Tord Vetle Gjertsen

## Oppgavebeskrivelse

 Beskrive samfunnets og/eller brukeres nytteverdi av applikasjonen. I tillegg skal dere utrede, velge og dokumentere systemets arkitektur.  Systemarkitekturen viser til alle nodene, som er involvert i systemet og beskriver deres funksjon.  


## Systemspesifikasjon

Formålet med applikasjonen er å gi brukeren en enkel og intuitiv måte å finne informasjon om nylige jordskjelv. Jordskjelv forutsetning er en egen gren innenfor seismologi. Informasjon gitt av applikasjonen som kraft og lokasjon er informasjon som blir brukt for å forutsi jordskjelv. Selv om applikasjonen ikke lagrer informasjonen på lang sikt, kan applikasjonen ha nytteverdi innenfor ikke bare jordskjelv forutsetning, men hele fagfeltet seismologi. Dette kan videre gi en nytteverdi til samfunnet gjennom for eksempel forvarsler av jordskjelv. Applikasjonen kan også være nyttig for brukere som bare har en generell interesse i jordskjelv.


## Systemarkitektur

Applikasjonen har en "single-page" web applikasjon arkitektur. Den lagrer eller oppdaterer ikke data automatisk, men henter alt gjennom en eneste innlasting. I dette systemet blir datamaskinen som kjører programmet både server og klient. Server i den forstand at det åpnes en lokal web server som lytter til port ":8080" når applikasjonen blir kjørt. Datamaskinen er klient når den sender en forespørsel om websiden som inneholder informasjon prosessert av web serveren. Applikasjonen henter også JSON basert data gjennom et API kall til en rekke av earthquake.usgs.gov sine API'er. Dette kallet sender en forespørsel til en server i datasenteret Akamai Technologies og får returnert data.

![in cmd](https://github.com/StavenX/IS-105_oppgaver/blob/master/Oblig4/images/web.png)
