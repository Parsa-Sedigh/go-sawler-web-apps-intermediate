# Section 09 Building Admin pages to manage purchases

## 111-001 Improving our front end and setting up an Admin menu
Now instead of determining if the user is logged in, using client side code, we set isAuthenticated on backend and conditionally render things
on page.

## 112-002 Setting up stub pages for sales and subscriptions

## 113-003 Updating migrations and resetting the database
Since we have some dirty development data, we want to reset the DB to a known good version.

```shell
soda generate fizz SeedWidgets
```

After writing the migration, run:
```shell
soda reset
```

## 114-00 4 Listing all sales_ database query
Since we we're doing a JOIN and our base table(FROM <table>) is orders table, let's put the data that the query returns, in the `Order` struct, like
`Widget` field and ... . Then create `GetAllOrders` func.

## 115-005 Listing all sales_ database function

## 116-006 Listing all sales_ writing the API handler and route

## 117-007 Listing all sales_ front end javascript

## 118-008 Displaying our results in a table

## 119-009 Making our table prettier, and adding some checks in JavaScript

## 120-010 Challenge_ Listing all Bronze Plan subscribers

## 121-011 Solution to challenge

## 122-012 Displaying a sale_ part 1

## 123-013 Displaying a sale_ part 2

## 124-014 Displaying a subscription
