package log

import (
	"context"
	"strings"

	"v2ray.com/core/common/serial"
)

type logKey int

const (
	accessMessageKey logKey = iota
)

type AccessStatus string

const (
	AccessAccepted = AccessStatus("accepted")
	AccessRejected = AccessStatus("rejected")
)

type AccessMessage struct {
	From   interface{}
	To     interface{}
	Status AccessStatus
	Reason interface{}
	Detour interface{}
}

func (m *AccessMessage) String() string {
	builder := strings.Builder{}
	builder.WriteString(serial.ToString(m.From))
	builder.WriteByte(' ')
	builder.WriteString(string(m.Status))
	builder.WriteByte(' ')
	builder.WriteString(serial.ToString(m.To))
	builder.WriteByte(' ')
	builder.WriteString(serial.ToString(m.Detour))
	builder.WriteByte(' ')
	builder.WriteString(serial.ToString(m.Reason))

	if len(m.Email) > 0 {
		builder.WriteString("email:")
		builder.WriteString(m.Email)
		builder.WriteByte(' ')
	}
	return builder.String()
}

func ContextWithAccessMessage(ctx context.Context, accessMessage *AccessMessage) context.Context {
	return context.WithValue(ctx, accessMessageKey, accessMessage)
}

func AccessMessageFromContext(ctx context.Context) *AccessMessage {
	if accessMessage, ok := ctx.Value(accessMessageKey).(*AccessMessage); ok {
		return accessMessage
	}
	return nil
}
