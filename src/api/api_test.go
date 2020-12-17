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
	request, _ := http.NewRequest("GET", "/api/piloto/hamilton", nil)

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
	request, _ := http.NewRequest("GET", "/api/escuderia/mercedes", nil)

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("A", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición no esperado, obtenido %s, esperado que contuviese: A",
			w.Body.String())
	}
}

func TestBuscaEstadisticas(t *testing.T) {

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/gp/australia/2019/estadisticas", nil)

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("\"Accidentes\":2", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición no esperado, obtenido %s, esperado que contuviese: \"Accidentes\":2",
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
func TestCambiaCircuito(t *testing.T) {

	w := httptest.NewRecorder()

	request, _ := http.NewRequest("PUT", "/api/circuito/australia",
		strings.NewReader("nombre=Le Castellet&pais=Francia"))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("\"Nombre\":\"Le Castellet\"", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición incorrecto, obtenido %s, esperado que contuviese \"Nombre\":\"Le Castellet\"",
			w.Body.String())
	}

	//Establecemos valor anterior
	w2 := httptest.NewRecorder()

	request2, _ := http.NewRequest("PUT", "/api/circuito/australia",
		strings.NewReader("nombre=C&pais=D"))
	request2.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w2, request2)
}

func TestCambiaEscuderia(t *testing.T) {

	w := httptest.NewRecorder()

	request, _ := http.NewRequest("PUT", "/api/escuderia/alfaromeo",
		strings.NewReader("nombre=A&piloto1=hamilton&piloto2=Francisco Beltrán&titulos=2&victorias=1&poles=6&vueltasrapidas=9"))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("\"Nombre\":\"A\"", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición incorrecto, obtenido %s, esperado que contuviese \"Nombre\":\"A\"",
			w.Body.String())
	}

	//Establecemos valor anterior
	w2 := httptest.NewRecorder()

	request2, _ := http.NewRequest("PUT", "/api/escuderia/alfaromeo",
		strings.NewReader("nombre=Alfa Romeo Sauber F1 Team&piloto1=raikkonen&piloto2=giovinazzi&titulos=0&victorias=0&poles=0&vueltasrapidas=2"))
	request2.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w2, request2)
}

func TestCambiaPiloto(t *testing.T) {

	w := httptest.NewRecorder()

	request, _ := http.NewRequest("PUT", "/api/piloto/raikkonen",
		strings.NewReader("nombre=Valtteri Bottas&victorias=10&poles=20&vueltasrapidas=5&mundiales=0"))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("\"Nombre\":\"Valtteri Bottas\"", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición incorrecto, obtenido %s, esperado que contuviese \"Nombre\":\"Valtteri Bottas\"",
			w.Body.String())
	}

	//Establecemos valor anterior
	w2 := httptest.NewRecorder()

	request2, _ := http.NewRequest("PUT", "/api/piloto/raikkonen",
		strings.NewReader("nombre=Kimi Raikkonen&victorias=21&poles=10&vueltasrapidas=20&mundiales=1"))
	request2.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w2, request2)
}

func TestCambiaEstadisticas(t *testing.T) {

	w := httptest.NewRecorder()

	request, _ := http.NewRequest("PUT", "/api/gp/australia/2019/estadisticas",
		strings.NewReader("accidentes=2&numerosafety=0&adelantamientos=2&banderasamarillas=0&banderasrojas=1&sanciones=5&mejortiempo=1:23:456"))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("\"Adelantamientos\":2", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición incorrecto, obtenido %s, esperado que contuviese \"Adelantamientos\":200",
			w.Body.String())
	}

	//Establecemos valor anterior
	w2 := httptest.NewRecorder()

	request2, _ := http.NewRequest("PUT", "/api/gp/australia/2019/estadisticas",
		strings.NewReader("accidentes=2&numerosafety=0&adelantamientos=200&banderasamarillas=0&banderasrojas=1&sanciones=5&mejortiempo=1:23:456"))
	request2.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w2, request2)
}

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

/****************************** TEST POSTS *********************************************/

func TestCreaCircuito(t *testing.T) {

	w := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", "/api/circuito",
		strings.NewReader("nombre=Le Castellet&pais=Francia"))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 201 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 201)
	}

	matched, _ := regexp.MatchString("\"Nombre\":\"Le Castellet\"", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición incorrecto, obtenido %s, esperado que contuviese \"Nombre\":\"Le Castellet\"",
			w.Body.String())
	}

}

func TestCreaEscuderia(t *testing.T) {

	w := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", "/api/escuderia",
		strings.NewReader("nombre=Prueba&piloto1=hamilton&piloto2=Francisco Beltrán&titulos=2&victorias=1&poles=6&vueltasrapidas=9"))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 201 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 201)
	}

	matched, _ := regexp.MatchString("\"Nombre\":\"Prueba\"", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición incorrecto, obtenido %s, esperado que contuviese \"Nombre\":\"Prueba\"",
			w.Body.String())
	}
}

func TestCreaPiloto(t *testing.T) {

	w := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", "/api/piloto",
		strings.NewReader("nombre=Lastname&victorias=10&poles=20&vueltasrapidas=5&mundiales=0"))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 201 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 201)
	}

	matched, _ := regexp.MatchString("\"Nombre\":\"Lastname\"", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición incorrecto, obtenido %s, esperado que contuviese \"Nombre\":\"Lastname\"",
			w.Body.String())
	}
}

func TestCreaGranPremio(t *testing.T) {

	w := httptest.NewRecorder()

	request, _ := http.NewRequest("POST", "/api/gp",
		strings.NewReader("nombre=Le Castellet&fp1=2:45:678|1:23:456#hamilton#1:23:456&fp2=2:45:678|1:23:456#hamilton#1:23:456&fp3=2:45:678|1:23:456#hamilton#1:23:456&clasificacion=2:45:678|1:23:456#hamilton#1:23:456&carrera=2:45:678|1:23:456#hamilton#1:23:456&ganador=hamilton&poleman=hamilton&podio=hamilton#ocon#perez&temporada=2012&estadisticas=2#1#2#4#0#10#2:22:000"))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 201 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 201)
	}

	matched, _ := regexp.MatchString("\"Ganador\":{\"Nombre\":\"Lewis Hamilton\"", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición incorrecto, obtenido %s, esperado que contuviese \"Ganador\":{\"Nombre\":\"Lewis Hamilton\"",
			w.Body.String())
	}
}

/************************************* TEST DELETE ************************************/

func TestEliminaCircuito(t *testing.T) {

	w := httptest.NewRecorder()

	request, _ := http.NewRequest("DELETE", "/api/circuito/Le Castellet", nil)

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("Registro Eliminado", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición incorrecto, obtenido %s, esperado que contuviese Registro Eliminado",
			w.Body.String())
	}
}

func TestEliminaPiloto(t *testing.T) {

	w := httptest.NewRecorder()

	request, _ := http.NewRequest("DELETE", "/api/piloto/Lastname", nil)

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("Registro Eliminado", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición incorrecto, obtenido %s, esperado que contuviese Registro Eliminado",
			w.Body.String())
	}
}

func TestEliminaEscuderia(t *testing.T) {

	w := httptest.NewRecorder()

	request, _ := http.NewRequest("DELETE", "/api/escuderia/Prueba", nil)

	r.ServeHTTP(w, request)

	fmt.Println(w.Code)
	if w.Code != 200 {
		t.Errorf("Código de estado no esperado, obtenido %d esperado %d", w.Code, 200)
	}

	matched, _ := regexp.MatchString("Registro Eliminado", w.Body.String())

	if !matched {
		t.Errorf("Cuerpo de la petición incorrecto, obtenido %s, esperado que contuviese Registro Eliminado",
			w.Body.String())
	}
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
