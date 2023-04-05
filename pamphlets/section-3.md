# Section 03 Building a virtual credit card terminal
## 11-001 What we're going to build
For stripe:
- a credit card that is always successful: `4242 4242 4242 4242`
- to get error: `4000 0000 0000 0002`

## 12-002 Setting up a (trivial) web application
We're gonna produce 2 different binaries with this single codebase. One for frontend and one for backend. Why?
For example we might wanna split the load between two applications, perhaps on two different servers.

With `cssVersion` constant, we're gonna append it to any external css or JS files and when we increment it, that will force most
browsers to bring the new version down. So we won't have to clear the cache when things aren't working the way we expect.

**Tip:** In order to get the stripe keys(both private and public), we don't want to read those from command line flags. Because when
someone is logged into the server, they could type `ps -aux` and possibly get that secret key and we don't want that to be available
anywhere. So instead, we read both keys from environment variables.

Install: `github.com/go-chi/chi/v5`.

To test, run this from root of project:
```shell
go run ./cmd/web
```

## 13-003 Setting up routes and building a render function
Go to `localhost:4000/virtual-terminal` to test the handler.

The `go:embed` directive allows us to compile our app, including all of it's associated templates into a single binary.

Many go templates are comprised of the template itself, some partials that we want to include in that template and the entire template is
governed by some base layer.

## 14-004 Displaying one page
To enable syntax highlighting with .tmpl , go to view>appearance>show status bar then at bottom right, click on `Plain Text` and relate
`Go Template` to this extension.

`air`(check it on github) is an application that will automatically recompile(a temporary binary) and reload our app anytime we
make a change to the source code.

Now to run your app instead of `go run ./cmd/web`, just type `air` at root level.

## 15-005 A better extension for Go templates and VS Code
Uninstall `go template support` extension and instead install `gotemplate-syntax`. You **could** change some settings to have good
highlighting for `tmpl` files, but let's use `gohtml`. 

## 16-006 Creating the form
We have to use `novalidate` on `<form>` if we're going to use bootstrap's validation.

## 17-007 Connecting our form to stripe


## 18-008 Client side validation
## 19-009 Getting the paymentIntent - setting up the back end package
## 20-010 Getting the paymentIntent - starting work on the back end api
## 21-011 Getting the paymentIntent - setting up a route and handler, and using make
## 22-012 Getting the paymentIntent - finishing up our handler
## 23-013 Updating the front end JavaScript to call our paymentIntent handler
## 24-014 Getting the payment intent, and completing the transaction
## 25-015 Generating a receipt
## 26-016 Cleaning up the API url and Stripe Publishable Key on our form



014 Air
https://github.com/cosmtrek/air

014 Bootstrap
https://getbootstrap.com

020 go-chi-cors
https://github.com/go-chi/cors
