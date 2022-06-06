package engine

func NewCatCommand(arg1, arg2 string) *catCommand {
	return &catCommand{
		arg1: arg1,
		arg2: arg2,
	}
}

type catCommand struct {
	arg1, arg2 string
}

func (cat *catCommand) Execute(hand Handler) {
	result := cat.arg1 + cat.arg2
	hand.Post(NewPrintCommand(result))
}
