package danskenergi

import (
	"errors"
	"math"
	"net/http"
	"strings"

	"github.com/nyaruka/phonenumbers"
)

func (c *Client) Incidents() ([]Incident, error) {

	var result []Incident
	resp, err := c.R().SetResult(&result).EnableTrace().Get("incidents")
	if err != nil {
		return result, err
	}

	if resp.StatusCode() != http.StatusOK {
		return result, errors.New(resp.Status())
	}

	for i := 0; i < len(result); i++ {
		zcs := strings.Split(result[i].ZipCodesRaw, ",")
		sliceTo := len(zcs)

		if zcs[len(zcs)-1] == "" {
			sliceTo--
		}
		result[i].ZipCodes = zcs[:sliceTo]

		num, err := phonenumbers.Parse(strings.Split(result[i].SupplierPhoneRaw, ")")[1], "DK")
		if err != nil {
			continue
		}
		result[i].SupplierPhone = num
	}

	return result, nil
}

// Calculates the distance from a given coordinate in kilometers
func (i *Incident) BirdFlyDistance(lat, lng float64) float64 {
	radlat1 := float64(math.Pi * i.CenterLat / 180)
	radlat2 := float64(math.Pi * lat / 180)

	theta := float64(i.CenterLng - lng)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515
	dist = dist * 1.609344

	return dist
}
