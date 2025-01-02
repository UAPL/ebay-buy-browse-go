# ShipToLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | The two-letter &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html \&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 3166&lt;/a&gt; standard of the country for where the item is to be shipped. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/ba:CountryCodeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**PostalCode** | Pointer to **string** | The zip code (postal code) for where the item is to be shipped. | [optional] 

## Methods

### NewShipToLocation

`func NewShipToLocation() *ShipToLocation`

NewShipToLocation instantiates a new ShipToLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipToLocationWithDefaults

`func NewShipToLocationWithDefaults() *ShipToLocation`

NewShipToLocationWithDefaults instantiates a new ShipToLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *ShipToLocation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ShipToLocation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ShipToLocation) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ShipToLocation) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPostalCode

`func (o *ShipToLocation) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ShipToLocation) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ShipToLocation) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ShipToLocation) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


