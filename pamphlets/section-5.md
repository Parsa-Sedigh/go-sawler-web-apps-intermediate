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
Add a new row to widgets table as our bronze plan and name it `Bronze Plan`. It's price will be 20$ every month so 2000 cents in DB.

Then in stripe dashboard, store API ID of the plan you created that is listed in Pricing table.

```shell
soda generate fizz AddColsToWidgets
```
and add these to it's `up` file:
`
add_column("widgets", "is_recurring", "bool", {"default": 0})
add_column("widgets", "plan_id", "string", {"default": ""})
`
and for `down` file:
`
drop_column("widgets", "is_recurring");
drop_column("widgets", "plan_id");
`

After writing the migrations, run:
```shell
soda migrate
```

Copy the ID of the plan you created in stripe(which is in Pricing table) and paste it to the `plan_id` column of bronze plan row in
`widgets` table.

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
Currently, if you test things, you get: status: 400 with `No such plan` as message. It's because we copied the wrong plan id. We
should copy the one that is in `Pricing list` not in `Details` in stripe dashboard. So the plan_id starts with `price_` not `prod_`.

You can see the subscriptions in `subscriptions` page of stripe.

## 64-010 Saving transaction & customer information to the database
We want to save the transaction, customer and order locally.

With our approach, we're only gonna communicate with backend, we're not gonna POST anything to frontend and then redirect the user
and ... , we'll handle everything from backend.

We do have `SaveCustomer` in frontend code and not on backend, but we need it in both places. So we could move it to it's own
package and call the same function from both places. **But** we're gonna assume that things might diverge over time, that I might
be doing slightly different things on frontend for SaveCustomer than we're on backend. So I'm just gonna duplicate the code for it
instead of having it in one place.

We don't get a bank return code on frontend, because when somebody subscribes to a plan in stripe, it creates an invoice and it lets
that invoice sit for a while for some reason. So we don't get a bank return code at this point. So we leave it when creating
a transaction(`txn`) in `CreateCustomerAndSubscribeToPlan`.

## 65-011 Saving transaction & customer information II

## 66-012 Displaying a receipt page for the Bronze Plan
How are we gonna save all of the info we want on the receipt page of a plan subscription between the `bronze-plan.page` and receipt page?
There are number of ways for doing it. We can store some values in JS session, redirect the user to another page and then read those values
from session(`sessionStorage`).