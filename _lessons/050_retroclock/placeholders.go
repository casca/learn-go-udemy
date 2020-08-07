package main

type placeholder [5]string

var zero = placeholder{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

var one = placeholder{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

var two = placeholder{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

var three = placeholder{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

var four = placeholder{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

var five = placeholder{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

var six = placeholder{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

var seven = placeholder{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

var eight = placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

var nine = placeholder{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

var colon = placeholder{
	"   ",
	" ░ ",
	"   ",
	" ░ ",
	"   ",
}

var dot = placeholder{
	"   ",
	"   ",
	"   ",
	"   ",
	" ░ ",
}

var digits = [...]placeholder{zero, one, two, three, four, five, six, seven, eight, nine}

var blank = placeholder{
	"   ",
	"   ",
	"   ",
	"   ",
	"   ",
}

var a = placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"█ █",
}

var l = placeholder{
	"█  ",
	"█  ",
	"█  ",
	"█  ",
	"███",
}

var r = placeholder{
	"██ ",
	"█ █",
	"██ ",
	"█ █",
	"█ █",
}

var m = placeholder{
	"█ █",
	"███",
	"█ █",
	"█ █",
	"█ █",
}

var exclamation = placeholder{
	" █ ",
	" █ ",
	" █ ",
	"   ",
	" █ ",
}

var alarm = [...]placeholder{
	blank,
	a, l, a, r, m, exclamation,
	blank, blank, blank,
}
