/*
Browse API

The Browse API has the following resources:<ul><li><b>item_summary:</b><br>Allows shoppers to search for specific items by keyword, GTIN, category, charity, product, image, or item aspects and refine the results by using filters, such as aspects, compatibility, and fields values, or UI parameters.</li><li><b>item:</b><br>Allows shoppers to retrieve the details of a specific item or all items in an item group, which is an item with variations such as color and size and check if a product is compatible with the specified item, such as if a specific car is compatible with a specific part.<br><br>This resource also provides a bridge between the eBay legacy APIs, such as the <a href=\"/api-docs/user-guides/static/finding-user-guide-landing.html\" target=\"_blank\">Finding</b>, and the RESTful APIs, which use different formats for the item IDs.</li></ul>The <b>item_summary</b>, <b>search_by_image</b>, and <b>item</b> resource calls require an <a href=\"/api-docs/static/oauth-client-credentials-grant.html\" target=\"_blank\">Application access token</a>.

API version: v1.19.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package buybrowse

import (
	"encoding/json"
)

// checks if the TargetLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetLocation{}

// TargetLocation The type that defines the fields for the distance between the item location and the buyer's location. 
type TargetLocation struct {
	// This value shows the unit of measurement used to measure the distance between the location of the item and the buyer's location. This value is typically <code> mi</code> or <code> km</code>.
	UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`
	// This value indicates the distance (measured in the measurement unit in the <b> unitOfMeasure</b>  field) between the item location and the buyer's location.
	Value *string `json:"value,omitempty"`
}

// NewTargetLocation instantiates a new TargetLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetLocation() *TargetLocation {
	this := TargetLocation{}
	return &this
}

// NewTargetLocationWithDefaults instantiates a new TargetLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetLocationWithDefaults() *TargetLocation {
	this := TargetLocation{}
	return &this
}

// GetUnitOfMeasure returns the UnitOfMeasure field value if set, zero value otherwise.
func (o *TargetLocation) GetUnitOfMeasure() string {
	if o == nil || IsNil(o.UnitOfMeasure) {
		var ret string
		return ret
	}
	return *o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetLocation) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil || IsNil(o.UnitOfMeasure) {
		return nil, false
	}
	return o.UnitOfMeasure, true
}

// HasUnitOfMeasure returns a boolean if a field has been set.
func (o *TargetLocation) HasUnitOfMeasure() bool {
	if o != nil && !IsNil(o.UnitOfMeasure) {
		return true
	}

	return false
}

// SetUnitOfMeasure gets a reference to the given string and assigns it to the UnitOfMeasure field.
func (o *TargetLocation) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TargetLocation) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetLocation) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TargetLocation) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TargetLocation) SetValue(v string) {
	o.Value = &v
}

func (o TargetLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UnitOfMeasure) {
		toSerialize["unitOfMeasure"] = o.UnitOfMeasure
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTargetLocation struct {
	value *TargetLocation
	isSet bool
}

func (v NullableTargetLocation) Get() *TargetLocation {
	return v.value
}

func (v *NullableTargetLocation) Set(val *TargetLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetLocation(val *TargetLocation) *NullableTargetLocation {
	return &NullableTargetLocation{value: val, isSet: true}
}

func (v NullableTargetLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


