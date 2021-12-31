# AspectValueDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalizedAspectValue** | Pointer to **string** | The value of an aspect. For example, Red is a value for the aspect Color. | [optional] 
**MatchCount** | Pointer to **int32** | The number of items with this aspect. | [optional] 
**RefinementHref** | Pointer to **string** | A HATEOAS reference for this aspect. | [optional] 

## Methods

### NewAspectValueDistribution

`func NewAspectValueDistribution() *AspectValueDistribution`

NewAspectValueDistribution instantiates a new AspectValueDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAspectValueDistributionWithDefaults

`func NewAspectValueDistributionWithDefaults() *AspectValueDistribution`

NewAspectValueDistributionWithDefaults instantiates a new AspectValueDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalizedAspectValue

`func (o *AspectValueDistribution) GetLocalizedAspectValue() string`

GetLocalizedAspectValue returns the LocalizedAspectValue field if non-nil, zero value otherwise.

### GetLocalizedAspectValueOk

`func (o *AspectValueDistribution) GetLocalizedAspectValueOk() (*string, bool)`

GetLocalizedAspectValueOk returns a tuple with the LocalizedAspectValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedAspectValue

`func (o *AspectValueDistribution) SetLocalizedAspectValue(v string)`

SetLocalizedAspectValue sets LocalizedAspectValue field to given value.

### HasLocalizedAspectValue

`func (o *AspectValueDistribution) HasLocalizedAspectValue() bool`

HasLocalizedAspectValue returns a boolean if a field has been set.

### GetMatchCount

`func (o *AspectValueDistribution) GetMatchCount() int32`

GetMatchCount returns the MatchCount field if non-nil, zero value otherwise.

### GetMatchCountOk

`func (o *AspectValueDistribution) GetMatchCountOk() (*int32, bool)`

GetMatchCountOk returns a tuple with the MatchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCount

`func (o *AspectValueDistribution) SetMatchCount(v int32)`

SetMatchCount sets MatchCount field to given value.

### HasMatchCount

`func (o *AspectValueDistribution) HasMatchCount() bool`

HasMatchCount returns a boolean if a field has been set.

### GetRefinementHref

`func (o *AspectValueDistribution) GetRefinementHref() string`

GetRefinementHref returns the RefinementHref field if non-nil, zero value otherwise.

### GetRefinementHrefOk

`func (o *AspectValueDistribution) GetRefinementHrefOk() (*string, bool)`

GetRefinementHrefOk returns a tuple with the RefinementHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefinementHref

`func (o *AspectValueDistribution) SetRefinementHref(v string)`

SetRefinementHref sets RefinementHref field to given value.

### HasRefinementHref

`func (o *AspectValueDistribution) HasRefinementHref() bool`

HasRefinementHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


