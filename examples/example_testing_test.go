package examples

import (
	"testing"

	examplesmock "uber-mockgen-example/examples/mock"

	"go.uber.org/mock/gomock"
)

func Test_prefixCounter_AddCountPrefix(t *testing.T) {
	type fields struct {
		counter counter
		prefix  prefix
	}
	type args struct {
		i      int
		prefix string
		str    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "test",
			fields: fields{},
			args: args{
				i:      5,
				prefix: "PRE_",
				str:    "string",
			},
			want: "PRE_PRE_PRE_PRE_PRE_string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockCounter := examplesmock.NewMockCounterInterface(ctrl)
			mockPrefix := examplesmock.NewMockPrefixInterface(ctrl)

			mockCounter.EXPECT().PlusOne(gomock.Any()).Return(5)
			mockPrefix.EXPECT().AddPrefix(gomock.Any(), gomock.Any()).Return(tt.want)

			p := prefixCounter{
				counter: mockCounter,
				prefix:  mockPrefix,
			}
			if got := p.AddCountPrefix(tt.args.i, tt.args.prefix, tt.args.str); got != tt.want {
				t.Errorf("AddCountPrefix() = %v, want %v", got, tt.want)
			}

		})
	}
}
