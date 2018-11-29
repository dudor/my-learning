package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var result RootElement
	result.ResourceString = append(result.ResourceString, ChildElement{
		StringName: "user1",
		InnerText:  "description",
	})
	fmt.Println(result)
	b, _ := xml.Marshal(result)
	b = append([]byte(xml.Header), b...)
	ioutil.WriteFile("./users.xml", b, os.ModePerm)

	/*
		gxml, err := ioutil.ReadFile("./users.xml")
		if err != nil {
			panic(err)
		}
		var result RootElement
		fmt.Println("序列化之前的result=", result)
		err = xml.Unmarshal(gxml, &result)
		if err != nil {
			panic(err)
		}
		fmt.Println("序列化之后的result=", result)
	*/
}

type RootElement struct {
	XMLName        xml.Name       `xml:"UserInformation"`
	ResourceString []ChildElement `xml:"string"`
}
type ChildElement struct {
	XMLName    xml.Name `xml:"string"`
	StringName string   `xml:"age,attr"`
	InnerText  string   `xml:",innerxml"`
}
