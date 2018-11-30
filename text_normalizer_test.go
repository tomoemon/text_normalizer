package text_normalizer

import (
	"testing"
)

func TestNewTextNormalizer(t *testing.T) {
	tests := []struct {
		options []TextReplaceOption
		name    string
		input   string
		want    string
	}{
		{
			name:    "半角数字を全角数字に変換",
			options: []TextReplaceOption{HankakuNumberToZenkaku},
			input:   "abcABC123!#$ａｂｃＡＢＣ１２３！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC１２３!#$ａｂｃＡＢＣ１２３！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
		},
		{
			name:    "全角数字を半角数字に変換",
			options: []TextReplaceOption{ZenkakuNumberToHankaku},
			input:   "abcABC123!#$ａｂｃＡＢＣ１２３！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
		},
		{
			name:    "半角カタカナを全角カタカナに変換",
			options: []TextReplaceOption{HankakuKatakanaToZenkaku},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄アイウアイウあいうガガがパパぱ' '　'",
		},
		{
			name:    "半角カタカナを濁点なし全角カタカナに変換",
			options: []TextReplaceOption{HankakuKatakanaToZenkaku, RemoveDakuten},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄アイウアイウあいうカガがハパぱ' '　'",
		},
		{
			name:    "全角カタカナを半角カタカナに変換",
			options: []TextReplaceOption{ZenkakuKatakanaToHankaku},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳｱｲｳあいうｶﾞｶﾞがﾊﾟﾊﾟぱ' '　'",
		},
		{
			name:    "全角カタカナを濁点なし半角カタカナに変換",
			options: []TextReplaceOption{ZenkakuKatakanaToHankaku, RemoveDakuten},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳｱｲｳあいうｶﾞｶがﾊﾟﾊぱ' '　'",
		},
		{
			name:    "全角、半角カタカナをひらがなに変換",
			options: []TextReplaceOption{KatakanaToHiragana},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄あいうあいうあいうがががぱぱぱ' '　'",
		},
		{
			name:    "全角、半角カタカナを濁点なしひらがなに変換",
			options: []TextReplaceOption{KatakanaToHiragana, RemoveDakuten},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄あいうあいうあいうかかがははぱ' '　'",
		},
		{
			name:    "ひらがなを全角カタカナに変換",
			options: []TextReplaceOption{HiraganaToZenkakuKatakana},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウアイウｶﾞガガﾊﾟパパ' '　'",
		},
		{
			name:    "ひらがなを濁点なし全角カタカナに変換",
			options: []TextReplaceOption{HiraganaToZenkakuKatakana, RemoveDakuten},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウアイウｶﾞガカﾊﾟパハ' '　'",
		},
		{
			name:    "ひらがなを半角カタカナに変換",
			options: []TextReplaceOption{HiraganaToHankakuKatakana},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウｱｲｳｶﾞガｶﾞﾊﾟパﾊﾟ' '　'",
		},
		{
			name:    "ひらがなを濁点なし半角カタカナに変換",
			options: []TextReplaceOption{HiraganaToHankakuKatakana, RemoveDakuten},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウｱｲｳｶﾞガｶﾊﾟパﾊ' '　'",
		},
		{
			name:    "全角、半角カタカナ、ひらがなを半角カタカナに変換",
			options: []TextReplaceOption{KanaToHankakuKatakana},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳｱｲｳｱｲｳｶﾞｶﾞｶﾞﾊﾟﾊﾟﾊﾟ' '　'",
		},
		{
			name:    "全角、半角カタカナ、ひらがなを濁点なし半角カタカナに変換",
			options: []TextReplaceOption{KanaToHankakuKatakana, RemoveDakuten},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳｱｲｳｱｲｳｶｶｶﾊﾊﾊ' '　'",
		},
		{
			name:    "全角、半角カタカナ、ひらがなを全角カタカナに変換",
			options: []TextReplaceOption{KanaToZenkakuKatakana},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄アイウアイウアイウガガガパパパ' '　'",
		},
		{
			name:    "全角、半角カタカナ、ひらがなを濁点なし全角カタカナに変換",
			options: []TextReplaceOption{KanaToZenkakuKatakana, RemoveDakuten},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄アイウアイウアイウカカカハハハ' '　'",
		},
		{
			name:    "全角、半角カタカナ、ひらがなをひらがなに変換",
			options: []TextReplaceOption{KanaToHiragana},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄あいうあいうあいうがががぱぱぱ' '　'",
		},
		{
			name:    "全角、半角カタカナ、ひらがなを濁点なしひらがなに変換",
			options: []TextReplaceOption{KanaToHiragana, RemoveDakuten},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄あいうあいうあいうかかかははは' '　'",
		},
		{
			name:    "英字を全角大文字に変換",
			options: []TextReplaceOption{AlphabetToUpperZenkaku},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "ＡＢＣＡＢＣ123!#$ＡＢＣＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
		},
		{
			name:    "英字を半角大文字に変換",
			options: []TextReplaceOption{AlphabetToUpperHankaku},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "ABCABC123!#$ABCABC123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
		},
		{
			name:    "英字を全角小文字に変換",
			options: []TextReplaceOption{AlphabetToLowerZenkaku},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "ａｂｃａｂｃ123!#$ａｂｃａｂｃ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
		},
		{
			name:    "英字を半角小文字に変換",
			options: []TextReplaceOption{AlphabetToLowerHankaku},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcabc123!#$abcabc123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
		},
		{
			name:    "英字を全角に変換",
			options: []TextReplaceOption{AlphabetToZenkaku},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "ａｂｃＡＢＣ123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
		},
		{
			name:    "英字を半角に変換",
			options: []TextReplaceOption{AlphabetToHankaku},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$abcABC123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
		},
		{
			name:    "半角記号を全角に変換",
			options: []TextReplaceOption{HankakuSignToZenkaku},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123！＃＄ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ’ ’　’",
		},
		{
			name:    "全角記号を半角に変換",
			options: []TextReplaceOption{ZenkakuSignToHankaku},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123!#$ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
		},
		{
			name:    "半角スペースを全角に変換",
			options: []TextReplaceOption{HankakuSpaceToZenkaku},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ'　'　'",
		},
		{
			name:    "全角スペースを半角に変換",
			options: []TextReplaceOption{ZenkakuSpaceToHankaku},
			input:   "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcABC123!#$ａｂｃＡＢＣ123！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' ' '",
		},
		{
			name:    "全角数字を半角に、全角記号を半角に、英字を半角小文字に、かな文字を全角カタカナに変換",
			options: []TextReplaceOption{ZenkakuNumberToHankaku, ZenkakuSignToHankaku, AlphabetToLowerHankaku, KanaToZenkakuKatakana},
			input:   "abcABC123!#$ａｂｃＡＢＣ１２３！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcabc123!#$abcabc123!#$アイウアイウアイウガガガパパパ' '　'",
		},
		{
			name:    "全角数字を半角に、全角記号を半角に、英字を半角小文字に、かな文字を濁点なし全角カタカナに変換",
			options: []TextReplaceOption{ZenkakuNumberToHankaku, ZenkakuSignToHankaku, AlphabetToLowerHankaku, KanaToZenkakuKatakana, RemoveDakuten},
			input:   "abcABC123!#$ａｂｃＡＢＣ１２３！＃＄ｱｲｳアイウあいうｶﾞガがﾊﾟパぱ' '　'",
			want:    "abcabc123!#$abcabc123!#$アイウアイウアイウカカカハハハ' '　'",
		},
		{
			name:    "マッピング先がないものは変換時に無視される",
			options: []TextReplaceOption{KanaToHiragana},
			input:   "アイウﾜﾞヷ",
			want:    "あいうﾜﾞヷ",
		},
		{
			name:    "マッピング先がないものは変換時に消される",
			options: []TextReplaceOption{KanaToHiragana, RemoveNoMapping},
			input:   "アイウﾜﾞヷ",
			want:    "あいう",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			n := NewTextNormalizer(test.options...)
			got := n.Replace(test.input)
			if test.want != got {
				t.Errorf("want: %s, got: %s (options: %+v)", test.want, got, test.options)
			}
		})
	}
}
