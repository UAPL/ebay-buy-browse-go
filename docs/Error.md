# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | This string value indicates the error category. There are three categories of errors: &lt;i&gt;request errors&lt;/i&gt;, &lt;i&gt;application errors&lt;/i&gt;, and &lt;i&gt;system errors&lt;/i&gt;. | [optional] 
**Domain** | Pointer to **string** | The name of the primary system where the error occurred. This is relevant for application errors. | [optional] 
**ErrorId** | Pointer to **int32** | A unique code that identifies the particular error or warning that occurred. Your application can use error codes as identifiers in your customized error-handling algorithms. | [optional] 
**InputRefIds** | Pointer to **[]string** | An array of reference IDs that identify the specific request elements most closely associated to the error or warning, if any. | [optional] 
**LongMessage** | Pointer to **string** | A detailed description of the condition that caused the error or warning, and information on what to do to correct the problem. | [optional] 
**Message** | Pointer to **string** | A description of the condition that caused the error or warning. | [optional] 
**OutputRefIds** | Pointer to **[]string** | An array of reference IDs that identify the specific response elements most closely associated to the error or warning, if any. | [optional] 
**Parameters** | Pointer to [**[]ErrorParameter**](ErrorParameter.md) | An array of warning and error messages that return one or more variables contextual information about the error or warning. This is often the field or value that triggered the error or warning. | [optional] 
**Subdomain** | Pointer to **string** | The name of the subdomain in which the error or warning occurred. | [optional] 

## Methods

### NewError

`func NewError() *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *Error) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Error) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Error) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Error) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDomain

`func (o *Error) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Error) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Error) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Error) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetErrorId

`func (o *Error) GetErrorId() int32`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *Error) GetErrorIdOk() (*int32, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *Error) SetErrorId(v int32)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *Error) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### GetInputRefIds

`func (o *Error) GetInputRefIds() []string`

GetInputRefIds returns the InputRefIds field if non-nil, zero value otherwise.

### GetInputRefIdsOk

`func (o *Error) GetInputRefIdsOk() (*[]string, bool)`

GetInputRefIdsOk returns a tuple with the InputRefIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputRefIds

`func (o *Error) SetInputRefIds(v []string)`

SetInputRefIds sets InputRefIds field to given value.

### HasInputRefIds

`func (o *Error) HasInputRefIds() bool`

HasInputRefIds returns a boolean if a field has been set.

### GetLongMessage

`func (o *Error) GetLongMessage() string`

GetLongMessage returns the LongMessage field if non-nil, zero value otherwise.

### GetLongMessageOk

`func (o *Error) GetLongMessageOk() (*string, bool)`

GetLongMessageOk returns a tuple with the LongMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongMessage

`func (o *Error) SetLongMessage(v string)`

SetLongMessage sets LongMessage field to given value.

### HasLongMessage

`func (o *Error) HasLongMessage() bool`

HasLongMessage returns a boolean if a field has been set.

### GetMessage

`func (o *Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Error) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Error) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOutputRefIds

`func (o *Error) GetOutputRefIds() []string`

GetOutputRefIds returns the OutputRefIds field if non-nil, zero value otherwise.

### GetOutputRefIdsOk

`func (o *Error) GetOutputRefIdsOk() (*[]string, bool)`

GetOutputRefIdsOk returns a tuple with the OutputRefIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputRefIds

`func (o *Error) SetOutputRefIds(v []string)`

SetOutputRefIds sets OutputRefIds field to given value.

### HasOutputRefIds

`func (o *Error) HasOutputRefIds() bool`

HasOutputRefIds returns a boolean if a field has been set.

### GetParameters

`func (o *Error) GetParameters() []ErrorParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Error) GetParametersOk() (*[]ErrorParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Error) SetParameters(v []ErrorParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Error) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetSubdomain

`func (o *Error) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *Error) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *Error) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *Error) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


