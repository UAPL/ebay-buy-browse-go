# ConditionDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **string** | The text describing the condition of the item, such as &lt;i&gt;New&lt;/i&gt; or &lt;i&gt;Used&lt;/i&gt;. For a list of condition names, refer to &lt;a href&#x3D;\&quot;/devzone/finding/callref/enums/conditionIdList.html \&quot; target&#x3D;\&quot;_blank\&quot;&gt;Item Condition IDs and Names&lt;/a&gt;. | [optional] 
**ConditionId** | Pointer to **string** | The identifier of the condition. For example, &lt;code&gt;1000&lt;/code&gt; is the identifier for &lt;code&gt;NEW&lt;/code&gt;. | [optional] 
**MatchCount** | Pointer to **int32** | The number of items having the condition. | [optional] 
**RefinementHref** | Pointer to **string** | The HATEOAS reference of this condition. | [optional] 

## Methods

### NewConditionDistribution

`func NewConditionDistribution() *ConditionDistribution`

NewConditionDistribution instantiates a new ConditionDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionDistributionWithDefaults

`func NewConditionDistributionWithDefaults() *ConditionDistribution`

NewConditionDistributionWithDefaults instantiates a new ConditionDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *ConditionDistribution) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ConditionDistribution) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ConditionDistribution) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ConditionDistribution) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetConditionId

`func (o *ConditionDistribution) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ConditionDistribution) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ConditionDistribution) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ConditionDistribution) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### GetMatchCount

`func (o *ConditionDistribution) GetMatchCount() int32`

GetMatchCount returns the MatchCount field if non-nil, zero value otherwise.

### GetMatchCountOk

`func (o *ConditionDistribution) GetMatchCountOk() (*int32, bool)`

GetMatchCountOk returns a tuple with the MatchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCount

`func (o *ConditionDistribution) SetMatchCount(v int32)`

SetMatchCount sets MatchCount field to given value.

### HasMatchCount

`func (o *ConditionDistribution) HasMatchCount() bool`

HasMatchCount returns a boolean if a field has been set.

### GetRefinementHref

`func (o *ConditionDistribution) GetRefinementHref() string`

GetRefinementHref returns the RefinementHref field if non-nil, zero value otherwise.

### GetRefinementHrefOk

`func (o *ConditionDistribution) GetRefinementHrefOk() (*string, bool)`

GetRefinementHrefOk returns a tuple with the RefinementHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefinementHref

`func (o *ConditionDistribution) SetRefinementHref(v string)`

SetRefinementHref sets RefinementHref field to given value.

### HasRefinementHref

`func (o *ConditionDistribution) HasRefinementHref() bool`

HasRefinementHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


