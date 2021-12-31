# ShipToLocations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionExcluded** | Pointer to [**[]ShipToRegion**](ShipToRegion.md) | An array of containers that express the large geographical regions, countries, state/provinces, or special locations within a country where the seller is not willing to ship to. | [optional] 
**RegionIncluded** | Pointer to [**[]ShipToRegion**](ShipToRegion.md) | An array of containers that express the large geographical regions, countries, or state/provinces within a country where the seller is willing to ship to. Prospective buyers must look at the shipping regions under this container, as well as the shipping regions that are under the &lt;b&gt;regionExcluded&lt;/b&gt; to see where the seller is willing to ship items. Sellers can specify that they ship &#39;Worldwide&#39;, but then add several large geographical regions (e.g. Asia, Oceania, Middle East) to the exclusion list, or sellers can specify that they ship to Europe and Africa, but then add several individual countries to the exclusion list. | [optional] 

## Methods

### NewShipToLocations

`func NewShipToLocations() *ShipToLocations`

NewShipToLocations instantiates a new ShipToLocations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipToLocationsWithDefaults

`func NewShipToLocationsWithDefaults() *ShipToLocations`

NewShipToLocationsWithDefaults instantiates a new ShipToLocations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionExcluded

`func (o *ShipToLocations) GetRegionExcluded() []ShipToRegion`

GetRegionExcluded returns the RegionExcluded field if non-nil, zero value otherwise.

### GetRegionExcludedOk

`func (o *ShipToLocations) GetRegionExcludedOk() (*[]ShipToRegion, bool)`

GetRegionExcludedOk returns a tuple with the RegionExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionExcluded

`func (o *ShipToLocations) SetRegionExcluded(v []ShipToRegion)`

SetRegionExcluded sets RegionExcluded field to given value.

### HasRegionExcluded

`func (o *ShipToLocations) HasRegionExcluded() bool`

HasRegionExcluded returns a boolean if a field has been set.

### GetRegionIncluded

`func (o *ShipToLocations) GetRegionIncluded() []ShipToRegion`

GetRegionIncluded returns the RegionIncluded field if non-nil, zero value otherwise.

### GetRegionIncludedOk

`func (o *ShipToLocations) GetRegionIncludedOk() (*[]ShipToRegion, bool)`

GetRegionIncludedOk returns a tuple with the RegionIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionIncluded

`func (o *ShipToLocations) SetRegionIncluded(v []ShipToRegion)`

SetRegionIncluded sets RegionIncluded field to given value.

### HasRegionIncluded

`func (o *ShipToLocations) HasRegionIncluded() bool`

HasRegionIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


