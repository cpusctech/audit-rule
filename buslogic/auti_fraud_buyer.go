package buslogic

import "audit-rule/buslogic/payload"

func (bl *BusLogic) AntiBuyerFraud(buyerInfo payload.AuditBuyerRequest) (*payload.AntiFraudOutput, error) {
	return &payload.AntiFraudOutput{
		IsPassed: true,
		RefuseReasons: []string{buyerInfo.IdCard},
	}, nil
}
