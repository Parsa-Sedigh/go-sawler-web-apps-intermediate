# Section 14 Microservices

## 162-001 What are microservices

## 163-002 Setting up a simple microservice
We could create a new project BTW.

You can put your microservices in `micro` folder.

Microservices are very rarely exposed to the public internet.

In our case, there's absolutely no reason why that microservice should ever have it's listening port(5000 in case 
of our pdf generator microservice) exposed to the public internet. Instead, our app will call that microservice functions over local network behind the firewall.

## 164-003 Receiving data with our micrsoservice

## 165-004 Generating an invoice as a PDF
```shell
go get github.com/phpdave11/gofpdf
go get github.com/phpdave11/gofpdf/contrib/gofpdi
```

Put the pdf in course resources at `pdf-templates` folder.

## 166-005 Testing our PDF
## 167-006 Mailing the invoice
## 168-007 Call the microservice when a Widget is sold