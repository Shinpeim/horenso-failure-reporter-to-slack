package failurereporter

// SlackClient can post message to slack
type SlackClient interface {
	Post(string) error
}

type slackClientImpl struct {
}

func (c *slackClientImpl) Post(text string) error {
	return nil
}

// NewSlackClient returns new slackClient
func NewSlackClient() SlackClient {
	return &slackClientImpl{}
}
