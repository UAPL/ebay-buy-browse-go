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

// AspectDistribution The type that define the fields for the aspect information. Aspects are the variations of an item, such as color, size, etc.
type AspectDistribution struct {
	// An array of containers for the various values of the aspect and the match count and a HATEOAS reference (<b> refinementHref</b>) for this aspect.
	AspectValueDistributions *[]AspectValueDistribution `json:"aspectValueDistributions,omitempty"`
	// The name of an aspect, such as Brand, Color, etc.
	LocalizedAspectName *string `json:"localizedAspectName,omitempty"`
}

// NewAspectDistribution instantiates a new AspectDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAspectDistribution() *AspectDistribution {
	this := AspectDistribution{}
	return &this
}

// NewAspectDistributionWithDefaults instantiates a new AspectDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAspectDistributionWithDefaults() *AspectDistribution {
	this := AspectDistribution{}
	return &this
}

// GetAspectValueDistributions returns the AspectValueDistributions field value if set, zero value otherwise.
func (o *AspectDistribution) GetAspectValueDistributions() []AspectValueDistribution {
	if o == nil || o.AspectValueDistributions == nil {
		var ret []AspectValueDistribution
		return ret
	}
	return *o.AspectValueDistributions
}

// GetAspectValueDistributionsOk returns a tuple with the AspectValueDistributions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AspectDistribution) GetAspectValueDistributionsOk() (*[]AspectValueDistribution, bool) {
	if o == nil || o.AspectValueDistributions == nil {
		return nil, false
	}
	return o.AspectValueDistributions, true
}

// HasAspectValueDistributions returns a boolean if a field has been set.
func (o *AspectDistribution) HasAspectValueDistributions() bool {
	if o != nil && o.AspectValueDistributions != nil {
		return true
	}

	return false
}

// SetAspectValueDistributions gets a reference to the given []AspectValueDistribution and assigns it to the AspectValueDistributions field.
func (o *AspectDistribution) SetAspectValueDistributions(v []AspectValueDistribution) {
	o.AspectValueDistributions = &v
}

// GetLocalizedAspectName returns the LocalizedAspectName field value if set, zero value otherwise.
func (o *AspectDistribution) GetLocalizedAspectName() string {
	if o == nil || o.LocalizedAspectName == nil {
		var ret string
		return ret
	}
	return *o.LocalizedAspectName
}

// GetLocalizedAspectNameOk returns a tuple with the LocalizedAspectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AspectDistribution) GetLocalizedAspectNameOk() (*string, bool) {
	if o == nil || o.LocalizedAspectName == nil {
		return nil, false
	}
	return o.LocalizedAspectName, true
}

// HasLocalizedAspectName returns a boolean if a field has been set.
func (o *AspectDistribution) HasLocalizedAspectName() bool {
	if o != nil && o.LocalizedAspectName != nil {
		return true
	}

	return false
}

// SetLocalizedAspectName gets a reference to the given string and assigns it to the LocalizedAspectName field.
func (o *AspectDistribution) SetLocalizedAspectName(v string) {
	o.LocalizedAspectName = &v
}

func (o AspectDistribution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AspectValueDistributions != nil {
		toSerialize["aspectValueDistributions"] = o.AspectValueDistributions
	}
	if o.LocalizedAspectName != nil {
		toSerialize["localizedAspectName"] = o.LocalizedAspectName
	}
	return json.Marshal(toSerialize)
}

type NullableAspectDistribution struct {
	value *AspectDistribution
	isSet bool
}

func (v NullableAspectDistribution) Get() *AspectDistribution {
	return v.value
}

func (v *NullableAspectDistribution) Set(val *AspectDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullableAspectDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullableAspectDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAspectDistribution(val *AspectDistribution) *NullableAspectDistribution {
	return &NullableAspectDistribution{value: val, isSet: true}
}

func (v NullableAspectDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAspectDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

