package response

import "app/src/model"

type Common struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessWithUser struct {
	Code    int          `json:"code"`
	Status  string       `json:"status"`
	Message string       `json:"message"`
	User    []model.User `json:"user"`
}

type SuccessWithJobList struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Job     []model.Job `json:"job"`
}

type SuccessWithJob struct {
	Code    int       `json:"code"`
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Job     model.Job `json:"job"`
}

type SuccessWithCategory struct {
	Code     int              `json:"code"`
	Status   string           `json:"status"`
	Message  string           `json:"message"`
	Category []model.Category `json:"category"`
}

type SuccessWithTokens struct {
	Code    int        `json:"code"`
	Status  string     `json:"status"`
	Message string     `json:"message"`
	User    model.User `json:"user"`
	Tokens  Tokens     `json:"tokens"`
}

type SuccessWithLanguageList struct {
	Code      int                         `json:"code"`
	Status    string                      `json:"status"`
	Message   string                      `json:"message"`
	Languages []model.ProgrammingLanguage `json:"languages"`
}

type SuccessWithPaginate[T any] struct {
	Code         int    `json:"code"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	Results      []T    `json:"results"`
	Page         int    `json:"page"`
	Limit        int    `json:"limit"`
	TotalPages   int64  `json:"total_pages"`
	TotalResults int64  `json:"total_results"`
}

type ErrorDetails struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}
