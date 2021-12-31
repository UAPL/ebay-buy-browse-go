# PickupOptionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PickupLocationType** | Pointer to **string** | This container returns the local pickup options available to the buyer. Possible values are &lt;code&gt;ARRANGED_LOCATION&lt;/code&gt; and &lt;code&gt;STORE&lt;/code&gt;. | [optional] 

## Methods

### NewPickupOptionSummary

`func NewPickupOptionSummary() *PickupOptionSummary`

NewPickupOptionSummary instantiates a new PickupOptionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPickupOptionSummaryWithDefaults

`func NewPickupOptionSummaryWithDefaults() *PickupOptionSummary`

NewPickupOptionSummaryWithDefaults instantiates a new PickupOptionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPickupLocationType

`func (o *PickupOptionSummary) GetPickupLocationType() string`

GetPickupLocationType returns the PickupLocationType field if non-nil, zero value otherwise.

### GetPickupLocationTypeOk

`func (o *PickupOptionSummary) GetPickupLocationTypeOk() (*string, bool)`

GetPickupLocationTypeOk returns a tuple with the PickupLocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupLocationType

`func (o *PickupOptionSummary) SetPickupLocationType(v string)`

SetPickupLocationType sets PickupLocationType field to given value.

### HasPickupLocationType

`func (o *PickupOptionSummary) HasPickupLocationType() bool`

HasPickupLocationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


