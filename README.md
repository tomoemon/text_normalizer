Text Normalizer
-------------------

- ひらがな、全角カタカナ、半角カタカナ
- 全角数字、半角数字
- 全角英字、半角英字
- 全角記号、半角記号

上記、文字集合内で任意の方向に対する変換を自由に組み合わせて文字列を正規化するためのライブラリです。

一言で言うと、PHPの [mb_convert_kana](http://php.net/manual/ja/function.mb-convert-kana.php) のGo言語実装です。
まったく同じ仕様ではありませんが、ほぼ同様のことが実現可能です。


Usage
---------

任意の組み合わせで下記の変換オプションを指定して `NewTextNormalizer` を呼び出すと、`string.Replacer` インスタンスが返るので、それを使って任意の文字列を正規化できます。

実行例は [テストコード](/text_normalizer_test.go) をご覧ください。

共通オプション

	// 濁点なしの文字に変換する（カナ系文字への変換オプションと同時に使う）
	RemoveDakuten

	// 変換先の文字が存在しない場合に削除する（デフォルト：無視する）
	RemoveNoMapping

個別の変換オプション

	// 半角数字→全角数字
	HankakuNumberToZenkaku

	// 全角数字→半角数字
	ZenkakuNumberToHankaku

	// 半角カタカナ→全角カタカナ
	HankakuKatakanaToZenkaku

	// 全角カタカナ→半角カタカナ
	ZenkakuKatakanaToHankaku

	// (全角カタカナ、半角カタカナ）→ひらがな
	KatakanaToHiragana

	// ひらがな→全角カタカナ
	HiraganaToZenkakuKatakana

	// ひらがな→半角カタカナ
	HiraganaToHankakuKatakana

	// (ひらがな、全角カタカナ、半角カタカナ) →ひらがな
	KanaToHiragana

	// (ひらがな、全角カタカナ、半角カタカナ) →全角カタカナ
	KanaToZenkakuKatakana

	// (ひらがな、全角カタカナ、半角カタカナ) →半角カタカナ
	KanaToHankakuKatakana

	// (半角英字、全角英字) →全角大文字英字
	AlphabetToUpperZenkaku

	// (半角英字、全角英字) →半角大文字英字
	AlphabetToUpperHankaku

	// (半角英字、全角英字) →全角小文字英字
	AlphabetToLowerZenkaku

	// (半角英字、全角英字) →半角小文字英字
	AlphabetToLowerHankaku

	// (半角英字、全角英字) →全角英字（大文字小文字の区別は維持する）
	AlphabetToZenkaku

	// (半角英字、全角英字) →半角英字（大文字小文字の区別は維持する）
	AlphabetToHankaku

	// 半角記号→全角記号
	HankakuSignToZenkaku

	// 全角記号→半角記号
	ZenkakuSignToHankaku

	// 半角スペース→全角スペース
	HankakuSpaceToZenkaku

	// 全角スペース→半角スペース
	ZenkakuSpaceToHankaku



Comparison
------------

文字列の正規化を行う際に、Unicode のコードポイントの範囲で文字集合を表現し、変換元から変換先へ射影する方法もありますが、このライブラリではナイーブに変換用の文字集合テーブルを定義しています。

コードポイントベースの方が、フットプリントを小さくできますが、コードの理解とデバッグのしやすさ、拡張のしやすさを考えて、文字集合テーブルの方式を取っています。