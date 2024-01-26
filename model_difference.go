/**
 * Go SDK for OpenFGA
 *
 * API version: 0.1
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://discord.gg/8naAwJfWN6
 * License: [Apache-2.0](https://github.com/openfga/go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"bytes"

	"encoding/json"
)

// Difference struct for Difference
type Difference struct {
	Base     Userset `json:"base"yaml:"base"`
	Subtract Userset `json:"subtract"yaml:"subtract"`
}

// NewDifference instantiates a new Difference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDifference(base Userset, subtract Userset) *Difference {
	this := Difference{}
	this.Base = base
	this.Subtract = subtract
	return &this
}

// NewDifferenceWithDefaults instantiates a new Difference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDifferenceWithDefaults() *Difference {
	this := Difference{}
	return &this
}

// GetBase returns the Base field value
func (o *Difference) GetBase() Userset {
	if o == nil {
		var ret Userset
		return ret
	}

	return o.Base
}

// GetBaseOk returns a tuple with the Base field value
// and a boolean to check if the value has been set.
func (o *Difference) GetBaseOk() (*Userset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Base, true
}

// SetBase sets field value
func (o *Difference) SetBase(v Userset) {
	o.Base = v
}

// GetSubtract returns the Subtract field value
func (o *Difference) GetSubtract() Userset {
	if o == nil {
		var ret Userset
		return ret
	}

	return o.Subtract
}

// GetSubtractOk returns a tuple with the Subtract field value
// and a boolean to check if the value has been set.
func (o *Difference) GetSubtractOk() (*Userset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtract, true
}

// SetSubtract sets field value
func (o *Difference) SetSubtract(v Userset) {
	o.Subtract = v
}

func (o Difference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["base"] = o.Base
	toSerialize["subtract"] = o.Subtract
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableDifference struct {
	value *Difference
	isSet bool
}

func (v NullableDifference) Get() *Difference {
	return v.value
}

func (v *NullableDifference) Set(val *Difference) {
	v.value = val
	v.isSet = true
}

func (v NullableDifference) IsSet() bool {
	return v.isSet
}

func (v *NullableDifference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDifference(val *Difference) *NullableDifference {
	return &NullableDifference{value: val, isSet: true}
}

func (v NullableDifference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDifference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
