package calc

import "testing"

func Test_calculator_Calculate(t *testing.T) {
	type args struct {
		n1 float64
		m  string
		n2 float64
	}
	tests := []struct {
		name    string
		c       calculator
		args    args
		want    float64
		wantErr bool
	}{
		{
			"Сложение",
			*NewCalculator(),
			args{
				1,
				"+",
				2,
			},
			3,
			false,
		},
		{
			"Вычитание",
			*NewCalculator(),
			args{
				3,
				"-",
				2,
			},
			1,
			false,
		},
		{
			"Умножение",
			*NewCalculator(),
			args{
				3,
				"*",
				2,
			},
			6,
			false,
		},
		{
			"Деление",
			*NewCalculator(),
			args{
				8,
				"/",
				2,
			},
			4,
			false,
		},
		{
			"Деление на ноль",
			*NewCalculator(),
			args{
				8,
				"/",
				0,
			},
			0,
			true,
		},
		{
			"несуществующая операция",
			*NewCalculator(),
			args{
				8,
				"=",
				0,
			},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := calculator{}
			got, err := c.Calculate(tt.args.n1, tt.args.m, tt.args.n2)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculator.Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("calculator.Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
