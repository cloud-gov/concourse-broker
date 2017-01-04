package cf

type Details struct {
	OrgGUID   string
	OrgName   string
	SpaceGUID string
	SpaceName string
}

type Client interface {
	GetDetails(spaceGUID string, orgGUID string) (Details, error)
}

func NewClient() Client {
	return &cfClient{}
}

type cfClient struct{}

func (c *cfClient) GetDetails(spaceGUID string, orgGUID string) (Details, error) {
	return Details{}, nil
}
