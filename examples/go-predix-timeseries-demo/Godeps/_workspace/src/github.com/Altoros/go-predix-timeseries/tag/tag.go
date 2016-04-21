package tag

import (
	"encoding/json"
	"errors"
	"regexp"
	"time"

	"github.com/Altoros/go-predix-timeseries/datapoint"
	"github.com/Altoros/go-predix-timeseries/dataquality"
	"github.com/Altoros/go-predix-timeseries/measurement"
)

var IncorrectName = errors.New("Tag name can contain only alphanumeric characters, periods (.), slashes (/), dashes (-), and underscores (_).")
var IncorrectAttribute = errors.New("Attributes can contain only alphanumeric characters, periods (.), slashes (/), dashes (-), and underscores (_).")
var CorrectNameRE = regexp.MustCompile("^[\\w-./]+$")

type Tag struct {
	name       string
	datapoints []datapoint.Datapoint
	attributes map[string]string
}

func (t *Tag) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name       string                `json:"name"`
		Datapoints []datapoint.Datapoint `json:"datapoints"`
		Attributes map[string]string     `json:"attributes,omitempty"`
	}{
		Name:       t.name,
		Datapoints: t.datapoints,
		Attributes: t.attributes,
	})
}

func (t *Tag) AddDatapoint(measure measurement.Measurement, quality dataquality.Quality) {
	dp := datapoint.Datapoint{Measure: measure, Timestamp: time.Now(), Quality: quality}
	t.datapoints = append(t.datapoints, dp)
}

func (t *Tag) SetAttribute(attribute, value string) error {
	if CorrectNameRE.MatchString(attribute) && CorrectNameRE.MatchString(value) {
		t.attributes[attribute] = value
		return nil
	} else {
		return IncorrectAttribute
	}
}

func New(name string) *Tag {
	return &Tag{name: name, attributes: make(map[string]string)}
}
