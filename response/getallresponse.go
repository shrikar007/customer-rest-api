package response

import(
	"github.com/shrikar007/customer-rest-api/structs"
	"net/http"
)

type Getallstruct struct {
	 *structs.Customers
}

func (Getallstruct) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func Getallresponse(customers *structs.Customers) *Getallstruct{
	return &Getallstruct{Customers: customers}
}

