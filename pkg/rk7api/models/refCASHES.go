package models

import "encoding/xml"

// RK7QueryResult was generated 2024-02-02 07:40:04 by https://xml-to-go.github.io/ in Ukraine.
type RK7QueryResultGetRefDataCashes struct {
	XMLName         xml.Name `xml:"RK7QueryResult" json:"rk7queryresult,omitempty"`
	Text            string   `xml:",chardata" json:"text,omitempty"`
	ServerVersion   string   `xml:"ServerVersion,attr" json:"serverversion,omitempty"`
	XmlVersion      string   `xml:"XmlVersion,attr" json:"xmlversion,omitempty"`
	NetName         string   `xml:"NetName,attr" json:"netname,omitempty"`
	Status          string   `xml:"Status,attr" json:"status,omitempty"`
	CMD             string   `xml:"CMD,attr" json:"cmd,omitempty"`
	ErrorText       string   `xml:"ErrorText,attr" json:"errortext,omitempty"`
	DateTime        string   `xml:"DateTime,attr" json:"datetime,omitempty"`
	WorkTime        string   `xml:"WorkTime,attr" json:"worktime,omitempty"`
	Processed       string   `xml:"Processed,attr" json:"processed,omitempty"`
	ArrivalDateTime string   `xml:"ArrivalDateTime,attr" json:"arrivaldatetime,omitempty"`
	RK7Reference    struct {
		Text           string `xml:",chardata" json:"text,omitempty"`
		DataVersion    string `xml:"DataVersion,attr" json:"dataversion,omitempty"`
		ClassName      string `xml:"ClassName,attr" json:"classname,omitempty"`
		TotalItemCount string `xml:"TotalItemCount,attr" json:"totalitemcount,omitempty"`
		RIChildItems   string `xml:"RIChildItems"`
		Items          struct {
			Text string `xml:",chardata" json:"text,omitempty"`
			Item []struct {
				Text      string `xml:",chardata" json:"text,omitempty"`
				Ident     string `xml:"Ident,attr" json:"ident,omitempty"`
				ItemIdent string `xml:"ItemIdent,attr" json:"itemident,omitempty"`
				Name      string `xml:"Name,attr" json:"name,omitempty"`
				Code      string `xml:"Code,attr" json:"code,omitempty"`
				Status    string `xml:"Status,attr" json:"status,omitempty"`
				NetName   string `xml:"NetName,attr" json:"netname,omitempty"`
			} `xml:"Item" json:"item,omitempty"`
		} `xml:"Items" json:"items,omitempty"`
	} `xml:"RK7Reference" json:"rk7reference,omitempty"`
}
