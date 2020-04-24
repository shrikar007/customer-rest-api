package response

import (
	"github.com/shrikar007/customer-rest-api/structs"
	"net/http"
)




type GetOneStruct struct {
	*structs.Customer
}

func (GetOneStruct) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}



func Getoneresponse(customer *structs.Customer) *GetOneStruct {
	return &GetOneStruct{Customer: customer}

}
