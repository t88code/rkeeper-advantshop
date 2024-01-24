package models

import "encoding/xml"

type Transaction struct {
	XMLName           xml.Name `xml:"CHECK"`
	Text              string   `xml:",chardata"`
	Stationcode       string   `xml:"stationcode,attr"`
	Restaurantcode    string   `xml:"restaurantcode,attr"`
	Cashservername    string   `xml:"cashservername,attr"`
	Generateddatetime string   `xml:"generateddatetime,attr"`
	Chmode            string   `xml:"chmode,attr"`
	Locale            string   `xml:"locale,attr"`
	Shiftdate         string   `xml:"shiftdate,attr"`
	Shiftnum          string   `xml:"shiftnum,attr"`
	EXTINFO           struct {
		Reservation string `xml:"reservation,attr"`
		INTERFACES  struct {
			Current   string `xml:"current,attr"`
			INTERFACE struct {
				Type      string `xml:"type,attr"`
				ID        string `xml:"id,attr"`
				Mode      string `xml:"mode,attr"`
				Interface string `xml:"interface,attr"`
				HOLDERS   struct {
					ITEM []struct {
						Cardcode string `xml:"cardcode,attr"`
					} `xml:"ITEM"`
				} `xml:"HOLDERS"`
				ALLCARDS struct {
					ITEM []struct {
						Cardcode string `xml:"cardcode,attr"`
					} `xml:"ITEM"`
				} `xml:"ALLCARDS"`
			} `xml:"INTERFACE"`
		} `xml:"INTERFACES"`
	} `xml:"EXTINFO"`
	CHECKDATA struct {
		Checknum          string `xml:"checknum,attr"`
		Printnum          string `xml:"printnum,attr"`
		Fiscdocnum        string `xml:"fiscdocnum,attr"`
		Delprintnum       string `xml:"delprintnum,attr"`
		Delfiscdocnum     string `xml:"delfiscdocnum,attr"`
		Extfiscid         string `xml:"extfiscid,attr"`
		Tablename         string `xml:"tablename,attr"`
		Startservice      string `xml:"startservice,attr"` //TODO добавить обработку времени
		Closedatetime     string `xml:"closedatetime,attr"`
		Ordernum          string `xml:"ordernum,attr"`
		Guests            string `xml:"guests,attr"`
		Orderguid         string `xml:"orderguid,attr"`
		Checkguid         string `xml:"checkguid,attr"`
		OrderCat          string `xml:"order_cat,attr"`
		OrderType         string `xml:"order_type,attr"`
		Persistentcomment string `xml:"persistentcomment,attr"`
		CHECKPERSONS      struct {
			Count  string `xml:"count,attr"`
			PERSON []struct {
				ID   string `xml:"id,attr"`
				Name string `xml:"name,attr"`
				Code string `xml:"code,attr"`
				Role string `xml:"role,attr"`
			} `xml:"PERSON"`
		} `xml:"CHECKPERSONS"`
		CHECKLINES struct {
			Count string `xml:"count,attr"`
			LINE  []struct {
				ID          string `xml:"id,attr"`
				Code        string `xml:"code,attr"`
				Name        string `xml:"name,attr"`
				Uni         string `xml:"uni,attr"`
				Type        string `xml:"type,attr"`
				Price       int    `xml:"price,attr"`
				PrListSum   string `xml:"pr_list_sum,attr"`
				CategID     string `xml:"categ_id,attr"`
				ServprintID string `xml:"servprint_id,attr"`
				Quantity    int    `xml:"quantity,attr"`
				Sum         string `xml:"sum,attr"`
				LINETAXES   struct {
					Count string `xml:"count,attr"`
					TAX   []struct {
						ID  string `xml:"id,attr"`
						Sum string `xml:"sum,attr"`
					} `xml:"TAX"`
				} `xml:"LINETAXES"`
				DISCOUNTS struct {
					DISCOUNTPART []struct {
						ID          string `xml:"id,attr"`
						Disclineuni string `xml:"disclineuni,attr"`
						Sum         string `xml:"sum,attr"`
					} `xml:"DISCOUNTPART"`
				} `xml:"DISCOUNTS"`
				LINEPAYMENTS struct {
					LINEPAYMENT []struct {
						ID  string `xml:"id,attr"`
						Sum string `xml:"sum,attr"`
					} `xml:"LINEPAYMENT"`
				} `xml:"LINEPAYMENTS"`
			} `xml:"LINE"`
		} `xml:"CHECKLINES"`
		CHECKCATEGS struct {
			Count string `xml:"count,attr"`
			CATEG []struct {
				ID      string `xml:"id,attr"`
				Code    string `xml:"code,attr"`
				Name    string `xml:"name,attr"`
				Sum     string `xml:"sum,attr"`
				Discsum string `xml:"discsum,attr"`
			} `xml:"CATEG"`
		} `xml:"CHECKCATEGS"`
		CHECKDISCOUNTS struct {
			Count    string `xml:"count,attr"`
			DISCOUNT []struct {
				ID        string `xml:"id,attr"`
				Code      string `xml:"code,attr"`
				Name      string `xml:"name,attr"`
				Uni       string `xml:"uni,attr"`
				Interface string `xml:"interface,attr"`
				Cardcode  string `xml:"cardcode,attr"`
				Account   string `xml:"account,attr"`
				Sum       int    `xml:"sum,attr"`
			} `xml:"DISCOUNT"`
		} `xml:"CHECKDISCOUNTS"`
		CHECKPAYMENTS struct {
			Count   string `xml:"count,attr"`
			PAYMENT []struct {
				ID                 string `xml:"id,attr"`
				Code               string `xml:"code,attr"`
				Name               string `xml:"name,attr"`
				Uni                string `xml:"uni,attr"`
				Paytype            string `xml:"paytype,attr"`
				Interface          string `xml:"interface,attr"`
				Cardcode           string `xml:"cardcode,attr"`
				Account            string `xml:"account,attr"`
				Ownerinfo          string `xml:"ownerinfo,attr"`
				Exttransactioninfo string `xml:"exttransactioninfo,attr"`
				Bsum               string `xml:"bsum,attr"`
				Sum                int    `xml:"sum,attr"`
			} `xml:"PAYMENT"`
		} `xml:"CHECKPAYMENTS"`
		CHECKTAXES []struct {
			Count string `xml:"count,attr"`
			TAX   []struct {
				ID   string `xml:"id,attr"`
				Code string `xml:"code,attr"`
				Rate string `xml:"rate,attr"`
				Sum  string `xml:"sum,attr"`
				Name string `xml:"name,attr"`
			} `xml:"TAX"`
		} `xml:"CHECKTAXES"`
		CURRENCIES string `xml:"CURRENCIES"`
	} `xml:"CHECKDATA"`
}

type TransactionResult struct {
	TRRESPONSE struct {
		ErrorCode   int    `json:"error_code"`
		ErrText     string `json:"err_text"`
		TRANSACTION struct {
			ExtId    string `json:"ext_id"`
			Num      string `json:"num"`
			Cardcode string `json:"cardcode"`
			Slip     string `json:"slip"`
			Value    string `json:"value"`
		} `json:"TRANSACTION"`
	} `json:"TRRESPONSE"`
}
