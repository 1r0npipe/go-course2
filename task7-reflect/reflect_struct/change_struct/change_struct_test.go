package changestruct

import (
	"reflect"
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
		{
			got: person{
				FirstName: "john",
				Skills: []Skill{
					{Name: "c++", Exp: int64(2)}, // added field name, add other
					{Name: "python", Exp: int64(3)},
					{Name: "golang", Exp: int64(5)},
				},
			},
			want: person{
				FirstName: "Bob",
				Skills: []Skill{
					{Name: "test", Exp: int64(7)}, // add field names
					{Name: "test", Exp: int64(7)},
					{Name: "test", Exp: int64(7)},
				},
			},
			mapTest: mapTest,
			wantErr: nil,
		},
		{
			got: person{
				FirstName: "Bill",
				Skills:    []Skill{},
			},
			want: person{
				FirstName: "Bob",
				Skills:    []Skill{},
			},
			mapTest: mapTest,
			wantErr: nil,
		},
		{
			got: person{
				FirstName: "",
				Skills: []Skill{
					{Name: "РусскийТекст", Exp: -10},
				},
			},
			want: person{
				FirstName: "Bob",
				Skills: []Skill{
					{Name: "test", Exp: 7},
				},
			},
			mapTest: mapTest,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		if err := ChangeIntoStruct(&tt.got, tt.mapTest); err != nil {
			t.Errorf("ChangeIntoStruct() error = %v, wantErr %v", err, tt.wantErr)
		}
		if tt.got.FirstName != tt.want.FirstName {
			t.Errorf("Can't change filed: \"First name\"\n")
		}
		if reflect.TypeOf(tt.want.Skills).Kind() == reflect.Slice {
			for i := 0; i < len(tt.want.Skills); i += 1 {
				if tt.want.Skills[i].Name != tt.got.Skills[i].Name {
					t.Errorf("Can't change field: Name in Slice Skill, element: %d\n", i)
				}
				if tt.want.Skills[i].Exp != tt.got.Skills[i].Exp {
					t.Errorf("Can't change field: Exp in Slice Skill, element: %d\n", i)
				}
			}
		}
		if tt.wantErr != nil {
			t.Errorf("return unexpected error when trying to set the field in Interface\n")
		}
	}
}
