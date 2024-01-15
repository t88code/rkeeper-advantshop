package maxma

import "strconv"

type CustomersParam func(*CustomersService)

func Page(page int) CustomersParam {
	return func(s *CustomersService) {
		s.httpClient.QueryParam.Add("page", strconv.Itoa(page))
	}
}

func ItemsPerPage(itemsPerPage int) CustomersParam {
	return func(s *CustomersService) {
		s.httpClient.QueryParam.Add("itemsPerPage", strconv.Itoa(itemsPerPage))
	}
}

func Name(name string) CustomersParam {
	return func(s *CustomersService) {
		s.httpClient.QueryParam.Add("name", name)
	}
}
func Email(email string) CustomersParam {
	return func(s *CustomersService) {
		s.httpClient.QueryParam.Add("email", email)
	}
}
func Phone(phone string) CustomersParam {
	return func(s *CustomersService) {
		s.httpClient.QueryParam.Add("phone", phone)
	}
}
