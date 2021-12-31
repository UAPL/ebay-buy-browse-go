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

// Aspect The type that defines the fields for the name/value pairs for the aspects of the product. For example: BRAND/Apple
type Aspect struct {
	// The text representing the name of the aspect for the name/value pair, such as Brand.
	LocalizedName *string `json:"localizedName,omitempty"`
	// The text representing the value of the aspect for the name/value pair, such as Apple.
	LocalizedValues *[]string `json:"localizedValues,omitempty"`
}

// NewAspect instantiates a new Aspect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAspect() *Aspect {
	this := Aspect{}
	return &this
}

// NewAspectWithDefaults instantiates a new Aspect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAspectWithDefaults() *Aspect {
	this := Aspect{}
	return &this
}

// GetLocalizedName returns the LocalizedName field value if set, zero value otherwise.
func (o *Aspect) GetLocalizedName() string {
	if o == nil || o.LocalizedName == nil {
		var ret string
		return ret
	}
	return *o.LocalizedName
}

// GetLocalizedNameOk returns a tuple with the LocalizedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aspect) GetLocalizedNameOk() (*string, bool) {
	if o == nil || o.LocalizedName == nil {
		return nil, false
	}
	return o.LocalizedName, true
}

// HasLocalizedName returns a boolean if a field has been set.
func (o *Aspect) HasLocalizedName() bool {
	if o != nil && o.LocalizedName != nil {
		return true
	}

	return false
}

// SetLocalizedName gets a reference to the given string and assigns it to the LocalizedName field.
func (o *Aspect) SetLocalizedName(v string) {
	o.LocalizedName = &v
}

// GetLocalizedValues returns the LocalizedValues field value if set, zero value otherwise.
func (o *Aspect) GetLocalizedValues() []string {
	if o == nil || o.LocalizedValues == nil {
		var ret []string
		return ret
	}
	return *o.LocalizedValues
}

// GetLocalizedValuesOk returns a tuple with the LocalizedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aspect) GetLocalizedValuesOk() (*[]string, bool) {
	if o == nil || o.LocalizedValues == nil {
		return nil, false
	}
	return o.LocalizedValues, true
}

// HasLocalizedValues returns a boolean if a field has been set.
func (o *Aspect) HasLocalizedValues() bool {
	if o != nil && o.LocalizedValues != nil {
		return true
	}

	return false
}

// SetLocalizedValues gets a reference to the given []string and assigns it to the LocalizedValues field.
func (o *Aspect) SetLocalizedValues(v []string) {
	o.LocalizedValues = &v
}

func (o Aspect) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LocalizedName != nil {
		toSerialize["localizedName"] = o.LocalizedName
	}
	if o.LocalizedValues != nil {
		toSerialize["localizedValues"] = o.LocalizedValues
	}
	return json.Marshal(toSerialize)
}

type NullableAspect struct {
	value *Aspect
	isSet bool
}

func (v NullableAspect) Get() *Aspect {
	return v.value
}

func (v *NullableAspect) Set(val *Aspect) {
	v.value = val
	v.isSet = true
}

func (v NullableAspect) IsSet() bool {
	return v.isSet
}

func (v *NullableAspect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAspect(val *Aspect) *NullableAspect {
	return &NullableAspect{value: val, isSet: true}
}

func (v NullableAspect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAspect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


