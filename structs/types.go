package structs

type Customer struct {
	Customer_Id     int    `json:"customer_id"`
	Company_Name    string `json:"company_name"`
	Company_Type_Id int    `json:"comany_type"`
	Full_Name       string `json:"full_name"`
}

type Customers []Customer


