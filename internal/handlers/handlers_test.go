package handlers

import (
	"net/http"
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
