package set

type Set struct {
	data map[int] string
}

func (s * Set) New()  {
	s.data = make(map[int] string)
}

func (s * Set) ContainElement(item int) bool {
	var exist bool
	_, exist = s.data[item]
	
	return exist
}

func(s *Set) AddElement(item int) {
	if !s.ContainElement(item){
		s.data[item] = "Test"
	}
}

func(s *Set) DeleteElement(item int) {
	delete(s.data, item)
}

func(s *Set) Intersect(anotherSet *Set) *Set {
	var intersectSet = &Set{}
	intersectSet.New()


	for key, _ := range s.data{
		if(anotherSet.ContainElement(key)){
			intersectSet.AddElement(key)
		}
	}

	return intersectSet
}

func(s *Set) Union(anotherSet *Set) *Set {
	var unionSet = &Set{}
	unionSet.New()

	for value := range s.data{
		unionSet.AddElement(value)
	}

	for value := range anotherSet.data{
		unionSet.AddElement(value)
	}

	return unionSet
}