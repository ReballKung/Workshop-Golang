package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Organizational struct {
	ORG_Name        string
	ORG_Address     string
	ORG_Tel         string
	ORG_TotalPeople string
}

func main() {
	fmt.Println("=====================================")
	fmt.Printf("Ex 0 : ")
	num0()
	fmt.Println("=====================================")
	fmt.Printf("Ex 1 : ")
	num1()
	fmt.Println("=====================================")
	fmt.Printf("Ex 1.2 : ")
	num1_2(20, 2)
	fmt.Println("=====================================")
	fmt.Printf("Ex 2 : ")
	num2()
	fmt.Println("=====================================")
	fmt.Printf("Ex 3 : ")
	num3()
	fmt.Println("=====================================")
	fmt.Printf("Ex 3.1 : ")
	num3_1(10000)
	fmt.Println("=====================================")
	fmt.Printf("Ex 4 : ")
	num4()
	fmt.Println("\n=====================================")
	fmt.Printf("Ex 4.1 : ")
	num4_1("ine t")
	fmt.Println("\n=====================================")
	fmt.Println("Ex 5 : ")
	num5()
	fmt.Println("\n=====================================")
	fmt.Printf("Ex 6 : ")
	num6()
	fmt.Println("\n=====================================")
	Paramid()

}

func num0() {
	i := 2

	if i == 0 {
		fmt.Println("ZERO")
	} else if i == 1 {
		fmt.Println("ONE")
	} else if i == 2 {
		fmt.Println("Two")
	} else if i == 3 {
		fmt.Println("Three")
	} else {
		fmt.Println("Your i not in case.")
	}
}

func num1() {
	var j int

	fmt.Printf("\n[")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			j = j + 1
			fmt.Printf(" %d ", i)
		}
	}
	fmt.Printf("]")
	total := j
	fmt.Println()
	fmt.Println("โดยมีเลขที่หาร 3 ลงตัวทั้งหมดจำนวน : ", total)
}

func num1_2(number1, number2 int) {
	result := math.Pow(float64(number1), float64(number2))
	fmt.Println(result)
}

func num2() {
	x := []int{
		48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17,
	}
	calMaxMIN(x)
}

func calMaxMIN(values []int) {
	max := values[0]
	min := values[0]

	for _, v := range values {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	fmt.Println("\nMax : ", max)
	fmt.Println("Min : ", min)
}

func num3() {
	var total int
	var c = 0
	for i := 0; i < 1000; i++ {
		total = i + 1
		t := strconv.Itoa(total) //แปลงจาก int เป็น string
		for _, v := range t {
			if string(v) == "9" {
				c++
			}
		}
	}
	fmt.Println(c)
}

func num3_1(values int) {
	var total int
	var c = 0
	for i := 0; i < values; i++ {
		total = i + 1
		t := strconv.Itoa(total) //แปลงจาก int เป็น string
		for _, v := range t {
			if string(v) == "9" {
				c++
			}
		}
	}
	fmt.Println(c)
}

func num4() {
	var myWords = "AW SOME GO!"
	for _, v := range myWords {
		if string(v) == " " {
			strings.TrimSpace(string(v))
		} else {
			fmt.Printf("%s", string(v))
		}
	}
}

func num4_1(values string) {
	for _, v := range values {
		if string(v) == " " {
			strings.TrimSpace(string(v))
		} else {
			fmt.Printf("%s", string(v))
		}
	}
}

func num5() {
	employees := map[string]map[string]string{
		"emp_01": {
			//* KEY       VALUES
			"name":    "Peter Parker",
			"age":     "20",
			"address": "USA",
		},
		"emp_02": {
			"name":    "Lady gaga",
			"age":     "32 ",
			"address": "UK",
		},
		"emp_03": {
			"name":    "John Wick",
			"age":     "55",
			"address": "LA",
		},
		"emp_04": {
			"name":    "Alan Smith",
			"age":     "45",
			"address": "TH",
		},
	}
	for key, value := range employees {
		fmt.Println("----------------")
		fmt.Println("Employees ID : ", key)
		fmt.Printf("Name -:  %s (Age : %s)\n", value["name"], value["age"])
		fmt.Printf("Address -:  %s \n", value["address"])
	}
}

func num6() {
	org1 := new(Organizational)
	//=============================
	org1.ORG_Name = "ITVERTEX"
	org1.ORG_Address = "26/6 ถ.ราษฎร์ยินดี (ข้างศูนย์กิฟฟารีน) ต.หาดใหญ่ อ. หาดใหญ่ จ.สงขลา 90110"
	org1.ORG_Tel = "099-999-9999"
	org1.ORG_TotalPeople = "20 คน"
	//=============================
	fmt.Print(org1)
	//=============================
}

func Paramid() {
	for i := 0; i < 6; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
