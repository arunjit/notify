package main

type Notification struct {
	ToURL string             `json:"to"`
	Data  *Notification_Data `json:"data"`

	messageFormat string `json:-`
}

type Notification_Data struct {
	Message string `json:"message"`
}

func NewNotification(toURL, msgFmt string) *Notification {
	return &Notification{
		ToURL:         toURL,
		messageFormat: msgFmt,
	}
}

func (n *Notification) SetData(cs *CommandStatus) {
	n.Data = &Notification_Data{
		Message: cs.Format(n.messageFormat),
	}
}
