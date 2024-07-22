/*
Daytona Server API

Daytona Server API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the ExistingProjectConfigDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExistingProjectConfigDTO{}

// ExistingProjectConfigDTO struct for ExistingProjectConfigDTO
type ExistingProjectConfigDTO struct {
	Branch *string `json:"branch,omitempty"`
	Name   *string `json:"name,omitempty"`
}

// NewExistingProjectConfigDTO instantiates a new ExistingProjectConfigDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExistingProjectConfigDTO() *ExistingProjectConfigDTO {
	this := ExistingProjectConfigDTO{}
	return &this
}

// NewExistingProjectConfigDTOWithDefaults instantiates a new ExistingProjectConfigDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExistingProjectConfigDTOWithDefaults() *ExistingProjectConfigDTO {
	this := ExistingProjectConfigDTO{}
	return &this
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *ExistingProjectConfigDTO) GetBranch() string {
	if o == nil || IsNil(o.Branch) {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExistingProjectConfigDTO) GetBranchOk() (*string, bool) {
	if o == nil || IsNil(o.Branch) {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *ExistingProjectConfigDTO) HasBranch() bool {
	if o != nil && !IsNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *ExistingProjectConfigDTO) SetBranch(v string) {
	o.Branch = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExistingProjectConfigDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExistingProjectConfigDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExistingProjectConfigDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExistingProjectConfigDTO) SetName(v string) {
	o.Name = &v
}

func (o ExistingProjectConfigDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExistingProjectConfigDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableExistingProjectConfigDTO struct {
	value *ExistingProjectConfigDTO
	isSet bool
}

func (v NullableExistingProjectConfigDTO) Get() *ExistingProjectConfigDTO {
	return v.value
}

func (v *NullableExistingProjectConfigDTO) Set(val *ExistingProjectConfigDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableExistingProjectConfigDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableExistingProjectConfigDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExistingProjectConfigDTO(val *ExistingProjectConfigDTO) *NullableExistingProjectConfigDTO {
	return &NullableExistingProjectConfigDTO{value: val, isSet: true}
}

func (v NullableExistingProjectConfigDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExistingProjectConfigDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
