# \ShoppingCartApi

All URIs are relative to *https://api.ebay.com/buy/browse/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddItem**](ShoppingCartApi.md#AddItem) | **Post** /shopping_cart/add_item | 
[**GetShoppingCart**](ShoppingCartApi.md#GetShoppingCart) | **Get** /shopping_cart/ | 
[**RemoveItem**](ShoppingCartApi.md#RemoveItem) | **Post** /shopping_cart/remove_item | 
[**UpdateQuantity**](ShoppingCartApi.md#UpdateQuantity) | **Post** /shopping_cart/update_quantity | 



## AddItem

> RemoteShopcartResponse AddItem(ctx).AddCartItemInput(addCartItemInput).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    addCartItemInput := *openapiclient.NewAddCartItemInput() // AddCartItemInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ShoppingCartApi.AddItem(context.Background()).AddCartItemInput(addCartItemInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShoppingCartApi.AddItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddItem`: RemoteShopcartResponse
    fmt.Fprintf(os.Stdout, "Response from `ShoppingCartApi.AddItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addCartItemInput** | [**AddCartItemInput**](AddCartItemInput.md) |  | 

### Return type

[**RemoteShopcartResponse**](RemoteShopcartResponse.md)

### Authorization

[api_auth](../README.md#api_auth), [api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShoppingCart

> RemoteShopcartResponse GetShoppingCart(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ShoppingCartApi.GetShoppingCart(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShoppingCartApi.GetShoppingCart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShoppingCart`: RemoteShopcartResponse
    fmt.Fprintf(os.Stdout, "Response from `ShoppingCartApi.GetShoppingCart`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetShoppingCartRequest struct via the builder pattern


### Return type

[**RemoteShopcartResponse**](RemoteShopcartResponse.md)

### Authorization

[api_auth](../README.md#api_auth), [api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveItem

> RemoteShopcartResponse RemoveItem(ctx).RemoveCartItemInput(removeCartItemInput).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    removeCartItemInput := *openapiclient.NewRemoveCartItemInput() // RemoveCartItemInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ShoppingCartApi.RemoveItem(context.Background()).RemoveCartItemInput(removeCartItemInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShoppingCartApi.RemoveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveItem`: RemoteShopcartResponse
    fmt.Fprintf(os.Stdout, "Response from `ShoppingCartApi.RemoveItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeCartItemInput** | [**RemoveCartItemInput**](RemoveCartItemInput.md) |  | 

### Return type

[**RemoteShopcartResponse**](RemoteShopcartResponse.md)

### Authorization

[api_auth](../README.md#api_auth), [api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateQuantity

> RemoteShopcartResponse UpdateQuantity(ctx).UpdateCartItemInput(updateCartItemInput).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateCartItemInput := *openapiclient.NewUpdateCartItemInput() // UpdateCartItemInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ShoppingCartApi.UpdateQuantity(context.Background()).UpdateCartItemInput(updateCartItemInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShoppingCartApi.UpdateQuantity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateQuantity`: RemoteShopcartResponse
    fmt.Fprintf(os.Stdout, "Response from `ShoppingCartApi.UpdateQuantity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateQuantityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateCartItemInput** | [**UpdateCartItemInput**](UpdateCartItemInput.md) |  | 

### Return type

[**RemoteShopcartResponse**](RemoteShopcartResponse.md)

### Authorization

[api_auth](../README.md#api_auth), [api_auth](../README.md#api_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

