package player

type Action int

const (
	fold Action = iota
	check
	call
	bet
	raise
	allIn
)

var Actions = struct {
	Fold  Action
	Check Action
	Call  Action
	Bet   Action
	Raise Action
	AllIn Action
}{
	Fold:  fold,
	Check: check,
	Call:  call,
	Bet:   bet,
	Raise: raise,
	AllIn: allIn,
}
