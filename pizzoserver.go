package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const pROVABILITATDEPROBLEMES int = 60
const mAXINCIDENTS int = 3

type lloc struct {
	Nom      string `json:"nom" xml:"nom"`
	Pagament int    `json:"pagament" xml:"-"`
}

type cobramento struct {
	On        lloc     `xml:"lloc"`
	Cobrador  string   `xml:"cobrador"`
	Incidents []string `xml:"incidents>incident"`
}

var incidents []string

var llocs []lloc

var cobradors = []string{"Rocco", "Enzo", "Tonino", "Fredo"}

func readLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

// Converteix el fitxer de llocs i quantitats a cobrar
// al format correcte
func readLlocs(linies []string) (llocs []lloc, err error) {

	for i, linia := range linies {
		valors := strings.Split(linia, ",")

		// Comprovar que són 2
		if len(valors) != 2 {
			return nil, fmt.Errorf("ERROR de format en la línia %d del fitxer de llocs: %v", i, valors)
		}
		// Comprovar que el segon és un número
		aCobrar, err := strconv.Atoi(strings.TrimSpace(valors[1]))
		if err != nil {
			return nil, fmt.Errorf("ERROR en la línia %d del fitxer de llocs, el valor a cobrar està malament %s", i, valors[1])
		}

		llocs = append(llocs, lloc{Nom: valors[0], Pagament: aCobrar})
	}
	return llocs, nil
}

// Calcula si hi ha hagut algun incident. I llavors retorna
// "CAP" si no n'hi ha hagut cap
// Una llista d'incidents si n'hi ha hagut algun.
func getIncidents(incidents []string) []string {

	nousIncidents := []string{"CAP"}

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

	incidents, err := readLines("incidents.txt")
	if err != nil {
		fmt.Println("Fitxer d'incidents `incidents.txt` no trobat")
		return
	}
	if len(incidents) < mAXINCIDENTS {
		fmt.Printf("En el fitxer d'incidents hi ha d'haver almenys %d incidents\n", mAXINCIDENTS)
	}

	dadesLlocs, err := readLines("llocs.txt")
	if err != nil {
		fmt.Println("Fitxer de llocs a cobrar `llocs.txt` no trobat")
		return
	}
	llocs, err := readLlocs(dadesLlocs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

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
		resultat.Incidents = getIncidents(incidents)
		c.XML(http.StatusOK, resultat)
	})

	router.Run()
}
