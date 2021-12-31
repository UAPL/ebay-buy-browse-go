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

// AspectValueDistribution The container that defines the fields for the conditions refinements. This container is returned when <b> fieldgroups</b> is set to <code>ASPECT_REFINEMENTS</code> or <code>FULL</code> in the request.
type AspectValueDistribution struct {
	// The value of an aspect. For example, Red is a value for the aspect Color.
	LocalizedAspectValue *string `json:"localizedAspectValue,omitempty"`
	// The number of items with this aspect.
	MatchCount *int32 `json:"matchCount,omitempty"`
	// A HATEOAS reference for this aspect.
	RefinementHref *string `json:"refinementHref,omitempty"`
}

// NewAspectValueDistribution instantiates a new AspectValueDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAspectValueDistribution() *AspectValueDistribution {
	this := AspectValueDistribution{}
	return &this
}

// NewAspectValueDistributionWithDefaults instantiates a new AspectValueDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAspectValueDistributionWithDefaults() *AspectValueDistribution {
	this := AspectValueDistribution{}
	return &this
}

// GetLocalizedAspectValue returns the LocalizedAspectValue field value if set, zero value otherwise.
func (o *AspectValueDistribution) GetLocalizedAspectValue() string {
	if o == nil || o.LocalizedAspectValue == nil {
		var ret string
		return ret
	}
	return *o.LocalizedAspectValue
}

// GetLocalizedAspectValueOk returns a tuple with the LocalizedAspectValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AspectValueDistribution) GetLocalizedAspectValueOk() (*string, bool) {
	if o == nil || o.LocalizedAspectValue == nil {
		return nil, false
	}
	return o.LocalizedAspectValue, true
}

// HasLocalizedAspectValue returns a boolean if a field has been set.
func (o *AspectValueDistribution) HasLocalizedAspectValue() bool {
	if o != nil && o.LocalizedAspectValue != nil {
		return true
	}

	return false
}

// SetLocalizedAspectValue gets a reference to the given string and assigns it to the LocalizedAspectValue field.
func (o *AspectValueDistribution) SetLocalizedAspectValue(v string) {
	o.LocalizedAspectValue = &v
}

// GetMatchCount returns the MatchCount field value if set, zero value otherwise.
func (o *AspectValueDistribution) GetMatchCount() int32 {
	if o == nil || o.MatchCount == nil {
		var ret int32
		return ret
	}
	return *o.MatchCount
}

// GetMatchCountOk returns a tuple with the MatchCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AspectValueDistribution) GetMatchCountOk() (*int32, bool) {
	if o == nil || o.MatchCount == nil {
		return nil, false
	}
	return o.MatchCount, true
}

// HasMatchCount returns a boolean if a field has been set.
func (o *AspectValueDistribution) HasMatchCount() bool {
	if o != nil && o.MatchCount != nil {
		return true
	}

	return false
}

// SetMatchCount gets a reference to the given int32 and assigns it to the MatchCount field.
func (o *AspectValueDistribution) SetMatchCount(v int32) {
	o.MatchCount = &v
}

// GetRefinementHref returns the RefinementHref field value if set, zero value otherwise.
func (o *AspectValueDistribution) GetRefinementHref() string {
	if o == nil || o.RefinementHref == nil {
		var ret string
		return ret
	}
	return *o.RefinementHref
}

// GetRefinementHrefOk returns a tuple with the RefinementHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AspectValueDistribution) GetRefinementHrefOk() (*string, bool) {
	if o == nil || o.RefinementHref == nil {
		return nil, false
	}
	return o.RefinementHref, true
}

// HasRefinementHref returns a boolean if a field has been set.
func (o *AspectValueDistribution) HasRefinementHref() bool {
	if o != nil && o.RefinementHref != nil {
		return true
	}

	return false
}

// SetRefinementHref gets a reference to the given string and assigns it to the RefinementHref field.
func (o *AspectValueDistribution) SetRefinementHref(v string) {
	o.RefinementHref = &v
}

func (o AspectValueDistribution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LocalizedAspectValue != nil {
		toSerialize["localizedAspectValue"] = o.LocalizedAspectValue
	}
	if o.MatchCount != nil {
		toSerialize["matchCount"] = o.MatchCount
	}
	if o.RefinementHref != nil {
		toSerialize["refinementHref"] = o.RefinementHref
	}
	return json.Marshal(toSerialize)
}

type NullableAspectValueDistribution struct {
	value *AspectValueDistribution
	isSet bool
}

func (v NullableAspectValueDistribution) Get() *AspectValueDistribution {
	return v.value
}

func (v *NullableAspectValueDistribution) Set(val *AspectValueDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullableAspectValueDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullableAspectValueDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAspectValueDistribution(val *AspectValueDistribution) *NullableAspectValueDistribution {
	return &NullableAspectValueDistribution{value: val, isSet: true}
}

func (v NullableAspectValueDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAspectValueDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


