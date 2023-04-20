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
Acronyms in go are all capitalized.

If user changes the signed url(email or hash query params), it becomes invalid.

We have tamper proof URLs and the great thing about this is that at no point, up to sending the reset password email, we have done no DB lookups
other than to verify the email is valid when we received the reset-password form req. Everything else was handled by cryptography.

## 106-009 Setting up the reset password page
To test things:
```shell
make stop
make start_back
air
```

## 107-010 Creating a back end route to handle password resets

## 108-011 Setting an expiry for password reset emails
When you send a password reset email, we need a expiry.

Also one other security issue we currently have is when we're sending the payload of /reset-password req, we're sending email field. Now a malicious user
who examines the source code, can change the email in that payload and therefore he can change somebody else's password.

There are many ways of fixing this. One approach is sticking it in a session. Because the session for our frontend exists only for our frontend and
we have no access to it from backend. We can encrypt that email to a text value when we write it to that `reset-password.page` and on backend we decrypt it. 

## 109-012 Adding an encryption package
User can change the source code of payload for reset-password to some other email.

Create a new package so it's available to both frontend and backend apps, named encryption.

The encryption algorithms we're using, require a very specific length for the secret key. It needs to be exactly 32 characters longs.

Note: By changing the secret key, if you have existing links for resetting password, they don't work anymore the next time you run the app aftere changin the
secret key.

By defining and initializing the `Encryption` type, that's how we initiate the `Encryption` package.

Now we must encode the email on frontend and since we have the same secret key on backend, we should be able to decode it as well.

## 110-013 Using our encryption package to lock down password resets
By sending an encrypted version of email, nobody will be able to guess the correct encryption algorithm if they wanna change the emails.

Now the user can't see the email in the payload in the source code of website, because it's encrypted.

You can also store encrypted values in DB for sensitive information, using the encryption package.