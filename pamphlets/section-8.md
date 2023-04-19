# Section 08 Mail and Password Resets

## 98-001 Password resets
For password reset, we need to send an email with a signed link to the password reset to user 

## 99-002 Sending mail Part I
Install `github.com/xhit/go-simple-mail/v2`.

For right now, we're gonna assume that the mail that we send from frontend app might be different than the mail that is sent from the backend.
So we don't put the logic for it in a separate package to be used in both apps. We want to have separate functions for both apps. So in api folder,
create `mailer.go`.

Everytime we send email, we want to send it in html version and a plain text version.

## 100-003 Mailtrap.io
We need an SMTP server.

Go to `Inboxes`>`SMTP Settings`.

Under `use these settings to send messages from your email client or mail transfer agent`, there's a `show credentials` button and it has all
the info we're gonna need to be able to send email from our app.

Note: We want our email to be trapped somewhere and not reach the destination.

## 101-004 Sending mail Part II
**Note:** We don't want to hardcode the SMTP configs, instead we want them to be command line parameters or environment variables, so we can work in multiple
environments like development and not have to change the actual source code before changing the environment(like going to production).

## 102-005 Creating our mail templates and sending a test email
We want the reset password link to be secure. We don't want to make that open to anyone. Now obviously we can't have people log in to change the password
because they don't know what the password is at the first place! So what we need is some means of securing that to make sure that the only
people who have access to the link are the people who should have access to link.

There are number of ways of doing this.

One common approach is to append a unique token to the email or to give them a token they have to enter into a form or sth like that and OFC you
look up that token in the DB and verify that it's valid and ... . That seems a lot of work.

The tutor suggests instead we use signed emails, email which have a hash or some kind of code in the URL itself that will ensure that URL can't
be changed and we'll also make sure that an email expires after a set length of time.

```shell
go get github.com/bwmarrin/go-alone
```

Create a new internal package that uses the functionality from `go-alone` package.

When we want to sign sth either on frontend(server rendered frontend app) or on backend, we can use our `urlsigner` package.

To use a secret key, we can either use an env var or use a command line flag. The `secretkey` in config struct is the key that we use
to sign our URLs and `frontend` would be the address for our frontend.

We put them as command line flags so we don't have to write them everytime we want to run the application(they have sensible defaults, but can be
overrided through command line).

With `VerifyToken`, we verify the link people clicks on, hasn't been changed in any way.



## 103-006 Implementing signed links for our email message

## 104-007 Using our urlsigner package

## 105-008 Creating the reset password route and handler
## 106-009 Setting up the reset password page
## 107-010 Creating a back end route to handle password resets
## 108-011 Setting an expiry for password reset emails
## 109-012 Adding an encryption package
## 110-013 Using our encryption package to lock down password resets