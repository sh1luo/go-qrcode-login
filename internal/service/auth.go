package service

type RegRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

func (svc *Service) CheckAndCreate(regInfo *RegRequest) error {
	if err := svc.Dao.CreateAccount(regInfo.AppKey, regInfo.AppSecret); err != nil {
		return err
	}
	return nil
}
