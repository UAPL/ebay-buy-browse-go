# RatingHistogram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of user ratings that the product has received. | [optional] 
**Rating** | Pointer to **string** | This is the average rating for the product. As part of a product review, users rate the product. Products are rated from one star (terrible) to five stars (excellent), with each star having a corresponding point value - one star gets 1 point, two stars get 2 points, and so on. If a product had one four-star rating and one five-star rating, its average rating would be &lt;code&gt; 4.5&lt;/code&gt;, and this is the value that would appear in this field. | [optional] 

## Methods

### NewRatingHistogram

`func NewRatingHistogram() *RatingHistogram`

NewRatingHistogram instantiates a new RatingHistogram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatingHistogramWithDefaults

`func NewRatingHistogramWithDefaults() *RatingHistogram`

NewRatingHistogramWithDefaults instantiates a new RatingHistogram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RatingHistogram) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RatingHistogram) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RatingHistogram) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *RatingHistogram) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetRating

`func (o *RatingHistogram) GetRating() string`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *RatingHistogram) GetRatingOk() (*string, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *RatingHistogram) SetRating(v string)`

SetRating sets Rating field to given value.

### HasRating

`func (o *RatingHistogram) HasRating() bool`

HasRating returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


