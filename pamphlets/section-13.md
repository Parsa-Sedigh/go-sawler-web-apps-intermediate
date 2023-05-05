# Section 13 Managing Users

## 145-001 Setting up templates to manage users

## 146-002 Adding routes and handlers on the front end

## 147-003 Writing the database functions to manage users

## 148-004 Creating a handler and route for all users on the back end

## 149-005 Updating the front end to call AllUsers

## 150-006 Displaying the list of users

## 151-007 Creating a user add_edit form
We did not put required attr on password and verify password because when we're editing a user, if we don't enter the password and verify password
fields, then we're not changing the password, otherwise we are.


## 152-008 Call the api back end to get one user

## 153-009 Populating the user form, and a challenge
We shouldn't be able to delete our own user in the UI! So we need to fix this. We need to determine whether or not the user we're looking at(in
the form), is the currently logged in user. We should look at the fronted session.

## 154-010 Solution to challenge

So user should not see the delete button when the current page is for his own user.

## 155-011 Saving an edited user - part one

## 156-012 Saving an edited user - part two

## 157-013 Deleting a user
## 158-014 Removing the deleted users token from the database
## 159-015 Setting up websockets
## 160-016 Connecting to WebSockets from the browser
## 161-017 Logging the deleted user out over websockets