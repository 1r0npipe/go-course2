package changestruct

import (
	"testing"
)

type person struct {
	FirstName string
	Skills    []Skill
}
type Skill struct {
	Name string
	Exp  int64
}

func TestChangeIntoStruct(t *testing.T) {
	personBefore := &person{
		FirstName: "john",
		Skills: []Skill{
			{"c++", int64(2)},
			{"python", int64(3)},
			{"golang", int64(5)},
		},
	}
	personAfter := &person{
		FirstName: "Bob",
		Skills: []Skill{
			{"test", int64(7)},
			{"test", int64(7)},
			{"test", int64(7)},
		},
	}
	mapTest := make(map[string]interface{}, 1)
	mapTest["Exp"] = int64(7)
	mapTest["Name"] = "test"
	mapTest["FirstName"] = "Bob"

	tests := []struct {
		got     person
		want    person
		mapTest map[string]interface{}
		wantErr error
	}{
		{*personBefore, *personAfter, mapTest, nil},
	}
	for _, tt := range tests {
		if err := ChangeIntoStruct(&tt.got, tt.mapTest); err != nil {
			t.Errorf("ChangeIntoStruct() error = %v, wantErr %v", err, tt.wantErr)
		}
		if tt.got.FirstName != tt.want.FirstName {
			t.Errorf("Can't change filed: \"First name\"")
		} else {
			for i := 0; i < len(tt.want.Skills); i += 1 {
				if tt.want.Skills[i].Name != tt.got.Skills[i].Name {
					t.Errorf("Can't change field: Name in Slice Skill")
				}
				if tt.want.Skills[i].Exp != tt.got.Skills[i].Exp {
					t.Errorf("Can't change field: Exp in Slice Skill")
				}
			}
		}
		if tt.wantErr != nil {
			t.Errorf("return unexpected error when trying to set the field in Interface")
		}
	}
}
