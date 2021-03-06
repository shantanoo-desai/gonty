package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	composerMock      string = `{"Name":"Johann Gambolputty","FamilyName":["de von Ausfern","schplenden","schlitter","crasscrenbon","fried","digger","dingle","dangle","dongle","dungle","burstein","von","knacker","thrasher","apple","banger","horowitz","ticolensic","grander","knotty","spelltinkle","grandlich","grumblemeyer","spelterwasser","kurstlich","himbleeisen","bahnwagen","gutenabend","bitte","ein","nürnburger","bratwustle","gerspurten","mitzweimache","luber","hundsfut","gumberaber","shönendanker","kalbsfleisch","mittler","aucher von Hautkopft of Ulm"],"Genre":"Baroque","Nationality":"German"}`
	jokesMock         string = `[{"Delivery":"Wenn ist das Nunstück git und Slotermeyer? Ja! Beiherhund das Oder die Flipperwaldt gersput!","Deadly":true,"Language":"German"},{"Delivery":"There were zwei [two] peanuts walking down der strasse [street]. Und one was assaulted... peanut!","Deadly":false,"Language":"English"}]`
	funniestJokeMock string = `{"Delivery":"Wenn ist das Nunstück git und Slotermeyer? Ja! Beiherhund das Oder die Flipperwaldt gersput!","Deadly":true,"Language":"German"}`
	expectMock        string = `{"Who":"Spanish Inquisition","Expecting":false,"Weapons":["fear","surprise","ruthless efficiency","fanatical devotion to the pope","nice red uniforms"]}`
	noblemanMock      string = `{"Name":"Biggus","FamilyName":"Dickus","Empire":"Roman","RelatedTo":[{"Name":"Pontius","FamilyName":"Pilate","Empire":"Roman","RelatedTo":[{"Name":"Julius","FamilyName":"Caesar","Empire":"Roman","RelatedTo":null}]},{"Name":"Incontinetia","FamilyName":"Buttocks","Empire":"Roman","RelatedTo":null}]}`
)

var router = setupRouter()

var w = httptest.NewRecorder()

func TestHealthCheckRoute(t *testing.T) {

	var router = setupRouter()

	var w = httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}

func TestComposerOK(t *testing.T) {

	var router = setupRouter()

	var w = httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/composer/baroque", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, composerMock, w.Body.String())
}

func TestComposerNotFound(t *testing.T) {

	var router = setupRouter()

	var w = httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/composer/classical", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "This is all getting rather silly!", w.Body.String())

}

func TestAllJokes(t *testing.T) {

	var router = setupRouter()

	var w = httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/jokes", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, jokesMock, w.Body.String())
}

func TestFunniestJoke(t *testing.T) {

	var router = setupRouter()

	var w = httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/jokes?funniest=true", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, funniestJokeMock, w.Body.String())
}

func TestExpect(t *testing.T) {
	var router = setupRouter()

	var w = httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/expect", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectMock, w.Body.String())
}

func TestNobleman(t *testing.T) {
	var router = setupRouter()

	var w = httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/history/nobleman", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, noblemanMock, w.Body.String())
}
