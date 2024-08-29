// Code generated by github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli, DO NOT EDIT.

package notstreams

import (
	"encoding/json"
	"fmt"
)

type Feed struct {
	// Object containing two benchmark prices extracted from the fullReport.
	Price FeedPrice `json:"Price" yaml:"Price" mapstructure:"Price"`

	// This value is extracted from the fullReport. A unix timestamp represented as an
	// int64 value. Timestamp is captured at the time of report creation.
	Timestamp int `json:"Timestamp" yaml:"Timestamp" mapstructure:"Timestamp"`

	// Full report represented as bytes encoded as base64 string.
	FullReport string `json:"fullReport" yaml:"fullReport" mapstructure:"fullReport"`

	// Report context represented as bytes encoded as base64 string. This is required
	// to validate the signatures.
	ReportContext string `json:"reportContext" yaml:"reportContext" mapstructure:"reportContext"`

	// Signature over full report and report context represented as bytes encoded as
	// base64 string.
	Signatures []string `json:"signatures" yaml:"signatures" mapstructure:"signatures"`
}

// Object containing two benchmark prices extracted from the fullReport.
type FeedPrice struct {
	// Benchmark price A represented as bytes encoded as base64 string.
	PriceA string `json:"PriceA" yaml:"PriceA" mapstructure:"PriceA"`

	// Benchmark price B represented as bytes encoded as base64 string.
	PriceB string `json:"PriceB" yaml:"PriceB" mapstructure:"PriceB"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FeedPrice) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["PriceA"]; raw != nil && !ok {
		return fmt.Errorf("field PriceA in FeedPrice: required")
	}
	if _, ok := raw["PriceB"]; raw != nil && !ok {
		return fmt.Errorf("field PriceB in FeedPrice: required")
	}
	type Plain FeedPrice
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = FeedPrice(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Feed) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["Price"]; raw != nil && !ok {
		return fmt.Errorf("field Price in Feed: required")
	}
	if _, ok := raw["Timestamp"]; raw != nil && !ok {
		return fmt.Errorf("field Timestamp in Feed: required")
	}
	if _, ok := raw["fullReport"]; raw != nil && !ok {
		return fmt.Errorf("field fullReport in Feed: required")
	}
	if _, ok := raw["reportContext"]; raw != nil && !ok {
		return fmt.Errorf("field reportContext in Feed: required")
	}
	if _, ok := raw["signatures"]; raw != nil && !ok {
		return fmt.Errorf("field signatures in Feed: required")
	}
	type Plain Feed
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if plain.Signatures != nil && len(plain.Signatures) < 1 {
		return fmt.Errorf("field %s length: must be >= %d", "signatures", 1)
	}
	*j = Feed(plain)
	return nil
}

// Not Streams Trigger
type Trigger struct {
	// Config corresponds to the JSON schema field "config".
	Config TriggerConfig `json:"config" yaml:"config" mapstructure:"config"`

	// Outputs corresponds to the JSON schema field "outputs".
	Outputs *Feed `json:"outputs,omitempty" yaml:"outputs,omitempty" mapstructure:"outputs,omitempty"`
}

type TriggerConfig struct {
	// The interval in seconds after which a new trigger event is generated.
	MaxFrequencyMs int `json:"maxFrequencyMs" yaml:"maxFrequencyMs" mapstructure:"maxFrequencyMs"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TriggerConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["maxFrequencyMs"]; raw != nil && !ok {
		return fmt.Errorf("field maxFrequencyMs in TriggerConfig: required")
	}
	type Plain TriggerConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if 1 > plain.MaxFrequencyMs {
		return fmt.Errorf("field %s: must be >= %v", "maxFrequencyMs", 1)
	}
	*j = TriggerConfig(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Trigger) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["config"]; raw != nil && !ok {
		return fmt.Errorf("field config in Trigger: required")
	}
	type Plain Trigger
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Trigger(plain)
	return nil
}