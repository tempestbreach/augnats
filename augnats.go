package augnats

import(
    "github.com/nats-io/nats.go"
)

type Handler interface {
    HandleMsg(*nats.Msg)
}

type HandlerFunc func(*nats.Msg)

func(f HandlerFunc) HandleMsg(msg *nats.Msg) {
    f(msg)
}
