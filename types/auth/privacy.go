package authtypes

type Visibility string

const (
	Private Visibility = "private"
	Public  Visibility = "public"
	Friends Visibility = "friends"
)

func (p Visibility) GetAllowedVisibilities() []string {
	switch p {
	case Private:
		return []string{"private", "friends", "public"}
	case Friends:
		return []string{"friends", "public"}
	case Public:
		return []string{"public"}
	default:
		return []string{"public"}
	}
}
