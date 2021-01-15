package api

import (
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
	request, _ := http.NewRequest("GET", "/circuito/Albert Park", nil)

	r.ServeHTTP(w, request)

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
	request, _ := http.NewRequest("GET", "/piloto/Lewis Hamilton", nil)

	r.ServeHTTP(w, request)

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
	request, _ := http.NewRequest("GET", "/escuderia/Mercedes", nil)

	r.ServeHTTP(w, request)

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
	request, _ := http.NewRequest("GET", "/gp/Albert Park/2019/estadisticas", nil)

	r.ServeHTTP(w, request)

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
	request, _ := http.NewRequest("GET", "/gp/Albert Park/2019/sesion/fp1", nil)

	r.ServeHTTP(w, request)

	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("\"Tiempos\":", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición no esperado, obtenido %s, esperado que contuviese: \"Tiempos\":",
			w.Body.String())
	}
}

/****************************** TESTS PUTS *************************************/
func TestModificaCircuito(t *testing.T) {
	w := httptest.NewRecorder()

	request, _ := http.NewRequest("PUT", "/circuito/Albert Park",
		strings.NewReader("{\"Nombre\":\"A\",\"Pais\":\"B\"}"))
	request.Header.Add("Content-Type", "application/json")

	r.ServeHTTP(w, request)

	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("\"Nombre\":\"A\"", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición no esperado, obtenido %s, esperado que contuviese: \"Nombre\":A",
			w.Body.String())
	}

	request2, _ := http.NewRequest("PUT", "/circuito/Albert Park",
		strings.NewReader("Nombre=Albert Park&Pais=Australia"))
	request2.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, request2)
}

func TestModificaEscuderia(t *testing.T) {
	w := httptest.NewRecorder()

	request, _ := http.NewRequest("PUT", "/escuderia/Mercedes",
		strings.NewReader("{\"Nombre\":\"A\",\"Piloto1\":\"Lewis Hamilton\",\"Piloto2\":\"Sebastian Vettel\",\"Titulos\":8,\"Victorias\":170,\"Poles\":200,\"Vueltas Rapidas\":180}"))
	request.Header.Add("Content-Type", "application/json")

	r.ServeHTTP(w, request)

	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("\"Victorias\":170", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición no esperado, obtenido %s, esperado que contuviese: \"Nombre\":A",
			w.Body.String())
	}

	request2, _ := http.NewRequest("PUT", "/escuderia/Mercedes",
		strings.NewReader("{\"Nombre\":\"Mercedes\",\"Piloto1\":\"Lewis Hamilton\",\"Piloto2\":\"Valtteri Bottas\",\"Titulos\":8,\"Victorias\":170,\"Poles\":200,\"Vueltas Rapidas\":180}"))
	request2.Header.Add("Content-Type", "application/json")

	r.ServeHTTP(w, request2)

	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched2, _ := regexp.MatchString("\"Nombre\":\"Mercedes\"", w.Body.String())

	if !matched2 {
		t.Errorf("Cuerpo de la petición no esperado, obtenido %s, esperado que contuviese: \"Nombre\":Mercedes",
			w.Body.String())
	}
}

func TestModificaPiloto(t *testing.T) {
	w := httptest.NewRecorder()

	request, _ := http.NewRequest("PUT", "/piloto/Lewis Hamilton",
		strings.NewReader("{\"Nombre\":\"Lewis\",\"Mundiales\":7,\"Victorias\":92,\"Poles\":98,\"Vueltas Rapidas\":55}"))
	request.Header.Add("Content-Type", "application/json")

	r.ServeHTTP(w, request)

	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("\"Nombre\":\"Lewis\",", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición no esperado, obtenido %s, esperado que contuviese: \"Nombre\":Lewis",
			w.Body.String())
	}

	request2, _ := http.NewRequest("PUT", "/piloto/Lewis Hamilton",
		strings.NewReader("{\"Nombre\":\"Lewis Hamilton\",\"Mundiales\":7,\"Victorias\":92,\"Poles\":98,\"Vueltas Rapidas\":55}"))
	request2.Header.Add("Content-Type", "application/json")

	r.ServeHTTP(w, request2)

	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched2, _ := regexp.MatchString("\"Nombre\":\"Lewis Hamilton\",", w.Body.String())

	if !matched2 {
		t.Errorf("Cuerpo de la petición no esperado, obtenido %s, esperado que contuviese: \"Nombre\":Lewis Hamilton",
			w.Body.String())
	}
}

func TestCreaEscuderia(t *testing.T) {
	w := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", "/escuderia",
		strings.NewReader("{\"Nombre\":\"Ferrari\",\"Piloto1\":\"Lewis Hamilton\",\"Piloto2\":\"Sebastian Vettel\",\"Titulos\":8,\"Victorias\":170,\"Poles\":200,\"Vueltas Rapidas\":180}"))
	request.Header.Add("Content-Type", "application/json")

	r.ServeHTTP(w, request)

	if w.Code != 201 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 201)
	}

	matched, _ := regexp.MatchString("\"Nombre\":\"Ferrari\"", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición no esperado, obtenido %s, esperado que contuviese: \"Nombre\":Ferrari",
			w.Body.String())
	}
}

/****************************** TESTS POST *************************************/

func TestCreaCircuito(t *testing.T) {
	w := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", "/circuito",
		strings.NewReader("{\"Nombre\":\"Mónaco\",\"Pais\":\"Mónaco\"}"))
	request.Header.Add("Content-Type", "application/json")

	r.ServeHTTP(w, request)

	if w.Code != 201 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 201)
	}

	matched, _ := regexp.MatchString("\"Nombre\":\"Mónaco\"", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición no esperado, obtenido %s, esperado que contuviese: \"Nombre\":Mónaco",
			w.Body.String())
	}
}

func TestCreaPiloto(t *testing.T) {
	w := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", "/piloto",
		strings.NewReader("{\"Nombre\":\"Fernando Alonso\",\"Mundiales\":7,\"Victorias\":92,\"Poles\":98,\"Vueltas Rapidas\":55}"))
	request.Header.Add("Content-Type", "application/json")

	r.ServeHTTP(w, request)

	if w.Code != 201 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 201)
	}

	matched, _ := regexp.MatchString("\"Nombre\":\"Fernando Alonso\",", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición no esperado, obtenido %s, esperado que contuviese: \"Nombre\":Fernando Alonso",
			w.Body.String())
	}
}

/****************************** OTROS TESTS *************************************/

func TestObtieneMedia(t *testing.T) {

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/procesogp/Albert Park/media/accidentes", nil)

	r.ServeHTTP(w, request)

	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("\"Media\":10", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición incorrecto, obtenido %s, esperado que contuviese \"Media\":10",
			w.Body.String())
	}
}

func TestObtienePrediccion(t *testing.T) {

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/procesogp/Albert Park/prediccion/accidentes", nil)

	r.ServeHTTP(w, request)

	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("\"Posibles Accidentes\":10", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición incorrecto, obtenido %s, esperado que contuviese \"Posibles Accidentes\":10",
			w.Body.String())
	}
}
