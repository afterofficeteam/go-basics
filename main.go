package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"

	"github.com/afteroffice/go-basics/model"
)

func main() {
	fmt.Println("start main")
	// basics()

	// fmt.Println("sum3Number(1, 2, 3) =", sum3Number(1, 2, 3))
	// fmt.Println("sum3Number(3, 5, 7) =", sum3Number(3, 5, 7))
	// fmt.Println("sum3Number(-10, 6, 4) =", sum3Number(-10, 6, 4))
	// fmt.Println("=====")

	// fmt.Println("mean3Number(1, 2, 3) =", mean3Number(1, 2, 3))       // 2.0
	// fmt.Println("mean3Number(4, 4, 5) =", mean3Number(4, 4, 5))       // 4.333
	// fmt.Println("mean3Number(10, 20, 50) =", mean3Number(10, 20, 50)) // 26.667
	// fmt.Println("=====")

	// fmt.Println("mean(1, 2, 3) =", mean([]int{1, 2, 3}))                                    // 2.0
	// fmt.Println("mean(10, 20, 30, 40, 50) =", mean([]int{10, 20, 30, 40, 50}))              // 30.0
	// fmt.Println("mean(1,2,3,4,5,6,7,8,9,10) =", mean([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})) // 4.5
	// fmt.Println("=====")

	// fmt.Println(`isPalindrome("katak") =`, isPalindrome("katak"))
	// fmt.Println(`isPalindrome("golang") =`, isPalindrome("golang"))
	// fmt.Println(`isPalindrome("1234567890987654321") =`, isPalindrome("1234567890987654321"))
	// fmt.Println(`isPalindrome("1234567890887654321") =`, isPalindrome("1234567890887654321"))
	// fmt.Println("=====")

	// fmt.Println("findDuplicateNumber(1,2,3,3,4,5) =", findDuplicateNumber([]int{1, 2, 3, 3, 4, 5}))                      // 3
	// fmt.Println("findDuplicateNumber(1,2,3,4,5,6,7,8,9,10) =", findDuplicateNumber([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 5})) // 5
	// fmt.Println("=====")

	// fmt.Println("printTypeAndValue(12.34) =", printTypeAndValue(12))    // int: 12
	// fmt.Println("printTypeAndValue(12.34) =", printTypeAndValue(12.34)) // float64: 12.34
	// fmt.Println(`printTypeAndValue("tes") =`, printTypeAndValue("tes")) // string: tes
	// fmt.Println("printTypeAndValue(true) =", printTypeAndValue(true))   // bool: true
	// fmt.Println("=====")

	// students := []model.Student{
	// 	{Name: "Erwin", Score: 90},
	// 	{Name: "a", Score: 75},
	// 	{Name: "Jody", Score: 90},
	// 	{Name: "b", Score: 70},
	// 	{Name: "Irham", Score: 80},
	// 	{Name: "c", Score: 65},
	// 	{Name: "d", Score: 60},
	// 	{Name: "Fatan", Score: 80},
	// }
	// fmt.Println("findStudents() =", findStudents(students, 80, false)) // 4 nama tidak terurut
	// fmt.Println("findStudents() =", findStudents(students, 80, true))  // 4 nama terurut
	// fmt.Println("findStudents() =", findStudents(students, 70, false)) // 6 nama tidak terurut
	// fmt.Println("findStudents() =", findStudents(students, 70, true))  // 6 nama terurut
	// fmt.Println("=====")

	// st := students[0]
	// changeName(&st, "Ethan")
	// fmt.Println("st.Name =", st.Name)
}

func basics() {
	title := "golang basics"
	fmt.Println("Hello, welcome to", title)

	// basic data types
	angka := 987                   // int
	angkaBerkoma := 12.34567       // float64
	angka32 := int32(angkaBerkoma) // int32 value = 12
	fmt.Printf("%d %f %d\n", angka, angkaBerkoma, angka32)
	fmt.Printf("%.2f %.3f %.4f %.10f\n", angkaBerkoma, angkaBerkoma, angkaBerkoma, math.Pow(angkaBerkoma, 2))

	var semangatBelajar bool // bool = false (zero-value)
	semangatBelajar = true
	// semangatBelajar := true

	const PI = 3.14
	const ROLE_OWNER = "owner"
	const (
		STATUS_200 = "OK"
		STATUS_400 = "Bad Request"
	)

	fmt.Printf("%s %t\n", title, semangatBelajar)
	fmt.Printf("%f %s\n", PI, STATUS_200)
	fmt.Printf("%v %T\n", title, title) // print value, type
	fmt.Printf("%v\n", []byte(title))   // print []byte

	fmt.Println("=====")

	// convert int to string
	angkaString := strconv.Itoa(angka)
	// angkaString := fmt.Sprintf("%d", angka)
	fmt.Printf("%s\n", angkaString)

	// convert string to int
	angka, err := strconv.Atoi(angkaString)
	if err != nil {
		panic(err)
	}

	// convert string to float64
	angkaBerkoma, err = strconv.ParseFloat(angkaString, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d %f\n", angka, angkaBerkoma)
	fmt.Println("=====")
}

// sum3Number adalah jumlah dari 3 angka (basic function)
func sum3Number(a int, b int, c int) int {
	return a + b + c
}

// mean3Number adalah nilai rata-rata dari 3 angka (parameter & return berbeda)
func mean3Number(a, b, c int) float64 {
	return float64(a+b+c) / 3
}

// Cari nilai rata-rata data arr [1,2,3,4,5] = 3 (for-range)
func mean(arr []int) float64 {
	// write code here
	return 0
}

// isPalindrome check wether str is palindrome or not. "katak" = true. (for-i, if)
func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i

		if str[i] != str[j] {
			return false
		}
	}
	return true
}

// which number is duplicate? [1,2,3,4,2,5] = 2 (map, for-range-map)
func findDuplicateNumber(arr []int) int {
	tempMap := map[int]bool{}
	for _, value := range arr {
		if _, found := tempMap[value]; found {
			return value
		}
	}
	return -1 // ga ketemu duplikat
}

// print Type and Value of data (learn interface{} as generic)
// contoh: "string: Hello" | "int: 123"
func printTypeAndValue(data interface{}) string {
	// Solution 1: use "type assertion"
	if value, ok := data.(int); ok {
		return fmt.Sprintf("int: %d", value)
	}
	if value, ok := data.(float64); ok {
		return fmt.Sprintf("float64: %f", value)
	}
	if value, ok := data.(string); ok {
		return fmt.Sprintf("string: %s", value)
	}
	if value, ok := data.(bool); ok {
		return fmt.Sprintf("bool: %t", value)
	}

	// Solution 2: use "reflect" package
	typeOf := reflect.TypeOf(data)
	switch typeOf.String() {
	case "int":
		return fmt.Sprintf("int: %d", data.(int))
	case "float64":
		return fmt.Sprintf("float64: %f", data.(float64))
	case "string":
		return fmt.Sprintf("string: %s", data.(string))
	case "bool":
		return fmt.Sprintf("bool: %t", data.(bool))
	}

	// Solution 3: utilize fmt.Sprintf's "formatting" capabilities, then return string
	return fmt.Sprintf("%T: %v", data, data)
}

// Sort & filter Students. exam score >= minScore, sort by name (using struct)
func findStudents(students []model.Student, minScore int, isSortByName bool) []model.Student {
	temp := []model.Student{}
	for _, st := range students {
		if st.Score >= minScore {
			temp = append(temp, st)
		}
	}

	if isSortByName {
		sort.Slice(temp, func(i, j int) bool {
			return temp[i].Name < temp[j].Name
		})
	}

	return temp
}

// mutate Student data (pointer)
func changeName(student *model.Student, name string) {
	student.Name = name
}
