# CategoryDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | Pointer to **string** | The unique identifier of the category. | [optional] 
**CategoryName** | Pointer to **string** | The name of the category, such as &lt;b&gt;Baby &amp;amp; Toddler Clothing&lt;/b&gt;. | [optional] 
**MatchCount** | Pointer to **int32** | The number of items in this category. | [optional] 
**RefinementHref** | Pointer to **string** | The HATEOAS reference of this category. | [optional] 

## Methods

### NewCategoryDistribution

`func NewCategoryDistribution() *CategoryDistribution`

NewCategoryDistribution instantiates a new CategoryDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryDistributionWithDefaults

`func NewCategoryDistributionWithDefaults() *CategoryDistribution`

NewCategoryDistributionWithDefaults instantiates a new CategoryDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *CategoryDistribution) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CategoryDistribution) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CategoryDistribution) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *CategoryDistribution) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetCategoryName

`func (o *CategoryDistribution) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *CategoryDistribution) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *CategoryDistribution) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *CategoryDistribution) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.

### GetMatchCount

`func (o *CategoryDistribution) GetMatchCount() int32`

GetMatchCount returns the MatchCount field if non-nil, zero value otherwise.

### GetMatchCountOk

`func (o *CategoryDistribution) GetMatchCountOk() (*int32, bool)`

GetMatchCountOk returns a tuple with the MatchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCount

`func (o *CategoryDistribution) SetMatchCount(v int32)`

SetMatchCount sets MatchCount field to given value.

### HasMatchCount

`func (o *CategoryDistribution) HasMatchCount() bool`

HasMatchCount returns a boolean if a field has been set.

### GetRefinementHref

`func (o *CategoryDistribution) GetRefinementHref() string`

GetRefinementHref returns the RefinementHref field if non-nil, zero value otherwise.

### GetRefinementHrefOk

`func (o *CategoryDistribution) GetRefinementHrefOk() (*string, bool)`

GetRefinementHrefOk returns a tuple with the RefinementHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefinementHref

`func (o *CategoryDistribution) SetRefinementHref(v string)`

SetRefinementHref sets RefinementHref field to given value.

### HasRefinementHref

`func (o *CategoryDistribution) HasRefinementHref() bool`

HasRefinementHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


