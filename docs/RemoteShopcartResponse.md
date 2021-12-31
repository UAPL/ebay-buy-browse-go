# RemoteShopcartResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CartItems** | Pointer to [**[]CartItem**](CartItem.md) | An array of the items in the member&#39;s eBay cart. | [optional] 
**CartSubtotal** | Pointer to [**Amount**](Amount.md) |  | [optional] 
**CartWebUrl** | Pointer to **string** | The URL of the member&#39;s eBay cart. | [optional] 
**UnavailableCartItems** | Pointer to [**[]CartItem**](CartItem.md) | An array of items in the cart that are unavailable. This can be for a variety of reasons such as, when the listing has ended or the item is out of stock. Because a cart never expires, these items will remain in the cart until they are removed. | [optional] 
**Warnings** | Pointer to [**[]Error**](Error.md) | An array of warning messages. These type of errors do not prevent the call from executing but should be checked. | [optional] 

## Methods

### NewRemoteShopcartResponse

`func NewRemoteShopcartResponse() *RemoteShopcartResponse`

NewRemoteShopcartResponse instantiates a new RemoteShopcartResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteShopcartResponseWithDefaults

`func NewRemoteShopcartResponseWithDefaults() *RemoteShopcartResponse`

NewRemoteShopcartResponseWithDefaults instantiates a new RemoteShopcartResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCartItems

`func (o *RemoteShopcartResponse) GetCartItems() []CartItem`

GetCartItems returns the CartItems field if non-nil, zero value otherwise.

### GetCartItemsOk

`func (o *RemoteShopcartResponse) GetCartItemsOk() (*[]CartItem, bool)`

GetCartItemsOk returns a tuple with the CartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItems

`func (o *RemoteShopcartResponse) SetCartItems(v []CartItem)`

SetCartItems sets CartItems field to given value.

### HasCartItems

`func (o *RemoteShopcartResponse) HasCartItems() bool`

HasCartItems returns a boolean if a field has been set.

### GetCartSubtotal

`func (o *RemoteShopcartResponse) GetCartSubtotal() Amount`

GetCartSubtotal returns the CartSubtotal field if non-nil, zero value otherwise.

### GetCartSubtotalOk

`func (o *RemoteShopcartResponse) GetCartSubtotalOk() (*Amount, bool)`

GetCartSubtotalOk returns a tuple with the CartSubtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartSubtotal

`func (o *RemoteShopcartResponse) SetCartSubtotal(v Amount)`

SetCartSubtotal sets CartSubtotal field to given value.

### HasCartSubtotal

`func (o *RemoteShopcartResponse) HasCartSubtotal() bool`

HasCartSubtotal returns a boolean if a field has been set.

### GetCartWebUrl

`func (o *RemoteShopcartResponse) GetCartWebUrl() string`

GetCartWebUrl returns the CartWebUrl field if non-nil, zero value otherwise.

### GetCartWebUrlOk

`func (o *RemoteShopcartResponse) GetCartWebUrlOk() (*string, bool)`

GetCartWebUrlOk returns a tuple with the CartWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartWebUrl

`func (o *RemoteShopcartResponse) SetCartWebUrl(v string)`

SetCartWebUrl sets CartWebUrl field to given value.

### HasCartWebUrl

`func (o *RemoteShopcartResponse) HasCartWebUrl() bool`

HasCartWebUrl returns a boolean if a field has been set.

### GetUnavailableCartItems

`func (o *RemoteShopcartResponse) GetUnavailableCartItems() []CartItem`

GetUnavailableCartItems returns the UnavailableCartItems field if non-nil, zero value otherwise.

### GetUnavailableCartItemsOk

`func (o *RemoteShopcartResponse) GetUnavailableCartItemsOk() (*[]CartItem, bool)`

GetUnavailableCartItemsOk returns a tuple with the UnavailableCartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableCartItems

`func (o *RemoteShopcartResponse) SetUnavailableCartItems(v []CartItem)`

SetUnavailableCartItems sets UnavailableCartItems field to given value.

### HasUnavailableCartItems

`func (o *RemoteShopcartResponse) HasUnavailableCartItems() bool`

HasUnavailableCartItems returns a boolean if a field has been set.

### GetWarnings

`func (o *RemoteShopcartResponse) GetWarnings() []Error`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *RemoteShopcartResponse) GetWarningsOk() (*[]Error, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *RemoteShopcartResponse) SetWarnings(v []Error)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *RemoteShopcartResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


