package types

type Boolean int

const (
	Enabled  Boolean = 1
	Disabled Boolean = 0
)

var BooleanToString = map[Boolean]string{
	Enabled:  "Enabled",
	Disabled: "Disabled",
}

func (c Boolean) String() string {
	return BooleanToString[c]
}

type State int

const (
	StateBeforeStart = 1
	StateRunning     = 2
)

var StateToString = map[State]string{
	StateBeforeStart: "BeforeStart",
	StateRunning:     "Started",
}

func (c State) String() string {
	return StateToString[c]
}

type Favorite int

const (
	FavoriteHighPriority = 1
	FavoriteLowPriority  = 2
)

var FavoriteToString = map[Favorite]string{
	FavoriteHighPriority: "High",
	FavoriteLowPriority:  "Low",
}

func (c Favorite) String() string {
	return FavoriteToString[c]
}
