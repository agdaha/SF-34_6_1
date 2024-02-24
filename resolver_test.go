package main

import (
	"reflect"
	"testing"
)

func TestResolver(t *testing.T) {
	type args struct {
		l []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Из задания на сайте",
			args: args{
				l: []string{"5+4=?",
					"9+3=?",
					"Сегодня прекрасная погода",
					"13+7=?",
					"4-2=?"},
			},
			want: []string{"5+4=9",
				"9+3=12",
				"13+7=20",
				"4-2=2",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Resolver(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("Resolver() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Resolver() = %v, want %v", got, tt.want)
			}
		})
	}
}
