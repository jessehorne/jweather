package common

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	characterHot      = "hot"
	characterCold     = "cold"
	characterModerate = "moderate"
)

type Forecast struct {
	Short     string `json:"short"`     // a short forecast for the area
	Character string `json:"character"` // a characterization such as "hot", "cold", or "moderate"
}

func (f *Forecast) ToJSON() []byte {
	j, err := json.Marshal(f)
	if err != nil {
		log.Println(err)
		return []byte("{}")
	}
	return j
}

type NWSForecastResponse struct {
	Properties NWSForecastResponseProperties `json:"properties"`
}

type NWSForecastResponseProperties struct {
	Periods []NWSPeriod `json:"periods"`
}

type NWSPeriod struct {
	Temperature   float64 `json:"temperature"`
	ShortForecast string  `json:"shortForecast"`
}

func GetForecastByURL(url string) (*Forecast, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "(github.com/jessehorne, j.horne2796@gmail.com")
	req.Header.Set("Accept", "application/geo+json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var f *NWSForecastResponse
	err = json.Unmarshal(body, &f)
	if err != nil {
		return nil, err
	}

	newForecast := &Forecast{}

	if len(f.Properties.Periods) > 0 {
		p := f.Properties.Periods[0]
		newForecast.Short = p.ShortForecast

		// assuming we're using fahrenheit and also these are my opinions
		if p.Temperature <= 32 {
			newForecast.Character = characterCold
		} else if p.Temperature > 90 {
			newForecast.Character = characterHot
		} else {
			newForecast.Character = characterModerate
		}
	}

	return newForecast, nil
}

type NWSPointsResponse struct {
	Properties NWSPointsResponseProperties `json:"properties"`
}

type NWSPointsResponseProperties struct {
	Forecast string `json:"forecast"`
}

// GetForecastByPoints gets the forecast from the NWS API and returns a *Forecast or error.
func GetForecastByPoints(lat, long string) (*Forecast, error) {
	url := fmt.Sprintf("https://api.weather.gov/points/%s,%s", lat, long)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "(github.com/jessehorne, j.horne2796@gmail.com")
	req.Header.Set("Accept", "application/geo+json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var r *NWSPointsResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return GetForecastByURL(r.Properties.Forecast)
}
