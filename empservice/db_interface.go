package empservice

import (
	"empstore/db"
	"empstore/empstoreapi"
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	accessKey          = "dummy"
	secretKey          = "dummy"
	region             = "us-east-1"
	endpoint           = "http://127.0.0.1:8000"
	tableName          = "Employees"
	hKey               = "HKey"
	rKey               = "RKey"
	readCapacityUnits  = 100
	writeCapacityUnits = 100
)

const (
	hkeyValEmployee = "employee/"
	hkeyValCounter  = "count/"
)

const (
	statusActive   = "Active"
	statusInActive = "InActive"
)

//////////////// Fucntions which are directly called by Service///////////////////////

//AddEmployee ...
func AddEmployee(addReq *empstoreapi.AddRequest) (empID string, err error) {
	//Generate New EmployeeID.
	empID, err = getNewEmployeeID()
	if err != nil {
		fmt.Printf("AddEmployee: getNewEmployeeID Failed. Error %v\n", err)
		return empID, err
	}

	empRecord := make(map[string]*dynamodb.AttributeValue)
	empRecord[db.HKeyName] = db.StrToAttr(hkeyValEmployee)
	empRecord[db.RKeyName] = db.StrToAttr(empID)
	empRecord["Id"] = db.StrToAttr(empID)
	empRecord["Name"] = db.StrToAttr(addReq.Name)
	if addReq.Department != "" {
		empRecord["Department"] = db.StrToAttr(addReq.Department)
	}
	if addReq.Address != "" {
		empRecord["Address"] = db.StrToAttr(addReq.Address)
	}
	if len(addReq.Skills) != 0 {
		empRecord["Skills"] = db.StrSetToAttr(addReq.Skills)
	}
	empRecord["Status"] = db.StrToAttr("Active")

	//Store employee into DB
	err = db.Put(tableName, empRecord)
	if err != nil {
		fmt.Printf("AddEmployee: db.Put Failed. Err: %v\n", err)
		return empID, err
	}

	return empID, nil
}

//Search ...
func Search(searchReq *empstoreapi.SearchRequest) (employees []*empstoreapi.Employee, err error) {
	//Bring in all the employees and filter out the unmatched records and then return.
	//TODO: This can be improved to pass the filters to dynamoDB itself.
	searchTerm := searchReq.Term
	empRecords, err := db.Query(tableName, hkeyValEmployee, nil)
	if err != nil {
		fmt.Printf("Search: db.Query failed. Err: %v\n", err)
		return nil, err
	}
	for _, empRec := range empRecords {
		//FilterOut the InActive Employee Records
		if aws.StringValue(empRec["Status"].S) == statusInActive {
			continue
		}

		emp := &empstoreapi.Employee{
			ID:   aws.StringValue(empRec["Id"].S),
			Name: aws.StringValue(empRec["Name"].S),
		}
		if attrVal, ok := empRec["Address"]; ok {
			emp.Address = aws.StringValue(attrVal.S)
		}
		if attrVal, ok := empRec["Department"]; ok {
			emp.Department = aws.StringValue(attrVal.S)
		}
		if attrVal, ok := empRec["Skills"]; ok {
			emp.Skills = aws.StringValueSlice(attrVal.SS)
		}

		//Check if the search term matches any of these employee Attributes
		matched := false
		matched = matched || (emp.ID == searchTerm)
		matched = matched || (emp.Name == searchTerm)
		matched = matched || (emp.Address == searchTerm)
		if !matched {
			for _, skill := range emp.Skills {
				matched = (skill == searchTerm)
				if matched {
					break
				}
			}
		}

		if matched {
			employees = append(employees, emp)
		}
	}
	return employees, nil
}

//List ...
//This matches the records with all the parameters in the ListRequest.
//FitersOut if anything doesnt match. - Basically an AND operation.
func List(listReq *empstoreapi.ListRequest) (employees []*empstoreapi.Employee, err error) {
	//Bring in all the employees and filter out the unmatched records and then return.
	//TODO: This can be improved to pass the filters to dynamoDB itself.

	empRecords, err := db.Query(tableName, hkeyValEmployee, nil)
	if err != nil {
		fmt.Printf("List: db.Query failed. Err: %v\n", err)
		return nil, err
	}
	bDone := false
	for _, empRec := range empRecords {
		if bDone {
			//bDone is set on the basis of EmpID getting matched with list request criteria.
			break
		}
		//FilterOut the InActive Employee Records
		if aws.StringValue(empRec["Status"].S) == statusInActive {
			continue
		}

		emp := &empstoreapi.Employee{
			ID:   aws.StringValue(empRec["Id"].S),
			Name: aws.StringValue(empRec["Name"].S),
		}
		if attrVal, ok := empRec["Address"]; ok {
			emp.Address = aws.StringValue(attrVal.S)
		}
		if attrVal, ok := empRec["Department"]; ok {
			emp.Department = aws.StringValue(attrVal.S)
		}
		if attrVal, ok := empRec["Skills"]; ok {
			emp.Skills = aws.StringValueSlice(attrVal.SS)
		}

		if listReq.ID != "" {
			if emp.ID == listReq.ID {
				//Then necessarily this is the only emp which can be listed.
				bDone = true
			} else {
				continue
			}
		}
		if listReq.Name != "" && emp.Name != listReq.Name {
			continue
		}
		if listReq.Department != "" && emp.Department != listReq.Department {
			continue
		}

		//Now the emp Record has matched all the critera. Add it into Result set.
		employees = append(employees, emp)

		//If we are done with Scaning. Break
		if bDone {
			break
		}
	}

	return employees, nil
}

//Update ...
func Update(updateReq *empstoreapi.UpdateRequest) (err error) {
	keys := map[string]*dynamodb.AttributeValue{
		db.HKeyName: db.StrToAttr(hkeyValEmployee),
		db.RKeyName: db.StrToAttr(updateReq.ID),
	}

	//TODO: Check if this emp exist then proceed. This can be take cabe by lease also

	//TODO: Load Lease here.

	updateInfo := make(map[string]*dynamodb.AttributeValue)
	if updateReq.Department != "" {
		updateInfo["Department"] = db.StrToAttr(updateReq.Department)
	}
	if updateReq.Address != "" {
		updateInfo["Address"] = db.StrToAttr(updateReq.Address)
	}
	if len(updateReq.Skills) != 0 {
		updateInfo["Skills"] = db.StrSetToAttr(updateReq.Skills)
	}
	err = db.Update(tableName, keys, updateInfo)
	return err
}

//Delete ...
func Delete(delReq *empstoreapi.DeleteRequest) (err error) {
	keys := map[string]*dynamodb.AttributeValue{
		db.HKeyName: db.StrToAttr(hkeyValEmployee),
		db.RKeyName: db.StrToAttr(delReq.ID),
	}
	//TODO: Load Lease here and releaes it at the end

	//Delete Employee if PermanentlyDelete is set
	//Else Deavtivate.
	if delReq.PermanentlyDelete {
		err = db.Delete(tableName, keys, nil)
		return err
	}
	updateInfo := map[string]*dynamodb.AttributeValue{
		"Status": db.StrToAttr(statusInActive),
	}
	err = db.Update(tableName, keys, updateInfo)
	return err
}

//Restore ...
func Restore(restoreReq *empstoreapi.RestoreRequest) (err error) {
	keys := map[string]*dynamodb.AttributeValue{
		db.HKeyName: db.StrToAttr(hkeyValEmployee),
		db.RKeyName: db.StrToAttr(restoreReq.ID),
	}
	//TODO: Load Lease here and releaes it at the end
	updateInfo := map[string]*dynamodb.AttributeValue{
		"Status": db.StrToAttr(statusActive),
	}
	err = db.Update(tableName, keys, updateInfo)
	return err
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//GetNewEmployeeID : Creates a unique id using Storage Counter and returns
func getNewEmployeeID() (string, error) {
	empID := ""

	keys := map[string]*dynamodb.AttributeValue{
		db.HKeyName: db.StrToAttr(hkeyValCounter),
		db.RKeyName: db.StrToAttr("emp"),
	}

	newCount, err := db.Increment(tableName, keys, "Counter", 1)
	if err != nil {
		fmt.Printf("GetNewEmployeeID Failed: %v\n", err)
		return "", err
	}
	empID = "Emp_" + strconv.FormatInt(int64(newCount), 10)
	return empID, nil
}

func initEmployeeCounter() error {
	counterRecord := map[string]*dynamodb.AttributeValue{
		db.HKeyName: db.StrToAttr(hkeyValCounter),
		db.RKeyName: db.StrToAttr("emp"),
		"Counter":   db.NumToAttr(0),
	}

	err := db.Put(tableName, counterRecord)
	if err != nil {
		fmt.Printf("initEmployeeCounter: db.Put Failed. Err: %v\n", err)
		return err
	}
	return nil
}

func init() {
	db.InitDBAPI(region, endpoint, accessKey, secretKey)
	fmt.Println("Initialized DB Session ...")
	exist, err := db.DoesTableExit(tableName)
	if err != nil {
		fmt.Printf("db.DoesTableExit Failed %v\n. Exitting....", err)
		os.Exit(1)
	}
	if exist {
		return
	}
	err = db.CreateTable(tableName, readCapacityUnits, writeCapacityUnits)
	if err != nil {
		fmt.Printf("db.CreateTable Failed %v\n. Exitting....", err)
		os.Exit(1)
	}
	err = initEmployeeCounter()
	if err != nil {
		fmt.Printf("initEmployeeCounter Failed %v\n. Exitting....", err)
		os.Exit(1)
	}
}
