package util

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

func DropUnwantedField() {

}

func TransferData(o1, o2 interface{}) error {
	// if reflect.TypeOf(o1) == reflect.TypeOf(o2) {
	err := copier.CopyWithOption(o2, o1, copier.Option{DeepCopy: true, IgnoreEmpty: true})
	// }
	// fmt.Println(reflect.TypeOf(o1))
	// fmt.Println(reflect.TypeOf(o2))
	return err //fmt.Errorf("we have an error util/util.go %v", "types doesnt matches")
}
func ToContextService(c echo.Context) context.Context {
	surround := make(map[string]interface{}, 0)
	id := c.Get("id")
	clientType := c.Get("client")
	client := c.Get("pgClient")
	surround["id"] = id
	surround["userType"] = clientType
	surround["pgClient"] = client
	return context.WithValue(context.Background(), "surround", surround)
}

func GetFromServiceCtx(c context.Context, key string) interface{} {
	//check if key is what we hold for now
	return c.Value("surround").(map[string]interface{})[key]
}

// func StringToObjectID(ctx context.Context) (primitive.ObjectID, error) {
// 	id := GetFromServiceCtx(ctx, "id").(string)
// 	oid, err := primitive.ObjectIDFromHex(id)
// 	return oid, err
// }

func Serialize(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}
func StringToInt(obj string) (int, error) {
	return strconv.Atoi(obj)
}
func DeSerialize(btArray []byte, obj interface{}) error {
	return json.Unmarshal(btArray, obj)
}
