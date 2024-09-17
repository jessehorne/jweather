package handlers

import (
	"github.com/jessehorne/jweather/internal/api/common"
	"net/http"
)

// PointsQueryHandler is the HTTP handler function for "GET /points" which requires the `long` and `lat` query parameters
// and returns a JSON response with the fields `short` (a short forecast for the specified area) and `character`, a
// characterization to describe the temperature, such as "cold", "hot" or "moderate".
func PointsQueryHandler(w http.ResponseWriter, req *http.Request) {
	lat := req.URL.Query().Get("lat")
	if lat == "" {
		common.APIResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": "missing lat",
		})
		return
	}

	long := req.URL.Query().Get("long")
	if long == "" {
		common.APIResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": "missing long",
		})
		return
	}

	points, err := common.GetForecastByPoints(lat, long)
	if err != nil {
		common.APIResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(points.ToJSON())
}
