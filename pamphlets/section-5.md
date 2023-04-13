# Section 5 - 05 Setting up and charging a recurring payment using Stripe Plans

## 55-001 What are we going to build in this section
You can look at the customers page of stripe to confirm user subscription.

In this case, we decided to have all of the logic on backend. So frontend app serves a webpage and nothing else. So our frontend
will call the backend which is a totally different application listening on a different port.

## 56-002 Creating a Plan on the Stripe Dashboard
If you want to sell a subscription, we need to set up stripe plan in stripe dashboard. Make sure the `viewing test data` is on.
Now we want to add a subscription plan. So go to `Products`>`Add product`. For name: Bronze Widget Plan. For description: Receive
three widgets for the price of two every month!

Choose `Recurring`.

## 57-003 Creating stubs for the front end page and handler

## 58-004 Setting up the form
TODO: DB part at the beginning

```shell
soda generate fizz AddColsToWidgets
```
After writing the migrations, run:
```shell
soda migrate
```

Copy the ID of the plan you created in stripe and paste it to the `plan_id` column of bronze plan row in `widgets` table.

Since you changed the structure of tables, you need to change the related models now.

## 59-005 Working on the JavaScript for plans
The way you charge a user for a recurring subscription in stripe, is a bit different than buying a product once. In buying a product,
we get a payment intent. But for recurring subscriptions, first we need to create a customer. Then we need to subscribe that 
stripe customer to a plan. When I say "create a customer", I don't mean create one and store it in our local database, I mean to create
a customer on the stripe backend and then use the customer obj from that and subscribe him to a plan.

## 60-006 Continuing with the Javascript for subscribing to a plan

## 61-007 Create a handler for the POST request after a user is subscribed

## 62-008 Create methods to create a Stripe customer and subscribe to a plan
## 63-009 Updating our handler to complete a subscription
## 64-010 Saving transaction & customer information to the database
## 65-011 Saving transaction & customer information II
## 66-012 Displaying a receipt page for the Bronze Plan