# CoreItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalImages** | Pointer to [**[]Image**](Image.md) | An array of containers with the URLs for the images that are in addition to the primary image.  The primary image is returned in the &lt;b&gt; image.imageUrl&lt;/b&gt; field. | [optional] 
**AdultOnly** | Pointer to **bool** | This indicates if the item is for  adults only. For more information about adult-only items on eBay, see &lt;a href&#x3D;\&quot;https://pages.ebay.com/help/policies/adult-only.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Adult items policy&lt;/a&gt; for sellers and &lt;a href&#x3D;\&quot;https://www.ebay.com/help/terms-conditions/default/searching-adult-items?id&#x3D;4661\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Adult-Only items on eBay&lt;/a&gt; for buyers. | [optional] 
**AgeGroup** | Pointer to **string** | (Primary Item Aspect) The age group for which the product is recommended. For example, newborn, infant, toddler, kids, adult, etc. All the item aspects, including this aspect, are returned in the &lt;b&gt; localizedAspects&lt;/b&gt; container. | [optional] 
**AuthenticityGuarantee** | Pointer to [**AuthenticityGuaranteeProgram**](AuthenticityGuaranteeProgram.md) |  | [optional] 
**AuthenticityVerification** | Pointer to [**AuthenticityVerificationProgram**](AuthenticityVerificationProgram.md) |  | [optional] 
**AvailableCoupons** | Pointer to [**[]AvailableCoupon**](AvailableCoupon.md) | A list of available coupons for the item. | [optional] 
**BidCount** | Pointer to **int32** | This integer value indicates the total number of bids that have been placed against an auction item. This field is returned only for auction items. | [optional] 
**Brand** | Pointer to **string** | (Primary Item Aspect) The name brand of the item, such as Nike, Apple, etc.  All the item aspects, including this aspect, are returned in the &lt;b&gt; localizedAspects&lt;/b&gt; container. | [optional] 
**BuyingOptions** | Pointer to **[]string** | A comma separated list of all the purchase options available for the item. The values returned are:  &lt;ul&gt; &lt;li&gt;&lt;code&gt;FIXED_PRICE&lt;/code&gt; - Indicates the buyer can purchase the item for a set price using the Buy It Now button. &lt;/li&gt;  &lt;li&gt;&lt;code&gt;AUCTION&lt;/code&gt; - Indicates the buyer can place a bid for the item. After the first bid is placed, this becomes a live auction item and is the only buying option for this item.&lt;/li&gt;  &lt;li&gt;&lt;code&gt;BEST_OFFER&lt;/code&gt; - Indicates the buyer can send the seller a price they&#39;re willing to pay for the item. The seller can accept, reject, or send a counter offer. For more information on how this works, see &lt;a href&#x3D;\&quot;https://www.ebay.com/help/buying/buy-now/making-best-offer?id&#x3D;4019\&quot;&gt;Making a Best Offer&lt;/a&gt;.&lt;/li&gt; &lt;/ul&gt; Code so that your app gracefully handles any future changes to this list. | [optional] 
**CategoryId** | Pointer to **string** | The ID of the leaf category for this item. A leaf category is the lowest level in that category and has no children. | [optional] 
**CategoryPath** | Pointer to **string** | Text that shows the category hierarchy of the item. For example: Computers/Tablets &amp;amp; Networking, Laptops &amp;amp; Netbooks, PC Laptops &amp;amp; Netbooks | [optional] 
**Color** | Pointer to **string** | (Primary Item Aspect) Text describing the color of the item.  All the item aspects, including this aspect, are returned in the &lt;b&gt; localizedAspects&lt;/b&gt; container. | [optional] 
**Condition** | Pointer to **string** | A short text description for the condition of the item, such as New or Used. For a list of condition names, see &lt;a href&#x3D;\&quot;https://developer.ebay.com/devzone/finding/callref/enums/conditionIdList.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Item Condition IDs and Names&lt;/a&gt;.  &lt;br /&gt;&lt;br /&gt;Code so that your app gracefully handles any future changes to this list. &lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt; Note: &lt;/b&gt;As of September 1, 2021, condition ID 2500 (&#39;Seller Refurbished&#39;) is no longer a valid item condition in the &lt;b&gt;Cell Phones &amp; Smartphones&lt;/b&gt; category (category ID 9355) for the following marketplaces: US, Canada, UK, Germany, and Australia. This refurbished item condition has been replaced by three new refurbished values, which include &#39;Excellent - Refurbished&#39; (condition ID 2010), &#39;Very Good - Refurbished&#39; (condition ID 2020), and &#39;Good - Refurbished&#39; (condition ID 2030).&lt;/span&gt; | [optional] 
**ConditionDescription** | Pointer to **string** | A full text description for the condition of the item. This field elaborates on the value specified in the &lt;b&gt;condition&lt;/b&gt; field and provides full details for the condition of the item. &lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt; Note: &lt;/b&gt;As of September 1, 2021, condition ID 2500 (&#39;Seller Refurbished&#39;) is no longer a valid item condition in the &lt;b&gt;Cell Phones &amp; Smartphones&lt;/b&gt; category (category ID 9355) for the following marketplaces: US, Canada, UK, Germany, and Australia. This refurbished item condition has been replaced by three new refurbished values, which include &#39;Excellent - Refurbished&#39; (condition ID 2010), &#39;Very Good - Refurbished&#39; (condition ID 2020), and &#39;Good - Refurbished&#39; (condition ID 2030).&lt;/span&gt; | [optional] 
**ConditionId** | Pointer to **string** | The identifier of the condition of the item. For example, 1000 is the identifier for NEW. For a list of condition names and IDs, see &lt;a href&#x3D;\&quot;https://developer.ebay.com/devzone/finding/callref/enums/conditionIdList.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Item Condition IDs and Names&lt;/a&gt;. &lt;br /&gt;&lt;br /&gt;Code so that your app gracefully handles any future changes to this list. &lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt; Note: &lt;/b&gt;As of September 1, 2021, condition ID 2500 (&#39;Seller Refurbished&#39;) is no longer a valid item condition in the &lt;b&gt;Cell Phones &amp; Smartphones&lt;/b&gt; category (category ID 9355) for the following marketplaces: US, Canada, UK, Germany, and Australia. This refurbished item condition has been replaced by three new refurbished values, which include &#39;Excellent - Refurbished&#39; (condition ID 2010), &#39;Very Good - Refurbished&#39; (condition ID 2020), and &#39;Good - Refurbished&#39; (condition ID 2030).&lt;/span&gt; | [optional] 
**CurrentBidPrice** | Pointer to [**ConvertedAmount**](ConvertedAmount.md) |  | [optional] 
**Description** | Pointer to **string** | The full description of the item that was created by the seller. This can be plain text or rich content and can be very large. | [optional] 
**EligibleForInlineCheckout** | Pointer to **bool** | This field indicates if the item can be purchased using the Buy &lt;a href&#x3D;\&quot;/api-docs/buy/order/resources/methods\&quot;&gt;Order API&lt;/a&gt;. &lt;ul&gt; &lt;li&gt;If the value of this field is &lt;code&gt;true&lt;/code&gt;, this indicates that the item can be purchased using the &lt;b&gt; Order API&lt;/b&gt;. &lt;/li&gt;  &lt;li&gt;If the value of this field is &lt;code&gt;false&lt;/code&gt;, this indicates that the item cannot be purchased using the &lt;b&gt; Order API&lt;/b&gt; and must be purchased on the eBay site.&lt;/li&gt; &lt;/ul&gt; | [optional] 
**EnabledForGuestCheckout** | Pointer to **bool** | This indicates if the item can be purchased using Guest Checkout in the &lt;a href&#x3D;\&quot;/api-docs/buy/order/resources/methods\&quot;&gt;Order API&lt;/a&gt;. You can use this flag to exclude items from your inventory that are not eligible for Guest Checkout, such as gift cards. | [optional] 
**EnergyEfficiencyClass** | Pointer to **string** | This indicates the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/European_Union_energy_label\&quot;&gt;European energy efficiency&lt;/a&gt; rating (EEK) of the item. This field is returned only if the seller specified the energy efficiency rating. &lt;br /&gt;&lt;br /&gt;The rating is a set of energy efficiency classes from A to G, where &#39;A&#39; is the most energy efficient and &#39;G&#39; is the least efficient. This rating helps buyers choose between various models. &lt;br /&gt;&lt;br /&gt;When the manufacturer&#39;s specifications for this item are available, the link to this information is returned in the &lt;b&gt; productFicheWebUrl&lt;/b&gt; field. | [optional] 
**Epid** | Pointer to **string** | An EPID is the eBay product identifier of a product from the eBay product catalog.  This indicates the product in which the item belongs. | [optional] 
**EstimatedAvailabilities** | Pointer to [**[]EstimatedAvailability**](EstimatedAvailability.md) | The estimated number of this item that are available for purchase. Because the quantity of an item can change several times within a second, it is impossible to return the exact quantity. So instead of returning quantity, the estimated availability of the item is returned. | [optional] 
**Gender** | Pointer to **string** | (Primary Item Aspect) The gender for the item. This is used for items that could vary by gender, such as clothing. For example: male, female, or unisex. All the item aspects, including this aspect, are returned in the &lt;b&gt; localizedAspects&lt;/b&gt; container. | [optional] 
**Gtin** | Pointer to **string** | The unique Global Trade Item number of the item as defined by &lt;a href&#x3D;\&quot;https://www.gtin.info\&quot; target&#x3D;\&quot;_blank\&quot;&gt;https://www.gtin.info&lt;/a&gt;. This can be a UPC (Universal Product Code), EAN (European Article Number), or an ISBN (International Standard Book Number) value. | [optional] 
**Image** | Pointer to [**Image**](Image.md) |  | [optional] 
**InferredEpid** | Pointer to **string** | The ePID (eBay Product ID of a product from the eBay product catalog) for the item, which has been programmatically determined by eBay using the item&#39;s title, aspects, and other data. &lt;br /&gt;&lt;br /&gt;If the seller provided an ePID for the item, the seller&#39;s value is returned in the &lt;b&gt; epid&lt;/b&gt; field. &lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt; Note: &lt;/b&gt; This field is returned only for authorized Partners.&lt;/span&gt; | [optional] 
**ItemAffiliateWebUrl** | Pointer to **string** | The URL of the View Item page of the item, which includes the affiliate tracking ID. This field is only returned if the eBay partner enables affiliate tracking for the item by including the  &lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#Headers\&quot;&gt;&lt;code&gt;X-EBAY-C-ENDUSERCTX&lt;/code&gt;&lt;/a&gt; request header in the method.  &lt;br /&gt; &lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note: &lt;/b&gt; eBay Partner Network, in order to be commissioned for your sales, you must use this URL to forward your buyer to the ebay.com site. &lt;/span&gt; | [optional] 
**ItemEndDate** | Pointer to **string** | This timestamp indicates the date and time up to which the item can be purchased.  This value is returned in UTC format (yyyy-MM-ddThh:mm:ss.sssZ), which you can convert into the local time of the buyer.&lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt; Note: &lt;/b&gt;This field is only returned for auction listings.&lt;/span&gt; | [optional] 
**ItemId** | Pointer to **string** | The unique RESTful identifier of the item. | [optional] 
**ItemLocation** | Pointer to [**Address**](Address.md) |  | [optional] 
**ItemWebUrl** | Pointer to **string** | The URL of the View Item page of the item. This enables you to include a \&quot;Report Item on eBay\&quot; link that takes the buyer to the View Item page on eBay. From there they can report any issues regarding this item to eBay. | [optional] 
**LegacyItemId** | Pointer to **string** | The unique identifier of the eBay listing that contains the item. This is the traditional/legacy ID that is often seen in the URL of the listing View Item page. | [optional] 
**LocalizedAspects** | Pointer to [**[]TypedNameValue**](TypedNameValue.md) | An array of containers that show the complete list of the aspect name/value pairs that describe the variation of the item. | [optional] 
**LotSize** | Pointer to **int32** | The number of items in a lot. In other words, a lot size is the number of items that are being sold together.  &lt;br /&gt;&lt;br /&gt;A lot is a set of two or more items included in a single listing that must be purchased together in a single order line item. All the items in the lot are the same but there can be multiple items in a single lot,  such as the package of batteries shown in the example below.   &lt;br /&gt;&lt;br /&gt;&lt;table border&#x3D;\&quot;1\&quot;&gt; &lt;tr&gt; &lt;tr&gt;  &lt;th&gt;Item&lt;/th&gt;  &lt;th&gt;Lot Definition&lt;/th&gt; &lt;th&gt;Lot Size&lt;/th&gt;&lt;/tr&gt;  &lt;tr&gt;  &lt;td&gt;A package of 24 AA batteries&lt;/td&gt;  &lt;td&gt;A box of 10 packages&lt;/td&gt;  &lt;td&gt;10  &lt;/td&gt; &lt;/tr&gt;  &lt;tr&gt;  &lt;td&gt;A P235/75-15 Goodyear tire &lt;/td&gt;  &lt;td&gt;4 tires  &lt;/td&gt;  &lt;td&gt;4  &lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;Fashion Jewelry Rings  &lt;/td&gt; &lt;td&gt;Package of 100 assorted rings  &lt;/td&gt; &lt;td&gt;100 &lt;/td&gt; &lt;/tr&gt;&lt;/table&gt;  &lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note: &lt;/b&gt;  Lots are not supported in all categories.  &lt;/span&gt; | [optional] 
**MarketingPrice** | Pointer to [**MarketingPrice**](MarketingPrice.md) |  | [optional] 
**Material** | Pointer to **string** | (Primary Item Aspect) Text describing what the item is made of. For example, silk. All the item aspects, including this aspect, are returned in the &lt;b&gt; localizedAspects&lt;/b&gt; container. | [optional] 
**MinimumPriceToBid** | Pointer to [**ConvertedAmount**](ConvertedAmount.md) |  | [optional] 
**Mpn** | Pointer to **string** | The manufacturer&#39;s part number, which is a unique number that identifies a specific product. To identify the product, this is always used along with brand. | [optional] 
**Pattern** | Pointer to **string** | (Primary Item Aspect) Text describing the pattern used on the item. For example, paisley. All the item aspects, including this aspect, are returned in the &lt;b&gt; localizedAspects&lt;/b&gt; container. | [optional] 
**PaymentMethods** | Pointer to [**[]PaymentMethod**](PaymentMethod.md) | The payment methods for the item, including the payment method types, brands, and instructions for the buyer. | [optional] 
**Price** | Pointer to [**ConvertedAmount**](ConvertedAmount.md) |  | [optional] 
**PriceDisplayCondition** | Pointer to **string** | Indicates when in the buying flow the item&#39;s price can appear for minimum advertised price (MAP) items, which is the lowest price a retailer can advertise/show for this item. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:PriceDisplayConditionEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**PrimaryItemGroup** | Pointer to [**ItemGroupSummary**](ItemGroupSummary.md) |  | [optional] 
**PrimaryProductReviewRating** | Pointer to [**ReviewRating**](ReviewRating.md) |  | [optional] 
**PriorityListing** | Pointer to **bool** | This field is returned as &lt;code&gt;true&lt;/code&gt; if the listing is part of a Promoted Listing campaign. Promoted Listings are available to Above Standard and Top Rated sellers with recent sales activity.&lt;br /&gt;&lt;br /&gt;For more information, see &lt;a href&#x3D;\&quot;https://pages.ebay.com/seller-center/listing-and-marketing/promoted-listings.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Promoted Listings&lt;/a&gt;. | [optional] 
**Product** | Pointer to [**Product**](Product.md) |  | [optional] 
**ProductFicheWebUrl** | Pointer to **string** | The URL of a page containing the manufacturer&#39;s specification of this item, which helps buyers make a purchasing decision. This information is available only for items that include the European energy efficiency rating (EEK) but is not available for &lt;em&gt; all&lt;/em&gt; items with an EEK rating and is returned only if this information is available. The EEK rating of the item is returned in the &lt;b&gt; energyEfficiencyClass&lt;/b&gt; field. | [optional] 
**QualifiedPrograms** | Pointer to **[]string** | An array of the qualified programs available for the item, such as EBAY_PLUS, AUTHENTICITY_GUARANTEE, and AUTHENTICITY_VERIFICATION.&lt;br /&gt;&lt;br /&gt;eBay Plus is a premium account option for buyers, which provides benefits such as fast free domestic shipping and free returns on selected items. Top-Rated eBay sellers must opt in to eBay Plus to be able to offer the program on qualifying listings. Sellers must commit to next-day delivery of those items.&lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note: &lt;/b&gt; eBay Plus is available only to buyers in Germany, Austria, and Australia marketplaces.&lt;/span&gt;&lt;br /&gt;&lt;br /&gt;The eBay &lt;a href&#x3D;\&quot;https://pages.ebay.com/authenticity-guarantee/\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Authenticity Guarantee&lt;/a&gt; program enables third-party authenticators to perform authentication verification inspections on items such as watches and sneakers. | [optional] 
**QuantityLimitPerBuyer** | Pointer to **int32** | The maximum number for a specific item that one buyer can purchase. | [optional] 
**ReservePriceMet** | Pointer to **bool** | This indicates if the reserve price of the item has been met. A reserve price is set by the seller and is the minimum amount the seller is willing to sell the item for. &lt;p&gt;If the highest bid is not equal to or higher than the reserve price when the auction ends, the listing ends and the item is not sold.&lt;/p&gt; &lt;p&gt;&lt;b&gt; Note: &lt;/b&gt;This is returned only for auctions that have a reserve price.&lt;/p&gt; | [optional] 
**ReturnTerms** | Pointer to [**ItemReturnTerms**](ItemReturnTerms.md) |  | [optional] 
**Seller** | Pointer to [**SellerDetail**](SellerDetail.md) |  | [optional] 
**SellerItemRevision** | Pointer to **string** | An identifier generated/incremented when a seller revises the item. There are two types of item revisions: &lt;ul&gt;&lt;li&gt;Seller changes, such as changing the title&lt;/li&gt;  &lt;li&gt;eBay system changes, such as changing the quantity when an item is purchased&lt;/li&gt;&lt;/ul&gt; This ID is changed &lt;em&gt; only&lt;/em&gt; when the seller makes a change to the item. This means you cannot use this value to determine if the quantity has changed. | [optional] 
**ShippingOptions** | Pointer to [**[]ShippingOption**](ShippingOption.md) | An array of shipping options containers that have the details about cost, carrier, etc. of one shipping option.  | [optional] 
**ShipToLocations** | Pointer to [**ShipToLocations**](ShipToLocations.md) |  | [optional] 
**ShortDescription** | Pointer to **string** | This text string is derived from the item condition and the item aspects (such as size, color, capacity, model, brand, etc.). | [optional] 
**Size** | Pointer to **string** | (Primary Item Aspect) The size of the item. For example, &#39;7&#39; for a size 7 shoe. All the item aspects, including this aspect, are returned in the &lt;b&gt; localizedAspects&lt;/b&gt; container. | [optional] 
**SizeSystem** | Pointer to **string** | (Primary Item Aspect) The sizing system of the country.  All the item aspects, including this aspect, are returned in the &lt;b&gt; localizedAspects&lt;/b&gt; container. &lt;br /&gt;&lt;br /&gt;&lt;b&gt; Valid Values: &lt;/b&gt; &lt;br /&gt;AU (Australia),  &lt;br /&gt;BR (Brazil), &lt;br /&gt;CN (China),  &lt;br /&gt;DE (Germany),  &lt;br /&gt;EU (European Union),  &lt;br /&gt; FR (France), &lt;br /&gt; IT (Italy),  &lt;br /&gt;JP (Japan), &lt;br /&gt;MX (Mexico),  &lt;br /&gt;US (USA), &lt;br /&gt; UK (United Kingdom) &lt;br /&gt;&lt;br /&gt;Code so that your app gracefully handles any future changes to this list.  | [optional] 
**SizeType** | Pointer to **string** | (Primary Item Aspect) Text describing a size group in which the item would be included, such as regular, petite, plus, big-and-tall or maternity. All the item aspects, including this aspect, are returned in the &lt;b&gt; localizedAspects&lt;/b&gt; container. | [optional] 
**Subtitle** | Pointer to **string** | A subtitle is optional and allows the seller to provide more information about the product, possibly including keywords that may assist with search results. | [optional] 
**Taxes** | Pointer to [**[]Taxes**](Taxes.md) | The container for the tax information for the item. | [optional] 
**Title** | Pointer to **string** | The seller-created title of the item. &lt;br&gt;&lt;br&gt;&lt;b&gt; Maximum Length: &lt;/b&gt; 80 characters | [optional] 
**TopRatedBuyingExperience** | Pointer to **bool** | This indicates if the item a top-rated plus item. There are three benefits of a top-rated plus item; a  minimum 30-day money-back return policy, shipping the items in 1 business day with tracking provided, and the added comfort of knowing this item is from experienced sellers with the highest buyer ratings. See the &lt;a href&#x3D;\&quot;https://pages.ebay.com/topratedplus/index.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;Top Rated Plus Items &lt;/a&gt; and  &lt;a href&#x3D;\&quot;https://pages.ebay.com/help/sell/top-rated.html target&#x3D;\&quot;_blank\&quot;&gt;Becoming a Top Rated Seller and qualifying for Top Rated Plus&lt;/a&gt; help topics for more information. | [optional] 
**TyreLabelImageUrl** | Pointer to **string** | The URL to the image that shows the information on the tyre label. | [optional] 
**UniqueBidderCount** | Pointer to **int32** | This integer value indicates the number of different eBay users who have placed one or more bids on an auction item. This field is only applicable to auction items. | [optional] 
**UnitPrice** | Pointer to [**ConvertedAmount**](ConvertedAmount.md) |  | [optional] 
**UnitPricingMeasure** | Pointer to **string** | The designation, such as size, weight, volume, count, etc., that was used to specify the quantity of the item.  This helps buyers compare prices. &lt;br /&gt;&lt;br /&gt;For example, the following tells the buyer that the item is 7.99 per 100 grams. &lt;br /&gt;&lt;br /&gt;&lt;code&gt;\&quot;unitPricingMeasure\&quot;: \&quot;100g\&quot;,&lt;br /&gt; \&quot;unitPrice\&quot;: {&lt;br /&gt;&amp;nbsp;&amp;nbsp;\&quot;value\&quot;: \&quot;7.99\&quot;,&lt;br /&gt;&amp;nbsp;&amp;nbsp;\&quot;currency\&quot;: \&quot;GBP\&quot;&lt;/code&gt; | [optional] 

## Methods

### NewCoreItem

`func NewCoreItem() *CoreItem`

NewCoreItem instantiates a new CoreItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoreItemWithDefaults

`func NewCoreItemWithDefaults() *CoreItem`

NewCoreItemWithDefaults instantiates a new CoreItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalImages

`func (o *CoreItem) GetAdditionalImages() []Image`

GetAdditionalImages returns the AdditionalImages field if non-nil, zero value otherwise.

### GetAdditionalImagesOk

`func (o *CoreItem) GetAdditionalImagesOk() (*[]Image, bool)`

GetAdditionalImagesOk returns a tuple with the AdditionalImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalImages

`func (o *CoreItem) SetAdditionalImages(v []Image)`

SetAdditionalImages sets AdditionalImages field to given value.

### HasAdditionalImages

`func (o *CoreItem) HasAdditionalImages() bool`

HasAdditionalImages returns a boolean if a field has been set.

### GetAdultOnly

`func (o *CoreItem) GetAdultOnly() bool`

GetAdultOnly returns the AdultOnly field if non-nil, zero value otherwise.

### GetAdultOnlyOk

`func (o *CoreItem) GetAdultOnlyOk() (*bool, bool)`

GetAdultOnlyOk returns a tuple with the AdultOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdultOnly

`func (o *CoreItem) SetAdultOnly(v bool)`

SetAdultOnly sets AdultOnly field to given value.

### HasAdultOnly

`func (o *CoreItem) HasAdultOnly() bool`

HasAdultOnly returns a boolean if a field has been set.

### GetAgeGroup

`func (o *CoreItem) GetAgeGroup() string`

GetAgeGroup returns the AgeGroup field if non-nil, zero value otherwise.

### GetAgeGroupOk

`func (o *CoreItem) GetAgeGroupOk() (*string, bool)`

GetAgeGroupOk returns a tuple with the AgeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeGroup

`func (o *CoreItem) SetAgeGroup(v string)`

SetAgeGroup sets AgeGroup field to given value.

### HasAgeGroup

`func (o *CoreItem) HasAgeGroup() bool`

HasAgeGroup returns a boolean if a field has been set.

### GetAuthenticityGuarantee

`func (o *CoreItem) GetAuthenticityGuarantee() AuthenticityGuaranteeProgram`

GetAuthenticityGuarantee returns the AuthenticityGuarantee field if non-nil, zero value otherwise.

### GetAuthenticityGuaranteeOk

`func (o *CoreItem) GetAuthenticityGuaranteeOk() (*AuthenticityGuaranteeProgram, bool)`

GetAuthenticityGuaranteeOk returns a tuple with the AuthenticityGuarantee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticityGuarantee

`func (o *CoreItem) SetAuthenticityGuarantee(v AuthenticityGuaranteeProgram)`

SetAuthenticityGuarantee sets AuthenticityGuarantee field to given value.

### HasAuthenticityGuarantee

`func (o *CoreItem) HasAuthenticityGuarantee() bool`

HasAuthenticityGuarantee returns a boolean if a field has been set.

### GetAuthenticityVerification

`func (o *CoreItem) GetAuthenticityVerification() AuthenticityVerificationProgram`

GetAuthenticityVerification returns the AuthenticityVerification field if non-nil, zero value otherwise.

### GetAuthenticityVerificationOk

`func (o *CoreItem) GetAuthenticityVerificationOk() (*AuthenticityVerificationProgram, bool)`

GetAuthenticityVerificationOk returns a tuple with the AuthenticityVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticityVerification

`func (o *CoreItem) SetAuthenticityVerification(v AuthenticityVerificationProgram)`

SetAuthenticityVerification sets AuthenticityVerification field to given value.

### HasAuthenticityVerification

`func (o *CoreItem) HasAuthenticityVerification() bool`

HasAuthenticityVerification returns a boolean if a field has been set.

### GetAvailableCoupons

`func (o *CoreItem) GetAvailableCoupons() []AvailableCoupon`

GetAvailableCoupons returns the AvailableCoupons field if non-nil, zero value otherwise.

### GetAvailableCouponsOk

`func (o *CoreItem) GetAvailableCouponsOk() (*[]AvailableCoupon, bool)`

GetAvailableCouponsOk returns a tuple with the AvailableCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCoupons

`func (o *CoreItem) SetAvailableCoupons(v []AvailableCoupon)`

SetAvailableCoupons sets AvailableCoupons field to given value.

### HasAvailableCoupons

`func (o *CoreItem) HasAvailableCoupons() bool`

HasAvailableCoupons returns a boolean if a field has been set.

### GetBidCount

`func (o *CoreItem) GetBidCount() int32`

GetBidCount returns the BidCount field if non-nil, zero value otherwise.

### GetBidCountOk

`func (o *CoreItem) GetBidCountOk() (*int32, bool)`

GetBidCountOk returns a tuple with the BidCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidCount

`func (o *CoreItem) SetBidCount(v int32)`

SetBidCount sets BidCount field to given value.

### HasBidCount

`func (o *CoreItem) HasBidCount() bool`

HasBidCount returns a boolean if a field has been set.

### GetBrand

`func (o *CoreItem) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CoreItem) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CoreItem) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *CoreItem) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetBuyingOptions

`func (o *CoreItem) GetBuyingOptions() []string`

GetBuyingOptions returns the BuyingOptions field if non-nil, zero value otherwise.

### GetBuyingOptionsOk

`func (o *CoreItem) GetBuyingOptionsOk() (*[]string, bool)`

GetBuyingOptionsOk returns a tuple with the BuyingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyingOptions

`func (o *CoreItem) SetBuyingOptions(v []string)`

SetBuyingOptions sets BuyingOptions field to given value.

### HasBuyingOptions

`func (o *CoreItem) HasBuyingOptions() bool`

HasBuyingOptions returns a boolean if a field has been set.

### GetCategoryId

`func (o *CoreItem) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CoreItem) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CoreItem) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *CoreItem) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetCategoryPath

`func (o *CoreItem) GetCategoryPath() string`

GetCategoryPath returns the CategoryPath field if non-nil, zero value otherwise.

### GetCategoryPathOk

`func (o *CoreItem) GetCategoryPathOk() (*string, bool)`

GetCategoryPathOk returns a tuple with the CategoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryPath

`func (o *CoreItem) SetCategoryPath(v string)`

SetCategoryPath sets CategoryPath field to given value.

### HasCategoryPath

`func (o *CoreItem) HasCategoryPath() bool`

HasCategoryPath returns a boolean if a field has been set.

### GetColor

`func (o *CoreItem) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CoreItem) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CoreItem) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CoreItem) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetCondition

`func (o *CoreItem) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CoreItem) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CoreItem) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *CoreItem) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetConditionDescription

`func (o *CoreItem) GetConditionDescription() string`

GetConditionDescription returns the ConditionDescription field if non-nil, zero value otherwise.

### GetConditionDescriptionOk

`func (o *CoreItem) GetConditionDescriptionOk() (*string, bool)`

GetConditionDescriptionOk returns a tuple with the ConditionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionDescription

`func (o *CoreItem) SetConditionDescription(v string)`

SetConditionDescription sets ConditionDescription field to given value.

### HasConditionDescription

`func (o *CoreItem) HasConditionDescription() bool`

HasConditionDescription returns a boolean if a field has been set.

### GetConditionId

`func (o *CoreItem) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *CoreItem) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *CoreItem) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *CoreItem) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### GetCurrentBidPrice

`func (o *CoreItem) GetCurrentBidPrice() ConvertedAmount`

GetCurrentBidPrice returns the CurrentBidPrice field if non-nil, zero value otherwise.

### GetCurrentBidPriceOk

`func (o *CoreItem) GetCurrentBidPriceOk() (*ConvertedAmount, bool)`

GetCurrentBidPriceOk returns a tuple with the CurrentBidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBidPrice

`func (o *CoreItem) SetCurrentBidPrice(v ConvertedAmount)`

SetCurrentBidPrice sets CurrentBidPrice field to given value.

### HasCurrentBidPrice

`func (o *CoreItem) HasCurrentBidPrice() bool`

HasCurrentBidPrice returns a boolean if a field has been set.

### GetDescription

`func (o *CoreItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CoreItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CoreItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CoreItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEligibleForInlineCheckout

`func (o *CoreItem) GetEligibleForInlineCheckout() bool`

GetEligibleForInlineCheckout returns the EligibleForInlineCheckout field if non-nil, zero value otherwise.

### GetEligibleForInlineCheckoutOk

`func (o *CoreItem) GetEligibleForInlineCheckoutOk() (*bool, bool)`

GetEligibleForInlineCheckoutOk returns a tuple with the EligibleForInlineCheckout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleForInlineCheckout

`func (o *CoreItem) SetEligibleForInlineCheckout(v bool)`

SetEligibleForInlineCheckout sets EligibleForInlineCheckout field to given value.

### HasEligibleForInlineCheckout

`func (o *CoreItem) HasEligibleForInlineCheckout() bool`

HasEligibleForInlineCheckout returns a boolean if a field has been set.

### GetEnabledForGuestCheckout

`func (o *CoreItem) GetEnabledForGuestCheckout() bool`

GetEnabledForGuestCheckout returns the EnabledForGuestCheckout field if non-nil, zero value otherwise.

### GetEnabledForGuestCheckoutOk

`func (o *CoreItem) GetEnabledForGuestCheckoutOk() (*bool, bool)`

GetEnabledForGuestCheckoutOk returns a tuple with the EnabledForGuestCheckout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledForGuestCheckout

`func (o *CoreItem) SetEnabledForGuestCheckout(v bool)`

SetEnabledForGuestCheckout sets EnabledForGuestCheckout field to given value.

### HasEnabledForGuestCheckout

`func (o *CoreItem) HasEnabledForGuestCheckout() bool`

HasEnabledForGuestCheckout returns a boolean if a field has been set.

### GetEnergyEfficiencyClass

`func (o *CoreItem) GetEnergyEfficiencyClass() string`

GetEnergyEfficiencyClass returns the EnergyEfficiencyClass field if non-nil, zero value otherwise.

### GetEnergyEfficiencyClassOk

`func (o *CoreItem) GetEnergyEfficiencyClassOk() (*string, bool)`

GetEnergyEfficiencyClassOk returns a tuple with the EnergyEfficiencyClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyEfficiencyClass

`func (o *CoreItem) SetEnergyEfficiencyClass(v string)`

SetEnergyEfficiencyClass sets EnergyEfficiencyClass field to given value.

### HasEnergyEfficiencyClass

`func (o *CoreItem) HasEnergyEfficiencyClass() bool`

HasEnergyEfficiencyClass returns a boolean if a field has been set.

### GetEpid

`func (o *CoreItem) GetEpid() string`

GetEpid returns the Epid field if non-nil, zero value otherwise.

### GetEpidOk

`func (o *CoreItem) GetEpidOk() (*string, bool)`

GetEpidOk returns a tuple with the Epid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpid

`func (o *CoreItem) SetEpid(v string)`

SetEpid sets Epid field to given value.

### HasEpid

`func (o *CoreItem) HasEpid() bool`

HasEpid returns a boolean if a field has been set.

### GetEstimatedAvailabilities

`func (o *CoreItem) GetEstimatedAvailabilities() []EstimatedAvailability`

GetEstimatedAvailabilities returns the EstimatedAvailabilities field if non-nil, zero value otherwise.

### GetEstimatedAvailabilitiesOk

`func (o *CoreItem) GetEstimatedAvailabilitiesOk() (*[]EstimatedAvailability, bool)`

GetEstimatedAvailabilitiesOk returns a tuple with the EstimatedAvailabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedAvailabilities

`func (o *CoreItem) SetEstimatedAvailabilities(v []EstimatedAvailability)`

SetEstimatedAvailabilities sets EstimatedAvailabilities field to given value.

### HasEstimatedAvailabilities

`func (o *CoreItem) HasEstimatedAvailabilities() bool`

HasEstimatedAvailabilities returns a boolean if a field has been set.

### GetGender

`func (o *CoreItem) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *CoreItem) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *CoreItem) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *CoreItem) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetGtin

`func (o *CoreItem) GetGtin() string`

GetGtin returns the Gtin field if non-nil, zero value otherwise.

### GetGtinOk

`func (o *CoreItem) GetGtinOk() (*string, bool)`

GetGtinOk returns a tuple with the Gtin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtin

`func (o *CoreItem) SetGtin(v string)`

SetGtin sets Gtin field to given value.

### HasGtin

`func (o *CoreItem) HasGtin() bool`

HasGtin returns a boolean if a field has been set.

### GetImage

`func (o *CoreItem) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CoreItem) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CoreItem) SetImage(v Image)`

SetImage sets Image field to given value.

### HasImage

`func (o *CoreItem) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetInferredEpid

`func (o *CoreItem) GetInferredEpid() string`

GetInferredEpid returns the InferredEpid field if non-nil, zero value otherwise.

### GetInferredEpidOk

`func (o *CoreItem) GetInferredEpidOk() (*string, bool)`

GetInferredEpidOk returns a tuple with the InferredEpid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferredEpid

`func (o *CoreItem) SetInferredEpid(v string)`

SetInferredEpid sets InferredEpid field to given value.

### HasInferredEpid

`func (o *CoreItem) HasInferredEpid() bool`

HasInferredEpid returns a boolean if a field has been set.

### GetItemAffiliateWebUrl

`func (o *CoreItem) GetItemAffiliateWebUrl() string`

GetItemAffiliateWebUrl returns the ItemAffiliateWebUrl field if non-nil, zero value otherwise.

### GetItemAffiliateWebUrlOk

`func (o *CoreItem) GetItemAffiliateWebUrlOk() (*string, bool)`

GetItemAffiliateWebUrlOk returns a tuple with the ItemAffiliateWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemAffiliateWebUrl

`func (o *CoreItem) SetItemAffiliateWebUrl(v string)`

SetItemAffiliateWebUrl sets ItemAffiliateWebUrl field to given value.

### HasItemAffiliateWebUrl

`func (o *CoreItem) HasItemAffiliateWebUrl() bool`

HasItemAffiliateWebUrl returns a boolean if a field has been set.

### GetItemEndDate

`func (o *CoreItem) GetItemEndDate() string`

GetItemEndDate returns the ItemEndDate field if non-nil, zero value otherwise.

### GetItemEndDateOk

`func (o *CoreItem) GetItemEndDateOk() (*string, bool)`

GetItemEndDateOk returns a tuple with the ItemEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemEndDate

`func (o *CoreItem) SetItemEndDate(v string)`

SetItemEndDate sets ItemEndDate field to given value.

### HasItemEndDate

`func (o *CoreItem) HasItemEndDate() bool`

HasItemEndDate returns a boolean if a field has been set.

### GetItemId

`func (o *CoreItem) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *CoreItem) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *CoreItem) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *CoreItem) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetItemLocation

`func (o *CoreItem) GetItemLocation() Address`

GetItemLocation returns the ItemLocation field if non-nil, zero value otherwise.

### GetItemLocationOk

`func (o *CoreItem) GetItemLocationOk() (*Address, bool)`

GetItemLocationOk returns a tuple with the ItemLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLocation

`func (o *CoreItem) SetItemLocation(v Address)`

SetItemLocation sets ItemLocation field to given value.

### HasItemLocation

`func (o *CoreItem) HasItemLocation() bool`

HasItemLocation returns a boolean if a field has been set.

### GetItemWebUrl

`func (o *CoreItem) GetItemWebUrl() string`

GetItemWebUrl returns the ItemWebUrl field if non-nil, zero value otherwise.

### GetItemWebUrlOk

`func (o *CoreItem) GetItemWebUrlOk() (*string, bool)`

GetItemWebUrlOk returns a tuple with the ItemWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemWebUrl

`func (o *CoreItem) SetItemWebUrl(v string)`

SetItemWebUrl sets ItemWebUrl field to given value.

### HasItemWebUrl

`func (o *CoreItem) HasItemWebUrl() bool`

HasItemWebUrl returns a boolean if a field has been set.

### GetLegacyItemId

`func (o *CoreItem) GetLegacyItemId() string`

GetLegacyItemId returns the LegacyItemId field if non-nil, zero value otherwise.

### GetLegacyItemIdOk

`func (o *CoreItem) GetLegacyItemIdOk() (*string, bool)`

GetLegacyItemIdOk returns a tuple with the LegacyItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyItemId

`func (o *CoreItem) SetLegacyItemId(v string)`

SetLegacyItemId sets LegacyItemId field to given value.

### HasLegacyItemId

`func (o *CoreItem) HasLegacyItemId() bool`

HasLegacyItemId returns a boolean if a field has been set.

### GetLocalizedAspects

`func (o *CoreItem) GetLocalizedAspects() []TypedNameValue`

GetLocalizedAspects returns the LocalizedAspects field if non-nil, zero value otherwise.

### GetLocalizedAspectsOk

`func (o *CoreItem) GetLocalizedAspectsOk() (*[]TypedNameValue, bool)`

GetLocalizedAspectsOk returns a tuple with the LocalizedAspects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedAspects

`func (o *CoreItem) SetLocalizedAspects(v []TypedNameValue)`

SetLocalizedAspects sets LocalizedAspects field to given value.

### HasLocalizedAspects

`func (o *CoreItem) HasLocalizedAspects() bool`

HasLocalizedAspects returns a boolean if a field has been set.

### GetLotSize

`func (o *CoreItem) GetLotSize() int32`

GetLotSize returns the LotSize field if non-nil, zero value otherwise.

### GetLotSizeOk

`func (o *CoreItem) GetLotSizeOk() (*int32, bool)`

GetLotSizeOk returns a tuple with the LotSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotSize

`func (o *CoreItem) SetLotSize(v int32)`

SetLotSize sets LotSize field to given value.

### HasLotSize

`func (o *CoreItem) HasLotSize() bool`

HasLotSize returns a boolean if a field has been set.

### GetMarketingPrice

`func (o *CoreItem) GetMarketingPrice() MarketingPrice`

GetMarketingPrice returns the MarketingPrice field if non-nil, zero value otherwise.

### GetMarketingPriceOk

`func (o *CoreItem) GetMarketingPriceOk() (*MarketingPrice, bool)`

GetMarketingPriceOk returns a tuple with the MarketingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingPrice

`func (o *CoreItem) SetMarketingPrice(v MarketingPrice)`

SetMarketingPrice sets MarketingPrice field to given value.

### HasMarketingPrice

`func (o *CoreItem) HasMarketingPrice() bool`

HasMarketingPrice returns a boolean if a field has been set.

### GetMaterial

`func (o *CoreItem) GetMaterial() string`

GetMaterial returns the Material field if non-nil, zero value otherwise.

### GetMaterialOk

`func (o *CoreItem) GetMaterialOk() (*string, bool)`

GetMaterialOk returns a tuple with the Material field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterial

`func (o *CoreItem) SetMaterial(v string)`

SetMaterial sets Material field to given value.

### HasMaterial

`func (o *CoreItem) HasMaterial() bool`

HasMaterial returns a boolean if a field has been set.

### GetMinimumPriceToBid

`func (o *CoreItem) GetMinimumPriceToBid() ConvertedAmount`

GetMinimumPriceToBid returns the MinimumPriceToBid field if non-nil, zero value otherwise.

### GetMinimumPriceToBidOk

`func (o *CoreItem) GetMinimumPriceToBidOk() (*ConvertedAmount, bool)`

GetMinimumPriceToBidOk returns a tuple with the MinimumPriceToBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPriceToBid

`func (o *CoreItem) SetMinimumPriceToBid(v ConvertedAmount)`

SetMinimumPriceToBid sets MinimumPriceToBid field to given value.

### HasMinimumPriceToBid

`func (o *CoreItem) HasMinimumPriceToBid() bool`

HasMinimumPriceToBid returns a boolean if a field has been set.

### GetMpn

`func (o *CoreItem) GetMpn() string`

GetMpn returns the Mpn field if non-nil, zero value otherwise.

### GetMpnOk

`func (o *CoreItem) GetMpnOk() (*string, bool)`

GetMpnOk returns a tuple with the Mpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpn

`func (o *CoreItem) SetMpn(v string)`

SetMpn sets Mpn field to given value.

### HasMpn

`func (o *CoreItem) HasMpn() bool`

HasMpn returns a boolean if a field has been set.

### GetPattern

`func (o *CoreItem) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *CoreItem) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *CoreItem) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *CoreItem) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *CoreItem) GetPaymentMethods() []PaymentMethod`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *CoreItem) GetPaymentMethodsOk() (*[]PaymentMethod, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *CoreItem) SetPaymentMethods(v []PaymentMethod)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *CoreItem) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetPrice

`func (o *CoreItem) GetPrice() ConvertedAmount`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CoreItem) GetPriceOk() (*ConvertedAmount, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CoreItem) SetPrice(v ConvertedAmount)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CoreItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceDisplayCondition

`func (o *CoreItem) GetPriceDisplayCondition() string`

GetPriceDisplayCondition returns the PriceDisplayCondition field if non-nil, zero value otherwise.

### GetPriceDisplayConditionOk

`func (o *CoreItem) GetPriceDisplayConditionOk() (*string, bool)`

GetPriceDisplayConditionOk returns a tuple with the PriceDisplayCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDisplayCondition

`func (o *CoreItem) SetPriceDisplayCondition(v string)`

SetPriceDisplayCondition sets PriceDisplayCondition field to given value.

### HasPriceDisplayCondition

`func (o *CoreItem) HasPriceDisplayCondition() bool`

HasPriceDisplayCondition returns a boolean if a field has been set.

### GetPrimaryItemGroup

`func (o *CoreItem) GetPrimaryItemGroup() ItemGroupSummary`

GetPrimaryItemGroup returns the PrimaryItemGroup field if non-nil, zero value otherwise.

### GetPrimaryItemGroupOk

`func (o *CoreItem) GetPrimaryItemGroupOk() (*ItemGroupSummary, bool)`

GetPrimaryItemGroupOk returns a tuple with the PrimaryItemGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryItemGroup

`func (o *CoreItem) SetPrimaryItemGroup(v ItemGroupSummary)`

SetPrimaryItemGroup sets PrimaryItemGroup field to given value.

### HasPrimaryItemGroup

`func (o *CoreItem) HasPrimaryItemGroup() bool`

HasPrimaryItemGroup returns a boolean if a field has been set.

### GetPrimaryProductReviewRating

`func (o *CoreItem) GetPrimaryProductReviewRating() ReviewRating`

GetPrimaryProductReviewRating returns the PrimaryProductReviewRating field if non-nil, zero value otherwise.

### GetPrimaryProductReviewRatingOk

`func (o *CoreItem) GetPrimaryProductReviewRatingOk() (*ReviewRating, bool)`

GetPrimaryProductReviewRatingOk returns a tuple with the PrimaryProductReviewRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryProductReviewRating

`func (o *CoreItem) SetPrimaryProductReviewRating(v ReviewRating)`

SetPrimaryProductReviewRating sets PrimaryProductReviewRating field to given value.

### HasPrimaryProductReviewRating

`func (o *CoreItem) HasPrimaryProductReviewRating() bool`

HasPrimaryProductReviewRating returns a boolean if a field has been set.

### GetPriorityListing

`func (o *CoreItem) GetPriorityListing() bool`

GetPriorityListing returns the PriorityListing field if non-nil, zero value otherwise.

### GetPriorityListingOk

`func (o *CoreItem) GetPriorityListingOk() (*bool, bool)`

GetPriorityListingOk returns a tuple with the PriorityListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityListing

`func (o *CoreItem) SetPriorityListing(v bool)`

SetPriorityListing sets PriorityListing field to given value.

### HasPriorityListing

`func (o *CoreItem) HasPriorityListing() bool`

HasPriorityListing returns a boolean if a field has been set.

### GetProduct

`func (o *CoreItem) GetProduct() Product`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *CoreItem) GetProductOk() (*Product, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *CoreItem) SetProduct(v Product)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *CoreItem) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProductFicheWebUrl

`func (o *CoreItem) GetProductFicheWebUrl() string`

GetProductFicheWebUrl returns the ProductFicheWebUrl field if non-nil, zero value otherwise.

### GetProductFicheWebUrlOk

`func (o *CoreItem) GetProductFicheWebUrlOk() (*string, bool)`

GetProductFicheWebUrlOk returns a tuple with the ProductFicheWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductFicheWebUrl

`func (o *CoreItem) SetProductFicheWebUrl(v string)`

SetProductFicheWebUrl sets ProductFicheWebUrl field to given value.

### HasProductFicheWebUrl

`func (o *CoreItem) HasProductFicheWebUrl() bool`

HasProductFicheWebUrl returns a boolean if a field has been set.

### GetQualifiedPrograms

`func (o *CoreItem) GetQualifiedPrograms() []string`

GetQualifiedPrograms returns the QualifiedPrograms field if non-nil, zero value otherwise.

### GetQualifiedProgramsOk

`func (o *CoreItem) GetQualifiedProgramsOk() (*[]string, bool)`

GetQualifiedProgramsOk returns a tuple with the QualifiedPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedPrograms

`func (o *CoreItem) SetQualifiedPrograms(v []string)`

SetQualifiedPrograms sets QualifiedPrograms field to given value.

### HasQualifiedPrograms

`func (o *CoreItem) HasQualifiedPrograms() bool`

HasQualifiedPrograms returns a boolean if a field has been set.

### GetQuantityLimitPerBuyer

`func (o *CoreItem) GetQuantityLimitPerBuyer() int32`

GetQuantityLimitPerBuyer returns the QuantityLimitPerBuyer field if non-nil, zero value otherwise.

### GetQuantityLimitPerBuyerOk

`func (o *CoreItem) GetQuantityLimitPerBuyerOk() (*int32, bool)`

GetQuantityLimitPerBuyerOk returns a tuple with the QuantityLimitPerBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityLimitPerBuyer

`func (o *CoreItem) SetQuantityLimitPerBuyer(v int32)`

SetQuantityLimitPerBuyer sets QuantityLimitPerBuyer field to given value.

### HasQuantityLimitPerBuyer

`func (o *CoreItem) HasQuantityLimitPerBuyer() bool`

HasQuantityLimitPerBuyer returns a boolean if a field has been set.

### GetReservePriceMet

`func (o *CoreItem) GetReservePriceMet() bool`

GetReservePriceMet returns the ReservePriceMet field if non-nil, zero value otherwise.

### GetReservePriceMetOk

`func (o *CoreItem) GetReservePriceMetOk() (*bool, bool)`

GetReservePriceMetOk returns a tuple with the ReservePriceMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservePriceMet

`func (o *CoreItem) SetReservePriceMet(v bool)`

SetReservePriceMet sets ReservePriceMet field to given value.

### HasReservePriceMet

`func (o *CoreItem) HasReservePriceMet() bool`

HasReservePriceMet returns a boolean if a field has been set.

### GetReturnTerms

`func (o *CoreItem) GetReturnTerms() ItemReturnTerms`

GetReturnTerms returns the ReturnTerms field if non-nil, zero value otherwise.

### GetReturnTermsOk

`func (o *CoreItem) GetReturnTermsOk() (*ItemReturnTerms, bool)`

GetReturnTermsOk returns a tuple with the ReturnTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnTerms

`func (o *CoreItem) SetReturnTerms(v ItemReturnTerms)`

SetReturnTerms sets ReturnTerms field to given value.

### HasReturnTerms

`func (o *CoreItem) HasReturnTerms() bool`

HasReturnTerms returns a boolean if a field has been set.

### GetSeller

`func (o *CoreItem) GetSeller() SellerDetail`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *CoreItem) GetSellerOk() (*SellerDetail, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *CoreItem) SetSeller(v SellerDetail)`

SetSeller sets Seller field to given value.

### HasSeller

`func (o *CoreItem) HasSeller() bool`

HasSeller returns a boolean if a field has been set.

### GetSellerItemRevision

`func (o *CoreItem) GetSellerItemRevision() string`

GetSellerItemRevision returns the SellerItemRevision field if non-nil, zero value otherwise.

### GetSellerItemRevisionOk

`func (o *CoreItem) GetSellerItemRevisionOk() (*string, bool)`

GetSellerItemRevisionOk returns a tuple with the SellerItemRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerItemRevision

`func (o *CoreItem) SetSellerItemRevision(v string)`

SetSellerItemRevision sets SellerItemRevision field to given value.

### HasSellerItemRevision

`func (o *CoreItem) HasSellerItemRevision() bool`

HasSellerItemRevision returns a boolean if a field has been set.

### GetShippingOptions

`func (o *CoreItem) GetShippingOptions() []ShippingOption`

GetShippingOptions returns the ShippingOptions field if non-nil, zero value otherwise.

### GetShippingOptionsOk

`func (o *CoreItem) GetShippingOptionsOk() (*[]ShippingOption, bool)`

GetShippingOptionsOk returns a tuple with the ShippingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingOptions

`func (o *CoreItem) SetShippingOptions(v []ShippingOption)`

SetShippingOptions sets ShippingOptions field to given value.

### HasShippingOptions

`func (o *CoreItem) HasShippingOptions() bool`

HasShippingOptions returns a boolean if a field has been set.

### GetShipToLocations

`func (o *CoreItem) GetShipToLocations() ShipToLocations`

GetShipToLocations returns the ShipToLocations field if non-nil, zero value otherwise.

### GetShipToLocationsOk

`func (o *CoreItem) GetShipToLocationsOk() (*ShipToLocations, bool)`

GetShipToLocationsOk returns a tuple with the ShipToLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToLocations

`func (o *CoreItem) SetShipToLocations(v ShipToLocations)`

SetShipToLocations sets ShipToLocations field to given value.

### HasShipToLocations

`func (o *CoreItem) HasShipToLocations() bool`

HasShipToLocations returns a boolean if a field has been set.

### GetShortDescription

`func (o *CoreItem) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *CoreItem) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *CoreItem) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *CoreItem) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetSize

`func (o *CoreItem) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CoreItem) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CoreItem) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CoreItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizeSystem

`func (o *CoreItem) GetSizeSystem() string`

GetSizeSystem returns the SizeSystem field if non-nil, zero value otherwise.

### GetSizeSystemOk

`func (o *CoreItem) GetSizeSystemOk() (*string, bool)`

GetSizeSystemOk returns a tuple with the SizeSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSystem

`func (o *CoreItem) SetSizeSystem(v string)`

SetSizeSystem sets SizeSystem field to given value.

### HasSizeSystem

`func (o *CoreItem) HasSizeSystem() bool`

HasSizeSystem returns a boolean if a field has been set.

### GetSizeType

`func (o *CoreItem) GetSizeType() string`

GetSizeType returns the SizeType field if non-nil, zero value otherwise.

### GetSizeTypeOk

`func (o *CoreItem) GetSizeTypeOk() (*string, bool)`

GetSizeTypeOk returns a tuple with the SizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeType

`func (o *CoreItem) SetSizeType(v string)`

SetSizeType sets SizeType field to given value.

### HasSizeType

`func (o *CoreItem) HasSizeType() bool`

HasSizeType returns a boolean if a field has been set.

### GetSubtitle

`func (o *CoreItem) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *CoreItem) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *CoreItem) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *CoreItem) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetTaxes

`func (o *CoreItem) GetTaxes() []Taxes`

GetTaxes returns the Taxes field if non-nil, zero value otherwise.

### GetTaxesOk

`func (o *CoreItem) GetTaxesOk() (*[]Taxes, bool)`

GetTaxesOk returns a tuple with the Taxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxes

`func (o *CoreItem) SetTaxes(v []Taxes)`

SetTaxes sets Taxes field to given value.

### HasTaxes

`func (o *CoreItem) HasTaxes() bool`

HasTaxes returns a boolean if a field has been set.

### GetTitle

`func (o *CoreItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CoreItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CoreItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CoreItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTopRatedBuyingExperience

`func (o *CoreItem) GetTopRatedBuyingExperience() bool`

GetTopRatedBuyingExperience returns the TopRatedBuyingExperience field if non-nil, zero value otherwise.

### GetTopRatedBuyingExperienceOk

`func (o *CoreItem) GetTopRatedBuyingExperienceOk() (*bool, bool)`

GetTopRatedBuyingExperienceOk returns a tuple with the TopRatedBuyingExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopRatedBuyingExperience

`func (o *CoreItem) SetTopRatedBuyingExperience(v bool)`

SetTopRatedBuyingExperience sets TopRatedBuyingExperience field to given value.

### HasTopRatedBuyingExperience

`func (o *CoreItem) HasTopRatedBuyingExperience() bool`

HasTopRatedBuyingExperience returns a boolean if a field has been set.

### GetTyreLabelImageUrl

`func (o *CoreItem) GetTyreLabelImageUrl() string`

GetTyreLabelImageUrl returns the TyreLabelImageUrl field if non-nil, zero value otherwise.

### GetTyreLabelImageUrlOk

`func (o *CoreItem) GetTyreLabelImageUrlOk() (*string, bool)`

GetTyreLabelImageUrlOk returns a tuple with the TyreLabelImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTyreLabelImageUrl

`func (o *CoreItem) SetTyreLabelImageUrl(v string)`

SetTyreLabelImageUrl sets TyreLabelImageUrl field to given value.

### HasTyreLabelImageUrl

`func (o *CoreItem) HasTyreLabelImageUrl() bool`

HasTyreLabelImageUrl returns a boolean if a field has been set.

### GetUniqueBidderCount

`func (o *CoreItem) GetUniqueBidderCount() int32`

GetUniqueBidderCount returns the UniqueBidderCount field if non-nil, zero value otherwise.

### GetUniqueBidderCountOk

`func (o *CoreItem) GetUniqueBidderCountOk() (*int32, bool)`

GetUniqueBidderCountOk returns a tuple with the UniqueBidderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueBidderCount

`func (o *CoreItem) SetUniqueBidderCount(v int32)`

SetUniqueBidderCount sets UniqueBidderCount field to given value.

### HasUniqueBidderCount

`func (o *CoreItem) HasUniqueBidderCount() bool`

HasUniqueBidderCount returns a boolean if a field has been set.

### GetUnitPrice

`func (o *CoreItem) GetUnitPrice() ConvertedAmount`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *CoreItem) GetUnitPriceOk() (*ConvertedAmount, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *CoreItem) SetUnitPrice(v ConvertedAmount)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *CoreItem) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetUnitPricingMeasure

`func (o *CoreItem) GetUnitPricingMeasure() string`

GetUnitPricingMeasure returns the UnitPricingMeasure field if non-nil, zero value otherwise.

### GetUnitPricingMeasureOk

`func (o *CoreItem) GetUnitPricingMeasureOk() (*string, bool)`

GetUnitPricingMeasureOk returns a tuple with the UnitPricingMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPricingMeasure

`func (o *CoreItem) SetUnitPricingMeasure(v string)`

SetUnitPricingMeasure sets UnitPricingMeasure field to given value.

### HasUnitPricingMeasure

`func (o *CoreItem) HasUnitPricingMeasure() bool`

HasUnitPricingMeasure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


