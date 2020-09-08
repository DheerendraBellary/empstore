package empservice

import (
	"empstore/empstoreapi"
	"errors"
	"net/http"
)

/////////////////-------------Some Utility Functions-------------////////////

func writeResponse(w http.ResponseWriter, jsonResp []byte) {
	w.Header().Add("content-type", "application/json")
	w.Header().Add("charset", "utf-8")
	w.Write(jsonResp)
}

func writeErrorResponse(w http.ResponseWriter, httpStatus int, errMsg string) {
	w.Header().Add("errormsg", errMsg)
	w.WriteHeader(httpStatus)

}

//validateAddReq ...
func validateAddReq(req *empstoreapi.AddRequest) error {
	if req.Name == "" {
		return errors.New("ValidateAddReq: Name Cannot be Empty")
	}
	return nil
}

//validateSearchReq ...
func validateSearchReq(req *empstoreapi.SearchRequest) error {
	if req.Term == "" {
		return errors.New("ValidateSeachReq: Term Cannot be Empty")
	}
	return nil
}

//validateUpdateReq ...
func validateUpdateReq(req *empstoreapi.UpdateRequest) error {
	if req.ID == "" {
		return errors.New("validateUpdateReq: ID Cannot be Empty")
	}
	if req.Department == "" && req.Address == "" && len(req.Skills) == 0 {
		return errors.New("validateUpdateReq: Alteast need one parameter Department/Address/Skills")
	}
	return nil
}

//validateDeleteReq ...
func validateDeleteReq(req *empstoreapi.DeleteRequest) error {
	if req.ID == "" {
		return errors.New("validateDeleteReq: ID Cannot be Empty")
	}
	return nil
}

//validateRestoreReq ...
func validateRestoreReq(req *empstoreapi.RestoreRequest) error {
	if req.ID == "" {
		return errors.New("validateRestoreReq: ID Cannot be Empty")
	}
	return nil
}
