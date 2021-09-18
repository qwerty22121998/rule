package rule

type validIf struct {
	con  Rule
	then Rule
	el   Rule
}

func (v validIf) Validate() error {
	if v.con.Validate() == nil {
		return v.then.Validate()
	}
	return v.el.Validate()
}

func If(rule Rule) validIf {
	return validIf{
		con:  rule,
		then: Skip(),
		el:   Skip(),
	}
}

func (v validIf) Then(rule Rule) validIf {
	v.then = rule
	return v
}

func (v validIf) Else(rule Rule) validIf {
	v.el = rule
	return v
}
