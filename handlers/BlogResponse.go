package handlers

type BlogResponse struct {
	Success bool				`json:"success"`
	Code int					`json:"code"`
	Message string				`json:"message"`
	Data map[string]interface{}	`json:"data"`
}