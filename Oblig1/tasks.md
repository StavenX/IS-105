<p><span style="font-size: 18pt;"><strong>Intro:</strong></span></p>
<p>Obligatorisk oppgave er en gruppeoppgave som er en del av mappeinnleveringen og teller mot sluttkarakter.</p>
<p>Oppgavene skal l&oslash;ses og lastes opp p&aring; Github.<br />Oppgavene skal besvares i en .md (markdown) fil.</p>

<h1><strong>1. Fyll ut manglende tall i tabell</strong></h1>
<table style="width: 687px;" border="1">
<tbody>
<tr>
<td style="width: 265px;">
<p>Bin&aelig;re tall (mest signifikant bit til venstre</p>
</td>
<td style="width: 204px;">
<p>Hexadesimaltall</p>
</td>
<td style="width: 195px;">
<p>Desimaltall</p>
</td>
</tr>
<tr>
<td style="width: 265px; text-align: center;">
<p>1101</p>
</td>
<td style="width: 204px; text-align: center;">
<p>0xD</p>
</td>
<td style="width: 195px; text-align: center;">
<p>13</p>
</td>
</tr>
<tr>
<td style="width: 265px; text-align: center;">
<p>110111101010</p>
</td>
<td style="width: 204px; text-align: center;">
<p>&nbsp;</p>
</td>
<td style="width: 195px; text-align: center;">
<p>&nbsp;</p>
</td>
</tr>
<tr>
<td style="width: 265px; text-align: center;">
<p>&nbsp;</p>
</td>
<td style="width: 204px; text-align: center;">
<p>0xAF34</p>
</td>
<td style="width: 195px; text-align: center;">
<p>&nbsp;</p>
</td>
</tr>
<tr>
<td style="width: 265px; text-align: center;">
<p>&nbsp;</p>
</td>
<td style="width: 204px; text-align: center;">
<p>&nbsp;</p>
</td>
<td style="width: 195px; text-align: center;">
<p>65535</p>
</td>
</tr>
<tr>
<td style="width: 265px; text-align: center;">
<p>&nbsp;</p>
</td>
<td style="width: 204px; text-align: center;">
<p>&nbsp;</p>
</td>
<td style="width: 195px; text-align: center;">
<p>71562</p>
</td>
</tr>
</tbody>
</table>
<h2><span style="font-size: 18pt;"></span></h2>
<h2><span style="font-size: 18pt;"><strong>Oppgave A</strong></span></h2>
<p>Beskriv kort metode for &aring; g&aring; fra bin&aelig;re tall til hexadesimale tall og motsatt. Beskriv kort metoden for &aring; g&aring; fra bin&aelig;re tall til desimaltall og motsatt.</p>
<h2><span style="font-size: 18pt;"><strong>Oppgave B</strong></span></h2>
<p>Beskriv kort metoden for &aring; g&aring; fra hexadesimale tall til desimaltall og motsatt.</p>
<p>&nbsp;</p>
<h1><strong>2. Forst&aring; algoritmer og utf&oslash;re &ldquo;benchmark&rdquo;-tester p&aring; koden</strong></h1>
<p>Programmeringsoppgave: - ta utgangspunkt i pakken algorithms i oblig1</p>
<h2><span style="font-size: 18pt;"><strong>Oppgave A</strong></span></h2>
<p>Skriv en modifisert bubble-sort funksjon benchmarkBSortModified basert p&aring; eksempel-funksjon Bubble_sort i filen sorting.go (se for tips)</p>
<h2><span style="font-size: 18pt;"><strong>Oppgave&nbsp;B</strong></span></h2>
<p>Skriv "benchmark"-tester for benchmarkBSortModified funksjonen basert p&aring; eksempel-funksjon benchmarkBSort i filen go</p>
<h2><span style="font-size: 18pt;"><strong>Oppgave C</strong></span></h2>
<p>Det finnes ogs&aring; en implementasjon av Quicksort algoritme i sorting.go og tilsvarende implementasjon av tester i go; utf&oslash;r alle benchmark- testene med kommando &ldquo;go test -bench=.&rdquo; og presenter resultatene grafisk</p>
<p>Hva kan du si om big-O for alle 3 algoritmene, som du har testet?</p>
<p>&nbsp;</p>
<p><em>Ressurser:</em></p>
<p><em>Bubble sort: <a href="https://en.wikipedia.org/wiki/Bubble_sort">https://en.wikipedia.org/wiki/Bubble_sort</a></em></p>
<p><em>Big-O <a href="http://bigocheatsheet.com">http://bigocheatsheet.com</a> &nbsp;</em></p>
<p>&nbsp;</p>
<h1><strong>3. Forst&aring; prosessadministrajon p&aring; et platform</strong></h1>
<p>Skriv et program som best&aring;r av en evig l&oslash;kke. Hvor mye minne og CPU bruker programmet n&aring;r det kj&oslash;rer. Programmet skal skrive ut en avslutningsmeld- ing n&aring;r programmet mottar et SIGINT signal. Generer ulike avslutningssig- naler til prosessen og dokumenter hvilke avslutningskommandoer programmet h&aring;ndterer og som trigger avslutningsmeldingen.</p>
<p>Ressurs: <br /> <span><a href="https://en.wikipedia.org/wiki/Signal_(IPC)">https://en.wikipedia.org/wiki/Signal_(IPC)</a></span></p>
<p>&nbsp;</p>
<p>&nbsp;</p>
<h1><strong>4. Typografiske symboler</strong></h1>
<p>Form&aring;l:</p>
<ul>
<li>Bli kjent med ISO/IEC 8859 serier for 8-bits koding av typografiske sym- boler.</li>
<li>Illustrere forskjell p&aring; ASCII og utvidet ASCII kode gjennom golang ram- meverk for behandling av tekststrenger (p&aring; engelsk brukes det nesten alltid begrepet &ldquo;strings&rdquo; for tekststrenger).</li>
</ul>
<p>Hva trenger du &aring; vite for &aring; klare denne oppgaven:</p>
<ul>
<li>du m&aring; forst&aring; &ldquo;slices&rdquo; i golang</li>
<li>at de f&oslash;rste 256 &ldquo;code points&rdquo; (kodepunkter, som tilsvarer typen &ldquo;rune&rdquo; i golang, varierer fra platform til platform og fra program til program; det finnes 15 forskjellige deler av ISO/IEC 8859 serier for 8-bits koder)</li>
</ul>
<p>&nbsp;</p>
<h2><span style="font-size: 18pt;"><strong>Oppgave A</strong></span></h2>
<p>Bruk filen ascii.go i Oblig1 mappen og lag en funksjon som itererer (g&aring;r i en l&oslash;kke over)&nbsp; over tegn med byte-verdier fra 0x80 til 0xFF, dvs. det utvidede ASCII settet.</p>
<h3>Kravspesifikasjon</h3>
<p>Funksjonsnavn skal v&aelig;re iterateOverExtendedASCIIStringLiteral(...) og den skal ta et argument,&nbsp; som skal v&aelig;re av&nbsp; type string (&ldquo;string literal&rdquo;).&nbsp;&nbsp; Dere&nbsp;&nbsp; m&aring; generere / deklarere en string med alle de 128 heksadesimale verdiene ( &ldquo;\x80\x81. . . \xFF&rdquo; ) som funksjonen kan ta som argument. Funksjonen trenger ikke &aring; returnere noe eksplisitt.</p>
<p>Utskriftsformatet skal v&aelig;re f&oslash;lgende:</p>
<p>[utvidet-ascii-kode heksadesimalt med store bokstaver A-F][mellomrom][symbol for utvidet-ascii-kode] [mellomrom][utvidet-ascii-kode bin&aelig;rt][linjeskift]</p>
<p>Eksempel :</p>
<table style="width: 299px;">
<tbody>
<tr>
<td style="width: 115px; text-align: center;">
<p>3E</p>
</td>
<td style="width: 115px; text-align: center;">
<p>&nbsp;&gt;</p>
</td>
<td style="width: 170px; text-align: center;">
<p>111110</p>
</td>
</tr>
<tr>
<td style="width: 115px; text-align: center;">
<p>BA</p>
</td>
<td style="width: 115px; text-align: center;">
<p>&nbsp;&ordm;</p>
</td>
<td style="width: 170px; text-align: center;">
<p>10111010</p>
</td>
</tr>
<tr>
<td style="width: 115px; text-align: center;">
<p>BB</p>
</td>
<td style="width: 115px; text-align: center;">
<p>&nbsp;&raquo;</p>
</td>
<td style="width: 170px; text-align: center;">
<p>10111110</p>
</td>
</tr>
<tr>
<td style="width: 115px; text-align: center;">
<p>BF</p>
</td>
<td style="width: 115px; text-align: center;">
<p>&iquest;</p>
</td>
<td style="width: 170px; text-align: center;">
<p>10111111</p>
</td>
</tr>
<tr>
<td style="width: 115px; text-align: center;">
<p>C0</p>
</td>
<td style="width: 115px; text-align: center;">
<p>&Agrave;</p>
</td>
<td style="width: 170px; text-align: center;">
<p>11000000</p>
</td>
</tr>
<tr>
<td style="width: 115px; text-align: center;">
<p>C1</p>
</td>
<td style="width: 115px; text-align: center;">
<p>&Aacute;</p>
</td>
<td style="width: 170px; text-align: center;">
<p>11000001</p>
</td>
</tr>
<tr>
<td style="width: 115px;">
<p>&nbsp;</p>
</td>
<td style="width: 115px;">
<p>. . .</p>
</td>
<td style="width: 170px;">
<p>&nbsp;</p>
</td>
</tr>
</tbody>
</table>
<p>&nbsp;</p>
<ul>
<li>analyser utskriften (spesielt for bytes fra 0x80 til 0x9F)</li>
<li>utf&oslash;r programmet p&aring; alle gruppemedlemmers datamaskin og analyser</li>
</ul>
<p>&nbsp;</p>
<h2><span style="font-size: 18pt;"><strong>Oppgave B</strong></span></h2>
<p>Lag en funksjonen ExtendedASCIIText () i samme filen iso.go, som skriver ut: " &euro; &divide; &frac34; dollar "</p>
<h3>Kravspesifikasjon</h3>
<ul>
<li>Funksjonen skal generere en utskrift fra en sekvens av bytes, dvs. av typen []bytes (det betyr at du m&aring; finne den heksadesimale eller bin&aelig;re representasjonen av alle tegn i strengen, som skal skrives ut (inkludert anf&oslash;rselstegn eller &ldquo;double quotes&rdquo; p&aring; engelsk).</li>
<li>Funksjonen ExtendedASCIIText () skal returnere en variabel av typen string, som inneholder tegn fra Extended</li>
<li>Utf&oslash;r programmet p&aring; forskjellige platformer(mac,windows,linux) eller forskjellig software(terminal, bash, powershell) og analyser resultater</li>
</ul>
<p>&nbsp;</p>
<h2><span style="font-size: 18pt;"><strong>Oppgave C</strong></span></h2>
<p>Implementer en test for funksjonen ExtendedASCIIText(String) i egen fil iso_test.go, som tester om input-verdier (av type string) inneholder kun tegn fra en Extended ASCII.</p>
<p>&nbsp;</p>
<p><em>Ressurser:</em></p>
<p><em>Runes: <a href="https://en.wikipedia.org/wiki/Runes">https://en.wikipedia.org/wiki/Runes</a></em></p>
<p><em>ASCII Table: <a href="http://www.asciitable.com/">http://www.asciitable.com/</a></em></p>
<p><em>Golang Stings, Bytes, Runes og Characters <a href="https://blog.golang.org/strings">https://blog.golang.org/strings</a></em></p>
<p><em>Golang FMT package: <a href="https://golang.org/pkg/fmt/">https://golang.org/pkg/fmt/</a></em></p>
<p><em>Golang Bytes package: <a href="https://golang.org/pkg/bytes/">https://golang.org/pkg/bytes/</a></em></p>
<p><em>Golang Slices: <a href="https://blog.golang.org/slices">https://blog.golang.org/slices</a></em></p>