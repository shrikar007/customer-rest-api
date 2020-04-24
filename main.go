package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shrikar007/customer-rest-api/crud_interface"
	"github.com/shrikar007/customer-rest-api/dberror"
	"github.com/shrikar007/customer-rest-api/requests"
	"github.com/shrikar007/customer-rest-api/response"
	"github.com/shrikar007/customer-rest-api/structs"

	"log"
	"net/http"
)

type Mysql struct {
	Db *gorm.DB
}


func main() {

	dba, err := gorm.Open("mysql", "root:root@tcp(sqldb:3306)/")
	dba.Exec("CREATE DATABASE IF NOT EXISTS"+" customer")
	dba.Close()

	db, err := gorm.Open("mysql", "root:root@tcp(sqldb:3306)/customer?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Println(err)
	}
	if (!db.HasTable(&structs.Customer{})) {
		db.AutoMigrate(&structs.Customer{})
	}
	set := &Mysql{db}
	Init(set)
}
func Init(d crud_interface.Database) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/customers", func(r chi.Router) {
		r.Post("/", d.CreateCustomer)
		r.Get("/", d.GetAll)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(d.CustomerCtx)
			r.Get("/", d.GetId)
		})
	})
	log.Fatal(http.ListenAndServe(":8086", r))
}

func (db *Mysql) CustomerCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var temp structs.Customer
		customerID := chi.URLParam(r, "id")
		DB:= db.Db.Table("customers").Where("id = ?", customerID).Find(&temp)
		//fmt.Println(temp)
		if DB.RowsAffected == 0{
			err:=errors.New("ID not Found")
			render.Render(w, r, dberror.ErrRender(err))
			return
		} else{
			ctx := context.WithValue(r.Context(), "customer", temp)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func (db *Mysql)  CreateCustomer(writer http.ResponseWriter, request *http.Request) {
	//var expense structs.Expense

	var req requests.CreateCustomerRequest
	err := render.Bind(request, &req)
	if err != nil {
		log.Println(err)
		return
	}
	customer := *req.Customer
	db1 := db.Db.Create(&customer)
	if db1.RowsAffected != 0 {
		_, _ = fmt.Fprintln(writer, `{"success": true}`)
		return

	} else {
		err := errors.New("Unable to update")
		render.Render(writer, request, dberror.ErrRender(err))
		return
	}
	//db.Db.Close()

}

func (db *Mysql) GetId(writer http.ResponseWriter, request *http.Request) {

	custo:= request.Context().Value("customer").(structs.Customer)
	fmt.Println(db.Db.Value)

		render.Render(writer, request, response.Getoneresponse(&custo))

}

func (db *Mysql) GetAll(writer http.ResponseWriter, request *http.Request) {
	var customers structs.Customers

	//Db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/expense?charset=utf8&parseTime=True")

	db1 := db.Db.Find(&customers)
	if db1.RowsAffected != 0 {
		_, _ = fmt.Fprintln(writer, `{"success": true}`)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		render.Render(writer, request, response.Getallresponse(&customers))
		return
	} else {
		err := errors.New("Unable to update")
		render.Render(writer, request, dberror.ErrRender(err))
		return
	}
}

