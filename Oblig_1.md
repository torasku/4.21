### Oppgave 1

| Binære tall         | Hexa          | Desimal  |
|:-------------------:|:-------------:| --------:|
|       1101          |       D       |    13    |
| 1101 1110 1010      |      DEA      |   3562   |
| 1010 1111 0011 0100 |      AF34     |   44852  |
| 1111 1111 1111 1111 |      FFFF     |   65535  |
| 1000 1011 1100 01010|     1178A     |   71562  |

#### Oppgave 1A
##### Binær til hexa
Man velger ut 4 binære tall, opphøyer de i riktig tall (bakfra blir det 1, 2, 4 og 8). 
Eks: 1101 = 1^8 + 1^4 + 0*2 + 1^1 = 8 + 4 + 0 + 1 = 13
Tallet 13 tilsvarer hexa D. 

##### Hexa til binær
Her kan man benytte en tabell med hexa fra 0 - F med tilsvarende binære tall, og skrive om. 
Eks: DAB --> 1101 1010 1011

##### Binær til desimal
Man kan sette det opp på følgende måte: 
1101 = 1*2^3 + 1*2^2 + 0*2^1 + 1*2^0
     = 8 + 4 + 0 + 1
     = 13
     
##### Desimal til binær
Her benytter man heltallsdivisjon, og skriver ned restledd. Prossessen repeteres til resultat av heltallsdivisjon = 0. 
Eks: 
13/2 = 6 | 13 % 2 = 1
6/2 = 3  | 6 % 2 = 0
3/2 = 1  | 3 % 2 = 1
1/2 = 0  | 1 % 2 = 1

     Resultat = 1101 

#### Oppgave 1B
##### Hexa til desimal 
Først multipliseres hexa tallet med 16 opphøyd med 0 (starter med 0 fra bakerste og øker med 1). Resultatet legges sammen og man får desimaltallet. 
Eks: C9 = 9 * 16^0 + 12 * 16^1 = 192 + 9 = 201

##### Desimal til hexa
Her må man dele desimaltallet på 16 helt til man er kommet til 0. 
Eks: 3000
     3000 / 16 = 187 | rest = 8 
     187 / 16 = 11   | rest = B       
     11 / 16 = 0     | rest = B
     
     Resultat =  BB8
     
#### Oppgave 2C

|  Test                            |    Loop runs       |    Result      |
|:--------------------------------:|:------------------:|:--------------:|
|BenchmarkBSortModified100-4       |         50000      |   25986 ns/op  |
|BenchmarkBSortModified1000-4      |         1000       |  1806256 ns/op |
|BenchmarkBSortModified10000-4     |             5      | 251170600 ns/op|
|BenchmarkBSort100-4               |         50000      |   24887 ns/op  |
|BenchmarkBSort1000-4              |          1000      |  1445854 ns/op |
|BenchmarkBSort10000-4             |             5      | 225405700 ns/op|
|BenchmarkQSort100-4               |        300000      |   5245 ns/op   |
|BenchmarkQSort1000-4              |         20000      |   66977 ns/op  |
|BenchmarkQSort10000-4             |          2000      |   875689 ns/op |

Utifra Big-O notasjon for bubbleSort og quickSort, skal værste tilfellet være likt for begge algoritmene. Alikevel ser vi at i gjennomsnitt så vil quickSort være noe raskere med en Big-O på O(n log(n)). Hvis vi ser på resultatene fra benchmark testene ovenfor, vil vi få en bekreftelse på det.

### Oppgave 3
Når programmet kjører den evige løkken, brukes 7,2MB minne og 0 CPU. Fordi den evige løkken ikke gjør noen operasjoner, brukes det kun litt minne, men ingen prosesser blir kjørt som bruker noe CPU. 

### Oppgave 4
#### Oppgave 4A og B
Funksjonen er testet på Mac OS og Windows, hvor både terminal og Powershell er benyttet. Det var kun Windows Powershell ISE (x86) som skrev ut riktige symboler for det utvidede ASCII settet. Jeg antar dette er fordi denne versjonen av Powershell har full støtte for UNICODE tegn, som ikke Windows eller MacOS har default støtte for. 




