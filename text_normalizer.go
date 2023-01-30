package text_normalizer

import "strings"

type TextReplaceOption int

/*
用語の定義
Hankaku…半角文字
Zenkaku…全角文字
Katakana…カタカナ
Hiragana…ひらがな
Kana…カタカナ、ひらがな
Ascii…ASCIIコード表で定義される文字
Alphabet…英字
Sign…記号
Space…空白文字
*/

const (
	// 以下、変換プロセスに影響を与える共通オプション

	// RemoveDakuten は濁点なしの文字に変換する（カナ系文字への変換オプションと同時に使う）
	RemoveDakuten TextReplaceOption = iota + 1

	// RemoveNoMapping は変換先の文字が存在しない場合に削除する（デフォルト：無視する）
	RemoveNoMapping

	// 以下、個別の変換オプション

	// HankakuNumberToZenkaku は半角数字→全角数字
	HankakuNumberToZenkaku

	// ZenkakuNumberToHankaku は全角数字→半角数字
	ZenkakuNumberToHankaku

	// HankakuKatakanaToZenkaku は半角カタカナ→全角カタカナ
	HankakuKatakanaToZenkaku

	// ZenkakuKatakanaToHankaku は全角カタカナ→半角カタカナ
	ZenkakuKatakanaToHankaku

	// KatakanaToHiragana は(全角カタカナ、半角カタカナ）→ひらがな
	KatakanaToHiragana

	// HiraganaToZenkakuKatakana はひらがな→全角カタカナ
	HiraganaToZenkakuKatakana

	// HiraganaToHankakuKatakana はひらがな→半角カタカナ
	HiraganaToHankakuKatakana

	// KanaToHiragana は(ひらがな、全角カタカナ、半角カタカナ) →ひらがな
	KanaToHiragana

	// KanaToZenkakuKatakana は(ひらがな、全角カタカナ、半角カタカナ) →全角カタカナ
	KanaToZenkakuKatakana

	// KanaToHankakuKatakana は(ひらがな、全角カタカナ、半角カタカナ) →半角カタカナ
	KanaToHankakuKatakana

	// AlphabetToUpperZenkaku は(半角英字、全角英字) →全角大文字英字
	AlphabetToUpperZenkaku

	// AlphabetToUpperHankaku は(半角英字、全角英字) →半角大文字英字
	AlphabetToUpperHankaku

	// AlphabetToLowerZenkaku は(半角英字、全角英字) →全角小文字英字
	AlphabetToLowerZenkaku

	// AlphabetToLowerHankaku は(半角英字、全角英字) →半角小文字英字
	AlphabetToLowerHankaku

	// AlphabetToZenkaku は(半角英字、全角英字) →全角英字（大文字小文字の区別は維持する）
	AlphabetToZenkaku

	// AlphabetToHankaku は(半角英字、全角英字) →半角英字（大文字小文字の区別は維持する）
	AlphabetToHankaku

	// HankakuSignToZenkaku は半角記号→全角記号
	HankakuSignToZenkaku

	// ZenkakuSignToHankaku は全角記号→半角記号
	ZenkakuSignToHankaku

	// HankakuSpaceToZenkaku は半角スペース→全角スペース
	HankakuSpaceToZenkaku

	// ZenkakuSpaceToHankaku は全角スペース→半角スペース
	ZenkakuSpaceToHankaku
)

func hasOption(flags []TextReplaceOption, search TextReplaceOption) bool {
	for _, f := range flags {
		if f == search {
			return true
		}
	}
	return false
}

func NewTextNormalizer(flags ...TextReplaceOption) *strings.Replacer {
	noDakutenSlide := 0
	if hasOption(flags, RemoveDakuten) {
		noDakutenSlide = 3
	}
	removeNoMapping := false
	if hasOption(flags, RemoveNoMapping) {
		removeNoMapping = true
	}

	projectAsciiNumber := asciiNumberMap.projectFuncWithOption(removeNoMapping)
	projectKana := kanaMap.projectFuncWithOption(removeNoMapping)
	projectAsciiAlphabet := asciiAlphabetMap.projectFuncWithOption(removeNoMapping)
	projectAsciiSign := asciiSignMap.projectFuncWithOption(removeNoMapping)
	projectSpace := spaceMap.projectFuncWithOption(removeNoMapping)

	var replacerMapping []string
	for _, f := range flags {
		switch f {
		case HankakuNumberToZenkaku:
			replacerMapping = append(replacerMapping, projectAsciiNumber(0, 1)...)
		case ZenkakuNumberToHankaku:
			replacerMapping = append(replacerMapping, projectAsciiNumber(1, 0)...)
		case HankakuKatakanaToZenkaku:
			replacerMapping = append(replacerMapping, projectKana(0, 1+noDakutenSlide)...)
		case ZenkakuKatakanaToHankaku:
			replacerMapping = append(replacerMapping, projectKana(1, 0+noDakutenSlide)...)
		case KatakanaToHiragana:
			replacerMapping = append(replacerMapping, projectKana(0, 2+noDakutenSlide)...)
			replacerMapping = append(replacerMapping, projectKana(1, 2+noDakutenSlide)...)
		case HiraganaToZenkakuKatakana:
			replacerMapping = append(replacerMapping, projectKana(2, 1+noDakutenSlide)...)
		case HiraganaToHankakuKatakana:
			replacerMapping = append(replacerMapping, projectKana(2, 0+noDakutenSlide)...)
		case KanaToHankakuKatakana:
			replacerMapping = append(replacerMapping, projectKana(0, 0+noDakutenSlide)...)
			replacerMapping = append(replacerMapping, projectKana(1, 0+noDakutenSlide)...)
			replacerMapping = append(replacerMapping, projectKana(2, 0+noDakutenSlide)...)
		case KanaToZenkakuKatakana:
			replacerMapping = append(replacerMapping, projectKana(0, 1+noDakutenSlide)...)
			replacerMapping = append(replacerMapping, projectKana(1, 1+noDakutenSlide)...)
			replacerMapping = append(replacerMapping, projectKana(2, 1+noDakutenSlide)...)
		case KanaToHiragana:
			replacerMapping = append(replacerMapping, projectKana(0, 2+noDakutenSlide)...)
			replacerMapping = append(replacerMapping, projectKana(1, 2+noDakutenSlide)...)
			replacerMapping = append(replacerMapping, projectKana(2, 2+noDakutenSlide)...)
		case AlphabetToUpperZenkaku:
			replacerMapping = append(replacerMapping, projectAsciiAlphabet(0, 3)...)
			replacerMapping = append(replacerMapping, projectAsciiAlphabet(1, 3)...)
			replacerMapping = append(replacerMapping, projectAsciiAlphabet(2, 3)...)
		case AlphabetToUpperHankaku:
			replacerMapping = append(replacerMapping, projectAsciiAlphabet(0, 1)...)
			replacerMapping = append(replacerMapping, projectAsciiAlphabet(2, 1)...)
			replacerMapping = append(replacerMapping, projectAsciiAlphabet(3, 1)...)
		case AlphabetToLowerZenkaku:
			replacerMapping = append(replacerMapping, projectAsciiAlphabet(0, 2)...)
			replacerMapping = append(replacerMapping, projectAsciiAlphabet(1, 2)...)
			replacerMapping = append(replacerMapping, projectAsciiAlphabet(3, 2)...)
		case AlphabetToLowerHankaku:
			replacerMapping = append(replacerMapping, projectAsciiAlphabet(1, 0)...)
			replacerMapping = append(replacerMapping, projectAsciiAlphabet(2, 0)...)
			replacerMapping = append(replacerMapping, projectAsciiAlphabet(3, 0)...)
		case AlphabetToZenkaku:
			replacerMapping = append(replacerMapping, projectAsciiAlphabet(0, 2)...)
			replacerMapping = append(replacerMapping, projectAsciiAlphabet(1, 3)...)
		case AlphabetToHankaku:
			replacerMapping = append(replacerMapping, projectAsciiAlphabet(2, 0)...)
			replacerMapping = append(replacerMapping, projectAsciiAlphabet(3, 1)...)
		case HankakuSignToZenkaku:
			replacerMapping = append(replacerMapping, projectAsciiSign(0, 1)...)
		case ZenkakuSignToHankaku:
			replacerMapping = append(replacerMapping, projectAsciiSign(1, 0)...)
		case HankakuSpaceToZenkaku:
			replacerMapping = append(replacerMapping, projectSpace(0, 1)...)
		case ZenkakuSpaceToHankaku:
			replacerMapping = append(replacerMapping, projectSpace(1, 0)...)
		}
	}
	// 同じマッピング元が複数回定義される場合、最初に定義されたものが使用される
	// eg. [[".", "。"], [".", "．"]]
	// というマッピングが渡されたとき、"hoge." は "hoge。" になる
	return strings.NewReplacer(replacerMapping...)
}

type charMap struct {
	data [][]string
}

func (c *charMap) projectFuncWithOption(removeUnknown bool) func(int, int) []string {
	return func(from, to int) []string {
		return c.project(from, to, removeUnknown)
	}
}

func (c *charMap) project(from, to int, removeUnknown bool) []string {
	if len(c.data) == 0 {
		return nil
	}
	var result []string
	for _, cols := range c.data {
		v1 := cols[from]
		v2 := cols[to]
		if v1 == "" {
			continue
		}
		if v2 == "" && !removeUnknown {
			v2 = v1
		}
		result = append(result, v1, v2)
	}
	return result
}

var asciiNumberMap = charMap{
	data: [][]string{
		// 半角数字, 全角数字
		{"0", "０"},
		{"1", "１"},
		{"2", "２"},
		{"3", "３"},
		{"4", "４"},
		{"5", "５"},
		{"6", "６"},
		{"7", "７"},
		{"8", "８"},
		{"9", "９"},
	},
}

var kanaMap = charMap{
	data: [][]string{
		// 半角カタカナ, 全角カタカナ, 全角ひらがな、濁点なし半角カタカナ、濁点なし全角カタカナ、濁点なし全角ひらがな
		// マッピングは上にあるものから順に適用される
		{"ｶﾞ", "ガ", "が", "ｶ", "カ", "か"},
		{"ｷﾞ", "ギ", "ぎ", "ｷ", "キ", "き"},
		{"ｸﾞ", "グ", "ぐ", "ｸ", "ク", "く"},
		{"ｹﾞ", "ゲ", "げ", "ｹ", "ケ", "け"},
		{"ｺﾞ", "ゴ", "ご", "ｺ", "コ", "こ"},
		{"ｻﾞ", "ザ", "ざ", "ｻ", "サ", "さ"},
		{"ｼﾞ", "ジ", "じ", "ｼ", "シ", "し"},
		{"ｽﾞ", "ズ", "ず", "ｽ", "ス", "す"},
		{"ｾﾞ", "ゼ", "ぜ", "ｾ", "セ", "せ"},
		{"ｿﾞ", "ゾ", "ぞ", "ｿ", "ソ", "そ"},
		{"ﾀﾞ", "ダ", "だ", "ﾀ", "タ", "た"},
		{"ﾁﾞ", "ヂ", "ぢ", "ﾁ", "チ", "ち"},
		{"ﾂﾞ", "ヅ", "づ", "ﾂ", "ツ", "つ"},
		{"ﾃﾞ", "デ", "で", "ﾃ", "テ", "て"},
		{"ﾄﾞ", "ド", "ど", "ﾄ", "ト", "と"},
		{"ﾊﾞ", "バ", "ば", "ﾊ", "ハ", "は"},
		{"ﾋﾞ", "ビ", "び", "ﾋ", "ヒ", "ひ"},
		{"ﾌﾞ", "ブ", "ぶ", "ﾌ", "フ", "ふ"},
		{"ﾍﾞ", "ベ", "べ", "ﾍ", "へ", "へ"},
		{"ﾎﾞ", "ボ", "ぼ", "ﾎ", "ホ", "ほ"},
		{"ﾊﾟ", "パ", "ぱ", "ﾊ", "ハ", "は"},
		{"ﾋﾟ", "ピ", "ぴ", "ﾋ", "ヒ", "ひ"},
		{"ﾌﾟ", "プ", "ぷ", "ﾌ", "フ", "ふ"},
		{"ﾍﾟ", "ペ", "ぺ", "ﾍ", "ヘ", "へ"},
		{"ﾎﾟ", "ポ", "ぽ", "ﾎ", "ホ", "ほ"},
		{"ｳﾞ", "ヴ", "ゔ", "ｳ", "ウ", "う"},
		{"ﾜﾞ", "ヷ", "", "ﾜ", "ワ", "わ"},
		{"ｦﾞ", "ヺ", "", "ｦ", "ヲ", "を"},
		{"ｱ", "ア", "あ", "ｱ", "ア", "あ"},
		{"ｲ", "イ", "い", "ｲ", "イ", "い"},
		{"ｳ", "ウ", "う", "ｳ", "ウ", "う"},
		{"ｴ", "エ", "え", "ｴ", "エ", "え"},
		{"ｵ", "オ", "お", "ｵ", "オ", "お"},
		{"ｶ", "カ", "か", "ｶ", "カ", "か"},
		{"ｷ", "キ", "き", "ｷ", "キ", "き"},
		{"ｸ", "ク", "く", "ｸ", "ク", "く"},
		{"ｹ", "ケ", "け", "ｹ", "ケ", "け"},
		{"ｺ", "コ", "こ", "ｺ", "コ", "こ"},
		{"ｻ", "サ", "さ", "ｻ", "サ", "さ"},
		{"ｼ", "シ", "し", "ｼ", "シ", "し"},
		{"ｽ", "ス", "す", "ｽ", "ス", "す"},
		{"ｾ", "セ", "せ", "ｾ", "セ", "せ"},
		{"ｿ", "ソ", "そ", "ｿ", "ソ", "そ"},
		{"ﾀ", "タ", "た", "ﾀ", "タ", "た"},
		{"ﾁ", "チ", "ち", "ﾁ", "チ", "ち"},
		{"ﾂ", "ツ", "つ", "ﾂ", "ツ", "つ"},
		{"ﾃ", "テ", "て", "ﾃ", "テ", "て"},
		{"ﾄ", "ト", "と", "ﾄ", "ト", "と"},
		{"ﾅ", "ナ", "な", "ﾅ", "ナ", "な"},
		{"ﾆ", "ニ", "に", "ﾆ", "ニ", "に"},
		{"ﾇ", "ヌ", "ぬ", "ﾇ", "ヌ", "ぬ"},
		{"ﾈ", "ネ", "ね", "ﾈ", "ネ", "ね"},
		{"ﾉ", "ノ", "の", "ﾉ", "ノ", "の"},
		{"ﾊ", "ハ", "は", "ﾊ", "ハ", "は"},
		{"ﾋ", "ヒ", "ひ", "ﾋ", "ヒ", "ひ"},
		{"ﾌ", "フ", "ふ", "ﾌ", "フ", "ふ"},
		{"ﾍ", "ヘ", "へ", "ﾍ", "ヘ", "へ"},
		{"ﾎ", "ホ", "ほ", "ﾎ", "ホ", "ほ"},
		{"ﾏ", "マ", "ま", "ﾏ", "マ", "ま"},
		{"ﾐ", "ミ", "み", "ﾐ", "ミ", "み"},
		{"ﾑ", "ム", "む", "ﾑ", "ム", "む"},
		{"ﾒ", "メ", "め", "ﾒ", "メ", "め"},
		{"ﾓ", "モ", "も", "ﾓ", "モ", "も"},
		{"ﾔ", "ヤ", "や", "ﾔ", "ヤ", "や"},
		{"ﾕ", "ユ", "ゆ", "ﾕ", "ユ", "ゆ"},
		{"ﾖ", "ヨ", "よ", "ﾖ", "ヨ", "よ"},
		{"ﾗ", "ラ", "ら", "ﾗ", "ラ", "ら"},
		{"ﾘ", "リ", "り", "ﾘ", "リ", "り"},
		{"ﾙ", "ル", "る", "ﾙ", "ル", "る"},
		{"ﾚ", "レ", "れ", "ﾚ", "レ", "れ"},
		{"ﾛ", "ロ", "ろ", "ﾛ", "ロ", "ろ"},
		{"ﾜ", "ワ", "わ", "ﾜ", "ワ", "わ"},
		{"ｦ", "ヲ", "を", "ｦ", "ヲ", "を"},
		{"ﾝ", "ン", "ん", "ﾝ", "ン", "ん"},
		{"ｧ", "ァ", "ぁ", "ｧ", "ァ", "ぁ"},
		{"ｨ", "ィ", "ぃ", "ｨ", "ィ", "ぃ"},
		{"ｩ", "ゥ", "ぅ", "ｩ", "ゥ", "ぅ"},
		{"ｪ", "ェ", "ぇ", "ｪ", "ェ", "ぇ"},
		{"ｫ", "ォ", "ぉ", "ｫ", "ォ", "ぉ"},
		{"ｯ", "ッ", "っ", "ｯ", "ッ", "っ"},
		{"ｬ", "ャ", "ゃ", "ｬ", "ャ", "ゃ"},
		{"ｭ", "ュ", "ゅ", "ｭ", "ュ", "ゅ"},
		{"ｮ", "ョ", "ょ", "ｮ", "ョ", "ょ"},
	},
}

var asciiAlphabetMap = charMap{
	data: [][]string{
		// 半角小文字, 半角大文字, 全角小文字, 全角大文字
		{"a", "A", "ａ", "Ａ"},
		{"b", "B", "ｂ", "Ｂ"},
		{"c", "C", "ｃ", "Ｃ"},
		{"d", "D", "ｄ", "Ｄ"},
		{"e", "E", "ｅ", "Ｅ"},
		{"f", "F", "ｆ", "Ｆ"},
		{"g", "G", "ｇ", "Ｇ"},
		{"h", "H", "ｈ", "Ｈ"},
		{"i", "I", "ｉ", "Ｉ"},
		{"j", "J", "ｊ", "Ｊ"},
		{"k", "K", "ｋ", "Ｋ"},
		{"l", "L", "ｌ", "Ｌ"},
		{"m", "M", "ｍ", "Ｍ"},
		{"n", "N", "ｎ", "Ｎ"},
		{"o", "O", "ｏ", "Ｏ"},
		{"p", "P", "ｐ", "Ｐ"},
		{"q", "Q", "ｑ", "Ｑ"},
		{"r", "R", "ｒ", "Ｒ"},
		{"s", "S", "ｓ", "Ｓ"},
		{"t", "T", "ｔ", "Ｔ"},
		{"u", "U", "ｕ", "Ｕ"},
		{"v", "V", "ｖ", "Ｖ"},
		{"w", "W", "ｗ", "Ｗ"},
		{"x", "X", "ｘ", "Ｘ"},
		{"y", "Y", "ｙ", "Ｙ"},
		{"z", "Z", "ｚ", "Ｚ"},
	},
}

var spaceMap = charMap{
	data: [][]string{
		// 半角スペース, 全角スペース
		{" ", "　"},
	},
}

var asciiSignMap = charMap{
	data: [][]string{
		// 半角記号, 全角記号
		{"!", "！"},
		{`"`, `”`},
		{"#", "＃"},
		{"$", "＄"},
		{"%", "％"},
		{"&", "＆"},
		{"'", "’"},
		{"'", "‘"},
		{"(", "（"},
		{")", "）"},
		{"*", "＊"},
		{"+", "＋"},
		{",", "，"},
		{",", "、"},
		{"-", "－"},
		{"-", "ー"},
		{".", "．"},
		{"。", "．"},
		{"/", "／"},
		{"/", "・"},
		{":", "："},
		{";", "；"},
		{"<", "＜"},
		{"=", "＝"},
		{">", "＞"},
		{"?", "？"},
		{"@", "＠"},
		{"[", "［"},
		{"\\", "＼"},
		{"\\", "￥"},
		{"]", "］"},
		{"^", "＾"},
		{"_", "＿"},
		{"`", "｀"},
		{"{", "｛"},
		{"|", "｜"},
		{"}", "｝"},
		{"~", "～"},
		{"~", "￣"},
	},
}
