# TargetLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitOfMeasure** | Pointer to **string** | This value shows the unit of measurement used to measure the distance between the location of the item and the buyer&#39;s location. This value is typically &lt;code&gt; mi&lt;/code&gt; or &lt;code&gt; km&lt;/code&gt;. | [optional] 
**Value** | Pointer to **string** | This value indicates the distance (measured in the measurement unit in the &lt;b&gt; unitOfMeasure&lt;/b&gt;  field) between the item location and the buyer&#39;s location. | [optional] 

## Methods

### NewTargetLocation

`func NewTargetLocation() *TargetLocation`

NewTargetLocation instantiates a new TargetLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetLocationWithDefaults

`func NewTargetLocationWithDefaults() *TargetLocation`

NewTargetLocationWithDefaults instantiates a new TargetLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitOfMeasure

`func (o *TargetLocation) GetUnitOfMeasure() string`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *TargetLocation) GetUnitOfMeasureOk() (*string, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *TargetLocation) SetUnitOfMeasure(v string)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.

### HasUnitOfMeasure

`func (o *TargetLocation) HasUnitOfMeasure() bool`

HasUnitOfMeasure returns a boolean if a field has been set.

### GetValue

`func (o *TargetLocation) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TargetLocation) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TargetLocation) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TargetLocation) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


