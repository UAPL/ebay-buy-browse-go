# ReviewRating

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageRating** | Pointer to **string** | The average rating given to a product based on customer reviews. | [optional] 
**RatingHistograms** | Pointer to [**[]RatingHistogram**](RatingHistogram.md) | An array of containers for the product rating histograms that shows the review counts and the product rating. | [optional] 
**ReviewCount** | Pointer to **int32** | The total number of reviews for the item. | [optional] 

## Methods

### NewReviewRating

`func NewReviewRating() *ReviewRating`

NewReviewRating instantiates a new ReviewRating object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewRatingWithDefaults

`func NewReviewRatingWithDefaults() *ReviewRating`

NewReviewRatingWithDefaults instantiates a new ReviewRating object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageRating

`func (o *ReviewRating) GetAverageRating() string`

GetAverageRating returns the AverageRating field if non-nil, zero value otherwise.

### GetAverageRatingOk

`func (o *ReviewRating) GetAverageRatingOk() (*string, bool)`

GetAverageRatingOk returns a tuple with the AverageRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRating

`func (o *ReviewRating) SetAverageRating(v string)`

SetAverageRating sets AverageRating field to given value.

### HasAverageRating

`func (o *ReviewRating) HasAverageRating() bool`

HasAverageRating returns a boolean if a field has been set.

### GetRatingHistograms

`func (o *ReviewRating) GetRatingHistograms() []RatingHistogram`

GetRatingHistograms returns the RatingHistograms field if non-nil, zero value otherwise.

### GetRatingHistogramsOk

`func (o *ReviewRating) GetRatingHistogramsOk() (*[]RatingHistogram, bool)`

GetRatingHistogramsOk returns a tuple with the RatingHistograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingHistograms

`func (o *ReviewRating) SetRatingHistograms(v []RatingHistogram)`

SetRatingHistograms sets RatingHistograms field to given value.

### HasRatingHistograms

`func (o *ReviewRating) HasRatingHistograms() bool`

HasRatingHistograms returns a boolean if a field has been set.

### GetReviewCount

`func (o *ReviewRating) GetReviewCount() int32`

GetReviewCount returns the ReviewCount field if non-nil, zero value otherwise.

### GetReviewCountOk

`func (o *ReviewRating) GetReviewCountOk() (*int32, bool)`

GetReviewCountOk returns a tuple with the ReviewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewCount

`func (o *ReviewRating) SetReviewCount(v int32)`

SetReviewCount sets ReviewCount field to given value.

### HasReviewCount

`func (o *ReviewRating) HasReviewCount() bool`

HasReviewCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


