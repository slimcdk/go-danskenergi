package danskenergi

import (
	"strings"
	"time"

	"github.com/nyaruka/phonenumbers"
)

type CustomTime time.Time

type Incident struct {
	ID                  int64       `json:"id"`
	Created             *CustomTime `json:"created"`
	Title               string      `json:"title"`
	Comment             string      `json:"comment"`
	Cause               string      `json:"cause"`
	IncidentTypeId      int64       `json:"incidentTypeId"`
	IncidentType        string      `json:"incidentType"`
	IncidentStatusId    int64       `json:"incidentStatusId"`
	IncidentStatus      string      `json:"incidentStatus"`
	SupplierId          int64       `json:"supplierId"`
	SupplierName        string      `json:"supplierName"`
	SupplierPhoneRaw    string      `json:"supplierPhone"`
	SupplierPhone       *phonenumbers.PhoneNumber
	SupplierLogo        string      `json:"supplierLogo"`
	SupplierWeb         string      `json:"supplierWeb"`
	PoweroutagesCount   int64       `json:"poweroutagesCount"`
	EffectedCustomers   int64       `json:"effectedCustomers"`
	StartDate           *CustomTime `json:"startDate"`
	EndDate             *CustomTime `json:"endDate"`
	ExpectedDowntime    *CustomTime `json:"expectedDowntime"`
	ZipCodesRaw         string      `json:"zipcodes"`
	ZipCodes            []string
	InternalReferenceId interface{} `json:"internalReferenceId"`
	SortOrder           int64       `json:"sortOrder"`
	ReportedViaApi      bool        `json:"reportedViaApi"`
	Radius              float64     `json:"radius"`
	CenterLat           float64     `json:"centerLat"`
	CenterLng           float64     `json:"centerLng"`
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" || value == "undefined" {
		return nil
	}

	t, err := time.Parse(time.RFC3339[:len(time.RFC3339)-6], value) // RFC3339 without timezone
	if err != nil {
		return err
	}
	*ct = CustomTime(t)
	return nil
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(ct).Format(time.RFC3339)), nil
}
