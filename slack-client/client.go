package slackclient

// Client can post message to slack
type Client interface {
	Post(string) error
}

type clientImpl struct {
}

func (c *clientImpl) Post(text string) error {
	return nil
}

// NewClient returns new slackClient
func NewClient() Client {
	return &clientImpl{}
}
