/*
Browse API

<p>The Browse API has the following resources:</p>   <ul> <li><b> item_summary: </b> Lets shoppers search for specific items by keyword, GTIN, category, charity, product, or item aspects and refine the results by using filters, such as aspects, compatibility, and fields values.</li>  <li><b> search_by_image: </b><a href=\"https://developer.ebay.com/api-docs/static/versioning.html#experimental\" target=\"_blank\"><img src=\"/cms/img/docs/experimental-icon.svg\" class=\"legend-icon experimental-icon\" alt=\"Experimental Release\" title=\"Experimental Release\" />&nbsp;(Experimental)</a> Lets shoppers search for specific items by image. You can refine the results by using URI parameters and filters.</li>   <li><b> item: </b> <ul><li>Lets you retrieve the details of a specific item or all the items in an item group, which is an item with variations such as color and size and check if a product is compatible with the specified item, such as if a specific car is compatible with a specific part.</li> <li>Provides a bridge between the eBay legacy APIs, such as <b> Finding</b>, and the RESTful APIs, which use different formats for the item IDs.</li>  </ul> </li>  <li> <b> shopping_cart: </b> <a href=\"https://developer.ebay.com/api-docs/static/versioning.html#experimental\" target=\"_blank\"><img src=\"/cms/img/docs/experimental-icon.svg\" class=\"legend-icon experimental-icon\" alt=\"Experimental Release\" title=\"Experimental Release\" />&nbsp;(Experimental)</a> <a href=\"https://developer.ebay.com/api-docs/static/versioning.html#limited\" target=\"_blank\"> <img src=\"/cms/img/docs/partners-api.svg\" class=\"legend-icon partners-icon\" title=\"Limited Release\"  alt=\"Limited Release\" />(Limited Release)</a> Provides the ability for eBay members to see the contents of their eBay cart, and add, remove, and change the quantity of items in their eBay cart.&nbsp;&nbsp;<b> Note: </b> This resource is not available in the eBay API Explorer.</li></ul>       <p>The <b> item_summary</b>, <b> search_by_image</b>, and <b> item</b> resource calls require an <a href=\"/api-docs/static/oauth-client-credentials-grant.html\">Application access token</a>. The <b> shopping_cart</b> resource calls require a <a href=\"/api-docs/static/oauth-authorization-code-grant.html\">User access token</a>.</p>

API version: v1.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package buybrowse

import (
	"encoding/json"
)

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
	if o == nil || o.UnitOfMeasure == nil {
		var ret string
		return ret
	}
	return *o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetLocation) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil || o.UnitOfMeasure == nil {
		return nil, false
	}
	return o.UnitOfMeasure, true
}

// HasUnitOfMeasure returns a boolean if a field has been set.
func (o *TargetLocation) HasUnitOfMeasure() bool {
	if o != nil && o.UnitOfMeasure != nil {
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
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetLocation) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TargetLocation) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TargetLocation) SetValue(v string) {
	o.Value = &v
}

func (o TargetLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnitOfMeasure != nil {
		toSerialize["unitOfMeasure"] = o.UnitOfMeasure
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
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

