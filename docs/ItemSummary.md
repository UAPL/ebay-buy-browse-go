# ItemSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalImages** | Pointer to [**[]Image**](Image.md) | An array of containers with the URLs for the images that are in addition to the primary image.  The primary image is returned in the &lt;b&gt; image.imageUrl&lt;/b&gt; field. | [optional] 
**AdultOnly** | Pointer to **bool** | This indicates if the item is for adults only. For more information about adult-only items on eBay, see &lt;a href&#x3D;\&quot;https://pages.ebay.com/help/policies/adult-only.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Adult items policy&lt;/a&gt; for sellers and &lt;a href&#x3D;\&quot;https://www.ebay.com/help/terms-conditions/default/searching-adult-items?id&#x3D;4661\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Adult-Only items on eBay&lt;/a&gt; for buyers. | [optional] 
**AvailableCoupons** | Pointer to **bool** | This boolean attribute indicates if coupons are available for the item. | [optional] 
**BidCount** | Pointer to **int32** | This integer value indicates the total number of bids that have been placed for an auction item. This field is only returned for auction items. | [optional] 
**BuyingOptions** | Pointer to **[]string** | A comma separated list of all the purchase options available for the item. &lt;br&gt;&lt;br&gt;&lt;b&gt; Values Returned: &lt;/b&gt;   &lt;ul&gt; &lt;li&gt;&lt;b&gt;FIXED_PRICE&lt;/b&gt; - Indicates the buyer can purchase the item for a set price using the Buy It Now button. &lt;/li&gt;  &lt;li&gt;&lt;b&gt; AUCTION&lt;/b&gt; - Indicates the buyer can place a bid for the item. After the first bid is placed, this becomes a live auction item and is the only buying option for this item.&lt;/li&gt;  &lt;li&gt;&lt;b&gt; BEST_OFFER&lt;/b&gt; - Items where the buyer can send the seller a price they&#39;re willing to pay for the item. The seller can accept, reject, or send a counter offer. For details about Best Offer, see &lt;a href&#x3D;\&quot;https://www.ebay.com/help/selling/listings/selling-buy-now/adding-best-offer-listing?id&#x3D;4144\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Best Offer&lt;/a&gt;.  &lt;/li&gt;&lt;/ul&gt; Code so that your app gracefully handles any future changes to this list. | [optional] 
**Categories** | Pointer to [**[]Category**](Category.md) | This container returns the primary category ID of the item (as well as the secondary category if the item was listed in two categories).  | [optional] 
**CompatibilityMatch** | Pointer to **string** | This indicates how well the item matches the &lt;b&gt;compatibility_filter&lt;/b&gt; product attributes.  &lt;br&gt;&lt;br&gt;&lt;b&gt; Valid Values: &lt;/b&gt; EXACT or POSSIBLE &lt;br /&gt;&lt;br /&gt;Code so that your app gracefully handles any future changes to this list. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:CompatibilityMatchEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**CompatibilityProperties** | Pointer to [**[]CompatibilityProperty**](CompatibilityProperty.md) | This container returns only the product attributes that are compatible with the item. These attributes were specified in the &lt;b&gt;compatibility_filter&lt;/b&gt; in the request. This means that if you passed in 5 attributes and only 4 are compatible, only those 4 are returned. If none of the attributes are compatible, this container is not returned. | [optional] 
**Condition** | Pointer to **string** | The text describing the condition of the item, such as New or Used. For a list of condition names, see &lt;a href&#x3D;\&quot;https://developer.ebay.com/devzone/finding/callref/enums/conditionIdList.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Item Condition IDs and Names&lt;/a&gt;.  &lt;br /&gt;&lt;br /&gt;Code so that your app gracefully handles any future changes to this list. &lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt; Note: &lt;/b&gt;As of September 1, 2021, condition ID 2500 (&#39;Seller Refurbished&#39;) is no longer a valid item condition in the &lt;b&gt;Cell Phones &amp; Smartphones&lt;/b&gt; category (category ID 9355) for the following marketplaces: US, Canada, UK, Germany, and Australia. This refurbished item condition has been replaced by three new refurbished values, which include &#39;Excellent - Refurbished&#39; (condition ID 2010), &#39;Very Good - Refurbished&#39; (condition ID 2020), and &#39;Good - Refurbished&#39; (condition ID 2030).&lt;/span&gt; | [optional] 
**ConditionId** | Pointer to **string** | The identifier of the condition of the item. For example, 1000 is the identifier for NEW. For a list of condition names and IDs, see &lt;a https://developer.ebay.com/devzone/finding/callref/enums/conditionIdList.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Item Condition IDs and Names&lt;/a&gt;. &lt;br /&gt;&lt;br /&gt;Code so that your app gracefully handles any future changes to this list.&lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt; Note: &lt;/b&gt;As of September 1, 2021, condition ID 2500 (&#39;Seller Refurbished&#39;) is no longer a valid item condition in the &lt;b&gt;Cell Phones &amp; Smartphones&lt;/b&gt; category (category ID 9355) for the following marketplaces: US, Canada, UK, Germany, and Australia. This refurbished item condition has been replaced by three new refurbished values, which include &#39;Excellent - Refurbished&#39; (condition ID 2010), &#39;Very Good - Refurbished&#39; (condition ID 2020), and &#39;Good - Refurbished&#39; (condition ID 2030).&lt;/span&gt; | [optional] 
**CurrentBidPrice** | Pointer to [**ConvertedAmount**](ConvertedAmount.md) |  | [optional] 
**DistanceFromPickupLocation** | Pointer to [**TargetLocation**](TargetLocation.md) |  | [optional] 
**EnergyEfficiencyClass** | Pointer to **string** | This indicates the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/European_Union_energy_label\&quot;&gt;European energy efficiency&lt;/a&gt; rating (EEK) of the item.  Energy efficiency ratings apply to products listed by commercial vendors in electronics categories only. &lt;br /&gt;&lt;br /&gt;Currently, this field is only applicable for the Germany site, and  is only returned if the seller specified the energy efficiency rating through item specifics at listing time. Rating values include &lt;code&gt;A+++&lt;/code&gt;, &lt;code&gt;A++&lt;/code&gt;, &lt;code&gt;A+&lt;/code&gt;, &lt;code&gt;A&lt;/code&gt;, &lt;code&gt;B&lt;/code&gt;, &lt;code&gt;C&lt;/code&gt;, &lt;code&gt;D&lt;/code&gt;, &lt;code&gt;E&lt;/code&gt;, &lt;code&gt;F&lt;/code&gt;, and &lt;code&gt;G&lt;/code&gt;. | [optional] 
**Epid** | Pointer to **string** | An ePID is the eBay product identifier of a product from the eBay product catalog.  This indicates the product in which the item belongs. | [optional] 
**Image** | Pointer to [**Image**](Image.md) |  | [optional] 
**ItemAffiliateWebUrl** | Pointer to **string** | The URL to the View Item page of the item, which includes the affiliate tracking ID. This field is only returned if the seller enables affiliate tracking for the item by including the &lt;code&gt;&lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#Headers\&quot;&gt;X-EBAY-C-ENDUSERCTX&lt;/a&gt;&lt;/code&gt; request header in the method.  &lt;br /&gt; &lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note: &lt;/b&gt; eBay Partner Network, in order to receive a commission for your sales, you must use this URL to forward your buyer to the ebay.com site. &lt;/span&gt; | [optional] 
**ItemEndDate** | Pointer to **string** | The date and time up to which the item can be purchased.  This value is returned in UTC format (yyyy-MM-ddThh:mm:ss.sssZ), which you can convert into the local time of the buyer.&lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt; Note: &lt;/b&gt;This field is not returned for Good &#39;Til Cancelled (GTC) listings.&lt;/span&gt; | [optional] 
**ItemGroupHref** | Pointer to **string** | The HATEOAS reference of the parent page of the item group. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc. &lt;br /&gt; &lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt; &lt;b&gt;  Note: &lt;/b&gt;This field is returned only for item groups.&lt;/span&gt; | [optional] 
**ItemGroupType** | Pointer to **string** | The indicates the item group type. An item group is an item that has various aspect differences, such as color, size, storage capacity, etc. &lt;br /&gt;&lt;br /&gt;Currently only the &lt;code&gt;SELLER_DEFINED_VARIATIONS&lt;/code&gt; is supported and indicates this is an item group created by the seller. &lt;br /&gt; &lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt; &lt;b&gt; Note: &lt;/b&gt;This field is returned only for item groups.&lt;/span&gt;&lt;br /&gt;&lt;br /&gt;Code so that your app gracefully handles any future changes to this list. | [optional] 
**ItemHref** | Pointer to **string** | The URI for the Browse API &lt;a href&#x3D;\&quot;/api-docs/buy/browse/resources/item/methods/getItem\&quot;&gt;getItem&lt;/a&gt; method, which can be used to retrieve more details about items in the search results.  | [optional] 
**ItemId** | Pointer to **string** | The unique RESTful identifier of the item. | [optional] 
**ItemLocation** | Pointer to [**ItemLocationImpl**](ItemLocationImpl.md) |  | [optional] 
**ItemWebUrl** | Pointer to **string** | The URL to the View Item page of the item.  This enables you to include a \&quot;Report Item on eBay\&quot; hyperlink that takes the buyer to the View Item page on eBay. From there they can report any issues regarding this item to eBay. | [optional] 
**LegacyItemId** | Pointer to **string** | The unique identifier of the eBay listing that contains the item. This is the traditional/legacy ID that is often seen in the URL of the listing View Item page. | [optional] 
**MarketingPrice** | Pointer to [**MarketingPrice**](MarketingPrice.md) |  | [optional] 
**PickupOptions** | Pointer to [**[]PickupOptionSummary**](PickupOptionSummary.md) | This container returns the local pickup options available to the buyer. This container is only returned if the user is searching for local pickup items and set the local pickup filters in the method request. | [optional] 
**Price** | Pointer to [**ConvertedAmount**](ConvertedAmount.md) |  | [optional] 
**PriceDisplayCondition** | Pointer to **string** | Indicates when in the buying flow the item&#39;s price can appear for minimum advertised price (MAP) items, which is the lowest price a retailer can advertise/show for this item. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:PriceDisplayConditionEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**PriorityListing** | Pointer to **bool** | This field is returned as &lt;code&gt;true&lt;/code&gt; if the listing is part of a Promoted Listing campaign. Promoted Listings are available to Above Standard and Top Rated sellers with recent sales activity.&lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; Priority Listing is returned only with a Best Match sort and will not be returned for other sort options.&lt;/span&gt; | [optional] 
**QualifiedPrograms** | Pointer to **[]string** | An array of the qualified programs available for the item, such as EBAY_PLUS, AUTHENTICITY_GUARANTEE, and AUTHENTICITY_VERIFICATION.&lt;br /&gt;&lt;br /&gt;eBay Plus is a premium account option for buyers, which provides benefits such as fast free domestic shipping and free returns on selected items. Top-Rated eBay sellers must opt in to eBay Plus to be able to offer the program on qualifying listings. Sellers must commit to next-day delivery of those items.&lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note: &lt;/b&gt; eBay Plus is available only to buyers in Germany, Austria, and Australia marketplaces.&lt;/span&gt;&lt;br /&gt;&lt;br /&gt;The eBay &lt;a href&#x3D;\&quot;https://pages.ebay.com/authenticity-guarantee/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Authenticity Guarantee&lt;/a&gt; program enables third-party authenticators to perform authentication verification inspections on items such as watches and sneakers. | [optional] 
**Seller** | Pointer to [**Seller**](Seller.md) |  | [optional] 
**ShippingOptions** | Pointer to [**[]ShippingOptionSummary**](ShippingOptionSummary.md) | This container returns the shipping options available to ship the item. | [optional] 
**ShortDescription** | Pointer to **string** | This text string is derived from the item condition and the item aspects (such as size, color, capacity, model, brand, etc.). Sometimes the title doesn&#39;t give enough information but the description is too big. Surfacing the &lt;b&gt;shortDescription&lt;/b&gt; can often provide buyers with the additional information that could help them make a buying decision.  &lt;br /&gt;&lt;br /&gt;For example: &lt;br /&gt;&lt;br /&gt;    &lt;code&gt;   \&quot;&lt;b&gt; title&lt;/b&gt;\&quot;: \&quot;Petrel U42W FPV Drone RC Quadcopter w/HD Camera Live Video One Key Off / Landing\&quot;, &lt;br /&gt;\&quot;&lt;b&gt;shortDescription&lt;/b&gt;\&quot;: \&quot;1 U42W Quadcopter. Syma X5SW-V3 Wifi FPV RC Drone Quadcopter 2.4Ghz 6-Axis Gyro with Headless Mode. Syma X20 Pocket Drone 2.4Ghz Mini RC Quadcopter Headless Mode Altitude Hold. One Key Take Off / Landing function: allow beginner to easy to fly the drone without any skill.\&quot;,&lt;/code&gt;       &lt;br /&gt;&lt;br /&gt;&lt;b&gt;Restriction: &lt;/b&gt; This field is returned by the &lt;b&gt; search&lt;/b&gt; method only when &lt;b&gt; fieldgroups&lt;/b&gt; &#x3D; &lt;code&gt;EXTENDED&lt;/code&gt;. | [optional] 
**ThumbnailImages** | Pointer to [**[]Image**](Image.md) | An array of thumbnail images for the item.  | [optional] 
**Title** | Pointer to **string** | The seller-created title of the item. &lt;br&gt;&lt;br&gt;&lt;b&gt;Maximum Length: &lt;/b&gt; 80 characters | [optional] 
**TopRatedBuyingExperience** | Pointer to **bool** | This indicates if the item is a top-rated plus item. There are three benefits of a top-rated plus item: a  minimum 30-day money-back return policy, shipping the item in 1 business day with tracking provided, and the added comfort of knowing that this item is from an experienced seller with the highest buyer ratings. See the &lt;a href&#x3D;\&quot;https://pages.ebay.com/topratedplus/index.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Top Rated Plus Items &lt;/a&gt; and  &lt;a href&#x3D;\&quot;https://pages.ebay.com/help/sell/top-rated.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Becoming a Top Rated Seller and qualifying for Top Rated Plus&lt;/a&gt; help topics for more information. | [optional] 
**TyreLabelImageUrl** | Pointer to **string** | The URL to the image that shows the information on the tyre label. | [optional] 
**UnitPrice** | Pointer to [**ConvertedAmount**](ConvertedAmount.md) |  | [optional] 
**UnitPricingMeasure** | Pointer to **string** | The designation, such as size, weight, volume, count, etc., that was used to specify the quantity of the item. This helps buyers compare prices. &lt;br /&gt;&lt;br /&gt;For example, the following tells the buyer that the item is 7.99 per 100 grams. &lt;br /&gt;&lt;br /&gt;&lt;code&gt;\&quot;unitPricingMeasure\&quot;: \&quot;100g\&quot;,&lt;br /&gt; \&quot;unitPrice\&quot;: {&lt;br /&gt;&amp;nbsp;&amp;nbsp;\&quot;value\&quot;: \&quot;7.99\&quot;,&lt;br /&gt;&amp;nbsp;&amp;nbsp;\&quot;currency\&quot;: \&quot;GBP\&quot;&lt;/code&gt; | [optional] 
**WatchCount** | Pointer to **int32** | The number of users that have added the item to their watch list.&lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt; &lt;strong&gt;Note:&lt;/strong&gt; This field is restricted to applications that have been granted permission to access this feature. You must submit an &lt;a href&#x3D;\&quot;https://developer.ebay.com/my/support/tickets?tab&#x3D;app-check\&quot;&gt;App Check ticket&lt;/a&gt; to request this access. In the App Check form, add a note to the &lt;b&gt;Application Title/Summary&lt;/b&gt; and/or &lt;b&gt;Application Details&lt;/b&gt; fields that you want access to Watch Count data in the Browse API.&lt;/span&gt; | [optional] 

## Methods

### NewItemSummary

`func NewItemSummary() *ItemSummary`

NewItemSummary instantiates a new ItemSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemSummaryWithDefaults

`func NewItemSummaryWithDefaults() *ItemSummary`

NewItemSummaryWithDefaults instantiates a new ItemSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalImages

`func (o *ItemSummary) GetAdditionalImages() []Image`

GetAdditionalImages returns the AdditionalImages field if non-nil, zero value otherwise.

### GetAdditionalImagesOk

`func (o *ItemSummary) GetAdditionalImagesOk() (*[]Image, bool)`

GetAdditionalImagesOk returns a tuple with the AdditionalImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalImages

`func (o *ItemSummary) SetAdditionalImages(v []Image)`

SetAdditionalImages sets AdditionalImages field to given value.

### HasAdditionalImages

`func (o *ItemSummary) HasAdditionalImages() bool`

HasAdditionalImages returns a boolean if a field has been set.

### GetAdultOnly

`func (o *ItemSummary) GetAdultOnly() bool`

GetAdultOnly returns the AdultOnly field if non-nil, zero value otherwise.

### GetAdultOnlyOk

`func (o *ItemSummary) GetAdultOnlyOk() (*bool, bool)`

GetAdultOnlyOk returns a tuple with the AdultOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdultOnly

`func (o *ItemSummary) SetAdultOnly(v bool)`

SetAdultOnly sets AdultOnly field to given value.

### HasAdultOnly

`func (o *ItemSummary) HasAdultOnly() bool`

HasAdultOnly returns a boolean if a field has been set.

### GetAvailableCoupons

`func (o *ItemSummary) GetAvailableCoupons() bool`

GetAvailableCoupons returns the AvailableCoupons field if non-nil, zero value otherwise.

### GetAvailableCouponsOk

`func (o *ItemSummary) GetAvailableCouponsOk() (*bool, bool)`

GetAvailableCouponsOk returns a tuple with the AvailableCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCoupons

`func (o *ItemSummary) SetAvailableCoupons(v bool)`

SetAvailableCoupons sets AvailableCoupons field to given value.

### HasAvailableCoupons

`func (o *ItemSummary) HasAvailableCoupons() bool`

HasAvailableCoupons returns a boolean if a field has been set.

### GetBidCount

`func (o *ItemSummary) GetBidCount() int32`

GetBidCount returns the BidCount field if non-nil, zero value otherwise.

### GetBidCountOk

`func (o *ItemSummary) GetBidCountOk() (*int32, bool)`

GetBidCountOk returns a tuple with the BidCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidCount

`func (o *ItemSummary) SetBidCount(v int32)`

SetBidCount sets BidCount field to given value.

### HasBidCount

`func (o *ItemSummary) HasBidCount() bool`

HasBidCount returns a boolean if a field has been set.

### GetBuyingOptions

`func (o *ItemSummary) GetBuyingOptions() []string`

GetBuyingOptions returns the BuyingOptions field if non-nil, zero value otherwise.

### GetBuyingOptionsOk

`func (o *ItemSummary) GetBuyingOptionsOk() (*[]string, bool)`

GetBuyingOptionsOk returns a tuple with the BuyingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyingOptions

`func (o *ItemSummary) SetBuyingOptions(v []string)`

SetBuyingOptions sets BuyingOptions field to given value.

### HasBuyingOptions

`func (o *ItemSummary) HasBuyingOptions() bool`

HasBuyingOptions returns a boolean if a field has been set.

### GetCategories

`func (o *ItemSummary) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ItemSummary) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ItemSummary) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ItemSummary) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCompatibilityMatch

`func (o *ItemSummary) GetCompatibilityMatch() string`

GetCompatibilityMatch returns the CompatibilityMatch field if non-nil, zero value otherwise.

### GetCompatibilityMatchOk

`func (o *ItemSummary) GetCompatibilityMatchOk() (*string, bool)`

GetCompatibilityMatchOk returns a tuple with the CompatibilityMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityMatch

`func (o *ItemSummary) SetCompatibilityMatch(v string)`

SetCompatibilityMatch sets CompatibilityMatch field to given value.

### HasCompatibilityMatch

`func (o *ItemSummary) HasCompatibilityMatch() bool`

HasCompatibilityMatch returns a boolean if a field has been set.

### GetCompatibilityProperties

`func (o *ItemSummary) GetCompatibilityProperties() []CompatibilityProperty`

GetCompatibilityProperties returns the CompatibilityProperties field if non-nil, zero value otherwise.

### GetCompatibilityPropertiesOk

`func (o *ItemSummary) GetCompatibilityPropertiesOk() (*[]CompatibilityProperty, bool)`

GetCompatibilityPropertiesOk returns a tuple with the CompatibilityProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibilityProperties

`func (o *ItemSummary) SetCompatibilityProperties(v []CompatibilityProperty)`

SetCompatibilityProperties sets CompatibilityProperties field to given value.

### HasCompatibilityProperties

`func (o *ItemSummary) HasCompatibilityProperties() bool`

HasCompatibilityProperties returns a boolean if a field has been set.

### GetCondition

`func (o *ItemSummary) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ItemSummary) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ItemSummary) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ItemSummary) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetConditionId

`func (o *ItemSummary) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *ItemSummary) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *ItemSummary) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *ItemSummary) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### GetCurrentBidPrice

`func (o *ItemSummary) GetCurrentBidPrice() ConvertedAmount`

GetCurrentBidPrice returns the CurrentBidPrice field if non-nil, zero value otherwise.

### GetCurrentBidPriceOk

`func (o *ItemSummary) GetCurrentBidPriceOk() (*ConvertedAmount, bool)`

GetCurrentBidPriceOk returns a tuple with the CurrentBidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBidPrice

`func (o *ItemSummary) SetCurrentBidPrice(v ConvertedAmount)`

SetCurrentBidPrice sets CurrentBidPrice field to given value.

### HasCurrentBidPrice

`func (o *ItemSummary) HasCurrentBidPrice() bool`

HasCurrentBidPrice returns a boolean if a field has been set.

### GetDistanceFromPickupLocation

`func (o *ItemSummary) GetDistanceFromPickupLocation() TargetLocation`

GetDistanceFromPickupLocation returns the DistanceFromPickupLocation field if non-nil, zero value otherwise.

### GetDistanceFromPickupLocationOk

`func (o *ItemSummary) GetDistanceFromPickupLocationOk() (*TargetLocation, bool)`

GetDistanceFromPickupLocationOk returns a tuple with the DistanceFromPickupLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceFromPickupLocation

`func (o *ItemSummary) SetDistanceFromPickupLocation(v TargetLocation)`

SetDistanceFromPickupLocation sets DistanceFromPickupLocation field to given value.

### HasDistanceFromPickupLocation

`func (o *ItemSummary) HasDistanceFromPickupLocation() bool`

HasDistanceFromPickupLocation returns a boolean if a field has been set.

### GetEnergyEfficiencyClass

`func (o *ItemSummary) GetEnergyEfficiencyClass() string`

GetEnergyEfficiencyClass returns the EnergyEfficiencyClass field if non-nil, zero value otherwise.

### GetEnergyEfficiencyClassOk

`func (o *ItemSummary) GetEnergyEfficiencyClassOk() (*string, bool)`

GetEnergyEfficiencyClassOk returns a tuple with the EnergyEfficiencyClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyEfficiencyClass

`func (o *ItemSummary) SetEnergyEfficiencyClass(v string)`

SetEnergyEfficiencyClass sets EnergyEfficiencyClass field to given value.

### HasEnergyEfficiencyClass

`func (o *ItemSummary) HasEnergyEfficiencyClass() bool`

HasEnergyEfficiencyClass returns a boolean if a field has been set.

### GetEpid

`func (o *ItemSummary) GetEpid() string`

GetEpid returns the Epid field if non-nil, zero value otherwise.

### GetEpidOk

`func (o *ItemSummary) GetEpidOk() (*string, bool)`

GetEpidOk returns a tuple with the Epid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpid

`func (o *ItemSummary) SetEpid(v string)`

SetEpid sets Epid field to given value.

### HasEpid

`func (o *ItemSummary) HasEpid() bool`

HasEpid returns a boolean if a field has been set.

### GetImage

`func (o *ItemSummary) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ItemSummary) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ItemSummary) SetImage(v Image)`

SetImage sets Image field to given value.

### HasImage

`func (o *ItemSummary) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetItemAffiliateWebUrl

`func (o *ItemSummary) GetItemAffiliateWebUrl() string`

GetItemAffiliateWebUrl returns the ItemAffiliateWebUrl field if non-nil, zero value otherwise.

### GetItemAffiliateWebUrlOk

`func (o *ItemSummary) GetItemAffiliateWebUrlOk() (*string, bool)`

GetItemAffiliateWebUrlOk returns a tuple with the ItemAffiliateWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemAffiliateWebUrl

`func (o *ItemSummary) SetItemAffiliateWebUrl(v string)`

SetItemAffiliateWebUrl sets ItemAffiliateWebUrl field to given value.

### HasItemAffiliateWebUrl

`func (o *ItemSummary) HasItemAffiliateWebUrl() bool`

HasItemAffiliateWebUrl returns a boolean if a field has been set.

### GetItemEndDate

`func (o *ItemSummary) GetItemEndDate() string`

GetItemEndDate returns the ItemEndDate field if non-nil, zero value otherwise.

### GetItemEndDateOk

`func (o *ItemSummary) GetItemEndDateOk() (*string, bool)`

GetItemEndDateOk returns a tuple with the ItemEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemEndDate

`func (o *ItemSummary) SetItemEndDate(v string)`

SetItemEndDate sets ItemEndDate field to given value.

### HasItemEndDate

`func (o *ItemSummary) HasItemEndDate() bool`

HasItemEndDate returns a boolean if a field has been set.

### GetItemGroupHref

`func (o *ItemSummary) GetItemGroupHref() string`

GetItemGroupHref returns the ItemGroupHref field if non-nil, zero value otherwise.

### GetItemGroupHrefOk

`func (o *ItemSummary) GetItemGroupHrefOk() (*string, bool)`

GetItemGroupHrefOk returns a tuple with the ItemGroupHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGroupHref

`func (o *ItemSummary) SetItemGroupHref(v string)`

SetItemGroupHref sets ItemGroupHref field to given value.

### HasItemGroupHref

`func (o *ItemSummary) HasItemGroupHref() bool`

HasItemGroupHref returns a boolean if a field has been set.

### GetItemGroupType

`func (o *ItemSummary) GetItemGroupType() string`

GetItemGroupType returns the ItemGroupType field if non-nil, zero value otherwise.

### GetItemGroupTypeOk

`func (o *ItemSummary) GetItemGroupTypeOk() (*string, bool)`

GetItemGroupTypeOk returns a tuple with the ItemGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemGroupType

`func (o *ItemSummary) SetItemGroupType(v string)`

SetItemGroupType sets ItemGroupType field to given value.

### HasItemGroupType

`func (o *ItemSummary) HasItemGroupType() bool`

HasItemGroupType returns a boolean if a field has been set.

### GetItemHref

`func (o *ItemSummary) GetItemHref() string`

GetItemHref returns the ItemHref field if non-nil, zero value otherwise.

### GetItemHrefOk

`func (o *ItemSummary) GetItemHrefOk() (*string, bool)`

GetItemHrefOk returns a tuple with the ItemHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemHref

`func (o *ItemSummary) SetItemHref(v string)`

SetItemHref sets ItemHref field to given value.

### HasItemHref

`func (o *ItemSummary) HasItemHref() bool`

HasItemHref returns a boolean if a field has been set.

### GetItemId

`func (o *ItemSummary) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemSummary) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemSummary) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ItemSummary) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetItemLocation

`func (o *ItemSummary) GetItemLocation() ItemLocationImpl`

GetItemLocation returns the ItemLocation field if non-nil, zero value otherwise.

### GetItemLocationOk

`func (o *ItemSummary) GetItemLocationOk() (*ItemLocationImpl, bool)`

GetItemLocationOk returns a tuple with the ItemLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLocation

`func (o *ItemSummary) SetItemLocation(v ItemLocationImpl)`

SetItemLocation sets ItemLocation field to given value.

### HasItemLocation

`func (o *ItemSummary) HasItemLocation() bool`

HasItemLocation returns a boolean if a field has been set.

### GetItemWebUrl

`func (o *ItemSummary) GetItemWebUrl() string`

GetItemWebUrl returns the ItemWebUrl field if non-nil, zero value otherwise.

### GetItemWebUrlOk

`func (o *ItemSummary) GetItemWebUrlOk() (*string, bool)`

GetItemWebUrlOk returns a tuple with the ItemWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemWebUrl

`func (o *ItemSummary) SetItemWebUrl(v string)`

SetItemWebUrl sets ItemWebUrl field to given value.

### HasItemWebUrl

`func (o *ItemSummary) HasItemWebUrl() bool`

HasItemWebUrl returns a boolean if a field has been set.

### GetLegacyItemId

`func (o *ItemSummary) GetLegacyItemId() string`

GetLegacyItemId returns the LegacyItemId field if non-nil, zero value otherwise.

### GetLegacyItemIdOk

`func (o *ItemSummary) GetLegacyItemIdOk() (*string, bool)`

GetLegacyItemIdOk returns a tuple with the LegacyItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyItemId

`func (o *ItemSummary) SetLegacyItemId(v string)`

SetLegacyItemId sets LegacyItemId field to given value.

### HasLegacyItemId

`func (o *ItemSummary) HasLegacyItemId() bool`

HasLegacyItemId returns a boolean if a field has been set.

### GetMarketingPrice

`func (o *ItemSummary) GetMarketingPrice() MarketingPrice`

GetMarketingPrice returns the MarketingPrice field if non-nil, zero value otherwise.

### GetMarketingPriceOk

`func (o *ItemSummary) GetMarketingPriceOk() (*MarketingPrice, bool)`

GetMarketingPriceOk returns a tuple with the MarketingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingPrice

`func (o *ItemSummary) SetMarketingPrice(v MarketingPrice)`

SetMarketingPrice sets MarketingPrice field to given value.

### HasMarketingPrice

`func (o *ItemSummary) HasMarketingPrice() bool`

HasMarketingPrice returns a boolean if a field has been set.

### GetPickupOptions

`func (o *ItemSummary) GetPickupOptions() []PickupOptionSummary`

GetPickupOptions returns the PickupOptions field if non-nil, zero value otherwise.

### GetPickupOptionsOk

`func (o *ItemSummary) GetPickupOptionsOk() (*[]PickupOptionSummary, bool)`

GetPickupOptionsOk returns a tuple with the PickupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupOptions

`func (o *ItemSummary) SetPickupOptions(v []PickupOptionSummary)`

SetPickupOptions sets PickupOptions field to given value.

### HasPickupOptions

`func (o *ItemSummary) HasPickupOptions() bool`

HasPickupOptions returns a boolean if a field has been set.

### GetPrice

`func (o *ItemSummary) GetPrice() ConvertedAmount`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ItemSummary) GetPriceOk() (*ConvertedAmount, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ItemSummary) SetPrice(v ConvertedAmount)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ItemSummary) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceDisplayCondition

`func (o *ItemSummary) GetPriceDisplayCondition() string`

GetPriceDisplayCondition returns the PriceDisplayCondition field if non-nil, zero value otherwise.

### GetPriceDisplayConditionOk

`func (o *ItemSummary) GetPriceDisplayConditionOk() (*string, bool)`

GetPriceDisplayConditionOk returns a tuple with the PriceDisplayCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDisplayCondition

`func (o *ItemSummary) SetPriceDisplayCondition(v string)`

SetPriceDisplayCondition sets PriceDisplayCondition field to given value.

### HasPriceDisplayCondition

`func (o *ItemSummary) HasPriceDisplayCondition() bool`

HasPriceDisplayCondition returns a boolean if a field has been set.

### GetPriorityListing

`func (o *ItemSummary) GetPriorityListing() bool`

GetPriorityListing returns the PriorityListing field if non-nil, zero value otherwise.

### GetPriorityListingOk

`func (o *ItemSummary) GetPriorityListingOk() (*bool, bool)`

GetPriorityListingOk returns a tuple with the PriorityListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityListing

`func (o *ItemSummary) SetPriorityListing(v bool)`

SetPriorityListing sets PriorityListing field to given value.

### HasPriorityListing

`func (o *ItemSummary) HasPriorityListing() bool`

HasPriorityListing returns a boolean if a field has been set.

### GetQualifiedPrograms

`func (o *ItemSummary) GetQualifiedPrograms() []string`

GetQualifiedPrograms returns the QualifiedPrograms field if non-nil, zero value otherwise.

### GetQualifiedProgramsOk

`func (o *ItemSummary) GetQualifiedProgramsOk() (*[]string, bool)`

GetQualifiedProgramsOk returns a tuple with the QualifiedPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedPrograms

`func (o *ItemSummary) SetQualifiedPrograms(v []string)`

SetQualifiedPrograms sets QualifiedPrograms field to given value.

### HasQualifiedPrograms

`func (o *ItemSummary) HasQualifiedPrograms() bool`

HasQualifiedPrograms returns a boolean if a field has been set.

### GetSeller

`func (o *ItemSummary) GetSeller() Seller`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *ItemSummary) GetSellerOk() (*Seller, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *ItemSummary) SetSeller(v Seller)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *ItemSummary) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetShippingOptions

`func (o *ItemSummary) GetShippingOptions() []ShippingOptionSummary`

GetShippingOptions returns the ShippingOptions field if non-nil, zero value otherwise.

### GetShippingOptionsOk

`func (o *ItemSummary) GetShippingOptionsOk() (*[]ShippingOptionSummary, bool)`

GetShippingOptionsOk returns a tuple with the ShippingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingOptions

`func (o *ItemSummary) SetShippingOptions(v []ShippingOptionSummary)`

SetShippingOptions sets ShippingOptions field to given value.

### HasShippingOptions

`func (o *ItemSummary) HasShippingOptions() bool`

HasShippingOptions returns a boolean if a field has been set.

### GetShortDescription

`func (o *ItemSummary) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *ItemSummary) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *ItemSummary) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *ItemSummary) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetThumbnailImages

`func (o *ItemSummary) GetThumbnailImages() []Image`

GetThumbnailImages returns the ThumbnailImages field if non-nil, zero value otherwise.

### GetThumbnailImagesOk

`func (o *ItemSummary) GetThumbnailImagesOk() (*[]Image, bool)`

GetThumbnailImagesOk returns a tuple with the ThumbnailImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailImages

`func (o *ItemSummary) SetThumbnailImages(v []Image)`

SetThumbnailImages sets ThumbnailImages field to given value.

### HasThumbnailImages

`func (o *ItemSummary) HasThumbnailImages() bool`

HasThumbnailImages returns a boolean if a field has been set.

### GetTitle

`func (o *ItemSummary) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemSummary) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemSummary) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ItemSummary) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTopRatedBuyingExperience

`func (o *ItemSummary) GetTopRatedBuyingExperience() bool`

GetTopRatedBuyingExperience returns the TopRatedBuyingExperience field if non-nil, zero value otherwise.

### GetTopRatedBuyingExperienceOk

`func (o *ItemSummary) GetTopRatedBuyingExperienceOk() (*bool, bool)`

GetTopRatedBuyingExperienceOk returns a tuple with the TopRatedBuyingExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopRatedBuyingExperience

`func (o *ItemSummary) SetTopRatedBuyingExperience(v bool)`

SetTopRatedBuyingExperience sets TopRatedBuyingExperience field to given value.

### HasTopRatedBuyingExperience

`func (o *ItemSummary) HasTopRatedBuyingExperience() bool`

HasTopRatedBuyingExperience returns a boolean if a field has been set.

### GetTyreLabelImageUrl

`func (o *ItemSummary) GetTyreLabelImageUrl() string`

GetTyreLabelImageUrl returns the TyreLabelImageUrl field if non-nil, zero value otherwise.

### GetTyreLabelImageUrlOk

`func (o *ItemSummary) GetTyreLabelImageUrlOk() (*string, bool)`

GetTyreLabelImageUrlOk returns a tuple with the TyreLabelImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTyreLabelImageUrl

`func (o *ItemSummary) SetTyreLabelImageUrl(v string)`

SetTyreLabelImageUrl sets TyreLabelImageUrl field to given value.

### HasTyreLabelImageUrl

`func (o *ItemSummary) HasTyreLabelImageUrl() bool`

HasTyreLabelImageUrl returns a boolean if a field has been set.

### GetUnitPrice

`func (o *ItemSummary) GetUnitPrice() ConvertedAmount`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *ItemSummary) GetUnitPriceOk() (*ConvertedAmount, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *ItemSummary) SetUnitPrice(v ConvertedAmount)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *ItemSummary) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetUnitPricingMeasure

`func (o *ItemSummary) GetUnitPricingMeasure() string`

GetUnitPricingMeasure returns the UnitPricingMeasure field if non-nil, zero value otherwise.

### GetUnitPricingMeasureOk

`func (o *ItemSummary) GetUnitPricingMeasureOk() (*string, bool)`

GetUnitPricingMeasureOk returns a tuple with the UnitPricingMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPricingMeasure

`func (o *ItemSummary) SetUnitPricingMeasure(v string)`

SetUnitPricingMeasure sets UnitPricingMeasure field to given value.

### HasUnitPricingMeasure

`func (o *ItemSummary) HasUnitPricingMeasure() bool`

HasUnitPricingMeasure returns a boolean if a field has been set.

### GetWatchCount

`func (o *ItemSummary) GetWatchCount() int32`

GetWatchCount returns the WatchCount field if non-nil, zero value otherwise.

### GetWatchCountOk

`func (o *ItemSummary) GetWatchCountOk() (*int32, bool)`

GetWatchCountOk returns a tuple with the WatchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchCount

`func (o *ItemSummary) SetWatchCount(v int32)`

SetWatchCount sets WatchCount field to given value.

### HasWatchCount

`func (o *ItemSummary) HasWatchCount() bool`

HasWatchCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


