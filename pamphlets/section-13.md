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
After a user has been deleted, we should log him out as well.

## 158-014 Removing the deleted users token from the database
We could set up a foreign key relationship and that would automatically delete the related token for the deleted user.

But we want to do this at the code-level and not at the database-level.

After deleting the user's token, we need to log the user out. To do this, in the middleware, we could do a DB-lookup at that point.
We would check for the existence of userID in session, then check the users table to see if that user still exists. But there are 2 drawbacks:
- it doesn't **instantly** log the user out as soon as that user is deleted
- every single time an authenticated user tries to access sth, we would do a DB look up and that's expensive 

Another approach is using websockets.

## 159-015 Setting up websockets
```shell
go get github.com/gorilla/websocket
```

When you connect using websockets, your connection to the web server is upgraded to permit 2-way communication.

With `CheckOrigin` we can secure our websocket connections. Now since we're not receiving anything from the frontend other than the initial connection,
we don't care about that function, so we always `return true` there. But if you're communicating from your browser to the server, you wanna do some
logic there.

In `WsEndpoint` first thing we wanna do is to upgrade the connection. So when that handler is hit, we need to upgrade the connection.

We need to run sth continuously in the background(using a goroutine that has a infinite for loop) to listen for 
websocket connections and we name that func, `ListenForWS`. Since we want to run it all the time, 

For connecting to websocket, we use a GET HTTP method.

## 160-016 Connecting to WebSockets from the browser

## 161-017 Logging the deleted user out over websockets
To test things, open a window with one user and another one on a private browser window in order to have 2 sessions and login with both.
Then with the first user, delete the second user and the second user should log out immediately(you can put the windows next to each other to see
things live).
