package advantshop

import (
	"encoding/json"
	"fmt"
	"time"
)

type OrdersService service

func (s *OrdersService) Add() {

}

type Order struct {
	Name      string    `json:"name,omitempty"`
	Notes     string    `json:"notes,omitempty"`
	Type      string    `json:"type,omitempty"`
	Status    string    `json:"status,omitempty"`
	DateStart time.Time `json:"date_start,omitempty"`
	//DateEnd         time.Time `json:"date_end,omitempty"`
	CardNo          string `json:"card_no,omitempty"`
	PromoCode       string `json:"promo_code,omitempty"`
	MarketSource    string `json:"market_source,omitempty"`
	MarketSegment   string `json:"market_segment,omitempty"`
	MarketTrackCode string `json:"market_track_code,omitempty"`
	MarketGeoCode   string `json:"market_geo_code,omitempty"`
	MarketOpenCode  string `json:"market_open_code,omitempty"`
	MarketExtra1    string `json:"market_extra1,omitempty"`
	MarketExtra2    string `json:"market_extra2,omitempty"`
	PointOfSaleId   int    `json:"point_of_sale_id,omitempty"`
	PosCode         string `json:"pos_code,omitempty"`
	CustomerId      int    `json:"customer_id,omitempty"`
	ExternalId      string `json:"external_id,omitempty"`
	//ExtDateCreated  time.Time `json:"ext_date_created,omitempty"`
	//ExtDateModified time.Time `json:"ext_date_modified,omitempty"`
	ParentId int `json:"parent_id,omitempty"`
	Customer *struct {
		ExternalId string `json:"external_id,omitempty"`
		//ExtDateCreated    time.Time `json:"ext_date_created,omitempty"`
		//ExtDateModified   time.Time `json:"ext_date_modified,omitempty"`
		FirstName  string `json:"first_name,omitempty"`
		LastName   string `json:"last_name,omitempty"`
		MiddleName string `json:"middle_name,omitempty"`
		Phone      string `json:"phone,omitempty"`
		//BirthDate         time.Time `json:"birth_date,omitempty"`
		Gender            string `json:"gender,omitempty"`
		Language          string `json:"language,omitempty"`
		Email             string `json:"email,omitempty"`
		Notes             string `json:"notes,omitempty"`
		AllowEmailContact bool   `json:"allow_email_contact,omitempty"`
		AllowPhoneContact bool   `json:"allow_phone_contact,omitempty"`
		IsClosed          bool   `json:"is_closed,omitempty"`
		IsVerified        bool   `json:"is_verified,omitempty"`
		PhoneVerified     bool   `json:"phone_verified,omitempty"`
		MergedToId        int    `json:"merged_to_id,omitempty"`
		UserId            int    `json:"user_id,omitempty"`
		Zip               string `json:"zip,omitempty"`
		CountryCode       string `json:"country_code,omitempty"`
		CountryName       string `json:"country_name,omitempty"`
		Region            string `json:"region,omitempty"`
		District          string `json:"district,omitempty"`
		SettlementType    string `json:"settlement_type,omitempty"`
		City              string `json:"city,omitempty"`
		Street            string `json:"street,omitempty"`
		HouseNo           string `json:"house_no,omitempty"`
		BuildingNo        string `json:"building_no,omitempty"`
		FlatNo            string `json:"flat_no,omitempty"`
		RoomNo            string `json:"room_no,omitempty"`
		RawAddress        string `json:"raw_address,omitempty"`
		DocType           string `json:"doc_type,omitempty"`
		DocIssuerInfo     string `json:"doc_issuer_info,omitempty"`
		DocSeries         string `json:"doc_series,omitempty"`
		DocNumber         string `json:"doc_number,omitempty"`
		DepartmentCode    string `json:"department_code,omitempty"`
		DepartmentName    string `json:"department_name,omitempty"`
		//DocIssueDate      time.Time `json:"doc_issue_date,omitempty"`
		//DocExpirationDate time.Time `json:"doc_expiration_date,omitempty"`
	} `json:"customer,omitempty"`
	Items []struct {
		Id         int    `json:"id,omitempty"`
		ExternalId string `json:"external_id,omitempty"`
		Name       string `json:"name,omitempty"`
		Notes      string `json:"notes,omitempty"`
		//DateCreated          time.Time `json:"date_created,omitempty"`
		//DateModified         time.Time `json:"date_modified,omitempty"`
		//ExtDateCreated       time.Time `json:"ext_date_created,omitempty"`
		//ExtDateModified      time.Time `json:"ext_date_modified,omitempty"`
		//Date                 time.Time `json:"date,omitempty"`
		GroupId              string  `json:"group_id,omitempty"`
		RevenueType          string  `json:"revenue_type,omitempty"`
		Code                 string  `json:"code,omitempty"`
		RateGroup            string  `json:"rate_group,omitempty"`
		Amount               float32 `json:"amount,omitempty"` // todo error
		AmountBeforeDiscount int     `json:"amount_before_discount,omitempty"`
		IncludedTaxAmount    int     `json:"included_tax_amount,omitempty"`
		Quantity             int     `json:"quantity,omitempty"`
		IsScheduled          bool    `json:"is_scheduled,omitempty"`
	} `json:"items,omitempty"`
	Payments    []*Item `json:"payments,omitempty"`
	ExtraFields []struct {
		BinaryData string `json:"binary_data,omitempty"`
		Id         int    `json:"id,omitempty"`
		Name       string `json:"name,omitempty"`
		Value      string `json:"value,omitempty"`
	} `json:"extra_fields,omitempty"`
}

type Item struct {
	Id         int    `json:"id,omitempty"`
	ExternalId string `json:"external_id,omitempty"`
	Name       string `json:"name,omitempty"`
	Notes      string `json:"notes,omitempty"`
	//DateCreated          time.Time `json:"date_created,omitempty"`
	//DateModified         time.Time `json:"date_modified,omitempty"`
	//ExtDateCreated       time.Time `json:"ext_date_created,omitempty"`
	//ExtDateModified      time.Time `json:"ext_date_modified,omitempty"`
	//Date                 time.Time `json:"date,omitempty"`
	GroupId              string  `json:"group_id,omitempty"`
	RevenueType          string  `json:"revenue_type,omitempty"`
	Code                 string  `json:"code,omitempty"`
	RateGroup            string  `json:"rate_group,omitempty"`
	Amount               float32 `json:"amount,omitempty"`
	AmountBeforeDiscount int     `json:"amount_before_discount,omitempty"`
	IncludedTaxAmount    int     `json:"included_tax_amount,omitempty"`
	Quantity             int     `json:"quantity,omitempty"`
	IsScheduled          bool    `json:"is_scheduled,omitempty"`
}

type OrderPostByPostCodeAndExternalId struct {
	Customer struct {
		ExtDateCreated    time.Time `json:"ext_date_created,omitempty"`
		ExtDateModified   time.Time `json:"ext_date_modified,omitempty"`
		Zip               string    `json:"zip,omitempty"`
		CountryCode       string    `json:"country_code,omitempty"`
		CountryName       string    `json:"country_name,omitempty"`
		Region            string    `json:"region,omitempty"`
		District          string    `json:"district,omitempty"`
		SettlementType    string    `json:"settlement_type,omitempty"`
		City              string    `json:"city,omitempty"`
		Street            string    `json:"street,omitempty"`
		HouseNo           string    `json:"house_no,omitempty"`
		BuildingNo        string    `json:"building_no,omitempty"`
		FlatNo            string    `json:"flat_no,omitempty"`
		RoomNo            string    `json:"room_no,omitempty"`
		RawAddres         string    `json:"raw_addres,omitempty"`
		DocType           string    `json:"doc_type,omitempty"`
		DocIssuerInfo     string    `json:"doc_issuer_info,omitempty"`
		DocSeries         string    `json:"doc_series,omitempty"`
		DocNumber         string    `json:"doc_number,omitempty"`
		DepartmentCode    string    `json:"department_code,omitempty"`
		DepartmentName    string    `json:"department_name,omitempty"`
		DocIssueDate      time.Time `json:"doc_issue_date,omitempty"`
		DocExpirationDate time.Time `json:"doc_expiration_date,omitempty"`
		Cards             []struct {
			CurrentTier struct {
				Id        int      `json:"id,omitempty"`
				Code      string   `json:"code,omitempty"`
				Name      string   `json:"name,omitempty"`
				MinPoints int      `json:"min_points,omitempty"`
				ProgramId int      `json:"program_id,omitempty"`
				CardImage string   `json:"card_image,omitempty"`
				Features  []string `json:"features,omitempty"`
			} `json:"current_tier,omitempty"`
			ExpectedTier struct {
				Id        int      `json:"id,omitempty"`
				Code      string   `json:"code,omitempty"`
				Name      string   `json:"name,omitempty"`
				MinPoints int      `json:"min_points,omitempty"`
				ProgramId int      `json:"program_id,omitempty"`
				CardImage string   `json:"card_image,omitempty"`
				Features  []string `json:"features,omitempty"`
			} `json:"expected_tier,omitempty"`
			Program struct {
				Id        int       `json:"id,omitempty"`
				Name      string    `json:"name,omitempty"`
				StartDate time.Time `json:"start_date,omitempty"`
				EndDate   time.Time `json:"end_date,omitempty"`
				IsDraft   bool      `json:"is_draft,omitempty"`
			} `json:"program,omitempty"`
			TierLogs []struct {
				DateCreated   time.Time `json:"date_created,omitempty"`
				EffectiveFrom MyTime    `json:"effective_from,omitempty"`
				EffectiveTill MyTime    `json:"effective_till,omitempty"`
				NewTier       struct {
					Id        int      `json:"id,omitempty"`
					Code      string   `json:"code,omitempty"`
					Name      string   `json:"name,omitempty"`
					MinPoints int      `json:"min_points,omitempty"`
					ProgramId int      `json:"program_id,omitempty"`
					CardImage string   `json:"card_image,omitempty"`
					Features  []string `json:"features,omitempty"`
				} `json:"new_tier,omitempty"`
				PreviousTier struct {
					Id        int      `json:"id,omitempty"`
					Code      string   `json:"code,omitempty"`
					Name      string   `json:"name,omitempty"`
					MinPoints int      `json:"min_points,omitempty"`
					ProgramId int      `json:"program_id,omitempty"`
					CardImage string   `json:"card_image,omitempty"`
					Features  []string `json:"features,omitempty"`
				} `json:"previous_tier,omitempty"`
			} `json:"tier_logs,omitempty"`
			Id              int       `json:"id,omitempty"`
			CardNo          string    `json:"card_no,omitempty"`
			Issuer          string    `json:"issuer,omitempty"`
			IssueDate       time.Time `json:"issue_date,omitempty"`
			DateCreated     time.Time `json:"date_created,omitempty"`
			DateModified    time.Time `json:"date_modified,omitempty"`
			ExtDateCreated  time.Time `json:"ext_date_created,omitempty"`
			ExtDateModified time.Time `json:"ext_date_modified,omitempty"`
			IsActive        bool      `json:"is_active,omitempty"`
			Account         struct {
				Id               int       `json:"id,omitempty"`
				BalanceDebit     int       `json:"balance_debit,omitempty"`
				BalanceExpirable int       `json:"balance_expirable,omitempty"`
				BalanceCredit    int       `json:"balance_credit,omitempty"`
				BalanceLevel     int       `json:"balance_level,omitempty"`
				Balance          int       `json:"balance,omitempty"`
				DateCreated      time.Time `json:"date_created,omitempty"`
				DateModified     time.Time `json:"date_modified,omitempty"`
				ExtDateCreated   time.Time `json:"ext_date_created,omitempty"`
				ExtDateModified  time.Time `json:"ext_date_modified,omitempty"`
				Attributes       []struct {
					Id    int    `json:"id,omitempty"`
					Name  string `json:"name,omitempty"`
					Value string `json:"value,omitempty"`
				} `json:"attributes,omitempty"`
			} `json:"account,omitempty"`
			CustomerId      int    `json:"customer_id,omitempty"`
			CurrentTierId   int    `json:"current_tier_id,omitempty"`
			CurrentTierCode string `json:"current_tier_code,omitempty"`
			CurrentTierName string `json:"current_tier_name,omitempty"`
			ExpectedTierId  int    `json:"expected_tier_id,omitempty"`
			ProgramId       int    `json:"program_id,omitempty"`
		} `json:"cards,omitempty"`
		ViewUrl string `json:"view_url,omitempty"`
		User    struct {
			Id          int       `json:"id,omitempty"`
			Username    string    `json:"username,omitempty"`
			FirstName   string    `json:"first_name,omitempty"`
			LastName    string    `json:"last_name,omitempty"`
			FullName    string    `json:"full_name,omitempty"`
			Email       string    `json:"email,omitempty"`
			LastLogin   time.Time `json:"last_login,omitempty"`
			DateJoined  time.Time `json:"date_joined,omitempty"`
			IsActive    bool      `json:"is_active,omitempty"`
			IsSuperuser bool      `json:"is_superuser,omitempty"`
			IsStaff     bool      `json:"is_staff,omitempty"`
		} `json:"user,omitempty"`
		Id                int       `json:"id,omitempty"`
		DateCreated       time.Time `json:"date_created,omitempty"`
		DateModified      time.Time `json:"date_modified,omitempty"`
		ExternalId        string    `json:"external_id,omitempty"`
		FirstName         string    `json:"first_name,omitempty"`
		LastName          string    `json:"last_name,omitempty"`
		MiddleName        string    `json:"middle_name,omitempty"`
		Phone             string    `json:"phone,omitempty"`
		BirthDate         time.Time `json:"birth_date,omitempty"`
		Gender            string    `json:"gender,omitempty"`
		Language          string    `json:"language,omitempty"`
		Email             string    `json:"email,omitempty"`
		Notes             string    `json:"notes,omitempty"`
		AllowEmailContact bool      `json:"allow_email_contact,omitempty"`
		AllowPhoneContact bool      `json:"allow_phone_contact,omitempty"`
		IsClosed          bool      `json:"is_closed,omitempty"`
		IsVerified        bool      `json:"is_verified,omitempty"`
		PhoneVerified     bool      `json:"phone_verified,omitempty"`
		MergedToId        int       `json:"merged_to_id,omitempty"`
		UserId            int       `json:"user_id,omitempty"`
	} `json:"customer,omitempty"`
	Payments []struct {
		Id              int       `json:"id,omitempty"`
		ExternalId      string    `json:"external_id,omitempty"`
		Name            string    `json:"name,omitempty"`
		Notes           string    `json:"notes,omitempty"`
		DateCreated     time.Time `json:"date_created,omitempty"`
		DateModified    time.Time `json:"date_modified,omitempty"`
		ExtDateCreated  time.Time `json:"ext_date_created,omitempty"`
		ExtDateModified time.Time `json:"ext_date_modified,omitempty"`
		Date            time.Time `json:"date,omitempty"`
		GroupId         string    `json:"group_id,omitempty"`
		Type            string    `json:"type,omitempty"`
		Code            string    `json:"code,omitempty"`
		Amount          int       `json:"amount,omitempty"`
		CardType        string    `json:"card_type,omitempty"`
		CardIssuer      string    `json:"card_issuer,omitempty"`
	} `json:"payments,omitempty"`
	ExtraFields []struct {
		Id    int    `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"extra_fields,omitempty"`
	Items []struct {
		Id                   int       `json:"id,omitempty"`
		ExternalId           string    `json:"external_id,omitempty"`
		Name                 string    `json:"name,omitempty"`
		Notes                string    `json:"notes,omitempty"`
		DateCreated          time.Time `json:"date_created,omitempty"`
		DateModified         time.Time `json:"date_modified,omitempty"`
		ExtDateCreated       time.Time `json:"ext_date_created,omitempty"`
		ExtDateModified      time.Time `json:"ext_date_modified,omitempty"`
		Date                 time.Time `json:"date,omitempty"`
		GroupId              string    `json:"group_id,omitempty"`
		RevenueType          string    `json:"revenue_type,omitempty"`
		Code                 string    `json:"code,omitempty"`
		RateGroup            string    `json:"rate_group,omitempty"`
		Amount               int       `json:"amount,omitempty"`
		AmountBeforeDiscount int       `json:"amount_before_discount,omitempty"`
		IncludedTaxAmount    int       `json:"included_tax_amount,omitempty"`
		Quantity             int       `json:"quantity,omitempty"`
		IsScheduled          bool      `json:"is_scheduled,omitempty"`
	} `json:"items,omitempty"`
	CustomerTier struct {
		Id        int      `json:"id,omitempty"`
		Code      string   `json:"code,omitempty"`
		Name      string   `json:"name,omitempty"`
		MinPoints int      `json:"min_points,omitempty"`
		ProgramId int      `json:"program_id,omitempty"`
		CardImage string   `json:"card_image,omitempty"`
		Features  []string `json:"features,omitempty"`
	} `json:"customer_tier,omitempty"`
	Id               int       `json:"id,omitempty"`
	DateCreated      time.Time `json:"date_created,omitempty"`
	DateModified     time.Time `json:"date_modified,omitempty"`
	Name             string    `json:"name,omitempty"`
	Notes            string    `json:"notes,omitempty"`
	Type             string    `json:"type,omitempty"`
	Status           string    `json:"status,omitempty"`
	DateStart        time.Time `json:"date_start,omitempty"`
	DateEnd          time.Time `json:"date_end,omitempty"`
	CardNo           string    `json:"card_no,omitempty"`
	PromoCode        string    `json:"promo_code,omitempty"`
	MarketSource     string    `json:"market_source,omitempty"`
	MarketSegment    string    `json:"market_segment,omitempty"`
	MarketTrackCode  string    `json:"market_track_code,omitempty"`
	MarketGeoCode    string    `json:"market_geo_code,omitempty"`
	MarketOpenCode   string    `json:"market_open_code,omitempty"`
	MarketExtra1     string    `json:"market_extra1,omitempty"`
	MarketExtra2     string    `json:"market_extra2,omitempty"`
	PointOfSaleId    int       `json:"point_of_sale_id,omitempty"`
	CustomerTierId   int       `json:"customer_tier_id,omitempty"`
	CustomerTierName string    `json:"customer_tier_name,omitempty"`
	CustomerId       int       `json:"customer_id,omitempty"`
	PosCode          string    `json:"pos_code,omitempty"`
	PosName          string    `json:"pos_name,omitempty"`
	ExternalId       string    `json:"external_id,omitempty"`
	ExtDateCreated   time.Time `json:"ext_date_created,omitempty"`
	ExtDateModified  time.Time `json:"ext_date_modified,omitempty"`
	ParentId         int       `json:"parent_id,omitempty"`
	ItemsAmount      int       `json:"items_amount,omitempty"`
	PaymentsAmount   int       `json:"payments_amount,omitempty"`
}

func (s *OrdersService) PostByPostCodeAndExternalId(PostCode *string, ExternalId *string, Order *Order) (orderPostByPostCodeAndExternalId *OrderPostByPostCodeAndExternalId, err error) {

	order, err := json.Marshal(Order)
	if err != nil {
		return nil, err
	}

	fmt.Println("==============")
	fmt.Println(string(order))
	fmt.Println("==============")

	r, err := s.httpClient.R().
		SetBody(order).
		Post(fmt.Sprintf("api/point-of-sales/%s/orders/eid/%s", *PostCode, *ExternalId))
	if err != nil {
		return
	}

	postByPostCodeAndExternalIdResult := new(PostByPostCodeAndExternalIdResult)
	if r.IsSuccess() {
		err = json.Unmarshal(r.Body(), &postByPostCodeAndExternalIdResult)
		fmt.Println(postByPostCodeAndExternalIdResult)
	}
	return
}

// todo переделать
type PostByPostCodeAndExternalIdResult struct {
	Customer struct {
		ExtDateCreated    time.Time `json:"ext_date_created"`
		ExtDateModified   time.Time `json:"ext_date_modified"`
		Zip               string    `json:"zip"`
		CountryCode       string    `json:"country_code"`
		CountryName       string    `json:"country_name"`
		Region            string    `json:"region"`
		District          string    `json:"district"`
		SettlementType    string    `json:"settlement_type"`
		City              string    `json:"city"`
		Street            string    `json:"street"`
		HouseNo           string    `json:"house_no"`
		BuildingNo        string    `json:"building_no"`
		FlatNo            string    `json:"flat_no"`
		RoomNo            string    `json:"room_no"`
		RawAddres         string    `json:"raw_addres"`
		DocType           string    `json:"doc_type"`
		DocIssuerInfo     string    `json:"doc_issuer_info"`
		DocSeries         string    `json:"doc_series"`
		DocNumber         string    `json:"doc_number"`
		DepartmentCode    string    `json:"department_code"`
		DepartmentName    string    `json:"department_name"`
		DocIssueDate      time.Time `json:"doc_issue_date"`
		DocExpirationDate time.Time `json:"doc_expiration_date"`
		Cards             []struct {
			CurrentTier struct {
				Id        int      `json:"id"`
				Code      string   `json:"code"`
				Name      string   `json:"name"`
				MinPoints int      `json:"min_points"`
				ProgramId int      `json:"program_id"`
				CardImage string   `json:"card_image"`
				Features  []string `json:"features"`
			} `json:"current_tier"`
			ExpectedTier struct {
				Id        int      `json:"id"`
				Code      string   `json:"code"`
				Name      string   `json:"name"`
				MinPoints int      `json:"min_points"`
				ProgramId int      `json:"program_id"`
				CardImage string   `json:"card_image"`
				Features  []string `json:"features"`
			} `json:"expected_tier"`
			Program struct {
				Id        int       `json:"id"`
				Name      string    `json:"name"`
				StartDate time.Time `json:"start_date"`
				EndDate   time.Time `json:"end_date"`
				IsDraft   bool      `json:"is_draft"`
			} `json:"program"`
			TierLogs []struct {
				DateCreated   time.Time `json:"date_created"`
				EffectiveFrom MyTime    `json:"effective_from"`
				EffectiveTill MyTime    `json:"effective_till"`
				NewTier       struct {
					Id        int      `json:"id"`
					Code      string   `json:"code"`
					Name      string   `json:"name"`
					MinPoints int      `json:"min_points"`
					ProgramId int      `json:"program_id"`
					CardImage string   `json:"card_image"`
					Features  []string `json:"features"`
				} `json:"new_tier"`
				PreviousTier struct {
					Id        int      `json:"id"`
					Code      string   `json:"code"`
					Name      string   `json:"name"`
					MinPoints int      `json:"min_points"`
					ProgramId int      `json:"program_id"`
					CardImage string   `json:"card_image"`
					Features  []string `json:"features"`
				} `json:"previous_tier"`
			} `json:"tier_logs"`
			Id              int       `json:"id"`
			CardNo          string    `json:"card_no"`
			Issuer          string    `json:"issuer"`
			IssueDate       time.Time `json:"issue_date"`
			DateCreated     time.Time `json:"date_created"`
			DateModified    time.Time `json:"date_modified"`
			ExtDateCreated  time.Time `json:"ext_date_created"`
			ExtDateModified time.Time `json:"ext_date_modified"`
			IsActive        bool      `json:"is_active"`
			Account         struct {
				Id               int       `json:"id"`
				BalanceDebit     float64   `json:"balance_debit"`
				BalanceExpirable float64   `json:"balance_expirable"`
				BalanceCredit    float64   `json:"balance_credit"`
				BalanceLevel     float64   `json:"balance_level"`
				Balance          float64   `json:"balance"`
				DateCreated      time.Time `json:"date_created"`
				DateModified     time.Time `json:"date_modified"`
				ExtDateCreated   time.Time `json:"ext_date_created"`
				ExtDateModified  time.Time `json:"ext_date_modified"`
				Attributes       []struct {
					Id    int    `json:"id"`
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"attributes"`
			} `json:"account"`
			CustomerId      int    `json:"customer_id"`
			CurrentTierId   int    `json:"current_tier_id"`
			CurrentTierCode string `json:"current_tier_code"`
			CurrentTierName string `json:"current_tier_name"`
			ExpectedTierId  int    `json:"expected_tier_id"`
			ProgramId       int    `json:"program_id"`
		} `json:"cards"`
		ViewUrl string `json:"view_url"`
		User    struct {
			Id          int       `json:"id"`
			Username    string    `json:"username"`
			FirstName   string    `json:"first_name"`
			LastName    string    `json:"last_name"`
			FullName    string    `json:"full_name"`
			Email       string    `json:"email"`
			LastLogin   time.Time `json:"last_login"`
			DateJoined  time.Time `json:"date_joined"`
			IsActive    bool      `json:"is_active"`
			IsSuperuser bool      `json:"is_superuser"`
			IsStaff     bool      `json:"is_staff"`
		} `json:"user"`
		Id                int       `json:"id"`
		DateCreated       time.Time `json:"date_created"`
		DateModified      time.Time `json:"date_modified"`
		ExternalId        string    `json:"external_id"`
		FirstName         string    `json:"first_name"`
		LastName          string    `json:"last_name"`
		MiddleName        string    `json:"middle_name"`
		Phone             string    `json:"phone"`
		BirthDate         time.Time `json:"birth_date"`
		Gender            string    `json:"gender"`
		Language          string    `json:"language"`
		Email             string    `json:"email"`
		Notes             string    `json:"notes"`
		AllowEmailContact bool      `json:"allow_email_contact"`
		AllowPhoneContact bool      `json:"allow_phone_contact"`
		IsClosed          bool      `json:"is_closed"`
		IsVerified        bool      `json:"is_verified"`
		PhoneVerified     bool      `json:"phone_verified"`
		MergedToId        int       `json:"merged_to_id"`
		UserId            int       `json:"user_id"`
	} `json:"customer"`
	Payments []struct {
		Id              int       `json:"id"`
		ExternalId      string    `json:"external_id"`
		Name            string    `json:"name"`
		Notes           string    `json:"notes"`
		DateCreated     time.Time `json:"date_created"`
		DateModified    time.Time `json:"date_modified"`
		ExtDateCreated  time.Time `json:"ext_date_created"`
		ExtDateModified time.Time `json:"ext_date_modified"`
		Date            time.Time `json:"date"`
		GroupId         string    `json:"group_id"`
		Type            string    `json:"type"`
		Code            string    `json:"code"`
		Amount          int       `json:"amount"`
		CardType        string    `json:"card_type"`
		CardIssuer      string    `json:"card_issuer"`
	} `json:"payments"`
	ExtraFields []struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"extra_fields"`
	Items []struct {
		Id                   int       `json:"id"`
		ExternalId           string    `json:"external_id"`
		Name                 string    `json:"name"`
		Notes                string    `json:"notes"`
		DateCreated          time.Time `json:"date_created"`
		DateModified         time.Time `json:"date_modified"`
		ExtDateCreated       time.Time `json:"ext_date_created"`
		ExtDateModified      time.Time `json:"ext_date_modified"`
		Date                 time.Time `json:"date"`
		GroupId              string    `json:"group_id"`
		RevenueType          string    `json:"revenue_type"`
		Code                 string    `json:"code"`
		RateGroup            string    `json:"rate_group"`
		Amount               int       `json:"amount"`
		AmountBeforeDiscount int       `json:"amount_before_discount"`
		IncludedTaxAmount    int       `json:"included_tax_amount"`
		Quantity             int       `json:"quantity"`
		IsScheduled          bool      `json:"is_scheduled"`
	} `json:"items"`
	CustomerTier struct {
		Id        int      `json:"id"`
		Code      string   `json:"code"`
		Name      string   `json:"name"`
		MinPoints int      `json:"min_points"`
		ProgramId int      `json:"program_id"`
		CardImage string   `json:"card_image"`
		Features  []string `json:"features"`
	} `json:"customer_tier"`
	Id               int       `json:"id"`
	DateCreated      time.Time `json:"date_created"`
	DateModified     time.Time `json:"date_modified"`
	Name             string    `json:"name"`
	Notes            string    `json:"notes"`
	Type             string    `json:"type"`
	Status           string    `json:"status"`
	DateStart        time.Time `json:"date_start"`
	DateEnd          time.Time `json:"date_end"`
	CardNo           string    `json:"card_no"`
	PromoCode        string    `json:"promo_code"`
	MarketSource     string    `json:"market_source"`
	MarketSegment    string    `json:"market_segment"`
	MarketTrackCode  string    `json:"market_track_code"`
	MarketGeoCode    string    `json:"market_geo_code"`
	MarketOpenCode   string    `json:"market_open_code"`
	MarketExtra1     string    `json:"market_extra1"`
	MarketExtra2     string    `json:"market_extra2"`
	PointOfSaleId    int       `json:"point_of_sale_id"`
	CustomerTierId   int       `json:"customer_tier_id"`
	CustomerTierName string    `json:"customer_tier_name"`
	CustomerId       int       `json:"customer_id"`
	PosCode          string    `json:"pos_code"`
	PosName          string    `json:"pos_name"`
	ExternalId       string    `json:"external_id"`
	ExtDateCreated   time.Time `json:"ext_date_created"`
	ExtDateModified  time.Time `json:"ext_date_modified"`
	ParentId         int       `json:"parent_id"`
	ItemsAmount      int       `json:"items_amount"`
	PaymentsAmount   int       `json:"payments_amount"`
}
