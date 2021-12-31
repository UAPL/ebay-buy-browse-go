# AspectDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AspectValueDistributions** | Pointer to [**[]AspectValueDistribution**](AspectValueDistribution.md) | An array of containers for the various values of the aspect and the match count and a HATEOAS reference (&lt;b&gt; refinementHref&lt;/b&gt;) for this aspect. | [optional] 
**LocalizedAspectName** | Pointer to **string** | The name of an aspect, such as Brand, Color, etc. | [optional] 

## Methods

### NewAspectDistribution

`func NewAspectDistribution() *AspectDistribution`

NewAspectDistribution instantiates a new AspectDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAspectDistributionWithDefaults

`func NewAspectDistributionWithDefaults() *AspectDistribution`

NewAspectDistributionWithDefaults instantiates a new AspectDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAspectValueDistributions

`func (o *AspectDistribution) GetAspectValueDistributions() []AspectValueDistribution`

GetAspectValueDistributions returns the AspectValueDistributions field if non-nil, zero value otherwise.

### GetAspectValueDistributionsOk

`func (o *AspectDistribution) GetAspectValueDistributionsOk() (*[]AspectValueDistribution, bool)`

GetAspectValueDistributionsOk returns a tuple with the AspectValueDistributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectValueDistributions

`func (o *AspectDistribution) SetAspectValueDistributions(v []AspectValueDistribution)`

SetAspectValueDistributions sets AspectValueDistributions field to given value.

### HasAspectValueDistributions

`func (o *AspectDistribution) HasAspectValueDistributions() bool`

HasAspectValueDistributions returns a boolean if a field has been set.

### GetLocalizedAspectName

`func (o *AspectDistribution) GetLocalizedAspectName() string`

GetLocalizedAspectName returns the LocalizedAspectName field if non-nil, zero value otherwise.

### GetLocalizedAspectNameOk

`func (o *AspectDistribution) GetLocalizedAspectNameOk() (*string, bool)`

GetLocalizedAspectNameOk returns a tuple with the LocalizedAspectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedAspectName

`func (o *AspectDistribution) SetLocalizedAspectName(v string)`

SetLocalizedAspectName sets LocalizedAspectName field to given value.

### HasLocalizedAspectName

`func (o *AspectDistribution) HasLocalizedAspectName() bool`

HasLocalizedAspectName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


