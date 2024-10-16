package storage

import "testing"

type testCase struct {
	name      string
	key       string
	value     string
	valueType string
}

func TestSet(t *testing.T) {
	cases := []testCase{
		{"test1", "key", "value", "S"},
		{"test2", "key", "42", "D"},
	}

	s, err := NewStorage()
	if err != nil {
		t.Errorf("new storage: %v", err)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s.Set(c.key, c.value)

			sValue := s.inner[c.key].stringValue

			if sValue != c.value {
				t.Errorf("values not equal")
			}
		})
	}
}

func TestGet(t *testing.T) {
	cases := []testCase{
		{"test1", "key", "value", "S"},
		{"test2", "key", "42", "D"},
	}

	s, err := NewStorage()
	if err != nil {
		t.Errorf("new storage: %v", err)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s.inner[c.key] = value{stringValue: c.value}

			sValue := s.Get(c.key)

			if *sValue != c.value {
				t.Errorf("values not equal")
			}
		})
	}
}

func TestGetType(t *testing.T) {
	cases := []testCase{
		{"test1", "key", "value", "S"},
		{"test2", "key", "42", "D"},
	}

	s, err := NewStorage()
	if err != nil {
		t.Errorf("new storage: %v", err)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s.Set(c.key, c.value)

			sType := s.GetType(c.key)

			if sType != c.valueType {
				t.Errorf("values not equal")
			}
		})
	}
}
