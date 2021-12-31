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

// PickupOptionSummary The type that defines the fields for the local pickup options that are available for the item. It is used by the <b>  pickupOptions</b>  container.
type PickupOptionSummary struct {
	// This container returns the local pickup options available to the buyer. Possible values are <code>ARRANGED_LOCATION</code> and <code>STORE</code>.
	PickupLocationType *string `json:"pickupLocationType,omitempty"`
}

// NewPickupOptionSummary instantiates a new PickupOptionSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPickupOptionSummary() *PickupOptionSummary {
	this := PickupOptionSummary{}
	return &this
}

// NewPickupOptionSummaryWithDefaults instantiates a new PickupOptionSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPickupOptionSummaryWithDefaults() *PickupOptionSummary {
	this := PickupOptionSummary{}
	return &this
}

// GetPickupLocationType returns the PickupLocationType field value if set, zero value otherwise.
func (o *PickupOptionSummary) GetPickupLocationType() string {
	if o == nil || o.PickupLocationType == nil {
		var ret string
		return ret
	}
	return *o.PickupLocationType
}

// GetPickupLocationTypeOk returns a tuple with the PickupLocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PickupOptionSummary) GetPickupLocationTypeOk() (*string, bool) {
	if o == nil || o.PickupLocationType == nil {
		return nil, false
	}
	return o.PickupLocationType, true
}

// HasPickupLocationType returns a boolean if a field has been set.
func (o *PickupOptionSummary) HasPickupLocationType() bool {
	if o != nil && o.PickupLocationType != nil {
		return true
	}

	return false
}

// SetPickupLocationType gets a reference to the given string and assigns it to the PickupLocationType field.
func (o *PickupOptionSummary) SetPickupLocationType(v string) {
	o.PickupLocationType = &v
}

func (o PickupOptionSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PickupLocationType != nil {
		toSerialize["pickupLocationType"] = o.PickupLocationType
	}
	return json.Marshal(toSerialize)
}

type NullablePickupOptionSummary struct {
	value *PickupOptionSummary
	isSet bool
}

func (v NullablePickupOptionSummary) Get() *PickupOptionSummary {
	return v.value
}

func (v *NullablePickupOptionSummary) Set(val *PickupOptionSummary) {
	v.value = val
	v.isSet = true
}

func (v NullablePickupOptionSummary) IsSet() bool {
	return v.isSet
}

func (v *NullablePickupOptionSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePickupOptionSummary(val *PickupOptionSummary) *NullablePickupOptionSummary {
	return &NullablePickupOptionSummary{value: val, isSet: true}
}

func (v NullablePickupOptionSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePickupOptionSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


