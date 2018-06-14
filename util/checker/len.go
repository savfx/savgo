package checker

func (ctx *Checker) Lgt(length, min int) *Checker {
	if length > min {
		return ctx
	}
	panic(ErrRuleLgt)
}

func (ctx *Checker) Lgte(length, min int) *Checker {
	if length >= min {
		return ctx
	}
	panic(ErrRuleLgt)
}

func (ctx *Checker) Llt(length, max int) *Checker {
	if length < max {
		return ctx
	}
	panic(ErrRuleLgt)
}

func (ctx *Checker) Llte(length, max int) *Checker {
	if length <= max {
		return ctx
	}
	panic(ErrRuleLgt)
}
