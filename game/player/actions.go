package player

type Action int

const (
	fold Action = iota
	check
	call
	bet
	raise
)

var Actions = struct {
	Fold  Action
	Check Action
	Call  Action
	Bet   Action
	Raise Action
}{
	Fold:  fold,
	Check: check,
	Call:  call,
	Bet:   bet,
	Raise: raise,
}
