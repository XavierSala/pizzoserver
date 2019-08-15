#FpInfor #DawMp03Uf04 #Programació

# No ens paguen la protecció

La Màfia està perdent pistonada i darrerament hi ha problemes per "cobrar la protecció" dels diferents locals. Els propietaris s'estan rebel·lant i a vegades costa massa cobrar... Per això el Capo Alfredo ha decidit que s'ha d'escarmentar algú per fer que els altres no es resisteixin...

![Mafia](README/mafia.png)

Necessita un programa que li permeti saber quantes vegades han deixat de pagar un determinat local.

El problema és que l'antic programador va fer que la informació que els cobradors entren en el sistema s'hagi de recuperar d'un servidor web que s'inicia executant el programa `pizzoserver`:

```bash
$ ./pizzoserver
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /llocs                    --> main.main.func1 (3 handlers)
[GIN-debug] GET    /cobrar                   --> main.main.func2 (3 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080

```

Com es pot veure en el logging el servidor web només té dues operacions:

- Veure la llista de llocs a cobrar: [http://localhost:8080/llocs](http://localhost:8080/llocs). Que retorna els llocs en un format de text on hi ha el nom del local i els diners a cobrar-hi.

  [(Bar Cendrassi,30),(Fira Dawsecondo,20),(Serveis Putoprofe,50),(Regals Noaprovare,10),(SuperMercat Carbonara,25),(Pizzes Puttanesca,12),(Colmado Ramone,5),(Supermercat Caprese,30),(Bar Pesto Rosso,30),(Tasca Bolognesa,12),(Caruso S.L.,10),(Bar Tonnata",20),(Colmado Parmigiano,25),(Peperoni e amicci,18),(La casa di Pomodoro,20),(Il Ragu napoletano,25)]

- Obtenir els cobraments un a un fent una petició al servidor [http://localhost:8080/cobrar](http://localhost:8080/cobrar). Aquesta petició retorna un fitxer XML en el que hi ha informació sobre el lloc on s'ha cobrat, qui ha sigut i si hi ha hagut incidents o no ('CAP' és que tot ha anat bé)

```xml
<cobramento>
   <lloc>
      <nom>Peperoni e amicci</nom>
    </lloc>
    <cobrador>Enzo</cobrador>
    <incidents>
       <incident>CAP</incident>
    </incidents>
</cobramento>
```

Després de veure com havia fet el servidor l'antic programador us han contractat a vosaltres perquè feu el programa client.

## Exercici

El capo us ha contractat perquè li feu un programa client que a partir de les dades que rep del servidor, digui en quin dels llocs protegits de la zona hi ha hagut més incidents en els darrers 100 cobraments però **que no siguin dels locals que paguen més**.

1. Dissenyeu els objectes necessaris per fer aquest programa
2. Feu un esquema de les passes que ha de fer el programa
3. Desenvolupeu el programa

![no patiu](https://raw.githubusercontent.com/utrescu/utrescu.github.io/master/images/capo4.png)

## Programa

El programa servidor està desenvolupat en `Go`. Per compilar-lo cal tenir el compilador instal·lat i executar:

```bash
go build
```

Generarà un executable de la plataforma de desenvolupament que es faci servir (Windows, Linux, ...). S'executa i ja tindrem el servidor en marxa:

```bash
./pizzoserver
```

Els locals i els diners a cobrar-hi estan en el fitxer `llocs.txt` i els possibles incidents (per aquest exercici no serveixen per res) estan a `incidents.txt`
