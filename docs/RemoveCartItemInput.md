# RemoveCartItemInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CartItemId** | Pointer to **string** | The identifier of the item in the cart to be removed. This ID is generated when the item was added to the cart.  | [optional] 

## Methods

### NewRemoveCartItemInput

`func NewRemoveCartItemInput() *RemoveCartItemInput`

NewRemoveCartItemInput instantiates a new RemoveCartItemInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveCartItemInputWithDefaults

`func NewRemoveCartItemInputWithDefaults() *RemoveCartItemInput`

NewRemoveCartItemInputWithDefaults instantiates a new RemoveCartItemInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCartItemId

`func (o *RemoveCartItemInput) GetCartItemId() string`

GetCartItemId returns the CartItemId field if non-nil, zero value otherwise.

### GetCartItemIdOk

`func (o *RemoveCartItemInput) GetCartItemIdOk() (*string, bool)`

GetCartItemIdOk returns a tuple with the CartItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemId

`func (o *RemoveCartItemInput) SetCartItemId(v string)`

SetCartItemId sets CartItemId field to given value.

### HasCartItemId

`func (o *RemoveCartItemInput) HasCartItemId() bool`

HasCartItemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


