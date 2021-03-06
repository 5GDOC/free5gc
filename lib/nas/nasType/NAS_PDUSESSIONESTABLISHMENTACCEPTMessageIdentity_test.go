//go:binary-only-package

package nasType_test

import (
	"free5gc/lib/nas"
	"free5gc/lib/nas/nasType"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNasTypeNewPDUSESSIONESTABLISHMENTACCEPTMessageIdentity(t *testing.T) {}

type nasTypePDUSESSIONESTABLISHMENTACCEPTMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONESTABLISHMENTACCEPTMessageIdentityMessageTypeTable = []nasTypePDUSESSIONESTABLISHMENTACCEPTMessageIdentityMessageType{
	{nas.MsgTypePDUSessionEstablishmentAccept, nas.MsgTypePDUSessionEstablishmentAccept},
}

func TestNasTypeGetSetPDUSESSIONESTABLISHMENTACCEPTMessageIdentityMessageType(t *testing.T) {}
