package udp

import (
	"time"
)

const (
	// UDPMsgFieldSeparator - UDP message field separator string
	//UDPMsgFieldSeparator = ":::" // Field separator
	msgSendListDelay      = time.Millisecond        // time delay between two send list items handling
	msgSendListEmptyDelay = 3 * time.Millisecond    // time delay for the empty send list
	msgSendListInterval   = 1500 * time.Millisecond // interval between repeated messages

	WaitResponseSecs   = 2 * time.Second
	MsgSendRepeatLimit = 3
)
