# Section 04 Selling a product online

## 27-001 What are we going to build
We're gonna use a test creditcard number which doesn't really charge a credit card, but it takes us through the steps as a real credit card
was entered.

## 28-002 Create the database
Sequel ACE as DB client. Create a new connection called widgets. Host is localhost(127.0.0.1), username is root(default user after you
install mariadb at least on mac), no password.

Create a new DB called widgets.

We don't want to connect to DB as root user, so run:
```sql
GRANT ALL ON widgets.* TO '<user>'@'%' IDENTIFIED BY '';
```
The `%` means anything, so we can connect from anywhere.

The value for `IDENTIFIED BY ''` is whatever we want for password.

## 29-003 Connecting to the database
In `internal` folder, create `driver` folder and create `driver.go`. That's where we connect to DB.

Get: `github.com/go-sql-driver/mysql`.

For now, we can read dsn(connection string) from the flag in command and if you wanna use an env var, that's fine too.
So write `flag.StringVar(&cfg.db.dsn)`.

```shell
make start_front
```

Now we need to do the same(connecting to DB) in our api.go .

Then run below to run both frontend and backend:
```shell
make start
```

## 30-004 Creating a product page
We could embed the static content into our project, the same way we embedded our templates, but it's not good. Let's serve them from an
external directory called `static`.

After creating the `fileServer`, we should be able to hit that image by going to our browser and typing in: `localhost:4000/static/widget.png`.

After creating a page template, we need a a route and a handler that serves it. In this case, we named the handler `ChargeOnce`. 

## 31-005 Creating the product form
We need to make the JS script for terminal.page , into it's own partial, so we can reuse it anywhere.

## 32-006 Moving JavaScript to a reusable file
Create `stripe-js.partial.gohtml`.

We need to have a fixed amount in buy-once.page . But a hidden field is not enough because someone could go in with their JS dev tools,
change the price to whatever they wanted like 1 cent and order it. 

## 33-007 Modifying the handler to take a struct
In `internal`, create `models` folder and we're putting it there because we want to share these models between the frontend and backend.

We don't consider price as decimal(in `Widget`) because we're not gonna store decimal values to avoid floating point errors.

10 dollars mean 1000(1000 cents).

## 34-008 Update the Widget page to use data passed to the template

## 35-009 Creating a formatCurrency template function

## 36-010 Testing the transaction functionality

## 37-011 Creating a database table for items for sale
We need to store transaction info in DB because at some point we might have to refund that transaction and it's a good practice to
keep track of things like credit card transactions.

To test things, run:
```shell
make stop
make start
```

Now in browser(since we wanna only test a GET req) go to: `localhost:4001/api/widget/1`

The fact that we have DB functions like `GetWidget` in internal/models/models.go which is our shared codebase between frontend and
backend, means we can get the content from DB and serve it as HTML or get the content from DB and serve it as JSON(API) or whatever
format we need to use.

## 38-012 Running database migrations
We'll use soda for DB migrations. It's developed as go buffalo framework.

Get `database.yml` from resources and fill the `user` and `password`.

Get rid of widgets table that you have created before.

Then run this command where the `migrations` folder and `database.yml` exists:
```shell
soda migrate
```

The `scehma_migration` table added by soda keeps track of our DB migrations.

## 39-013 Creating database models
Create `migrations` folder at root level and copy the migration files in course resources. Also copy `database.yml` to root level.

We need to add a column named image to widget table. For this, first, at root level you can run this command:
```shell
`soda generate fizz AddImageToWidgets`
```
The above command, `generate` is gonna create an up and down migration and with `fizz` it's gonna be in fizz format we wanna put them in and
we call that migration `AddImageToWidgets`.
This create 2 files.

Now copy this to up file: `add_column("widgets", "image", "string", {})`


Now for generated down migration: `drop_column("widgets", "image")`.

Now run:
```shell
soda migrate
```
It will add the column to table.

Add widget.png which is our static image for a test widget in image column of table.

## 40-014 Working on database functions


## 41-015 Inserting a new transaction
## 42-016 Inserting a new order
## 43-017 An aside_ fixing a problem with calculating the amount
## 44-018 Getting more information about a transaction
## 45-019 Customers
## 46-020 Getting started saving customer and transaction information
## 47-021 Create the save customer database method
## 48-022 Saving the customer, transaction, and order from the handler
## 49-023 Running a test transaction
## 50-024 Fixing a database error, and saving more details
## 51-025 Redirecting after post
## 52-026 Simplifying our PaymentSucceeded handler
## 53-027 Revising our Virtual Terminal
## 54-028 Fixing a mistake in the formatCurrency template function