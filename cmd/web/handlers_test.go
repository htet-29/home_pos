package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNavigateAllRoutesReturnOKStatus(t *testing.T) {
	tt := []struct {
		name           string
		url            string
		expectedStatus int
	}{
		{
			name:           "Navigate to / route",
			url:            "http://localhost:4000/",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Navigate to stock viewing",
			url:            "http://localhost:4000/pos/stock",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Navigate to stock creation",
			url:            "http://localhost:4000/pos/stock/create",
			expectedStatus: http.StatusOK,
		}}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(tc.url)
			if err != nil {
				t.Fatal("Cannot get to the url", err)
			}
			if resp.StatusCode != tc.expectedStatus {
				t.Errorf("expected status ok, got %d instead.\n", resp.StatusCode)
			}
		})
	}
}

func TestHomeRoute(t *testing.T) {
	templateCache, err := newTemplateCache()
	if err != nil {
		t.Fatalf("Cannot create template cache: %v", err)
	}

	app := &application{
		templateCache: templateCache,
	}

	ts := httptest.NewServer(app.routes())
	defer ts.Close()

	path := ts.URL + "/"
	rs, err := ts.Client().Get(path)
	if err != nil {
		t.Errorf("%v not found. %v", path, err)
	}

	if rs.StatusCode != http.StatusOK {
		t.Errorf("expected %d but got %d instead.", http.StatusOK, rs.StatusCode)
	}
}
