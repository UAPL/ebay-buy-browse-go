# HazardStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatementDescription** | Pointer to **string** | A description of the nature of the hazard, such as whether the material is toxic if swallowed. | [optional] 
**StatementId** | Pointer to **string** | The ID of the hazard statement. | [optional] 

## Methods

### NewHazardStatement

`func NewHazardStatement() *HazardStatement`

NewHazardStatement instantiates a new HazardStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHazardStatementWithDefaults

`func NewHazardStatementWithDefaults() *HazardStatement`

NewHazardStatementWithDefaults instantiates a new HazardStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatementDescription

`func (o *HazardStatement) GetStatementDescription() string`

GetStatementDescription returns the StatementDescription field if non-nil, zero value otherwise.

### GetStatementDescriptionOk

`func (o *HazardStatement) GetStatementDescriptionOk() (*string, bool)`

GetStatementDescriptionOk returns a tuple with the StatementDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescription

`func (o *HazardStatement) SetStatementDescription(v string)`

SetStatementDescription sets StatementDescription field to given value.

### HasStatementDescription

`func (o *HazardStatement) HasStatementDescription() bool`

HasStatementDescription returns a boolean if a field has been set.

### GetStatementId

`func (o *HazardStatement) GetStatementId() string`

GetStatementId returns the StatementId field if non-nil, zero value otherwise.

### GetStatementIdOk

`func (o *HazardStatement) GetStatementIdOk() (*string, bool)`

GetStatementIdOk returns a tuple with the StatementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementId

`func (o *HazardStatement) SetStatementId(v string)`

SetStatementId sets StatementId field to given value.

### HasStatementId

`func (o *HazardStatement) HasStatementId() bool`

HasStatementId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


