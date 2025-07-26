package mandi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

const (
	apiKey  = "579b464db66ec23bdd000001185c55b29a984645468d0f573fee9052"
	baseURL = "https://api.data.gov.in/resource/35985678-0d79-46b4-9ed6-6f13308a1d24"
)

func(m *mandiHandler) GetMandiBhav(c *gin.Context) {
	params := url.Values{}
	params.Add("api-key", apiKey)
	params.Add("format", "json")

	// Query filters
	if s := c.Query("state"); s != "" {
		params.Add("filters[State]", s)
	}
	if d := c.Query("district"); d != "" {
		params.Add("filters[District]", d)
	}
	if cm := c.Query("commodity"); cm != "" {
		params.Add("filters[Commodity]", cm)
	}
	if dt := c.Query("date"); dt != "" {
		params.Add("filters[Arrival_Date]", dt)
	}
	if l := c.Query("limit"); l != "" {
		params.Add("limit", l)
	}

	// Build full URL
	fullURL := baseURL + "?" + params.Encode()

	// Fetch JSON from data.gov.in
	resp, err := http.Get(fullURL)
	if err != nil || resp.StatusCode != 200 {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to fetch data", "details": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read data"})
		return
	}

	// Parse only the "records" field
	var raw map[string]interface{}
	if err := json.Unmarshal(body, &raw); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid JSON"})
		return
	}

	if records, ok := raw["records"]; ok {
		c.JSON(http.StatusOK, records)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "No records found"})
	}
}
