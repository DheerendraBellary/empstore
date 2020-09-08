package empservice

import (
	"empstore/empstoreapi"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//AddHandler ...
func AddHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddHandler: Received Add Request")
	switch method := r.Method; method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errMsg := fmt.Sprintf("AddHandler: Request Read Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusBadRequest, errMsg)
			return
		}
		addReq := &empstoreapi.AddRequest{}
		err = json.Unmarshal(body, addReq)
		if err != nil {
			errMsg := fmt.Sprintf("AddHandler: Request Processing Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusBadRequest, errMsg)
			return
		}

		err = validateAddReq(addReq)
		if err != nil {
			errMsg := fmt.Sprintf("AddHandler: Request Validation Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusBadRequest, errMsg)
			return
		}

		empID, err := AddEmployee(addReq)
		if err != nil {
			errMsg := fmt.Sprintf("AddHandler: AddEmployee Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusInternalServerError, errMsg)
			return
		}

		addResp := empstoreapi.AddResponse{ID: empID}
		resp, err := json.Marshal(addResp)
		if err != nil {
			errMsg := fmt.Sprintf("AddHandler: Response Building Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusInternalServerError, errMsg)
			return
		}

		fmt.Printf("Employee Added. ID: %v\n", empID)
		writeResponse(w, resp)

	default:
		errMsg := fmt.Sprintf("AddHandler: Invalide Request Method. %v\n", method)
		fmt.Printf(errMsg)
		writeErrorResponse(w, http.StatusBadRequest, errMsg)
		return
	}
}

//SearchHandler ...
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SearchHandler: Received Request")
	switch method := r.Method; method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errMsg := fmt.Sprintf("SearchHandler: Request Read Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusBadRequest, errMsg)
			return
		}
		searchReq := &empstoreapi.SearchRequest{}
		err = json.Unmarshal(body, searchReq)
		if err != nil {
			errMsg := fmt.Sprintf("SearchHandler: Request Processing Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusBadRequest, errMsg)
			return
		}

		err = validateSearchReq(searchReq)
		if err != nil {
			errMsg := fmt.Sprintf("SearchHandler: Request Validation Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusBadRequest, errMsg)
			return
		}

		employees, err := Search(searchReq)
		if err != nil {
			errMsg := fmt.Sprintf("SearchHandler: Search Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusInternalServerError, errMsg)
			return
		}

		searchResp := empstoreapi.SearchResponse{Employees: employees}
		resp, err := json.Marshal(searchResp)
		if err != nil {
			errMsg := fmt.Sprintf("SearchHandler: Response Building Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusInternalServerError, errMsg)
			return
		}

		fmt.Printf("Search Complete.\n")
		writeResponse(w, resp)

	default:
		errMsg := fmt.Sprintf("SearchHandler: Invalide Request Method. %v\n", method)
		fmt.Printf(errMsg)
		writeErrorResponse(w, http.StatusBadRequest, errMsg)
		return
	}
}

//ListHandler ...
func ListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ListHandler: Received Request")
	switch method := r.Method; method {
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errMsg := fmt.Sprintf("ListHandler: Request Read Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusBadRequest, errMsg)
			return
		}
		listReq := &empstoreapi.ListRequest{}
		err = json.Unmarshal(body, listReq)
		if err != nil {
			errMsg := fmt.Sprintf("ListHandler: Request Processing Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusBadRequest, errMsg)
			return
		}

		employees, err := List(listReq)
		if err != nil {
			errMsg := fmt.Sprintf("ListHandler: List Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusInternalServerError, errMsg)
			return
		}

		listResp := empstoreapi.ListResponse{Employees: employees}
		resp, err := json.Marshal(listResp)
		if err != nil {
			errMsg := fmt.Sprintf("ListHandler: Response Building Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusInternalServerError, errMsg)
			return
		}

		fmt.Printf("List Complete.\n")
		writeResponse(w, resp)

	default:
		errMsg := fmt.Sprintf("ListHandler: Invalide Request Method. %v\n", method)
		fmt.Printf(errMsg)
		writeErrorResponse(w, http.StatusBadRequest, errMsg)
		return
	}
}

//UpdateHandler ...
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateHandler: Received Request")
	switch method := r.Method; method {
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errMsg := fmt.Sprintf("UpdateHandler: Request Read Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusBadRequest, errMsg)
			return
		}
		updateReq := &empstoreapi.UpdateRequest{}
		err = json.Unmarshal(body, updateReq)
		if err != nil {
			errMsg := fmt.Sprintf("UpdateHandler: Request Processing Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusBadRequest, errMsg)
			return
		}

		err = validateUpdateReq(updateReq)
		if err != nil {
			errMsg := fmt.Sprintf("UpdateHandler: Request Validation Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusBadRequest, errMsg)
			return
		}

		err = Update(updateReq)
		if err != nil {
			errMsg := fmt.Sprintf("UpdateHandler: Update Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusInternalServerError, errMsg)
			return
		}

		fmt.Printf("Update Complete.\n")
		writeResponse(w, []byte{})

	default:
		errMsg := fmt.Sprintf("UpdateHandler: Invalide Request Method. %v\n", method)
		fmt.Printf(errMsg)
		writeErrorResponse(w, http.StatusBadRequest, errMsg)
		return
	}
}

//DeleteHandler ...
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteHandler: Received Request")
	switch method := r.Method; method {
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errMsg := fmt.Sprintf("DeleteHandler: Request Read Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusBadRequest, errMsg)
			return
		}
		deleteReq := &empstoreapi.DeleteRequest{}
		err = json.Unmarshal(body, deleteReq)
		if err != nil {
			errMsg := fmt.Sprintf("DeleteHandler: Request Processing Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusBadRequest, errMsg)
			return
		}

		err = validateDeleteReq(deleteReq)
		if err != nil {
			errMsg := fmt.Sprintf("DeleteHandler: Request Validation Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusBadRequest, errMsg)
			return
		}

		err = Delete(deleteReq)
		if err != nil {
			errMsg := fmt.Sprintf("DeleteHandler: Update Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusInternalServerError, errMsg)
			return
		}

		fmt.Printf("Delete Complete.\n")
		writeResponse(w, []byte{})

	default:
		errMsg := fmt.Sprintf("DeleteHandler: Invalide Request Method. %v\n", method)
		fmt.Printf(errMsg)
		writeErrorResponse(w, http.StatusBadRequest, errMsg)
		return
	}
}

//RestoreHandler ...
func RestoreHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RestoreHandler: Received Request")
	switch method := r.Method; method {
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errMsg := fmt.Sprintf("RestoreHandler: Request Read Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusBadRequest, errMsg)
			return
		}
		restoreReq := &empstoreapi.RestoreRequest{}
		err = json.Unmarshal(body, restoreReq)
		if err != nil {
			errMsg := fmt.Sprintf("RestoreHandler: Request Processing Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusBadRequest, errMsg)
			return
		}

		err = validateRestoreReq(restoreReq)
		if err != nil {
			errMsg := fmt.Sprintf("RestoreHandler: Request Validation Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusBadRequest, errMsg)
			return
		}

		err = Restore(restoreReq)
		if err != nil {
			errMsg := fmt.Sprintf("RestoreHandler: Restore Failed. Err: %v\n", err)
			fmt.Printf(errMsg)
			writeErrorResponse(w, http.StatusInternalServerError, errMsg)
			return
		}

		fmt.Printf("Restore Complete.\n")
		writeResponse(w, []byte{})

	default:
		errMsg := fmt.Sprintf("RestoreHandler: Invalide Request Method. %v\n", method)
		fmt.Printf(errMsg)
		writeErrorResponse(w, http.StatusBadRequest, errMsg)
		return
	}
}
