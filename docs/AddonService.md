# AddonService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Selection** | Pointer to **string** | This field indicates whether the add-on service must be selected for the item. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:AddonServiceSelectionEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**ServiceFee** | Pointer to [**ConvertedAmount**](ConvertedAmount.md) |  | [optional] 
**ServiceId** | Pointer to **string** | The ID number of the add-on service. | [optional] 
**ServiceType** | Pointer to **string** | The type of add-on service, such as &lt;code&gt;AUTHENTICITY_GUARANTEE&lt;/code&gt;. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:AddonServiceTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewAddonService

`func NewAddonService() *AddonService`

NewAddonService instantiates a new AddonService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddonServiceWithDefaults

`func NewAddonServiceWithDefaults() *AddonService`

NewAddonServiceWithDefaults instantiates a new AddonService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelection

`func (o *AddonService) GetSelection() string`

GetSelection returns the Selection field if non-nil, zero value otherwise.

### GetSelectionOk

`func (o *AddonService) GetSelectionOk() (*string, bool)`

GetSelectionOk returns a tuple with the Selection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelection

`func (o *AddonService) SetSelection(v string)`

SetSelection sets Selection field to given value.

### HasSelection

`func (o *AddonService) HasSelection() bool`

HasSelection returns a boolean if a field has been set.

### GetServiceFee

`func (o *AddonService) GetServiceFee() ConvertedAmount`

GetServiceFee returns the ServiceFee field if non-nil, zero value otherwise.

### GetServiceFeeOk

`func (o *AddonService) GetServiceFeeOk() (*ConvertedAmount, bool)`

GetServiceFeeOk returns a tuple with the ServiceFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFee

`func (o *AddonService) SetServiceFee(v ConvertedAmount)`

SetServiceFee sets ServiceFee field to given value.

### HasServiceFee

`func (o *AddonService) HasServiceFee() bool`

HasServiceFee returns a boolean if a field has been set.

### GetServiceId

`func (o *AddonService) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *AddonService) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *AddonService) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *AddonService) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceType

`func (o *AddonService) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *AddonService) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *AddonService) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *AddonService) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


