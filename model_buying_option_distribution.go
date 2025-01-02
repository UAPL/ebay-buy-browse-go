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

// checks if the BuyingOptionDistribution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuyingOptionDistribution{}

// BuyingOptionDistribution The container that defines the fields for the buying options refinements. This container is returned when <code>fieldgroups</code> is set to <code>BUYING_OPTION_REFINEMENTS</code> or <code>FULL</code> in the request.
type BuyingOptionDistribution struct {
	// The container that returns the buying option type. This will be AUCTION, FIXED_PRICE, CLASSIFIED_AD, or a combination of these options. For details, see <a href=\"/api-docs/buy/browse/resources/item_summary/methods/search#response.itemSummaries.buyingOptions\">buyingOptions</a>.
	BuyingOption *string `json:"buyingOption,omitempty"`
	// The number of items having this buying option.
	MatchCount *int32 `json:"matchCount,omitempty"`
	// The HATEOAS reference for this buying option.
	RefinementHref *string `json:"refinementHref,omitempty"`
}

// NewBuyingOptionDistribution instantiates a new BuyingOptionDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuyingOptionDistribution() *BuyingOptionDistribution {
	this := BuyingOptionDistribution{}
	return &this
}

// NewBuyingOptionDistributionWithDefaults instantiates a new BuyingOptionDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuyingOptionDistributionWithDefaults() *BuyingOptionDistribution {
	this := BuyingOptionDistribution{}
	return &this
}

// GetBuyingOption returns the BuyingOption field value if set, zero value otherwise.
func (o *BuyingOptionDistribution) GetBuyingOption() string {
	if o == nil || IsNil(o.BuyingOption) {
		var ret string
		return ret
	}
	return *o.BuyingOption
}

// GetBuyingOptionOk returns a tuple with the BuyingOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyingOptionDistribution) GetBuyingOptionOk() (*string, bool) {
	if o == nil || IsNil(o.BuyingOption) {
		return nil, false
	}
	return o.BuyingOption, true
}

// HasBuyingOption returns a boolean if a field has been set.
func (o *BuyingOptionDistribution) HasBuyingOption() bool {
	if o != nil && !IsNil(o.BuyingOption) {
		return true
	}

	return false
}

// SetBuyingOption gets a reference to the given string and assigns it to the BuyingOption field.
func (o *BuyingOptionDistribution) SetBuyingOption(v string) {
	o.BuyingOption = &v
}

// GetMatchCount returns the MatchCount field value if set, zero value otherwise.
func (o *BuyingOptionDistribution) GetMatchCount() int32 {
	if o == nil || IsNil(o.MatchCount) {
		var ret int32
		return ret
	}
	return *o.MatchCount
}

// GetMatchCountOk returns a tuple with the MatchCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyingOptionDistribution) GetMatchCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MatchCount) {
		return nil, false
	}
	return o.MatchCount, true
}

// HasMatchCount returns a boolean if a field has been set.
func (o *BuyingOptionDistribution) HasMatchCount() bool {
	if o != nil && !IsNil(o.MatchCount) {
		return true
	}

	return false
}

// SetMatchCount gets a reference to the given int32 and assigns it to the MatchCount field.
func (o *BuyingOptionDistribution) SetMatchCount(v int32) {
	o.MatchCount = &v
}

// GetRefinementHref returns the RefinementHref field value if set, zero value otherwise.
func (o *BuyingOptionDistribution) GetRefinementHref() string {
	if o == nil || IsNil(o.RefinementHref) {
		var ret string
		return ret
	}
	return *o.RefinementHref
}

// GetRefinementHrefOk returns a tuple with the RefinementHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuyingOptionDistribution) GetRefinementHrefOk() (*string, bool) {
	if o == nil || IsNil(o.RefinementHref) {
		return nil, false
	}
	return o.RefinementHref, true
}

// HasRefinementHref returns a boolean if a field has been set.
func (o *BuyingOptionDistribution) HasRefinementHref() bool {
	if o != nil && !IsNil(o.RefinementHref) {
		return true
	}

	return false
}

// SetRefinementHref gets a reference to the given string and assigns it to the RefinementHref field.
func (o *BuyingOptionDistribution) SetRefinementHref(v string) {
	o.RefinementHref = &v
}

func (o BuyingOptionDistribution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuyingOptionDistribution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuyingOption) {
		toSerialize["buyingOption"] = o.BuyingOption
	}
	if !IsNil(o.MatchCount) {
		toSerialize["matchCount"] = o.MatchCount
	}
	if !IsNil(o.RefinementHref) {
		toSerialize["refinementHref"] = o.RefinementHref
	}
	return toSerialize, nil
}

type NullableBuyingOptionDistribution struct {
	value *BuyingOptionDistribution
	isSet bool
}

func (v NullableBuyingOptionDistribution) Get() *BuyingOptionDistribution {
	return v.value
}

func (v *NullableBuyingOptionDistribution) Set(val *BuyingOptionDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullableBuyingOptionDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullableBuyingOptionDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuyingOptionDistribution(val *BuyingOptionDistribution) *NullableBuyingOptionDistribution {
	return &NullableBuyingOptionDistribution{value: val, isSet: true}
}

func (v NullableBuyingOptionDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuyingOptionDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


