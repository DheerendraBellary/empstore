package empservice

import (
	"empstore/db"
	"empstore/empstoreapi"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var testEmpDataSet []*empstoreapi.Employee

func TestGetNewEmployeeID(t *testing.T) {
	keys := map[string]*dynamodb.AttributeValue{
		db.HKeyName: db.StrToAttr("count/"),
		db.RKeyName: db.StrToAttr("emp"),
	}
	res, err := db.Get(tableName, keys)
	if err != nil {
		t.Fatalf("TestGetNewEmployeeID Get Failed. Err: %v", err)
		return
	}
	currentCount, err := db.AttrToNum(res["Counter"])
	if err != nil {
		t.Fatalf("TestGetNewEmployeeID Parse Counter Failed. Err: %v", err)
		return
	}
	empID, err := getNewEmployeeID()
	if err != nil {
		t.Fatalf("TestGetNewEmployeeID Failed. Err: %v", err)
		return
	}
	expectedEmpID := "Emp_" + strconv.FormatInt(int64(currentCount+1), 10)
	if empID != expectedEmpID {
		t.Fatalf("TestGetNewEmployeeID Failed. Expected EmpID: %v, Acutal: %v", expectedEmpID, empID)
		return
	}
}

func TestAddEmployee(t *testing.T) {
	emp := &empstoreapi.Employee{
		Name:       "Abc",
		Department: "Development",
		Address:    "HNo 1, pqr Apts, pqr Road",
		Skills:     []string{"go", "c++", "aws"},
	}
	empID, err := AddEmployee(emp)
	if err != nil {
		t.Fatalf("TestGetNewEmployeeID Failed. Err: %v", err)
		return
	}
	fmt.Printf("Employee Added. ID: %v\n", empID)
	fmt.Println("TestAddEmployee Done")
}

func TestSearch(t *testing.T) {
	for _, emp := range testEmpDataSet {
		AddEmployee(emp)
	}
	searchReq := &empstoreapi.SearchRequest{
		Term: "python",
	}
	res, err := Search(searchReq)
	if err != nil {
		t.Fatalf("TestSearch: Search Failed. Error: %v\n", err)
		return
	}
	if len(res) != 2 {
		t.Fatalf("TestSearch: Expected 2. Actual: %v\n", len(res))
		return
	}

	for _, resEmp := range res {
		fmt.Printf("%v\n", resEmp)
	}

}

func setup() {
	testEmpDataSet = []*empstoreapi.Employee{
		&empstoreapi.Employee{
			Name:       "Abc",
			Department: "Development",
			Address:    "HNo 1, pqr Apts, pqr Road",
			Skills:     []string{"go", "c++", "python", "aws"},
		},
		&empstoreapi.Employee{
			Name:       "Def",
			Department: "Testing",
			Address:    "HNo 2, tuv Apts, tuv Road",
			Skills:     []string{"python", "Selenium"},
		},
		&empstoreapi.Employee{
			Name:       "Lmn",
			Department: "Performance",
			Address:    "HNo 3, xyz Apts, xyz Road",
			Skills:     []string{"perftesting"},
		},
	}
}
func shutdown() {

}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
