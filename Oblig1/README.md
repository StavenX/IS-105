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
	Måten man går fra desimal til binær er å sammenligne tallet med det nærmeste totallsystem tallet som er lavere en eller det samme som desimaltallet (f.eks 1973 sitt nærmeste totallsystemtall er 1024 siden 2048 er høyere en 1973).
	Deretter tar man desmaltallet og trekker fra totallsystemtallet (1973-1024) og man setter et ett tall for å indikere når man har trukket fra.
	Så ser man på neste totallsystemtall som er lavere (512) og sjekker om det er høyere en vårt tall (som nå er 949).
	hvis desimaltallet er lavere enn totallsystemtallet setter man 0 og går til neste totallsystemtall helt til desimaltallet er større enn det neste totallsystemtallet. Slik fortsetter man til man kommer til det siste totallsystemtallet.
	Eksempel:
	1973 > 1024 -> 1973 - 1024 = 949 -> 1
	949 > 512 -> 949 - 512 = 437 -> 1
	437 > 256 -> 437 - 256 = 181 -> 1
	181 > 128 -> 181 - 128 = 53 -> 1
	53 < 64 -> 0
	53 > 32 -> 53 - 32 = 21 -> 1
	21 > 16 -> 21 - 16 = 5 -> 1
	5 < 8 -> 0
	5 > 4 -> 5 - 4 = 1 -> 1
	1 < 2 -> 0
	1 = 1 -> 1 - 1 = 0 -> 1

	1973 = 11110110101
----------------------------------------------------------------------------------------------------------


### - oppg B: Måten man går fra hexadesimale tall til desimal
----------------------------------------------------------------------------------------------------------
	DESIMAL TIL HEX
	Man tar desimaltallet og deler det på 16, så setter man det foran kommaet på desimaltallet man får til siden.
	Så ganger man det bak kommaet med 16. Da får du det første hexadesimaltallet (fra høyre til venstre).
	Deretter tar det man satt til siden tidligere, og deler det på 16.
	Man gjør denne prosessen helt til det foran kommaet er '0'. Ta da det siste sifferet bak komma, og gang det med 16.
	Dette er det siste hexadesimalet.
	Eksempel:2034
	2034 / 16 = 127.125
	0.125 * 16 = 2
	127 / 16 = 7.9375
	0.9375 * 16 = 15 = F
	7 / 16 = 0.4375
	0.4375 * 16 = 7
	2034 = 0x2F7
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
Modifisert Bubble-sort funksjon ligger i 'IS-105_obOppg_1/blob/master/Oblig1/src/algorithms/sorting.go', og heter Bubble_sort_modified.
### - oppg B:
Benchmark testene ligger i 'IS-105_obOppg_1/blob/master/Oblig1/src/algorithms/sortingtest.go', og bruker 'benchmarkBSortModified'. De andre benchmarktestene er 'BenchmarkMBSort100', 'BenchmarkMBSort1000' og 'BenchmarkMBSort100'.
### - oppg C:
Viser funksjon som er kjørt.
![cmd benchmark](https://raw.githubusercontent.com/StavenX/IS-105_obOppg_1/master/images/Benchmark-test.png "Benchmark")

## Oppgave 3: Forstå prosessadministrasjon på et platform
Selve loopen ligger i 'IS-105_obOppg_1/blob/master/Oblig1/src/loop/systemLoop.go', og heter main.
Vi fant en gjennomsnittlig 41% økning i CPU-bruk etter start av loopen.

Dette er Task-Manager statistikk før vi startet den uendelige loopen
![task manager statistikk](https://raw.githubusercontent.com/StavenX/IS-105_obOppg_1/master/images/status_idle.png "CPU før man starter loop")

Dette er Task-Manager statistikk etter vi startet den uendelige loopen
![task manager statistikk](https://raw.githubusercontent.com/StavenX/IS-105_obOppg_1/master/images/status_running.png "CPU etter man starter loop")

## Oppgave 4: Typografiske symboler
### - oppg A:
### - oppg B:
### - oppg C:
