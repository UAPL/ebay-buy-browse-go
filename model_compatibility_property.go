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

// checks if the CompatibilityProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompatibilityProperty{}

// CompatibilityProperty This container returns the product attribute name/value pairs that are compatible with the keyword. These attributes are submitted in the <code>compatibility_filter</code> request field.
type CompatibilityProperty struct {
	// The name of the product attribute that as been translated to the language of the site.
	LocalizedName *string `json:"localizedName,omitempty"`
	// The name of the product attribute, such as Make, Model, Year, etc.
	Name *string `json:"name,omitempty"`
	// The value for the <code>name</code> attribute, such as <b>BMW</b>, <b>R1200GS</b>, <b>2011</b>, etc.
	Value *string `json:"value,omitempty"`
}

// NewCompatibilityProperty instantiates a new CompatibilityProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompatibilityProperty() *CompatibilityProperty {
	this := CompatibilityProperty{}
	return &this
}

// NewCompatibilityPropertyWithDefaults instantiates a new CompatibilityProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompatibilityPropertyWithDefaults() *CompatibilityProperty {
	this := CompatibilityProperty{}
	return &this
}

// GetLocalizedName returns the LocalizedName field value if set, zero value otherwise.
func (o *CompatibilityProperty) GetLocalizedName() string {
	if o == nil || IsNil(o.LocalizedName) {
		var ret string
		return ret
	}
	return *o.LocalizedName
}

// GetLocalizedNameOk returns a tuple with the LocalizedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompatibilityProperty) GetLocalizedNameOk() (*string, bool) {
	if o == nil || IsNil(o.LocalizedName) {
		return nil, false
	}
	return o.LocalizedName, true
}

// HasLocalizedName returns a boolean if a field has been set.
func (o *CompatibilityProperty) HasLocalizedName() bool {
	if o != nil && !IsNil(o.LocalizedName) {
		return true
	}

	return false
}

// SetLocalizedName gets a reference to the given string and assigns it to the LocalizedName field.
func (o *CompatibilityProperty) SetLocalizedName(v string) {
	o.LocalizedName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CompatibilityProperty) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompatibilityProperty) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CompatibilityProperty) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CompatibilityProperty) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CompatibilityProperty) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompatibilityProperty) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CompatibilityProperty) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CompatibilityProperty) SetValue(v string) {
	o.Value = &v
}

func (o CompatibilityProperty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompatibilityProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LocalizedName) {
		toSerialize["localizedName"] = o.LocalizedName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCompatibilityProperty struct {
	value *CompatibilityProperty
	isSet bool
}

func (v NullableCompatibilityProperty) Get() *CompatibilityProperty {
	return v.value
}

func (v *NullableCompatibilityProperty) Set(val *CompatibilityProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableCompatibilityProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableCompatibilityProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompatibilityProperty(val *CompatibilityProperty) *NullableCompatibilityProperty {
	return &NullableCompatibilityProperty{value: val, isSet: true}
}

func (v NullableCompatibilityProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompatibilityProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


