package tools

import (
	"github.com/rs/xid"
)

func GetXid() string {
	guXid := xid.New()
	id := guXid.String()
	return id
}
