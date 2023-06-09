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
When you have an id in go, go likes it to be both capital, so: ID.

You can use backticks to put string in multiple lines like in sql statements.

## 41-015 Inserting a new transaction

## 42-016 Inserting a new order
The transaction we save, will change it's status throughout the lifecycle of the actual full transaction where we're charging a
credit card.

## 43-017 An aside_ fixing a problem with calculating the amount

## 44-018 Getting more information about a transaction
Under no circumstances we will ever store a credit card number(we just store the last 4 digits), but it's useful to be able to verify
that you're looking at the right transaction by saying: Do the last 4 digits match the ones I'm looking for and is the expiry date the same?
and that's enough info to verify that you have the right transaction, but not enough for someone to actually charge that credit card when
our DB ever gets compromised.

PaymentIntent changes during it's lifecycle, when you get it initially, you have some bit of info, when you get it later on, you might
have a different bit of info. So we need to get an **existing** paymentIntent, not create a new one, so for convenience, let's create
a method on *Card, named `RetrievePaymentIntent`. So we used the word retrieve, because we're getting an **existing** one.

After changes, stop the app and run it again:
```shell
make stop
make start
```

## 45-019 Customers
Right now, our handlers are doing nothing more than grabbing info from out transaction and from the form and displaying a receipt to the
end user. But we also want to save the transaction and order info.

Create a new table and add new columns to the transactions table using `soda`:
```shell
soda generate fizz CreateCustomerTable # will generate 2 migrations: up and down
```

Now for adding columns:
```shell
soda generate fizz AddColsToTransactions
```

Now run:
```shell
soda migrate
```

Now we need to change the types accordingly in models.go .

After adding customerID in Order type of models.go , we need to reflect this change in DB as well, so we need a migration:
```shell
soda generate fizz AddCustomerIDToOrders
```

Now run:
```shell
soda migrate
```

## 46-020 Getting started saving customer and transaction information
Currently, if you're on /payment-succeeded page and reload that page, it's gonna ask you: are you sure you want to resubmit the form?
And if you say yes, it will charge the user again(resubmit the form)!

We don't want this to happen.

To fix this, we need a session and we need to redirect. But before redirecting, we need to save some info.

Install: `github.com/alexedwards/scs/v2`.

A middleware, receives a http.Handler, modify it and return a http.Handler . 

## 47-021 Create the save customer database method

## 48-022 Saving the customer, transaction, and order from the handler
It's not good to get the customer name from cardholder name(maybe the customer is a 14 year old and using his mom's or dad's card, so the
name is not accurate). So add `first_name` and `last_name` fields to form when buying. 

It's good to use air when you're making changes to template files. But if you're not, use makefile commands.

Why we didn't call the DB directly in `SaveCustomer`? Why we didn't put the logic of `SaveCustomer` right in the handler?

We could have. But maybe we wanna do more in SaveCustomer function than inserting sth into the DB, maybe we wanna fire off an email
or an alert. Maybe we want to create an audit log, save every single transaction and all of the details associated in the DB.
The place for that is in `SaveCustomer`(in it's own function) and not at the DB level.

After creating a customer, we need to create a transaction and then create an order, because we need the transactionID before we can
save the order.

So we don't call the DB directly in `PaymentSucceeded`.

## 49-023 Running a test transaction

## 50-024 Fixing a database error, and saving more details
We need to store payment_intent and payment_method if we're gonna do a refund of transaction at some point. To update the tables, run:
```shell
soda generate fizz AddColsToTxn
```
And populate the generated files with appropriate logic and run:
```shell
soda migrate
```

Then add the new columns(changes to table structure) to models in models.go .

## 51-025 Redirecting after post
After someone's credit-card is charged, they'll go to a receipt page called `/payment-succeeded` but that page is actually a direct result of
a POST req(so if we refresh that page, it's gonna make POST reqs! which is not good) and **it's good practice to redirect people
somewhere else so they can't accidentally POST that data twice or multiple times**. So in `PaymentSucceeded` handler, we write data to
session and then redirect user to a new page.

Unless you're putting a primitive like a string or an int in the session, you need to register it's type. If you're coming from a language
that is not strongly typed, it may be confusing to you. For this, we're gonna use `gob` in main.go .

With this:
```go
gob.Register(map[string]interface{}{})
```
now, when we run the program, that type type will be registered and we can put the type `map[string]interface{}` into the session in our
code.

Now if someone tries to reload that page, since we removed info from the session, we'll get an error(you can do error checking),
but it won't POST the data again.

## 52-026 Simplifying our PaymentSucceeded handler
Create `TransactionData` type and `GetTransactionData` func.

## 53-027 Revising our Virtual Terminal
For amount, 10.00 is equal to 1000(both are 10 dollars - 1000 shows cents as well).

If we put 33.45 , we get a long decimal that even multiplying it by 100 doesn/t make it an int and on backend we expect an int.
To fix this, we use parseInt() for `charge_amount` onChange listener. So now we pass 3345 as 33.45 cents.

In reality, if we had one application running our frontend and one application running our backend, if you wanted to be a purist,
you wouldn't use POST reqs on the frontend at all. You'd handle everything using your backend API. But for monolith apps, the approach
we used till here, will work fine.

We need to protect /virtual-terminal page so not everybody can charge a credit card!

## 54-028 Fixing a mistake in the formatCurrency template function