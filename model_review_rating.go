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

// ReviewRating The type that defines the fields for the rating of a product review.
type ReviewRating struct {
	// The average rating given to a product based on customer reviews.
	AverageRating *string `json:"averageRating,omitempty"`
	// An array of containers for the product rating histograms that shows the review counts and the product rating.
	RatingHistograms *[]RatingHistogram `json:"ratingHistograms,omitempty"`
	// The total number of reviews for the item.
	ReviewCount *int32 `json:"reviewCount,omitempty"`
}

// NewReviewRating instantiates a new ReviewRating object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewRating() *ReviewRating {
	this := ReviewRating{}
	return &this
}

// NewReviewRatingWithDefaults instantiates a new ReviewRating object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewRatingWithDefaults() *ReviewRating {
	this := ReviewRating{}
	return &this
}

// GetAverageRating returns the AverageRating field value if set, zero value otherwise.
func (o *ReviewRating) GetAverageRating() string {
	if o == nil || o.AverageRating == nil {
		var ret string
		return ret
	}
	return *o.AverageRating
}

// GetAverageRatingOk returns a tuple with the AverageRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewRating) GetAverageRatingOk() (*string, bool) {
	if o == nil || o.AverageRating == nil {
		return nil, false
	}
	return o.AverageRating, true
}

// HasAverageRating returns a boolean if a field has been set.
func (o *ReviewRating) HasAverageRating() bool {
	if o != nil && o.AverageRating != nil {
		return true
	}

	return false
}

// SetAverageRating gets a reference to the given string and assigns it to the AverageRating field.
func (o *ReviewRating) SetAverageRating(v string) {
	o.AverageRating = &v
}

// GetRatingHistograms returns the RatingHistograms field value if set, zero value otherwise.
func (o *ReviewRating) GetRatingHistograms() []RatingHistogram {
	if o == nil || o.RatingHistograms == nil {
		var ret []RatingHistogram
		return ret
	}
	return *o.RatingHistograms
}

// GetRatingHistogramsOk returns a tuple with the RatingHistograms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewRating) GetRatingHistogramsOk() (*[]RatingHistogram, bool) {
	if o == nil || o.RatingHistograms == nil {
		return nil, false
	}
	return o.RatingHistograms, true
}

// HasRatingHistograms returns a boolean if a field has been set.
func (o *ReviewRating) HasRatingHistograms() bool {
	if o != nil && o.RatingHistograms != nil {
		return true
	}

	return false
}

// SetRatingHistograms gets a reference to the given []RatingHistogram and assigns it to the RatingHistograms field.
func (o *ReviewRating) SetRatingHistograms(v []RatingHistogram) {
	o.RatingHistograms = &v
}

// GetReviewCount returns the ReviewCount field value if set, zero value otherwise.
func (o *ReviewRating) GetReviewCount() int32 {
	if o == nil || o.ReviewCount == nil {
		var ret int32
		return ret
	}
	return *o.ReviewCount
}

// GetReviewCountOk returns a tuple with the ReviewCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewRating) GetReviewCountOk() (*int32, bool) {
	if o == nil || o.ReviewCount == nil {
		return nil, false
	}
	return o.ReviewCount, true
}

// HasReviewCount returns a boolean if a field has been set.
func (o *ReviewRating) HasReviewCount() bool {
	if o != nil && o.ReviewCount != nil {
		return true
	}

	return false
}

// SetReviewCount gets a reference to the given int32 and assigns it to the ReviewCount field.
func (o *ReviewRating) SetReviewCount(v int32) {
	o.ReviewCount = &v
}

func (o ReviewRating) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AverageRating != nil {
		toSerialize["averageRating"] = o.AverageRating
	}
	if o.RatingHistograms != nil {
		toSerialize["ratingHistograms"] = o.RatingHistograms
	}
	if o.ReviewCount != nil {
		toSerialize["reviewCount"] = o.ReviewCount
	}
	return json.Marshal(toSerialize)
}

type NullableReviewRating struct {
	value *ReviewRating
	isSet bool
}

func (v NullableReviewRating) Get() *ReviewRating {
	return v.value
}

func (v *NullableReviewRating) Set(val *ReviewRating) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewRating) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewRating) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewRating(val *ReviewRating) *NullableReviewRating {
	return &NullableReviewRating{value: val, isSet: true}
}

func (v NullableReviewRating) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewRating) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


