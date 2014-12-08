package tiny

func F() string { return "foo" }

func G() string { return F() }

var h = "hello"
