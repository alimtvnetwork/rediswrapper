package rediswrapper

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errbool"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"github.com/evatix-go/rediswrapper/internal/messages"
)

func (rw *Wrapper) Delete(key string) *errbool.Result {
	cmd := rw.client.Del(rw.ctx, key)

	if cmd == nil {
		return &errbool.Result{
			Value: false,
			ErrorWrapper: errnew.MessagesPtr(
				errtype.CommunicationFailed,
				messages.IntCmdNull),
		}
	}

	result, err := cmd.Result()

	if err != nil {
		return &errbool.Result{
			Value: false,
			ErrorWrapper: errnew.NewPtr(
				errtype.DeleteFailed,
				err),
		}
	}

	return &errbool.Result{
		Value:        result > 0,
		ErrorWrapper: errnew.EmptyPtr,
	}
}
