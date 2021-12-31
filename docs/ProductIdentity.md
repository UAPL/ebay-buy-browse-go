# ProductIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentifierType** | Pointer to **string** | The type of product identifier, such as UPC and EAN. | [optional] 
**IdentifierValue** | Pointer to **string** | The product identifier value. | [optional] 

## Methods

### NewProductIdentity

`func NewProductIdentity() *ProductIdentity`

NewProductIdentity instantiates a new ProductIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductIdentityWithDefaults

`func NewProductIdentityWithDefaults() *ProductIdentity`

NewProductIdentityWithDefaults instantiates a new ProductIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifierType

`func (o *ProductIdentity) GetIdentifierType() string`

GetIdentifierType returns the IdentifierType field if non-nil, zero value otherwise.

### GetIdentifierTypeOk

`func (o *ProductIdentity) GetIdentifierTypeOk() (*string, bool)`

GetIdentifierTypeOk returns a tuple with the IdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierType

`func (o *ProductIdentity) SetIdentifierType(v string)`

SetIdentifierType sets IdentifierType field to given value.

### HasIdentifierType

`func (o *ProductIdentity) HasIdentifierType() bool`

HasIdentifierType returns a boolean if a field has been set.

### GetIdentifierValue

`func (o *ProductIdentity) GetIdentifierValue() string`

GetIdentifierValue returns the IdentifierValue field if non-nil, zero value otherwise.

### GetIdentifierValueOk

`func (o *ProductIdentity) GetIdentifierValueOk() (*string, bool)`

GetIdentifierValueOk returns a tuple with the IdentifierValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierValue

`func (o *ProductIdentity) SetIdentifierValue(v string)`

SetIdentifierValue sets IdentifierValue field to given value.

### HasIdentifierValue

`func (o *ProductIdentity) HasIdentifierValue() bool`

HasIdentifierValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


