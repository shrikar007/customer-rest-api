package structs

type Customer struct {
	Customer_Id     int    `json:"customer_id"`
	Company_Name    string `json:"description"`
	Company_Type_Id int    `json:"typeofaccount"`
	Full_Name       string `json:"amount"`
}

type Customers []Customer


