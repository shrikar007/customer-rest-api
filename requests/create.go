package requests
import (
	"errors"
	"github.com/shrikar007/customer-rest-api/structs"

	"net/http"
)
type CreateCustomerRequest struct {
	*structs.Customer
}
func (c *CreateCustomerRequest) Bind(r *http.Request) error {
	if c.Company_Name== "" {
		return errors.New("Company name is either empty or invalid")
	}
	if c.Full_Name == "" {
		return errors.New("Full name is either empty or invalid")
	}

	return nil
}



