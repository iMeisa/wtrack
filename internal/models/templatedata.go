package models

import (
	"github.com/iMeisa/weed/internal/config"
	"github.com/justinas/nosurf"
	"net/http"
	"time"
)

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	Records   []Record
	CSRFToken string
	Flash     string // Some message
	Warning   string
	Error     string
}

// AddDefaultData adds data for all templates
func (data *TemplateData) AddDefaultData(r *http.Request, a *config.AppConfig) {
	// String map
	//fmt.Println("Adding default data")
	stringMap := make(map[string]string)
	if data.StringMap != nil {
		stringMap = data.StringMap
	}

	data.StringMap = stringMap

	data.CSRFToken = nosurf.Token(r)

	records := []Record{
		{
			Date:   time.Date(2022, 07, 20, 12, 00, 00, 00, time.UTC),
			Amount: 1.6,
			Strain: Strain{
				ID:     1,
				Name:   "Diamonds",
				Amount: 20,
			},
		},
	}

	data.Records = records

	// Separate types

	if data.Data == nil {
		data.Data = make(map[string]interface{})
	}
}

// GetDefaultData returns a TemplateData object with default data
func GetDefaultData(r *http.Request, a *config.AppConfig) TemplateData {
	data := TemplateData{}
	data.AddDefaultData(r, a)

	return data
}
