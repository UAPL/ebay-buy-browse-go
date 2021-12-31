# Refinement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AspectDistributions** | Pointer to [**[]AspectDistribution**](AspectDistribution.md) | An array of containers for the all the aspect refinements. | [optional] 
**BuyingOptionDistributions** | Pointer to [**[]BuyingOptionDistribution**](BuyingOptionDistribution.md) | An array of containers for the all the buying option refinements. | [optional] 
**CategoryDistributions** | Pointer to [**[]CategoryDistribution**](CategoryDistribution.md) | An array of containers for the all the category refinements. | [optional] 
**ConditionDistributions** | Pointer to [**[]ConditionDistribution**](ConditionDistribution.md) | An array of containers for the all the condition refinements. | [optional] 
**DominantCategoryId** | Pointer to **string** | The identifier of the category that most of the items are part of.  | [optional] 

## Methods

### NewRefinement

`func NewRefinement() *Refinement`

NewRefinement instantiates a new Refinement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefinementWithDefaults

`func NewRefinementWithDefaults() *Refinement`

NewRefinementWithDefaults instantiates a new Refinement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAspectDistributions

`func (o *Refinement) GetAspectDistributions() []AspectDistribution`

GetAspectDistributions returns the AspectDistributions field if non-nil, zero value otherwise.

### GetAspectDistributionsOk

`func (o *Refinement) GetAspectDistributionsOk() (*[]AspectDistribution, bool)`

GetAspectDistributionsOk returns a tuple with the AspectDistributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectDistributions

`func (o *Refinement) SetAspectDistributions(v []AspectDistribution)`

SetAspectDistributions sets AspectDistributions field to given value.

### HasAspectDistributions

`func (o *Refinement) HasAspectDistributions() bool`

HasAspectDistributions returns a boolean if a field has been set.

### GetBuyingOptionDistributions

`func (o *Refinement) GetBuyingOptionDistributions() []BuyingOptionDistribution`

GetBuyingOptionDistributions returns the BuyingOptionDistributions field if non-nil, zero value otherwise.

### GetBuyingOptionDistributionsOk

`func (o *Refinement) GetBuyingOptionDistributionsOk() (*[]BuyingOptionDistribution, bool)`

GetBuyingOptionDistributionsOk returns a tuple with the BuyingOptionDistributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyingOptionDistributions

`func (o *Refinement) SetBuyingOptionDistributions(v []BuyingOptionDistribution)`

SetBuyingOptionDistributions sets BuyingOptionDistributions field to given value.

### HasBuyingOptionDistributions

`func (o *Refinement) HasBuyingOptionDistributions() bool`

HasBuyingOptionDistributions returns a boolean if a field has been set.

### GetCategoryDistributions

`func (o *Refinement) GetCategoryDistributions() []CategoryDistribution`

GetCategoryDistributions returns the CategoryDistributions field if non-nil, zero value otherwise.

### GetCategoryDistributionsOk

`func (o *Refinement) GetCategoryDistributionsOk() (*[]CategoryDistribution, bool)`

GetCategoryDistributionsOk returns a tuple with the CategoryDistributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryDistributions

`func (o *Refinement) SetCategoryDistributions(v []CategoryDistribution)`

SetCategoryDistributions sets CategoryDistributions field to given value.

### HasCategoryDistributions

`func (o *Refinement) HasCategoryDistributions() bool`

HasCategoryDistributions returns a boolean if a field has been set.

### GetConditionDistributions

`func (o *Refinement) GetConditionDistributions() []ConditionDistribution`

GetConditionDistributions returns the ConditionDistributions field if non-nil, zero value otherwise.

### GetConditionDistributionsOk

`func (o *Refinement) GetConditionDistributionsOk() (*[]ConditionDistribution, bool)`

GetConditionDistributionsOk returns a tuple with the ConditionDistributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionDistributions

`func (o *Refinement) SetConditionDistributions(v []ConditionDistribution)`

SetConditionDistributions sets ConditionDistributions field to given value.

### HasConditionDistributions

`func (o *Refinement) HasConditionDistributions() bool`

HasConditionDistributions returns a boolean if a field has been set.

### GetDominantCategoryId

`func (o *Refinement) GetDominantCategoryId() string`

GetDominantCategoryId returns the DominantCategoryId field if non-nil, zero value otherwise.

### GetDominantCategoryIdOk

`func (o *Refinement) GetDominantCategoryIdOk() (*string, bool)`

GetDominantCategoryIdOk returns a tuple with the DominantCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDominantCategoryId

`func (o *Refinement) SetDominantCategoryId(v string)`

SetDominantCategoryId sets DominantCategoryId field to given value.

### HasDominantCategoryId

`func (o *Refinement) HasDominantCategoryId() bool`

HasDominantCategoryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


