package buslogic

import (
	"audit-rule/model"
)

type BusLogic struct {
	Model            *model.Model
}

func NewBusLogic() (*BusLogic, error) {
	m := model.NewModel()
	bl := &BusLogic{
		Model:            m,
	}
	return bl, nil
}

