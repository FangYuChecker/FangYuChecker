package Unit

import "testing"

func Test_monster_Store(t *testing.T) {
	type fields struct {
		Name  string
		Age   int
		Skill string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{"1",fields{
			Name:  "Tom",
			Age:   20,
			Skill: "1",
		},true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Monster{
				Name:  tt.fields.Name,
				Age:   tt.fields.Age,
				Skill: tt.fields.Skill,
			}
			if got := this.Store(); got != tt.want {
				t.Errorf("Store() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_monster_Restore(t *testing.T) {
	type fields struct {
		Name  string
		Age   int
		Skill string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{"1",fields{
			Name:  "Tom",
			Age:   20,
			Skill: "1",
		},true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Monster{
				Name:  tt.fields.Name,
				Age:   tt.fields.Age,
				Skill: tt.fields.Skill,
			}
			if got := this.Restore(); got != tt.want {
				t.Errorf("Restore() = %v, want %v", got, tt.want)
			}
		})
	}
}

