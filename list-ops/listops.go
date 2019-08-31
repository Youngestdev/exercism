package listops

type binFunc func(x, y int) int 
type IntList []int
type predFunc func(n int) bool
type unaryFunc func(x int) int

func (list IntList) Foldl(f binFunc, init int) int {
	if list == nil {
		return init
	}
	for _, v := range list {
		init = f(init, v)
	}
	return init
}

func (list IntList) Foldr(f binFunc, init int) int {
	if list == nil {
		return init
	}

	for i := range list {
		init = f(list[len(list)-1-i], init)
	}
	return init
}

func (list IntList) Filter(f predFunc) IntList {
	if len(list) == 0 {
		return list
	}
	var newList IntList
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func (list IntList) Length() int{
	return len(list)
}

func (list IntList) Map(f unaryFunc) IntList {
	if len(list) == 0 {
		return list
	}

	var newList IntList
	for _, v := range list {
		newList = append(newList, f(v))
	}
	return newList
}

func (list IntList) Reverse() IntList {
	if len(list) == 0 {
		return list
	}
	var newList IntList
	for _, v := range list {
		// Variadic looks coooool, lel.
		newList = append(IntList([]int{v}), newList...)
	}
	return newList
}

func (list IntList) Append(n []int) IntList {
	return append(list, n...)
}

func (list IntList) Concat(args []IntList) IntList {
	if len(list) == 0 {
		return list
	}
	newList := list
	for _, v := range args {
		newList = append(newList, v...)
	}
	return newList
}