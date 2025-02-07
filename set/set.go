package set

type set struct {
	m map[string]struct{}
}

func New() *set { // конструктор
	s := &set{
		m: make(map[string]struct{}, 0),
	}

	return s
}

func(s *set) Put(value string) {
	s.m[value] = struct{}{}
}

func(s *set) Exists(value string) bool {
	_, exists := s.m[value]
	return exists
}

func(s *set) Length() int {
	return len(s.m)
}

func(s *set) Delete(value string) {
	delete(s.m, value)
}



// Пакет set
// set.Put
// set.Pop
// set.Exists // значение существует в нашем множестве
// set.Length