package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const pROVABILITATDEPROBLEMES int = 60
const mAXINCIDENTS int = 2

type lloc struct {
	Nom      string `json:"nom" xml:"nom"`
	Pagament int    `json:"pagament" xml:"-"`
}

type cobramento struct {
	On        lloc     `xml:"lloc"`
	Cobrador  string   `xml:"cobrador"`
	Incidents []string `xml:"incidents>incident"`
}

var incidents = [...]string{
	"Li hem trencat un braç",
	"Ha fet falta amenaçar-lo",
	"Hem trencat la porta",
	"Li hem segrestat la dona",
	"Li he matat la mascota",
	"Se li ha hagut de tallar un dit",
	"Hem matat el seu gat .. d'un tret",
	"L'hem hagut d'apallissar una mica",
	"Li hem posat un cap de cavall al llit",
	"Se li ha trencat l'aparador",
	"Li hem clavat la mà esquerra a una taula",
	"Li hem clavat la mà dreta a una taula",
	"Li he hagut d'afaitar les parts",
	"Se li ha posat el cap en un forn",
	"Se l'ha convençut de pagar posant-li un mitxó a la boca",
	"Es negava a pagar fins que l'hem ruixat amb gasolina",
	"Ha pagat amb espècies ... i després amb diners",
	"He hagut de jugar a futbol amb el seu cap",
	"He destrossat la porta ... en realitat totes les portes",
	"Li hem calat foc al magatzem",
	"Se li han donat un parell de cops de porra perquè pagués més de pressa",
	"S'han trencat les finestres...",
}

var llocs = []lloc{
	lloc{"Bar Cendrassi", 30},
	lloc{"Fira Dawsecondo", 20},
	lloc{"Serveis Putoprofe", 50},
	lloc{"Regals Noaprovare", 10},
	lloc{"SuperMercat Carbonara", 25},
	lloc{"Pizzes Puttanesca", 12},
	lloc{"Colmado Ramone", 5},
	lloc{"Supermercat Caprese", 30},
	lloc{"Bar Pesto Rosso", 30},
	lloc{"Tasca Bolognesa", 12},
	lloc{"Caruso, S.L.", 10},
	lloc{"Bar Tonnata", 20},
	lloc{"Colmado Parmigiano", 25},
	lloc{"Peperoni e amicci", 18},
	lloc{"La casa di Pomodoro", 20},
	lloc{"Il Ragu napoletano", 25},
}

var cobradors = []string{"Rocco", "Enzo", "Tonino", "Fredo"}

// Calcula si hi ha hagut algun incident. I llavors retorna
// "CAP" si no n'hi ha hagut cap
// Una llista d'incidents si n'hi ha hagut algun.
func getIncidents() []string {
	nousIncidents := make([]string, 1)
	nousIncidents[0] = "CAP"

	hiHaIncidents := (rand.Intn(100) > pROVABILITATDEPROBLEMES)
	if hiHaIncidents {
		numIncidents := rand.Intn(mAXINCIDENTS) + 1
		nousIncidents = make([]string, numIncidents)
		for i := 0; i < numIncidents; i++ {
			nousIncidents[i] = incidents[rand.Intn(len(incidents))]
		}
	}
	return nousIncidents
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	router := gin.Default()

	router.GET("/llocs", func(c *gin.Context) {
		cadena := make([]string, len(llocs))
		for index, record := range llocs {
			cadena[index] = "(" + record.Nom + "," + strconv.Itoa(record.Pagament) + ")"
		}
		c.String(http.StatusOK, "["+strings.Join(cadena, ",")+"]")
	})

	router.GET("/cobrar", func(c *gin.Context) {
		resultat := new(cobramento)
		resultat.Cobrador = cobradors[rand.Intn(len(cobradors))]
		resultat.On = llocs[rand.Intn(len(llocs))]
		resultat.Incidents = getIncidents()
		c.XML(http.StatusOK, resultat)
	})

	router.Run()
}
