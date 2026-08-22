package repo

// MockDataProvider provides static/mock data for testing and development
// when database is unavailable
type MockDataProvider struct{}

// GetMockUser returns mock user data
func (m *MockDataProvider) GetMockUser() map[string]interface{} {
	return map[string]interface{}{
		"id":    1,
		"name":  "Demo User",
		"email": "demo@kisaansathi.com",
		"role":  "user",
		"phone": "+91-9999999999",
	}
}

// GetMockForecasts returns mock forecast data
func (m *MockDataProvider) GetMockForecasts() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"id":       1,
			"location": "Maharashtra",
			"crop":     "Wheat",
			"date":     "2026-04-16",
			"forecast": "Rain expected",
		},
		{
			"id":       2,
			"location": "Punjab",
			"crop":     "Rice",
			"date":     "2026-04-16",
			"forecast": "Clear skies",
		},
	}
}

// GetMockMandiRates returns mock mandi rates
func (m *MockDataProvider) GetMockMandiRates() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"id":     1,
			"mandi":  "APMC Nashik",
			"crop":   "Onion",
			"price":  "₹2500/quintal",
			"date":   "2026-04-16",
			"update": "Daily",
		},
		{
			"id":     2,
			"mandi":  "APMC Pune",
			"crop":   "Tomato",
			"price":  "₹1800/quintal",
			"date":   "2026-04-16",
			"update": "Daily",
		},
	}
}

// GetMockFeeds returns mock feed data
func (m *MockDataProvider) GetMockFeeds() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"id":      1,
			"title":   "New Crop Season Tips",
			"content": "Best practices for the upcoming season",
			"author":  "Agriculture Expert",
			"date":    "2026-04-16",
		},
		{
			"id":      2,
			"title":   "Water Management Guide",
			"content": "Efficient irrigation techniques",
			"author":  "Farm Specialist",
			"date":    "2026-04-15",
		},
	}
}

// NewMockDataProvider creates a new MockDataProvider instance
func NewMockDataProvider() *MockDataProvider {
	return &MockDataProvider{}
}
