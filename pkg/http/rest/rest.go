package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgparker/httpsrv"
	"github.com/google/uuid"
	"github.com/husobee/vestigo"
	"github.com/scottshotgg/yugabyte-test/db"
)

func New(db db.DB) *http.Server {
	return httpsrv.NewWithDefault(newRouter(db))
}

func newRouter(db db.DB) http.Handler {
	r := vestigo.NewRouter()

	r.Get("/ping", pingHandler())
	r.Get("/employee/:id", getEmployeeByIDHandler(db))
	r.Post("/employee", postEmployeeHandler(db))
	return r
}

func pingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"response": "pong"}`))
	}
}

func getEmployeeByIDHandler(db db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := vestigo.Param(r, "id")
		if idStr == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		res, err := db.GetEmployeeByID(id)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resBody, err := json.Marshal(res)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resBody)
	}
}

func postEmployeeHandler(yb db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyyy, err := ioutil.ReadAll(r.Body)
		if err != nil {
			defer r.Body.Close()
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Println("body", string(bodyyy))

		defer r.Body.Close()

		emp := &db.Employee{}

		err = json.Unmarshal(bodyyy, emp)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		emp.ID, err = yb.NewEmployee(emp)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		blob, err := json.Marshal(emp)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(blob))
	}
}
