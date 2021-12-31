# Go API client for buybrowse

<p>The Browse API has the following resources:</p>   <ul> <li><b> item_summary: </b> Lets shoppers search for specific items by keyword, GTIN, category, charity, product, or item aspects and refine the results by using filters, such as aspects, compatibility, and fields values.</li>  <li><b> search_by_image: </b><a href=\"https://developer.ebay.com/api-docs/static/versioning.html#experimental\" target=\"_blank\"><img src=\"/cms/img/docs/experimental-icon.svg\" class=\"legend-icon experimental-icon\" alt=\"Experimental Release\" title=\"Experimental Release\" />&nbsp;(Experimental)</a> Lets shoppers search for specific items by image. You can refine the results by using URI parameters and filters.</li>   <li><b> item: </b> <ul><li>Lets you retrieve the details of a specific item or all the items in an item group, which is an item with variations such as color and size and check if a product is compatible with the specified item, such as if a specific car is compatible with a specific part.</li> <li>Provides a bridge between the eBay legacy APIs, such as <b> Finding</b>, and the RESTful APIs, which use different formats for the item IDs.</li>  </ul> </li>  <li> <b> shopping_cart: </b> <a href=\"https://developer.ebay.com/api-docs/static/versioning.html#experimental\" target=\"_blank\"><img src=\"/cms/img/docs/experimental-icon.svg\" class=\"legend-icon experimental-icon\" alt=\"Experimental Release\" title=\"Experimental Release\" />&nbsp;(Experimental)</a> <a href=\"https://developer.ebay.com/api-docs/static/versioning.html#limited\" target=\"_blank\"> <img src=\"/cms/img/docs/partners-api.svg\" class=\"legend-icon partners-icon\" title=\"Limited Release\"  alt=\"Limited Release\" />(Limited Release)</a> Provides the ability for eBay members to see the contents of their eBay cart, and add, remove, and change the quantity of items in their eBay cart.&nbsp;&nbsp;<b> Note: </b> This resource is not available in the eBay API Explorer.</li></ul>       <p>The <b> item_summary</b>, <b> search_by_image</b>, and <b> item</b> resource calls require an <a href=\"/api-docs/static/oauth-client-credentials-grant.html\">Application access token</a>. The <b> shopping_cart</b> resource calls require a <a href=\"/api-docs/static/oauth-authorization-code-grant.html\">User access token</a>.</p>

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1.11.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./buybrowse"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.ebay.com/buy/browse/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ItemApi* | [**CheckCompatibility**](docs/ItemApi.md#checkcompatibility) | **Post** /item/{item_id}/check_compatibility | 
*ItemApi* | [**GetItem**](docs/ItemApi.md#getitem) | **Get** /item/{item_id} | 
*ItemApi* | [**GetItemByLegacyId**](docs/ItemApi.md#getitembylegacyid) | **Get** /item/get_item_by_legacy_id | 
*ItemApi* | [**GetItems**](docs/ItemApi.md#getitems) | **Get** /item/ | 
*ItemApi* | [**GetItemsByItemGroup**](docs/ItemApi.md#getitemsbyitemgroup) | **Get** /item/get_items_by_item_group | 
*ItemSummaryApi* | [**Search**](docs/ItemSummaryApi.md#search) | **Get** /item_summary/search | 
*SearchByImageApi* | [**SearchByImage**](docs/SearchByImageApi.md#searchbyimage) | **Post** /item_summary/search_by_image | 
*ShoppingCartApi* | [**AddItem**](docs/ShoppingCartApi.md#additem) | **Post** /shopping_cart/add_item | 
*ShoppingCartApi* | [**GetShoppingCart**](docs/ShoppingCartApi.md#getshoppingcart) | **Get** /shopping_cart/ | 
*ShoppingCartApi* | [**RemoveItem**](docs/ShoppingCartApi.md#removeitem) | **Post** /shopping_cart/remove_item | 
*ShoppingCartApi* | [**UpdateQuantity**](docs/ShoppingCartApi.md#updatequantity) | **Post** /shopping_cart/update_quantity | 


## Documentation For Models

 - [AddCartItemInput](docs/AddCartItemInput.md)
 - [AdditionalProductIdentity](docs/AdditionalProductIdentity.md)
 - [Address](docs/Address.md)
 - [Amount](docs/Amount.md)
 - [Aspect](docs/Aspect.md)
 - [AspectDistribution](docs/AspectDistribution.md)
 - [AspectGroup](docs/AspectGroup.md)
 - [AspectValueDistribution](docs/AspectValueDistribution.md)
 - [AttributeNameValue](docs/AttributeNameValue.md)
 - [AuthenticityGuaranteeProgram](docs/AuthenticityGuaranteeProgram.md)
 - [AuthenticityVerificationProgram](docs/AuthenticityVerificationProgram.md)
 - [AutoCorrections](docs/AutoCorrections.md)
 - [AvailableCoupon](docs/AvailableCoupon.md)
 - [BuyingOptionDistribution](docs/BuyingOptionDistribution.md)
 - [CartItem](docs/CartItem.md)
 - [Category](docs/Category.md)
 - [CategoryDistribution](docs/CategoryDistribution.md)
 - [CommonDescriptions](docs/CommonDescriptions.md)
 - [CompatibilityPayload](docs/CompatibilityPayload.md)
 - [CompatibilityProperty](docs/CompatibilityProperty.md)
 - [CompatibilityResponse](docs/CompatibilityResponse.md)
 - [ConditionDistribution](docs/ConditionDistribution.md)
 - [ConvertedAmount](docs/ConvertedAmount.md)
 - [CoreItem](docs/CoreItem.md)
 - [CouponConstraint](docs/CouponConstraint.md)
 - [Error](docs/Error.md)
 - [ErrorParameter](docs/ErrorParameter.md)
 - [EstimatedAvailability](docs/EstimatedAvailability.md)
 - [Image](docs/Image.md)
 - [Item](docs/Item.md)
 - [ItemGroup](docs/ItemGroup.md)
 - [ItemGroupSummary](docs/ItemGroupSummary.md)
 - [ItemLocationImpl](docs/ItemLocationImpl.md)
 - [ItemReturnTerms](docs/ItemReturnTerms.md)
 - [ItemSummary](docs/ItemSummary.md)
 - [Items](docs/Items.md)
 - [LegalAddress](docs/LegalAddress.md)
 - [MarketingPrice](docs/MarketingPrice.md)
 - [PaymentMethod](docs/PaymentMethod.md)
 - [PaymentMethodBrand](docs/PaymentMethodBrand.md)
 - [PickupOptionSummary](docs/PickupOptionSummary.md)
 - [Price](docs/Price.md)
 - [Product](docs/Product.md)
 - [ProductIdentity](docs/ProductIdentity.md)
 - [RatingHistogram](docs/RatingHistogram.md)
 - [Refinement](docs/Refinement.md)
 - [Region](docs/Region.md)
 - [RemoteShopcartResponse](docs/RemoteShopcartResponse.md)
 - [RemoveCartItemInput](docs/RemoveCartItemInput.md)
 - [ReviewRating](docs/ReviewRating.md)
 - [SearchByImageRequest](docs/SearchByImageRequest.md)
 - [SearchPagedCollection](docs/SearchPagedCollection.md)
 - [Seller](docs/Seller.md)
 - [SellerCustomPolicy](docs/SellerCustomPolicy.md)
 - [SellerDetail](docs/SellerDetail.md)
 - [SellerLegalInfo](docs/SellerLegalInfo.md)
 - [ShipToLocation](docs/ShipToLocation.md)
 - [ShipToLocations](docs/ShipToLocations.md)
 - [ShipToRegion](docs/ShipToRegion.md)
 - [ShippingOption](docs/ShippingOption.md)
 - [ShippingOptionSummary](docs/ShippingOptionSummary.md)
 - [TargetLocation](docs/TargetLocation.md)
 - [TaxJurisdiction](docs/TaxJurisdiction.md)
 - [Taxes](docs/Taxes.md)
 - [TimeDuration](docs/TimeDuration.md)
 - [TypedNameValue](docs/TypedNameValue.md)
 - [UpdateCartItemInput](docs/UpdateCartItemInput.md)
 - [VatDetail](docs/VatDetail.md)


## Documentation For Authorization



### api_auth


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **https://api.ebay.com/oauth/api_scope/buy.item.bulk**: Retrieve eBay items in bulk.
 - **https://api.ebay.com/oauth/api_scope**: View public data from eBay

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


### api_auth


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://auth.ebay.com/oauth2/authorize
- **Scopes**: 
 - **https://api.ebay.com/oauth/api_scope/buy.shopping.cart**:  This scope would allow signed in user to access shopping carts

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


