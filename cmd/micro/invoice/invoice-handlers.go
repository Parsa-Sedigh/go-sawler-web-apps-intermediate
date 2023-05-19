package main

import (
	"fmt"
	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
	"net/http"
	"time"
)

type Order struct {
	ID        int       `json:"id"`
	Quantity  int       `json:"quantity"`
	Amount    int       `json:"amount"`
	Product   string    `json:"product"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (app *application) CreateAndSendInvoice(w http.ResponseWriter, r *http.Request) {
	// receive json
	var order Order

	err := app.readJSON(w, r, &order)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	// generate a pdf invoice

	// create mail

	// send mail with attachment

	// send response
	var resp struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}
	resp.Error = false
	resp.Message = fmt.Sprintf("Invoice %d.pdf created and sent to %s", order.ID, order.Email)

	app.writeJSON(w, http.StatusCreated, resp)
}

// createInvoicePDF writes a pdf to disk
func (app *application) createInvoicePDF(order Order) error {
	pdf := gofpdf.New("P", "mm", "Letter", "")
	pdf.SetMargins(10, 13, 10)

	// doesn't matter for our case, but it's a good practice. Invoices will all be 1 page long, so you don't have to do this, but we did this to be consistent
	pdf.SetAutoPageBreak(true, 0)

	/* Create an importer that will allow us to import the pdf we want to write on, that way, we only have to put a few things in there
	programmatically.*/
	importer := gofpdi.NewImporter()

	/* t stands for template.

	Since we start the app from root level of project(using make commands), we say look at the pdf-templates at THAT LEVEL.*/
	t := importer.ImportPage(pdf, "./pdf-templates/invoice.pdf", 1, "/MediaBox")

	pdf.AddPage()

	importer.UseImportedTemplate(pdf, t, 0, 0, 215.9, 0)

	/* We want to write on that template, but we need to tell it where we wanna write? Think about X and Y axis's.*/

	// write info
	pdf.SetY(50)
	pdf.SetX(10)
	pdf.SetFont("Times", "", 11)
	pdf.CellFormat(97, 8, fmt.Sprintf("Attention: %s %s", order.FirstName, order.LastName), "", 0, "L", false, 0, "")

	pdf.Ln(5)

	// exactly the same location at previous writing that we did, but down one line(with the call to pdf.Ln(5) that we did in previous line)
	pdf.CellFormat(97, 8, order.Email, "", 0, "L", false, 0, "")

	pdf.Ln(5)

	pdf.CellFormat(97, 8, order.CreatedAt.Format("2006-01-02"), "", 0, "L", false, 0, "")

	// populate the table(we measured X and Y using a ruler)
	pdf.SetX(58)
	pdf.SetY(93)

	pdf.CellFormat(155, 8, order.Product, "", 0, "L", false, 0, "")

	// move to the right to print the next thing:
	pdf.SetX(166)

	pdf.CellFormat(20, 8, fmt.Sprintf("%d", order.Quantity), "", 0, "C", false, 0, "")

	pdf.SetX(185)

	// alignment is R so if there are more items in the future, the decimal points will all lineup nicely
	pdf.CellFormat(20, 8, fmt.Sprintf("$%.2f", float32(order.Amount/100.0)), "", 0, "R", false, 0, "")

	/* We need to export this as a PDF, but we need to put it somewhere, right? So we're gonna assume that there's a folder at the root level of our app
	called `invoices`.*/
	invoicePath := fmt.Sprintf("./invoices/%d.pdf", order.ID)
	err := pdf.OutputFileAndClose(invoicePath)
	if err != nil {
		return err
	}

	return nil
}
