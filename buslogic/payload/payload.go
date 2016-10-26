package payload

/**********const************/
const CATEGORY_ANTI_FRAUD_BUYER  = 1
const CATEGORY_ANTI_FRAUD_CAR    = 2

/*********Request***********/
type AuditBuyerRequest struct {
	Phone       string      `json:"phone" validate:"nonzero"`
	IdCard      string      `json:"id_card" validate:"nonzero"`
	Name        string      `json:"name" validate:"nonzero"`
	AppointInfo AppointInfo `json:"appoint_info"` //带看信息
	AddressInfo AddressInfo `json:"address_info"` //地址信息
	WorkInfo    WorkInfo    `json:"work_info"`    //工作信息
}

type AppointInfo struct {
	SeeTimes int `json:"see_times"`
}

type AddressInfo struct {
	ProvinceId int `json:"province_id"`
	CityId     int `json:"city_id"`
}

type WorkInfo struct {
	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
}

type AuditSellerRequest struct {
	Phone  string `json:"phone"`
	IdCard string `json:"id_card"`
	Name   string `json:"name"`
}

/************Output*************/
type AntiFraudOutput struct {
	IsPassed      bool      `json:"is_passed"`
	RefuseReasons []string `json:"refuse_reasons"`
}

type EstimationOutput struct {
	IsPassed   bool    `json:"is_passed"`
	estimation float32 `json:"estimation"`
}

type ScoreOutput struct {
	score float32 `json:"score"`
}
