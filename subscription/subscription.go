package subscription

import (
	"fmt"
	"time"
)

// JSONTime is present because The JSON time format expected is not the default used by go
// Let's implement the interface Marshaler and override the format
type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	const layout = "2006-01-02T15:04:05.000Z"
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(layout))
	return []byte(stamp), nil

}

// SubscriptionRequest represents the payload to be send to csa
// https:// <CSAFQDN>:<port>/csa/api/mpp/mpp-request/<offeringId>
// method: POST
// Content-Type: multipart/form-data; boundary=Abcdefgh
// Accept: application/json
type SubscriptionRequest struct {
	CatalogID               string                 `json:"catalogId"`
	CategoryName            string                 `json:"categoryName"`
	SubscriptionName        string                 `json:"subscriptionName"`
	SubscriptionDescription string                 `json:"subscriptionDescription"`
	StartDate               JSONTime               `json:"startDate"`
	EndDate                 JSONTime               `json:"endDate"`
	Fields                  map[string]interface{} `json:"fields"`
	Action                  string                 `json:"action"`
}

type InitPrice struct {
	Currency string      `json:"currency,omitempty" yaml:"-"`
	Price    interface{} `json:"price,omitempty" yaml:"-"`
}
type RecurringPrice struct {
	BasedOn  string      `json:"basedOn,omitempty" yaml:"-"`
	Currency string      `json:"currency,omitempty" yaml:"-"`
	Price    interface{} `json:"price,omitempty" yaml:"-"`
}
type SubscriptionField struct {
	Confidential   bool            `json:"confidential,omitempty" yaml:"-"`
	Description    string          `json:"description,omitempty,omitempty" yaml:"-"`
	Disabled       bool            `json:"disabled,omitempty" yaml:"-"`
	DisplayName    string          `json:"displayName,omitempty" yaml:"display_name"`
	Encrypted      bool            `json:"encrypted,omitempty" yaml:"-"`
	Hidden         bool            `json:"hidden,omitempty" yaml:"-"`
	ID             string          `json:"id,omitempty" yaml:"id"`
	InitPrice      *InitPrice      `json:"initPrice,omitempty" yaml:"-"`
	Name           string          `json:"name,omitempty" yaml:"name"`
	RecurringPrice *RecurringPrice `json:"recurringPrice,omitempty" yaml:"-"`
	Required       bool            `json:"required,omitempty" yaml:"-"`
	Value          interface{}     `json:"value,omitempty" yaml:"value"`
	Visible        bool            `json:"visible,omitempty" yaml:"-"`
	MaxValue       int             `json:"maxValue,omitempty,omitempty" yaml:"-"`
	MinValue       int             `json:"minValue,omitempty,omitempty" yaml:"-"`
	Needed         bool            `json:"-,omitempty" yaml:"needed"` // for the API

}

type SubscriptionDetail struct {
	ID               string `json:"id,omitempty" yaml:"id"`
	ApprovalRequired bool   `json:"approvalRequired,omitempty" yaml:"-"`
	CatalogID        string `json:"catalogId,omitempty" yaml:"catalog_id"`
	Category         struct {
		DisplayName string `json:"displayName,omitempty" yaml:"-"`
		Name        string `json:"name,omitempty" yaml:"name"`
	} `json:"category,omitempty" yaml:"category"`
	Description        string              `json:"description,omitempty" yaml:"Description"`
	DisplayName        string              `json:"displayName,omitempty" yaml:"-"`
	Fields             []SubscriptionField `json:"fields,omitempty" yaml:"fields"`
	HideInitialPrice   bool                `json:"hideInitialPrice,omitempty" yaml:"-"`
	HideRecurringPrice bool                `json:"hideRecurringPrice,omitempty" yaml:"-"`
	Image              string              `json:"image,omitempty" yaml:"-"`
	InitPrice          *InitPrice          `json:"initPrice,omitempty" yaml:"-"`
	Layout             []struct {
		Description  string `json:"description,omitempty" yaml:"-"`
		DisplayClass string `json:"displayClass,omitempty" yaml:"-"`
		DisplayName  string `json:"displayName,omitempty" yaml:"-"`
		DisplayType  string `json:"displayType,omitempty" yaml:"-"`
		Image        string `json:"image,omitempty" yaml:"-"`
		Layout       []struct {
			FieldID  string      `json:"fieldId,omitempty" yaml:"-"`
			ImageURL interface{} `json:"imageUrl,omitempty" yaml:"-"`
			Name     string      `json:"name,omitempty" yaml:"-"`
			Type     string      `json:"type,omitempty" yaml:"-"`
		} `json:"layout,omitempty" yaml:"-"`
		Name string `json:"name,omitempty" yaml:"-"`
		Type string `json:"type,omitempty" yaml:"-"`
	} `json:"layout,omitempty" yaml:"-"`
	Name            string          `json:"name,omitempty" yaml:"-"`
	OfferingVersion string          `json:"offeringVersion,omitempty" yaml:"-"`
	PublishedDate   time.Time       `json:"publishedDate,omitempty" yaml:"-"`
	RecurringPrice  *RecurringPrice `json:"recurringPrice,omitempty" yaml:"-"`
	State           string          `json:"state,omitempty" yaml:"-"`
}
