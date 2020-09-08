package lease

import (
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//Lease ...
type Lease struct {
	timeStamp int64
}

//Load ...
func Load(key map[string]*dynamodb.AttributeValue) (ls *Lease, err error) {
	curTime := time.Now().Unix()

	return ls, nil
}
