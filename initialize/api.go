package initialize

import "github.com/13842727496/gin-vue-admin-backend/repository"

// ApiGroup general api.
type ApiGroup struct {
	LoginRepository repository.LoginRepository
}

// NewApiGroup create new ApiGroup instance.
func NewApiGroup() *ApiGroup {
	return &ApiGroup{}
}
