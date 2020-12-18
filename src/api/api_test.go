package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func TestMain(m *testing.M) {
	r = DiseñoRutas()

	os.Exit(m.Run())
}

/************************************** TEST GETS *********************************************/
func TestBuscaCircuito(t *testing.T) {

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/circuito/Albert Park", nil)

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("\"Nombre\":\"Albert Park\"", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición incorrecto, obtenido %s, esperado que contuviese \"Nombre\":\"Albert Park\"",
			w.Body.String())
	}
}

func TestBuscaPiloto(t *testing.T) {

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/piloto/Lewis Hamilton", nil)

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("Lewis", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición no esperado, obtenido %s, esperado que contuviese: Lewis",
			w.Body.String())
	}
}

func TestBuscaEscuderia(t *testing.T) {

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/escuderia/Mercedes", nil)

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("Mercedes", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición no esperado, obtenido %s, esperado que contuviese: Mercedes",
			w.Body.String())
	}
}

func TestBuscaEstadisticas(t *testing.T) {

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/gp/Albert Park/2019/estadisticas", nil)

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("\"Accidentes\":10", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición no esperado, obtenido %s, esperado que contuviese: \"Accidentes\":10",
			w.Body.String())
	}
}

func TestBuscaResultadoSesion(t *testing.T) {

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/gp/australia/2019/sesion/fp1", nil)

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("\"Tiempos\":", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición no esperado, obtenido %s, esperado que contuviese: \"Tiempos\":",
			w.Body.String())
	}
}

/************************************** TEST PUTS *********************************************/

func TestCambiaSesion(t *testing.T) {

	w := httptest.NewRecorder()

	request, _ := http.NewRequest("PUT", "/api/gp/australia/2019/sesion/fp1",
		strings.NewReader("sesion=2:45:678|1:23:456#verstappen#1:23:456"))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("Max Verstappen", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición incorrecto, obtenido %s, esperado que contuviese: Max Verstappen",
			w.Body.String())
	}

	//Establecemos valor anterior
	w2 := httptest.NewRecorder()

	request2, _ := http.NewRequest("PUT", "/api/gp/australia/2019/sesion/fp1",
		strings.NewReader("sesion=2:45:678|1:23:456#hamilton#1:23:456"))
	request2.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w2, request2)
}

/****************************** OTROS TESTS *************************************/

func TestObtieneMedia(t *testing.T) {

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/procesogp/australia/media/accidentes", nil)

	r.ServeHTTP(w, request)

	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("\"Media\":2", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición incorrecto, obtenido %s, esperado que contuviese \"Media\":2",
			w.Body.String())
	}
}

func TestObtienePrediccion(t *testing.T) {

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/procesogp/australia/prediccion/accidentes", nil)

	r.ServeHTTP(w, request)

	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("\"Posibles Accidentes\":2", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición incorrecto, obtenido %s, esperado que contuviese \"Posibles Accidentes\":2",
			w.Body.String())
	}
}
