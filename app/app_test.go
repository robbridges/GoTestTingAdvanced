package app_test

import (
	"GoTestingAdvanced/app"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	app.Home(w, req)

	resp := w.Result()

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("io.ReadAll() err = %s, want nil", err)
	}
	got := string(body)
	want := "<h1>Welcome!</h1>"
	if got != want {
		t.Errorf("GET / =%s, want %s", got, want)
	}
}

func TestApp_v1(t *testing.T) {
	server := httptest.NewServer(&app.Server{})
	defer server.Close()

	resp, err := http.Get(server.URL)
	t.Log(server.URL)
	if err != nil {
		t.Fatalf("GET / err = %s; want nil", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	t.Log(body)
	if err != nil {
		t.Errorf("ioutil.ReadAll() err = %s; want nil", err)
	}
	got := string(body)
	want := "<h1>Welcome!</h1>"
	if got != want {
		t.Errorf("GET / = %s; want %s", got, want)
	}
}
