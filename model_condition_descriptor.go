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

// checks if the ConditionDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConditionDescriptor{}

// ConditionDescriptor This type displays additional information about the condition of an item in a structured format.
type ConditionDescriptor struct {
	// The name of a condition descriptor. The value(s) for this condition descriptor is returned in the associated <b>values</b> array.
	Name *string `json:"name,omitempty"`
	// This array displays the value(s) for a condition descriptor (denoted by the associated <b>name</b> field), as well as any other additional information about the condition of the item.
	Values []ConditionDescriptorValue `json:"values,omitempty"`
}

// NewConditionDescriptor instantiates a new ConditionDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionDescriptor() *ConditionDescriptor {
	this := ConditionDescriptor{}
	return &this
}

// NewConditionDescriptorWithDefaults instantiates a new ConditionDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionDescriptorWithDefaults() *ConditionDescriptor {
	this := ConditionDescriptor{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConditionDescriptor) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionDescriptor) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConditionDescriptor) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConditionDescriptor) SetName(v string) {
	o.Name = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ConditionDescriptor) GetValues() []ConditionDescriptorValue {
	if o == nil || IsNil(o.Values) {
		var ret []ConditionDescriptorValue
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionDescriptor) GetValuesOk() ([]ConditionDescriptorValue, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ConditionDescriptor) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []ConditionDescriptorValue and assigns it to the Values field.
func (o *ConditionDescriptor) SetValues(v []ConditionDescriptorValue) {
	o.Values = v
}

func (o ConditionDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableConditionDescriptor struct {
	value *ConditionDescriptor
	isSet bool
}

func (v NullableConditionDescriptor) Get() *ConditionDescriptor {
	return v.value
}

func (v *NullableConditionDescriptor) Set(val *ConditionDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionDescriptor(val *ConditionDescriptor) *NullableConditionDescriptor {
	return &NullableConditionDescriptor{value: val, isSet: true}
}

func (v NullableConditionDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


