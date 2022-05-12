package singleton

import (
	"gateserver/component"
	"sync"
)

var VerifyInstance component.Verify
var verifyOnce sync.Once

func NewVerify(verify component.Verify) component.Verify {
	verifyOnce.Do(func() {
		VerifyInstance = verify
	})

	return VerifyInstance
}
