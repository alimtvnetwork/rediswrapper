package rediserrwrapper

import (
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"

	"github.com/evatix-go/rediswrapper/internal/messages"
)

var (
	NullEmptyStringCmd = errnew.MessagesPtr(
		errtype.NullOrEmptyReference,
		messages.StringCmdNull)
)
