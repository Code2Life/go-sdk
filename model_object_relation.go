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

// ObjectRelation struct for ObjectRelation
type ObjectRelation struct {
	Object   *string `json:"object,omitempty"yaml:"object,omitempty"`
	Relation *string `json:"relation,omitempty"yaml:"relation,omitempty"`
}

// NewObjectRelation instantiates a new ObjectRelation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectRelation() *ObjectRelation {
	this := ObjectRelation{}
	return &this
}

// NewObjectRelationWithDefaults instantiates a new ObjectRelation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectRelationWithDefaults() *ObjectRelation {
	this := ObjectRelation{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *ObjectRelation) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectRelation) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *ObjectRelation) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *ObjectRelation) SetObject(v string) {
	o.Object = &v
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *ObjectRelation) GetRelation() string {
	if o == nil || o.Relation == nil {
		var ret string
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectRelation) GetRelationOk() (*string, bool) {
	if o == nil || o.Relation == nil {
		return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *ObjectRelation) HasRelation() bool {
	if o != nil && o.Relation != nil {
		return true
	}

	return false
}

// SetRelation gets a reference to the given string and assigns it to the Relation field.
func (o *ObjectRelation) SetRelation(v string) {
	o.Relation = &v
}

func (o ObjectRelation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Relation != nil {
		toSerialize["relation"] = o.Relation
	}
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	err := enc.Encode(toSerialize)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type NullableObjectRelation struct {
	value *ObjectRelation
	isSet bool
}

func (v NullableObjectRelation) Get() *ObjectRelation {
	return v.value
}

func (v *NullableObjectRelation) Set(val *ObjectRelation) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectRelation) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectRelation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectRelation(val *ObjectRelation) *NullableObjectRelation {
	return &NullableObjectRelation{value: val, isSet: true}
}

func (v NullableObjectRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectRelation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
