# ErrorParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | This is the name of input field that caused an issue with the call request. | [optional] 
**Value** | Pointer to **string** | This is the actual value that was passed in for the element specified in the &lt;code&gt;name&lt;/code&gt; field. | [optional] 

## Methods

### NewErrorParameter

`func NewErrorParameter() *ErrorParameter`

NewErrorParameter instantiates a new ErrorParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorParameterWithDefaults

`func NewErrorParameterWithDefaults() *ErrorParameter`

NewErrorParameterWithDefaults instantiates a new ErrorParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ErrorParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ErrorParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ErrorParameter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ErrorParameter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ErrorParameter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorParameter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorParameter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ErrorParameter) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


