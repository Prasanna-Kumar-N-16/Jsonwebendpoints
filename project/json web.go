package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Comp struct {
	Cname      string      `json:"cname"`
	Ckri       interface{} `json:"ckra-kri"`
	Department department  `json:"department"`
}
type department struct {
	Dname      int         `json:"did"`
	Dkri       interface{} `json:"dkra-kri"`
	Individual Individual  `json:"Individual"`
}
type Individual struct {
	Iname int         `json:"iname"`
	Ikri  interface{} `json:"kra-kri"`
}

func Company(w http.ResponseWriter, r *http.Request) {
	plan, _ := ioutil.ReadFile("c:/Users/radha/Desktop/gop.json")
	var data map[string]interface{}
	err := json.Unmarshal(plan, &data)
	if err != nil {
		panic(err)
	}
	//s := make([]string, 1)
	//s=append(s,"Param")
	kr := []string{"KRA 1", "KRA 2", "KRA 3"}
	ki := []int{11, 22, 33}
	m := make(map[string]int)
	for i := 0; i < len(kr); i++ {
		m[kr[i]] = ki[i]
	}
	names := "Param"
	/*nam := Comp{Cname: names, Ckri: m}
	jsonString, _ := json.Marshal(nam)
	ioutil.WriteFile("c:/Users/radha/Desktop/gop.json", jsonString, os.ModePerm)
	json.NewEncoder(w).Encode(nam)*/
	if data["cname"] == names {
		fmt.Fprintf(w, "Company exists")
	} else {
		nam := Comp{Cname: names, Ckri: m}
		jsonString, _ := json.Marshal(nam)
		ioutil.WriteFile("c:/Users/radha/Desktop/gop.json", jsonString, os.ModePerm)
		json.NewEncoder(w).Encode(nam)
	}
}
func Department(w http.ResponseWriter, r *http.Request) {
	plan, _ := ioutil.ReadFile("c:/Users/radha/Desktop/gop.json")
	var data map[string]interface{}
	err := json.Unmarshal(plan, &data)
	if err != nil {
		panic(err)
	}
	tempcname := "Param"
	did := 200
	kr := []string{"DKRA 1", "DKRA 2", "DKRA 3"}
	ki := []int{111, 222, 333}
	m := make(map[string]int)
	for i := 0; i < len(kr); i++ {
		m[kr[i]] = ki[i]
	}
	fmt.Println(data["dname"])
	//n,_:=data["department"]
	if data["cname"] == tempcname {
		if data["dname"] == nil {
			cd := department{Dname: did, Dkri: m}
			tempcomp := Comp{Cname: tempcname, Ckri: data["ckra-kri"], Department: cd}
			jsonString, _ := json.Marshal(tempcomp)
			ioutil.WriteFile("c:/Users/radha/Desktop/gop.json", jsonString, os.ModePerm)
			json.NewEncoder(w).Encode(tempcomp)
		} else if data["dname"] == did {
			fmt.Fprintf(w, "Department alreadyś exist")
		}
	} else {
		fmt.Fprintf(w, "Department doesn't exist as company doesn't exists")
	}
}
func Individ(w http.ResponseWriter, r *http.Request) {
	plan, _ := ioutil.ReadFile("c:/Users/radha/Desktop/gop.json")
	var data map[string]interface{}
	err := json.Unmarshal(plan, &data)
	if err != nil {
		panic(err)
	}
	tname := "Param"
	deptid := 200
	empname := 1234
	kr := []string{"IKRA 1", "IKRA 2", "IKRA 3"}
	ki := []int{1111, 2222, 3333}
	n := make(map[string]int)
	for i := 0; i < len(kr); i++ {
		n[kr[i]] = ki[i]
	}
	if data["cname"] == tname {
		//fmt.Println(data["cname"])
		if data["did"] == deptid {
			id := Individual{Iname: empname, Ikri: n}
			cd := department{Dname: deptid, Dkri: data["dkra-kri"], Individual: id}
			tempcomp := Comp{Cname: tname, Ckri: data["ckra-kri"], Department: cd}
			jsonString, _ := json.Marshal(tempcomp)
			ioutil.WriteFile("c:/Users/radha/Desktop/gop.json", jsonString, os.ModePerm)
			json.NewEncoder(w).Encode(tempcomp)
		}
	}
}

func main() {
	http.HandleFunc("/cmp", Company)
	http.HandleFunc("/deptid", Department)
	http.HandleFunc("/indiv", Individ)
	http.ListenAndServe(":10000", nil)
}
