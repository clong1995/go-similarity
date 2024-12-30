package similarity

import (
	"testing"
)

func TestNGramFeatures(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "NGram Features",
			args: args{
				s: "谢谢夸奖！😊 我很高兴能帮助到你！如果你有更多问题或者其他想法，随时告诉我，我会尽力提供支持和解决方案的！你已经有很棒的想法了，相信你能做得更好！",
				n: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NGramFeatures(tt.args.s, tt.args.n)
			t.Logf("NGramFeatures() = %v", got)
		})
	}
}
