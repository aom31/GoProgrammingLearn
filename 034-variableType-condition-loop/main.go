package main
import(
	"fmt"
	"encoding/json"
)
func main(){

	//1. Literal types
	// literal type copy by value
	theString := "this is variable type string"
	theInt := 1234
	theBool := true
	fmt.Printf("show type and value :\n %T = %v , %T = %v , %T = %v \n",theString,theString,theInt,theInt,theBool,theBool)

	//2 Type in golang are static type , and error will show in step complie

	//3 interface{} type can accept/store all type
	printInterface(theString)
	printInterface(theBool)
	printInterface(theInt)

	//4 map type use map[type-key]type-value
	//theMap := map[string]interface{}{} //interface{} is type if interface{}{} is declare instance to interface
	
	theMap := make(map[string]interface{})
	theMap["first_name"] = "Thamakorn"
	theMap["nick_name"] = "aom"
	theMap["age"] = 23
	theMap["isMale"] = true

	//should convert map with json
	//for json.Marshal , Marshal will convert any object (like struct , slice ,string,array) to json string
	// but for json.Unmarshal will convert json string to any object

	//this is anonymous function
	showJson := func(input interface{}){
		responseJsonString, _ := json.Marshal(input)
		fmt.Println("map object to json string = ", string(responseJsonString)) //json marshal will retuen []byte convert it to string
	} 
	showJson(theMap)

	//when you want any value from map need to use 2 variable

	valueKey, okCheckKeyhaveWithbool := theMap["aom"]
	if okCheckKeyhaveWithbool {
		fmt.Println("have value with this key is",valueKey)
	}else if !okCheckKeyhaveWithbool { //check false
		fmt.Println("not have this key in map")
	}

	//5 Slice type (dynamic array)
	//Note: slices pass by pointer of slice (like map)

	//declare slice with []type{}
	//theSlices := []string{}
	theSlices := make([]string, 0)
	//add value to slice by append
	theSlices = append(theSlices, "add-item1", "add-item2")
	theSlices = append(theSlices,"add-item3")
	fmt.Println("show value in slices = ", theSlices)
	showJson(theSlices)

	//6 struct type use for reporesent object

	//create struct with instance
	// use pointer of struct for change value in address
	// So declare struct with instance value by pointer
	theUserInfo := &UserInterface{
		//initial value of struct field
		FirstName : "thamakorn",
		NickName : "aom",
		Age : 23,

	}
	showJson(theUserInfo)

	//7 Enum is just constant in golang
	// คือ การจำกัดค่าคงที่ที่เป็นไปได้ทั้งหมดของข้อมูลชุดนั้นเอาไว้ โดย
	// 7.1 declare enum type คือการประกาศtype enum ให้เป็นข้อมูลคงที่รูปแบบไหน 
	type GenderType string
	//7.2 declare const of enum type
	const (
		//varName enumType = value-constant
		Male GenderType = "MALE"
		Famale GenderType = "Famale"
		Another GenderType = "Another"
	)

	var theGender GenderType

	//theGender = Male
	//if want use swith state just use enum
	switch theGender {
	case Male:
		fmt.Println("gender type is male")
	case Famale:
		fmt.Println("gender type is famale")
		
	}

}

//declare Struct
type UserInterface struct {
	FirstName string `json:"first_name"`
	NickName string `json:"nick_name"`
	Age int `json:"age"`
}

//receive value with interface{} type ,So interface type can receive every type like string, bool,int
func printInterface(input interface{}){
	fmt.Printf("show interface value = %v \n",input)
	

}