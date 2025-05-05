package brigg

func Limit(fun func(*Style) bool, howmany int) func(s *Style) bool {
	var count int
	return func(s *Style) bool {
		if count >= howmany {
			return true
		}
		count++
		return fun(s)
	}
}

func Once(fun func(*Style) bool) func(s *Style) bool {
	return Limit(fun, 1)
}
