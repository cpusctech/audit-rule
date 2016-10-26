package service

import (
	"github.com/kataras/iris"
	"audit-rule/buslogic/payload"
	ae "audit-rule/error"
)

func (s *Service) test(ctx *iris.Context) (*Response, error) {
	var request payload.AuditBuyerRequest
	return s.jsonRequest(ctx, &request, func() (*Response, error) {
		auditResult, err := s.bl.AntiBuyerFraud(request)
		if err != nil {
			return s.NewResponse(ae.ERROR_ANTI_FRAUD_BUYER, err.Error(), err.Error()), nil
		}
		return s.NewApiOkResponse(auditResult), nil
	})
}
