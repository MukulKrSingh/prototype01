# GraphQL Operation Naming Conventions

This document outlines the standard naming conventions for GraphQL queries and mutations in our e-commerce application.

## General Principles

1. **Consistency**: Use consistent naming patterns across all operations
2. **Descriptiveness**: Names should clearly communicate what the operation does
3. **Specificity**: Include resource names in operation names
4. **Camel Case**: Use camelCase for all operation names (first letter lowercase)

## Query Naming Patterns

| Pattern | Usage | Examples |
|---------|-------|----------|
| `get<Resource>` | Retrieve a single resource by ID | `getProduct`, `getOrder`, `getCategory` |
| `list<Resources>` | Retrieve multiple resources | `listProducts`, `listCategories`, `listOrders` |
| `search<Resources>` | Search for resources with filters | `searchProducts`, `searchOrders` |
| `count<Resources>` | Count the number of resources | `countProducts`, `countActiveUsers` |
| `get<Resource><Aspect>` | Get a specific aspect of a resource | `getUserOrders`, `getProductReviews` |

## Mutation Naming Patterns

| Pattern | Usage | Examples |
|---------|-------|----------|
| `create<Resource>` | Create a new resource | `createProduct`, `createOrder`, `createCategory` |
| `update<Resource>` | Update an existing resource | `updateProduct`, `updateOrder`, `updateUserProfile` |
| `delete<Resource>` | Delete a resource | `deleteProduct`, `deleteCategory`, `deleteReview` |
| `add<Item>To<Collection>` | Add an item to a collection | `addProductToCart`, `addProductToWishlist` |
| `remove<Item>From<Collection>` | Remove an item from a collection | `removeProductFromCart`, `removeItemFromWishlist` |
| `<verb><Resource>` | Other specific actions | `cancelOrder`, `processPayment`, `validateCoupon` |
| `<verb><Resource><Aspect>` | Act on a specific aspect of a resource | `updateProductInventory`, `updateOrderStatus` |

## Authorization Operations

| Pattern | Usage | Examples |
|---------|-------|----------|
| `login<UserType>` | Log in a user | `loginUser`, `loginAdmin` |
| `register<UserType>` | Register a new user | `registerUser`, `registerSeller` |
| `changePassword` | Change user password | `changePassword` |
| `resetPassword` | Reset user password | `requestPasswordReset`, `confirmPasswordReset` |
| `verify<Aspect>` | Verify a user aspect | `verifyEmail`, `verifyPhoneNumber` |

## Collection Management

| Pattern | Usage | Examples |
|---------|-------|----------|
| `<verb>All<Resources>` | Operation affecting all resources | `deleteAllCartItems`, `markAllNotificationsAsRead` |
| `<verb>Selected<Resources>` | Operation affecting selected resources | `deleteSelectedProducts`, `updateSelectedOrders` |
| `bulk<Verb><Resources>` | Bulk operations | `bulkCreateProducts`, `bulkUpdatePrices` |

## Examples of Well-Named Operations

```graphql
query GetProduct($id: ID!) { ... }
query ListProductsByCategory($categoryId: ID!) { ... }
query SearchProductsByKeyword($keyword: String!) { ... }

mutation CreateProduct($input: CreateProductInput!) { ... }
mutation UpdateProductPrice($id: ID!, $price: Float!) { ... }
mutation DeleteProduct($id: ID!) { ... }
mutation AddProductToWishlist($productId: ID!) { ... }
mutation UpdateOrderStatus($orderId: ID!, $status: OrderStatus!) { ... }
```

## Avoid These Common Mistakes

❌ **Too Generic**: `get`, `list`, `create`, `update`  
❌ **Inconsistent Verbs**: Mixing `fetch`/`get`/`retrieve` for similar operations  
❌ **Inconsistent Pluralization**: Using both `listOrder` and `listProducts`  
❌ **Non-descriptive**: `doThing`, `process`, `handle`  
❌ **Using HTTP Methods**: `getProduct`, not `fetchProduct` or `retrieveProduct`  

## Advanced Patterns for Complex Operations

For complex operations that don't fit the standard patterns, use these naming strategies:

1. **Process-oriented**: `initiateCheckout`, `completePayment`, `processRefund`
2. **State transitions**: `cancelOrder`, `approveReview`, `publishProduct`
3. **Grouped operations**: `updateUserProfileAndPreferences`
