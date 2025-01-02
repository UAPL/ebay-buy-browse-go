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

// checks if the Seller type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Seller{}

// Seller The type that defines the fields for basic information about the seller of the item returned by the <code>item_summary</code> resource.
type Seller struct {
	// The percentage of the total positive feedback.
	FeedbackPercentage *string `json:"feedbackPercentage,omitempty"`
	// The feedback score of the seller. This value is based on the ratings from eBay members that bought items from this seller.
	FeedbackScore *int32 `json:"feedbackScore,omitempty"`
	// Indicates if the seller is a business or an individual. This is determined when the seller registers with eBay:<ul><li>If they register for a business account, this value will be <code>BUSINESS<code>.</li><li>If they register for a private account, this value will be <code>INDIVIDUAL</code>.</li></ul>This designation is required by the tax laws in some countries.<br><br>This field is returned only on the following sites:<br><br>EBAY_AT, EBAY_BE, EBAY_CH, EBAY_DE, EBAY_ES, EBAY_FR, EBAY_GB, EBAY_IE, EBAY_IT, EBAY_PL<br><br><b>Valid Values:</b> <code>BUSINESS</code> or <code>INDIVIDUAL</code>
	SellerAccountType *string `json:"sellerAccountType,omitempty"`
	// The user name created by the seller for use on eBay.
	Username *string `json:"username,omitempty"`
}

// NewSeller instantiates a new Seller object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeller() *Seller {
	this := Seller{}
	return &this
}

// NewSellerWithDefaults instantiates a new Seller object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSellerWithDefaults() *Seller {
	this := Seller{}
	return &this
}

// GetFeedbackPercentage returns the FeedbackPercentage field value if set, zero value otherwise.
func (o *Seller) GetFeedbackPercentage() string {
	if o == nil || IsNil(o.FeedbackPercentage) {
		var ret string
		return ret
	}
	return *o.FeedbackPercentage
}

// GetFeedbackPercentageOk returns a tuple with the FeedbackPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Seller) GetFeedbackPercentageOk() (*string, bool) {
	if o == nil || IsNil(o.FeedbackPercentage) {
		return nil, false
	}
	return o.FeedbackPercentage, true
}

// HasFeedbackPercentage returns a boolean if a field has been set.
func (o *Seller) HasFeedbackPercentage() bool {
	if o != nil && !IsNil(o.FeedbackPercentage) {
		return true
	}

	return false
}

// SetFeedbackPercentage gets a reference to the given string and assigns it to the FeedbackPercentage field.
func (o *Seller) SetFeedbackPercentage(v string) {
	o.FeedbackPercentage = &v
}

// GetFeedbackScore returns the FeedbackScore field value if set, zero value otherwise.
func (o *Seller) GetFeedbackScore() int32 {
	if o == nil || IsNil(o.FeedbackScore) {
		var ret int32
		return ret
	}
	return *o.FeedbackScore
}

// GetFeedbackScoreOk returns a tuple with the FeedbackScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Seller) GetFeedbackScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.FeedbackScore) {
		return nil, false
	}
	return o.FeedbackScore, true
}

// HasFeedbackScore returns a boolean if a field has been set.
func (o *Seller) HasFeedbackScore() bool {
	if o != nil && !IsNil(o.FeedbackScore) {
		return true
	}

	return false
}

// SetFeedbackScore gets a reference to the given int32 and assigns it to the FeedbackScore field.
func (o *Seller) SetFeedbackScore(v int32) {
	o.FeedbackScore = &v
}

// GetSellerAccountType returns the SellerAccountType field value if set, zero value otherwise.
func (o *Seller) GetSellerAccountType() string {
	if o == nil || IsNil(o.SellerAccountType) {
		var ret string
		return ret
	}
	return *o.SellerAccountType
}

// GetSellerAccountTypeOk returns a tuple with the SellerAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Seller) GetSellerAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SellerAccountType) {
		return nil, false
	}
	return o.SellerAccountType, true
}

// HasSellerAccountType returns a boolean if a field has been set.
func (o *Seller) HasSellerAccountType() bool {
	if o != nil && !IsNil(o.SellerAccountType) {
		return true
	}

	return false
}

// SetSellerAccountType gets a reference to the given string and assigns it to the SellerAccountType field.
func (o *Seller) SetSellerAccountType(v string) {
	o.SellerAccountType = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *Seller) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Seller) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *Seller) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *Seller) SetUsername(v string) {
	o.Username = &v
}

func (o Seller) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Seller) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeedbackPercentage) {
		toSerialize["feedbackPercentage"] = o.FeedbackPercentage
	}
	if !IsNil(o.FeedbackScore) {
		toSerialize["feedbackScore"] = o.FeedbackScore
	}
	if !IsNil(o.SellerAccountType) {
		toSerialize["sellerAccountType"] = o.SellerAccountType
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableSeller struct {
	value *Seller
	isSet bool
}

func (v NullableSeller) Get() *Seller {
	return v.value
}

func (v *NullableSeller) Set(val *Seller) {
	v.value = val
	v.isSet = true
}

func (v NullableSeller) IsSet() bool {
	return v.isSet
}

func (v *NullableSeller) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeller(val *Seller) *NullableSeller {
	return &NullableSeller{value: val, isSet: true}
}

func (v NullableSeller) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeller) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


