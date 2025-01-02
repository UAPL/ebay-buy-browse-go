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

// checks if the ConditionDescriptorValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConditionDescriptorValue{}

// ConditionDescriptorValue This type displays the value(s) associated with the specified condition descriptor name, as well as any additional information about a condition descriptor. 
type ConditionDescriptorValue struct {
	// Additional information about the condition of an item as it relates to a condition descriptor. This array elaborates on the value specified in the <b>content</b> field and provides additional details about the condition of an item.
	AdditionalInfo []string `json:"additionalInfo,omitempty"`
	// The value for the condition descriptor indicated in the associated <b>name</b> field.
	Content *string `json:"content,omitempty"`
}

// NewConditionDescriptorValue instantiates a new ConditionDescriptorValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionDescriptorValue() *ConditionDescriptorValue {
	this := ConditionDescriptorValue{}
	return &this
}

// NewConditionDescriptorValueWithDefaults instantiates a new ConditionDescriptorValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionDescriptorValueWithDefaults() *ConditionDescriptorValue {
	this := ConditionDescriptorValue{}
	return &this
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *ConditionDescriptorValue) GetAdditionalInfo() []string {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret []string
		return ret
	}
	return o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionDescriptorValue) GetAdditionalInfoOk() ([]string, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *ConditionDescriptorValue) HasAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given []string and assigns it to the AdditionalInfo field.
func (o *ConditionDescriptorValue) SetAdditionalInfo(v []string) {
	o.AdditionalInfo = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ConditionDescriptorValue) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionDescriptorValue) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ConditionDescriptorValue) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ConditionDescriptorValue) SetContent(v string) {
	o.Content = &v
}

func (o ConditionDescriptorValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionDescriptorValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

type NullableConditionDescriptorValue struct {
	value *ConditionDescriptorValue
	isSet bool
}

func (v NullableConditionDescriptorValue) Get() *ConditionDescriptorValue {
	return v.value
}

func (v *NullableConditionDescriptorValue) Set(val *ConditionDescriptorValue) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionDescriptorValue) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionDescriptorValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionDescriptorValue(val *ConditionDescriptorValue) *NullableConditionDescriptorValue {
	return &NullableConditionDescriptorValue{value: val, isSet: true}
}

func (v NullableConditionDescriptorValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionDescriptorValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

