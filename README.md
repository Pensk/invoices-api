# invoices-api
offers endpoints for creating and viewing invoices

Requires docker/make

## setup

Make setup may be unreliable if mysql takes time to init.
``` sh
$ make setup
```

Run the setup commands in order manually
``` sh
$ make build
$ make mod-tidy
$ make up
$ make migrate
$ make seed-db
```

Run tests
``` sh
$ make test
```

### first time run

seed-db adds seed data for companies/clients. This can be used for a test endpoint to create a user:

``` sh
curl -X POST -H "Content-Type: application/json" -d '{"email":"test@example.com","password":"password","company_id":1, "name": "user name"}' http://localhost:8080/api/users/signup -i
```

You can then use that to authenticate
``` json
curl -X POST -H "Content-Type: application/json" -d '{"email":"test@example.com","password":"password"}' http://localhost:8080/api/users/authenticate -i
```

Put the access token in the "Authorization": "Bearer {token}" header for api access.

Invoices is under /api/invoices.

``` go
// post /api/invoices
type CreateInvoiceRequest struct {
	ClientID      int    `json:"client_id"`
	IssueDate     string `json:"issue_date"`
	PaymentAmount uint64 `json:"payment_amount"`
	DueDate       string `json:"due_date"`
}

// get /api/invoices
type ListInvoiceRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	PerPage   int    `json:"per_page"`
	Page      int    `json:"page"`
}
```
