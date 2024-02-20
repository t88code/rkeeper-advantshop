package models

import "encoding/xml"

// RK7QueryResultGetRefDataRestaurants was generated 2024-02-12 10:43:57 by https://xml-to-go.github.io/ in Ukraine.
type RK7QueryResultGetRefDataRestaurants struct {
	XMLName         xml.Name `xml:"RK7QueryResult,omitempty"`
	ServerVersion   string   `xml:"ServerVersion,attr,omitempty"`
	XmlVersion      string   `xml:"XmlVersion,attr,omitempty"`
	NetName         string   `xml:"NetName,attr,omitempty"`
	Status          string   `xml:"Status,attr,omitempty"`
	CMD             string   `xml:"CMD,attr,omitempty"`
	ErrorText       string   `xml:"ErrorText,attr,omitempty"`
	DateTime        string   `xml:"DateTime,attr,omitempty"`
	WorkTime        string   `xml:"WorkTime,attr,omitempty"`
	Processed       string   `xml:"Processed,attr,omitempty"`
	ArrivalDateTime string   `xml:"ArrivalDateTime,attr,omitempty"`
	RK7Reference    struct {
		DataVersion         string `xml:"DataVersion,attr,omitempty"`
		ClassName           string `xml:"ClassName,attr,omitempty"`
		Name                string `xml:"Name,attr,omitempty"`
		MinIdent            string `xml:"MinIdent,attr,omitempty"`
		MaxIdent            string `xml:"MaxIdent,attr,omitempty"`
		ViewRight           string `xml:"ViewRight,attr,omitempty"`
		UpdateRight         string `xml:"UpdateRight,attr,omitempty"`
		ChildRight          string `xml:"ChildRight,attr,omitempty"`
		DeleteRight         string `xml:"DeleteRight,attr,omitempty"`
		XMLExport           string `xml:"XMLExport,attr,omitempty"`
		XMLMask             string `xml:"XMLMask,attr,omitempty"`
		LeafCollectionCount string `xml:"LeafCollectionCount,attr,omitempty"`
		TotalItemCount      string `xml:"TotalItemCount,attr,omitempty"`
		Items               struct {
			Item []struct {
				Ident                 string `xml:"Ident,attr,omitempty"`
				ItemIdent             string `xml:"ItemIdent,attr,omitempty"`
				SourceIdent           string `xml:"SourceIdent,attr,omitempty"`
				GUIDString            string `xml:"GUIDString,attr,omitempty"`
				AssignChildsOnServer  string `xml:"AssignChildsOnServer,attr,omitempty"`
				ActiveHierarchy       string `xml:"ActiveHierarchy,attr,omitempty"`
				Name                  string `xml:"Name,attr,omitempty"`
				AltName               string `xml:"AltName,attr,omitempty"`
				Code                  string `xml:"Code,attr,omitempty"`
				MainParentIdent       string `xml:"MainParentIdent,attr,omitempty"`
				Status                string `xml:"Status,attr,omitempty"`
				ExtCode               string `xml:"ExtCode,attr,omitempty"`
				EditRight             string `xml:"EditRight,attr,omitempty"`
				Owner                 string `xml:"Owner,attr,omitempty"`
				Concept               string `xml:"Concept,attr,omitempty"`
				Region                string `xml:"Region,attr,omitempty"`
				BasicCurrency         string `xml:"BasicCurrency,attr,omitempty"`
				NationalCurrency      string `xml:"NationalCurrency,attr,omitempty"`
				RightLvl              string `xml:"RightLvl,attr,omitempty"`
				Address               string `xml:"Address,attr,omitempty"`
				SH4TradeGroup         string `xml:"SH4TradeGroup,attr,omitempty"`
				SH4Price              string `xml:"SH4Price,attr,omitempty"`
				StoreHouseCode        string `xml:"StoreHouseCode,attr,omitempty"`
				OpeningDate           string `xml:"OpeningDate,attr,omitempty"`
				Franchise             string `xml:"Franchise,attr,omitempty"`
				LocationLat           string `xml:"LocationLat,attr,omitempty"`
				LocationLong          string `xml:"LocationLong,attr,omitempty"`
				AddressGUID           string `xml:"AddressGUID,attr,omitempty"`
				OperationalHours      string `xml:"OperationalHours,attr,omitempty"`
				FullRestaurantCode    string `xml:"FullRestaurantCode,attr,omitempty"`
				FinancialAccStartDate string `xml:"FinancialAccStartDate,attr,omitempty"`
				PrinterAssigns        struct {
					ClassName string `xml:"ClassName,attr,omitempty"`
					Items     string `xml:"Items,omitempty"`
				} `xml:"PrinterAssigns,omitempty"`
				DeviceLicenses struct {
					ClassName string `xml:"ClassName,attr,omitempty"`
					Items     struct {
						Item struct {
							Ident      string `xml:"Ident,attr,omitempty"`
							LicenseTxt string `xml:"LicenseTxt,attr,omitempty"`
							VerifyNeed string `xml:"VerifyNeed,attr,omitempty"`
							DLLLicKind string `xml:"DLLLicKind,attr,omitempty"`
							ExpiresAT  string `xml:"ExpiresAT,attr,omitempty"`
						} `xml:"Item,omitempty"`
					} `xml:"Items,omitempty"`
				} `xml:"DeviceLicenses,omitempty"`
				Childs struct {
					ClassName string `xml:"ClassName,attr,omitempty"`
					Child     []struct {
						ChildIdent string `xml:"ChildIdent,attr,omitempty"`
						IsTerminal string `xml:"IsTerminal,attr,omitempty"`
					} `xml:"Child,omitempty"`
				} `xml:"Childs,omitempty"`
			} `xml:"Item,omitempty"`
		} `xml:"Items,omitempty"`
	} `xml:"RK7Reference,omitempty"`
}
