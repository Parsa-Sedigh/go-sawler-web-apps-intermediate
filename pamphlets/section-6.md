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
## 75-009 Making sure that everything works
## 76-010 Create a function to generate a token
## 77-011 Generating and sending back a token
## 78-012 Saving the token to the database
## 79-013 Saving the token to local storage
## 80-014 Changing the login link based on authentication status
## 81-015 Checking authentication on the back end
## 82-016 A bit of housekeeping
## 83-017 Creating stub functions to validate a token
## 84-018 Extracting the token from the authorization header
## 85-019 Validating the token on the back end
## 86-020 Testing out our token validation
## 87-021 Challenge_ Checking for expiry
## 88-022 Solution to challenge
## 89-023 Implementing middleware to protect specfic routes
## 90-024 Trying out a protected route
## 90-025 Converting the Virtual Terminal post to use the back end
## 90-026 Changing the virtual terminal page to use fetch
## 90-027 Verifying the saved transaction