# UpdateCartItemInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CartItemId** | Pointer to **string** | The identifier of the item in the cart to be updated. This ID is generated when the item was added to the cart. | [optional] 
**Quantity** | Pointer to **int32** | The new quantity for the item that is being updated. | [optional] 

## Methods

### NewUpdateCartItemInput

`func NewUpdateCartItemInput() *UpdateCartItemInput`

NewUpdateCartItemInput instantiates a new UpdateCartItemInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCartItemInputWithDefaults

`func NewUpdateCartItemInputWithDefaults() *UpdateCartItemInput`

NewUpdateCartItemInputWithDefaults instantiates a new UpdateCartItemInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCartItemId

`func (o *UpdateCartItemInput) GetCartItemId() string`

GetCartItemId returns the CartItemId field if non-nil, zero value otherwise.

### GetCartItemIdOk

`func (o *UpdateCartItemInput) GetCartItemIdOk() (*string, bool)`

GetCartItemIdOk returns a tuple with the CartItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemId

`func (o *UpdateCartItemInput) SetCartItemId(v string)`

SetCartItemId sets CartItemId field to given value.

### HasCartItemId

`func (o *UpdateCartItemInput) HasCartItemId() bool`

HasCartItemId returns a boolean if a field has been set.

### GetQuantity

`func (o *UpdateCartItemInput) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UpdateCartItemInput) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UpdateCartItemInput) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UpdateCartItemInput) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


