# MalGO | Obligatorisk oppgave 1
IS-105 @V2018

## Oppgave 1: Fyll ut manglende tall i tabell

Binary|Hexadeximal|Decimal
---|---|---
1101|0xD|13
110111101010|0xDEA|3562
1010111100110100|0xAF34|44852
1111111111111111|0xFFFF|65535
00010001011110001010|0x1178A|71562

### - oppg A: Beskriv kort metode for å gå fra binære tall til hexadesimale tall og motsatt. Beskriv kort metoden for å gå fra binære tall til desimaltall og motsatt.

----------------------------------------------------------------------------------------------------------
	BINÆR TIL HEX
	Først få binærverdien til å gå opp i 4 (sette på ekstra nuller foran, f.eks 111010 = 0011-1010). Etter dette kan man splitte binærtallet opp i 4 (f.eks 11110010 blir 1111-0010). 
	Her kan vi regne ut de seperate segmentene.
	1111 = 15
	0010 = 2
	Deretter gjør vi disse verdiene om til HEX, og plasserer de etter hverandre.
	15 = F
	2 = 2
	Resultat blir F2 (0xF2).
----------------------------------------------------------------------------------------------------------
	HEX TIL BINÆR
	Man tar hvert individuelt tall i en hex-verdi (f.eks 0xDEAA). Deretter gjør vi hvert siffer om til binært med fire siffre.

	D = 1101
	E = 1110
	A = 1010
	A = 1010

	Resultatet blir binærsiffrene satt opp etter hverandre: D(1101),E(1110),A(1010),A(1010)
	Eller 1101111010101010.
----------------------------------------------------------------------------------------------------------
	BINÆR TIL DESIMAL
	Måten man går fra binære tall til desimaltall, er å gjøre en mattefunksjon på hvert siffer i binærtallet. 
	Først, ser man hvilken plass sifferet står på. 
	Eksempeltall: 110010110
	Hvert av disse tallene skal ganges med 2 (siden vi bruker binærsystem / totallsystem). Deretter skal vi opphøye dette i plassen sifferet står på.
	Siden binærtallet består av 9 siffre, tar vi 8 som et utgangspunkt for lengde (siden datamaskiner teller fra 0, ikke 1). Det første tallet får plass 8, det andre plass 7, osv. 
	Vi plusser deretter funksjonene sammen, og får svaret som desimaltall.
	1 * 2^8
	+ 1 * 2^7
	+ 0 * 2^6
	+ 0 * 2^5
	+ 1 * 2^4
	+ 0 * 2^3
	+ 1 * 2^2
	+ 1 * 2^1
	+ 1 * 2^0
	= 406
----------------------------------------------------------------------------------------------------------
	DESIMAL TIL BINÆR
	Måten man går fra desimal til binær er å trekke fra nærmeste, og laveste verdien i totallsystemet. Deretter krysser man '1' hvis dette går,
	og går til verdien man har igjen etter første regnestykke. Deretter gjør man samme prosess (krysser av 0 hvis det ikke går opp), helt til man
	får en rekke av '1'er og '0'er.
----------------------------------------------------------------------------------------------------------


### - oppg B: 

	Måten man går fra hexadesimale tall til desimal
----------------------------------------------------------------------------------------------------------
	DESIMAL TIL HEX
	Man tar desimaltallet og deler det på 16, så setter man det foran kommaet på desimaltallet man får til siden. 
	Så ganger man det bak kommaet med 16. Da får du det første hexadesimaltallet (fra høyre til venstre).
	Deretter tar det man satt til siden tidligere, og deler det på 16. 
	Man gjør denne prosessen helt til det foran kommaet er '0'. Ta da det siste sifferet bak komma, og gang det med 16. 
	Dette er det siste hexadesimalet. 
----------------------------------------------------------------------------------------------------------
	HEX TIL DESIMAL
	Først tar man hex-verdien (f.eks 0xDEA) og gjør desimaltall. Deretter ganger man tallet med 16 (pga hex er sekstentallsystem),
	for å så opphøye det i hvilken plass sifferet står på. Etter man har gjort denne prosessen individuelt på hvert siffer, plusser 
	man sammen tallene man står igjen med.

	0xDEA = 13 * (16^2), + 14 * (16^1), + 10 * (16^0).
	= 3328 + 224 + 10 = 3562.
----------------------------------------------------------------------------------------------------------

## Oppgave 2: Forstå algoritmer og utføre "benchmark"-tester på koden

### - oppg A:
	

Titallsystem
123456 //10

1 * 10^5
2 * 10^4
3 * 10^3
4 * 10^2
5 * 10^1
6 * 10^0

Binær til desimaltall 

For å 
110010110 //2

1 * 2^8
1 * 2^7
0 * 2^6
0 * 2^5
1 * 2^4
0 * 2^3
1 * 2^2
1 * 2^1
1 * 2^0
