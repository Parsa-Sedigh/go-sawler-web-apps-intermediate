# Section 06 Authentication

## 67-001 Introduction
Authentication: how we ensure our users are valid.

Frontend: session auth. For a website that has login and doesn't have an API, we tend to use sessions. We set sth like `isAuth = true` and
until that session expires or the user logs out, they are authenticated. We validate that session variable on every request.

backend: tokens. There are lots of different kinds of tokens:
- HTTP basic: Simple to implement, very slow. Typically not an ideal solution if you're gonna have any amount of traffic.
- tokens: We validate whether or not the token is valid on the backend with every req to backend.
- stateful tokens: We store the token and it's expiry date in DB. If we want to invalidate a user, we can remove that token in
DB or mark it as expired.
- stateless tokens(JWT): it's expiry date is stored in the token itself. The big disadvantage is that you can't revoke a token on a per token basis.
If for example you have some user that's gonna rogue, the only way to ensure that you get that user off the system instantly, is to revoke
all of the tokens all at once and that's not ideal. The other problem with JWT is requires a lot of client side code to refresh the token and ... .
- API keys: Like github API keys
- OAuth 2.0: Authenticate a user using for example his google account(user has to have an account on that system like google, facebook or ...)

We're gonna use stateful tokens.

## 68-002 Creating a login page

## 69-003 Writing the stub javascript to authenticate against the back end

## 70-004 Create a route and handler for authentication
In readJSON, since we receive the data arg as a reference to a variable, we're just changing a pointer value.

## 71-005 Create a writeJSON helper function

## 72-006 Starting the authentication process

## 73-007 Creating an invalidCredentials helper function

## 74-008 Creating a passwordMatches helper function
Vscode has difficulty importing bcrypt package. To import it, you can go to go.mod and hover over `module <module name>` and use quick fix.

## 75-009 Making sure that everything works
The user we had from the beginning in users table, his password is `password`.

Note: The email address is case-insensitive(so if user enters all uppercase, we will convert it to lowercase).

## 76-010 Create a function to generate a token
Create `tokens.go` and put it in `models` folder because we might at some point in the future, need to have access to the functionality in this new file
from another application.

We declared `ScopeAuthentication` because when you're working with an API, you'll have different kinds of scope and it's generally considered 
good practice to identify the scope for some particular piece of your code.

## 77-011 Generating and sending back a token
When we generate the token, we need to save it in the DB(so we need migrations and also changes to models).

## 78-012 Saving the token to the database
```shell
soda generate fizz CreateTokensTable
```

In `create_tokens_table` migrations we know `hash` is a reserved word in sql, so we named that column `token_hash`.

```shell
soda migrate
```

## 79-013 Saving the token to local storage
We don't want to keep the old tokens. So once someone logs in, we need to get rid of any preexisting tokens for that user id.

## 80-014 Changing the login link based on authentication status

## 81-015 Checking authentication on the back end
Currently, a logged out user can type a page that we want only logged in users to see, and see that page!

There are a couple ways of doing this:

If we were to do everything completely divorced from the frontend, in other words, all of our authentication is handled by our API server, we can do
that byt it's kinda awkward because we have a couple of options:

- One is to make the URLs that appear in our menu items and a logged out user shouldn't see some of them, make those 
not-standard hrefs(`<a>`), so instead make them call some JS function which first goes and verifies against the backend that you're authenticated and if
it is, then it redirects you to that intended page. This is one option. But it gets awkward fast.
- Allow authentication to the backend like we have right now, but also when we authenticate, to also authenticate through the frontend and store
some session variable which we can check on every req to see if the user is authenticated or not. If this was a SPA, this wouldn't be an issue,
because with SPA we only have one page and we could do all of our authentications through backend without any difficultly.

We'll do both ways.

**First way:**

By calling `checkAuth` function in every page, if the user is not logged in, we see the page for a brief second and that's why this kind of
authentication check using only the backend is not ideal for apps that are not SPAs.

Note: The next lines are from vid #20

If you want to use this approach, move `checkAuth` function definition to `<header>` of html in base.layout.gohtml and in pages that are calling this
function, in a new `block`(`in-head`) that would be where the comment I written in base.layout for `check auth function`, exists.

**Reason:** The way the browsers work is that as a page is being rendered, as soon as it finds that `<script>` tag in `<head>`, everything comes
to halt, it executes that JS and then it moves on(that's the way browsers work right now at least).
So if we put checkAuth() in <head>, the user is not gonna see that page until that JS is executed and if the auth fails(like if the token user sent
either doesn't exist or is expired or is invalid), he gets redirected to /login before seeing anything. This is checking for auth in frontend(for a
server-side rendered app which is not great as you can see).

Another form of checking authentication that is handled server-side with sessions.

## 82-016 A bit of housekeeping

## 83-017 Creating stub functions to validate a token

## 84-018 Extracting the token from the authorization header
Although we're specifying some error messages in `authenticateToken`, we're actually not gonna use them! It's better to give as little info as
possible to users who are trying to authenticate when sth goes wrong. We don't want to give any indication as to what is missing.

So in authentication, make the error messages as bare minimum.

## 85-019 Validating the token on the back end
Our token is a string of 26 characters long.

Q: Why we didn't just store the token in DB? Why we stored a hash of it?

Because we never want to store a valid token in DB. If that ever gets compromised, the hacker has access to **all** the users(since all the tokens
are just there! he can do whatever he wants).

**Just as you never store a password in the DB as plain text, you never store a token in DB, we store a hash of these.**

To convert a slice to an array, write: a[:] . Here a is a slice.

## 86-020 Testing out our token validation

## 87-021 Challenge_ Checking for expiry

## 88-022 Solution to challenge
```shell
soda generate fizz AddExpiryToTokens
```
After changes(adding the column), run:
```shell
soda migrate
```

Now delete the entries in tokens table because their expiry is 0000 , so that we can get new tokens with the correct expiry.

After getting a new token, to test the expiry, in DB change the expiry to some time in past so that we can test the logic for expiry.

## 89-023 Implementing middleware to protect specific routes
We wanna write a middleware to protect access to certain routes on the backend.

`mux.Route()` allows us to create a new mux and apply middleware to it and to group certain kinds of routes logically into one location.
Now any calls to middlewares in the closure we pass as second arg of `mux.Route()`, will **only** apply to the routes that are included in
the routes defined in that group.

## 90-024 Trying out a protected route

## 91-025 Converting the Virtual Terminal post to use the back end
We want to convert the form submission style of sending req in /virtual-terminal(which sends the form submission data to our frontend go app) to
instead an api call to a protected route on our backend go app. So in the backend handler(`VirtualTerminalPaymentSucceeded`) we're not
gonna handle a form post, since data is sent as JSON as the body of POST req.

## 92-026 Changing the virtual terminal page to use fetch
We want the virtual-terminal form to instead of calling a route on frontend go app, we want to use fetch and call a route on the backend.

To do this, get rid of `action` attr on `<form>` and remove `stripe-js` block in that page. Because we won't be calling stripe on frontend anymore since
that will be done on backend.

In this approach, we don't need to submit the form and set the hidden fields in the <form>(by getting the element and setting it's 
`value` property), instead, we just set the values of those hidden fields in JSON payload of ajax req.

## 93-027 Verifying the saved transaction
Currently the payment_intent and payment_method columns in transactions table are stored empty when we submit the virtual terminal page req(charging
a credit card). Those columns should be empty when we subscribe to a plan but not when we're charging a credit card. The problem is we're not
populating those fields in `VirtualTerminalPaymentSucceeded` handler.
