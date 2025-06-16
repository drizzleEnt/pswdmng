package checker

import "pswdmng/internal/service"

var _ service.Checker = (*checkerService)(nil)

type checkerService struct {
}
