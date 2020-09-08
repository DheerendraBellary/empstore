/*
 * package empstore defines the json apis for http empstore service.
 */

package empstoreapi

//Employee ...
type Employee struct {
	ID         string   `json:"id,omitempty"`
	Name       string   `json:"name"`
	Department string   `json:"department,omitempty"`
	Address    string   `json:"address,omitempty"`
	Skills     []string `json:"skills,omitempty"`
}

//AddRequest ...
type AddRequest struct {
	Name       string   `json:"name"`
	Department string   `json:"department,omitempty"`
	Address    string   `json:"address,omitempty"`
	Skills     []string `json:"skills,omitempty"`
}

//AddResponse ...
type AddResponse struct {
	ID string `json:"id,omitempty"`
}

//SearchRequest ...
type SearchRequest struct {
	Term string `json:"term"`
}

//SearchResponse ...
type SearchResponse struct {
	Employees []*Employee `json:"employees,omitempty"`
}

//ListRequest ...
type ListRequest struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name"`
	Department string `json:"department,omitempty"`
}

//ListResponse ...
type ListResponse struct {
	Employees []*Employee `json:"employees,omitempty"`
}

//UpdateRequest ...
type UpdateRequest struct {
	ID         string   `json:"id"`
	Department string   `json:"department,omitempty"`
	Address    string   `json:"address,omitempty"`
	Skills     []string `json:"skills,omitempty"`
}

//DeleteRequest ...
type DeleteRequest struct {
	ID                string `json:"id"`
	PermanentlyDelete bool   `json:"permanentlyDelete,omitempty"`
}

//RestoreRequest ...
type RestoreRequest struct {
	ID string `json:"id"`
}
