package entity

type UserStatus int

const (
	INACTIVE UserStatus = iota
	ACTIVE
	BLOCKED
	UNVERIFIED
)

func (status UserStatus) Value() string {
	switch status {
	case ACTIVE:
		return "active"
	case INACTIVE:
		return "inactive"
	case BLOCKED:
		return "blocked"
	case UNVERIFIED:
		return "unverified"
	default:
		return "-"
	}
}
