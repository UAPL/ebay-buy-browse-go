# ShipToRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionId** | Pointer to **string** | The unique identifier of the shipping region. The value returned here is dependent on the corresponding &lt;b&gt;regionType&lt;/b&gt; value. The &lt;b&gt;regionId&lt;/b&gt; value for a region does not vary based on the eBay marketplace. However, the corresponding &lt;b&gt;regionName&lt;/b&gt; value for a region is a localized, text-based description of the shipping region. &lt;br&gt;&lt;br&gt; If the &lt;b&gt;regionType&lt;/b&gt; value is &lt;code&gt;WORLDWIDE&lt;/code&gt;, the &lt;b&gt;regionId&lt;/b&gt; value will also be &lt;code&gt;WORLDWIDE&lt;/code&gt;.&lt;br&gt;&lt;br&gt; If the &lt;b&gt;regionType&lt;/b&gt; value is &lt;code&gt;WORLD_REGION&lt;/code&gt;, the &lt;b&gt;regionId&lt;/b&gt; value will be one of the following: &lt;code&gt;AFRICA&lt;/code&gt;, &lt;code&gt;AMERICAS&lt;/code&gt;, &lt;code&gt;ASIA&lt;/code&gt;, &lt;code&gt;AUSTRALIA&lt;/code&gt;, &lt;code&gt;CENTRAL_AMERICA_AND_CARIBBEAN&lt;/code&gt;, &lt;code&gt;EUROPE&lt;/code&gt;, &lt;code&gt;EUROPEAN_UNION&lt;/code&gt;, &lt;code&gt;GREATER_CHINA&lt;/code&gt;, &lt;code&gt;MIDDLE_EAST&lt;/code&gt;, &lt;code&gt;NORTH_AMERICA&lt;/code&gt;, &lt;code&gt;OCEANIA&lt;/code&gt;, &lt;code&gt;SOUTH_AMERICA&lt;/code&gt;, &lt;code&gt;SOUTHEAST_ASIA&lt;/code&gt; or &lt;code&gt;CHANNEL_ISLANDS&lt;/code&gt;.&lt;br&gt;&lt;br&gt;If the &lt;b&gt;regionType&lt;/b&gt; value is &lt;code&gt;COUNTRY&lt;/code&gt;, the &lt;b&gt;regionId&lt;/b&gt; value will be the two-letter code for the country, as defined in the &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html \&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 3166&lt;/a&gt; standard.&lt;br&gt;&lt;br&gt;If the &lt;b&gt;regionType&lt;/b&gt; value is &lt;code&gt;STATE_OR_PROVINCE&lt;/code&gt;, the &lt;b&gt;regionId&lt;/b&gt; value will either be the two-letter code for US states and DC (as defined on this &lt;a href&#x3D;\&quot;https://www.ssa.gov/international/coc-docs/states.html \&quot; target&#x3D;\&quot;_blank\&quot;&gt;Social Security Administration&lt;/a&gt; page), or the two-letter code for Canadian provinces (as defined by this &lt;a href&#x3D;\&quot;https://www.canadapost.ca/tools/pg/manual/PGaddress-e.asp?ecid&#x3D;murl10006450#1442131 \&quot; target&#x3D;\&quot;_blank\&quot;&gt;Canada Post&lt;/a&gt; page).&lt;br&gt;&lt;br&gt;If the &lt;b&gt;regionType&lt;/b&gt; value is &lt;code&gt;COUNTRY_REGION&lt;/code&gt;, the &lt;b&gt;regionId&lt;/b&gt; value may be one of following: &lt;code&gt;_AH&lt;/code&gt; (if a seller is not willing to ship to Alaska/Hawaii), &lt;code&gt;_PR&lt;/code&gt; (if the seller is not willing to ship to US Protectorates), &lt;code&gt;_AP&lt;/code&gt; (if seller is not willing to ship to a US Army or Fleet Post Office), and &lt;code&gt;PO_BOX&lt;/code&gt; (if the seller is not willing to ship to a Post Office Box). | [optional] 
**RegionName** | Pointer to **string** | A localized text string that indicates the name of the shipping region. The value returned here is dependent on the corresponding &lt;b&gt;regionType&lt;/b&gt; value. &lt;br&gt;&lt;br&gt; If the &lt;b&gt;regionType&lt;/b&gt; value is &lt;code&gt;WORLDWIDE&lt;/code&gt;, the &lt;b&gt;regionName&lt;/b&gt; value will show &lt;code&gt;Worldwide&lt;/code&gt;.&lt;br&gt;&lt;br&gt; If the &lt;b&gt;regionType&lt;/b&gt; value is &lt;code&gt;WORLD_REGION&lt;/code&gt;, the &lt;b&gt;regionName&lt;/b&gt; value will be a localized text string for one of the following large geographical regions: Africa, Americas, Asia, Australia, Central America and Caribbean, Europe, European Union, Greater China, Middle East, North America, Oceania, South America, Southeast Asia, or Channel Islands.&lt;br&gt;&lt;br&gt;If the &lt;b&gt;regionType&lt;/b&gt; value is &lt;code&gt;COUNTRY&lt;/code&gt;, the &lt;b&gt;regionName&lt;/b&gt; value will be a localized text string for any country in the world.&lt;br&gt;&lt;br&gt;If the &lt;b&gt;regionType&lt;/b&gt; value is &lt;code&gt;STATE_OR_PROVINCE&lt;/code&gt;, the &lt;b&gt;regionName&lt;/b&gt; value will be a localized text string for any US state or Canadian province. &lt;br&gt;&lt;br&gt;If the &lt;b&gt;regionType&lt;/b&gt; value is &lt;code&gt;COUNTRY_REGION&lt;/code&gt;, the &lt;b&gt;regionName&lt;/b&gt; value may be a localized version of one of the following: Alaska/Hawaii, US Protectorates, APO/FPO (Army or Fleet Post Office), or PO BOX. | [optional] 
**RegionType** | Pointer to **string** | An enumeration value that indicates the level or type of shipping region. &lt;br&gt;&lt;br&gt;&lt;b&gt; Valid Values: &lt;/b&gt; &lt;ul&gt;&lt;li&gt;&lt;b&gt; COUNTRY_REGION &lt;/b&gt; - Indicates the region is a domestic region or special location within a country.&lt;/li&gt;&lt;li&gt;&lt;b&gt; STATE_OR_PROVINCE &lt;/b&gt; - Indicates the region is a state or province within a country, such as California or New York in the US, or Ontario or Alberta in Canada.&lt;/li&gt;&lt;li&gt;&lt;b&gt; COUNTRY &lt;/b&gt; - Indicates the region is a single country.&lt;/li&gt;&lt;li&gt;&lt;b&gt; WORLD_REGION &lt;/b&gt; - Indicates the region is a world region, such as Africa, the Middle East, or Southeast Asia.&lt;/li&gt;&lt;li&gt;&lt;b&gt; WORLDWIDE &lt;/b&gt; - Indicates the region is the entire world. This value is only applicable for included shiping regions, and not excluded shipping regions.&lt;/li&gt;&lt;/ul&gt; For more detail on the actual &lt;b&gt;regionName&lt;/b&gt;/&lt;b&gt;regionId&lt;/b&gt; values that will be returned based on the &lt;b&gt;regionType&lt;/b&gt; value, see the &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item/methods/getItem#response.shipToLocations.regionExcluded.regionId\&quot;&gt;regionId&lt;/a&gt; and/or &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item/methods/getItem#response.shipToLocations.regionExcluded.regionName\&quot;&gt;regionName&lt;/a&gt; field descriptions.&lt;br&gt;&lt;br&gt; Code so that your app gracefully handles any future changes to this list. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/ba:RegionTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewShipToRegion

`func NewShipToRegion() *ShipToRegion`

NewShipToRegion instantiates a new ShipToRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipToRegionWithDefaults

`func NewShipToRegionWithDefaults() *ShipToRegion`

NewShipToRegionWithDefaults instantiates a new ShipToRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionId

`func (o *ShipToRegion) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *ShipToRegion) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *ShipToRegion) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *ShipToRegion) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

### GetRegionName

`func (o *ShipToRegion) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ShipToRegion) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ShipToRegion) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ShipToRegion) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetRegionType

`func (o *ShipToRegion) GetRegionType() string`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *ShipToRegion) GetRegionTypeOk() (*string, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *ShipToRegion) SetRegionType(v string)`

SetRegionType sets RegionType field to given value.

### HasRegionType

`func (o *ShipToRegion) HasRegionType() bool`

HasRegionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


