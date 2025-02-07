package brackets

import (
	"github.com/AnnaVolchkova/tasks/linkedlist"
	
)

func IsValid(a string) bool { // a = "[]"
	list := linkedlist.New() // экземпляр связанного списка, с которым мы можем  работать через методы
	// list := {"["}

	if a == "" { // false
		return false 
	}
// "[]["

	for _, char := range a { // char == ']'
		charstr := string(char) // charstr == "]"

		switch charstr {
		case "(", "{", "[":
			list.Push(charstr) // list == "["
		case ")", "}", "]":
			value, exists := list.Pop() // value, exists == "[", true
			if !exists { // false
				return false
			}
			// if value != charstr { // "[" != "]" // true
			// return false
			// }
			if value == "[" && charstr == "]" || value == "(" && charstr == ")" || value == "{" && charstr == "}" {
				continue
			}
			
			return false
		}
	}

    return list.Length() == 0

	// if list.Length() != 0 { // синоним
	// 	return false
	//   }
	  
	//   return true
}

