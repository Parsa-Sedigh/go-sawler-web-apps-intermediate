# Section 07 Protecting routes on the Front End and improving authentication

## 94-001 Writing middleware on the front end to check authentication
We want to authenticate our MPA.

If you're building a SPA, this section is not relevant.

## 95-002 Protecting routes on the front end
Anytime someone logs in or logs out, we need to regenerate or renew his token in session by saying:
```go
app.Session.RenewToken(r.Context())
```
You can write code without doing this but it's a good idea to do this from a security POV.

## 96-003 Logging out from the front end
Currently, by default our session package uses cookies as session store and that's gonna be annoying because everytime we
start and stop our app, we lose all of our sessions. Now this is not the case for the tokens because we're storing them in DB, but we
still have the token in localStorage of browser and therefore there will be inconsistency in UI and after a page reload, the user will be logged out
by server.

We'll take advantage of the database store for `scs` package to store the sessions in DB.

## 97-004 Saving sessions in the database
Add this package which is an addon to scs package and it's named `github.com/alexedrwards/scs/mysqlstore`. It has some steps to set up.
It needs a sessions table, so we need a sql(not fizz) migration:
```shell
soda generate sql CreateSessionsTable
```
Paste the SQL from docs to up migration.

```shell
soda migrate
```

```shell
go get github.com/alexedwards/scs/mysqlstore
```

Now after login, if you stop the app and start it again, we're still logged in. 

---

097 mysqlstore-for-sessions
https://github.com/alexedwards/scs/tree/master/mysqlstore