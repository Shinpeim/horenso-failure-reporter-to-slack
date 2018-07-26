package failurereporter

// SlackClient can post horenso result to slack
type SlackClient interface {
	Post(*horensoOut) error
}

type slackClientImpl struct {
}

// NewSlackClient returns new slack client
func NewSlackClient() SlackClient {
	return &slackClientImpl{}
}

func (c *slackClientImpl) Post(ho *horensoOut) error {
	return nil
}
