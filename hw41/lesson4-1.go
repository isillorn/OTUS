package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type freqDistr = map[string]int

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sortPairs(words map[string]int) PairList {

	pl := make(PairList, len(words))
	i := 0
	for k, v := range words {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func top10(p PairList) PairList {

	var res = PairList{}

	if len(p) > 10 {
		res = make(PairList, 10)
		copy(res, p[0:10])
	} else {
		res = make(PairList, len(p))
		copy(res, p[0:])
	}

	return res
}

// top 10 of frequent words in text
func freqAnalyze(text string) freqDistr {
	words := map[string]int{}

	spaceReplace := func(r rune) rune {
		switch r {

		case '.', ',', '!', '?', ':', ';', '-', '—', '\r', '\n', '\t':
			return ' '
		}
		return unicode.ToLower(r)
	}

	splits := strings.Split(strings.Map(spaceReplace, text), " ")

	for _, s := range splits {
		if s != "" {
			words[s]++
		}
	}
	return words
}

func main() {

	text := []string{"", "раз!,два; три3...",
		`
	Мороз и солнце; день чудесный!
	Еще ты дремлешь, друг прелестный —
	Пора, красавица, проснись:
	Открой сомкнуты негой взоры
	Навстречу северной Авроры,
	Звездою севера явись!
	
	Вечор, ты помнишь, вьюга злилась,
	На мутном небе мгла носилась;
	Луна, как бледное пятно,
	Сквозь тучи мрачные желтела,
	И ты печальная сидела —
	А нынче… погляди в окно:
	
	Под голубыми небесами
	Великолепными коврами,
	Блестя на солнце, снег лежит;
	Прозрачный лес один чернеет,
	И ель сквозь иней зеленеет,
	И речка подо льдом блестит.
	
	Вся комната янтарным блеском
	Озарена. Веселым треском
	Трещит затопленная печь.
	Приятно думать у лежанки.
	Но знаешь: не велеть ли в санки
	Кобылку бурую запречь?
	
	Скользя по утреннему снегу,
	Друг милый, предадимся бегу
	Нетерпеливого коня
	И навестим поля пустые,
	Леса, недавно столь густые,
	И берег, милый для меня.`,
		`
	我 独 自 一 人 走 到 大 路 上 ，
	一 条 石 子 路 在 雾 中 发 亮 。
	夜 很 静 。荒 原 面 对 着 太 空 ，
	星 星 与 星 星 互 诉 衷 肠 。
	天 空 是 多 么 庄 严 而 神 异 ！
	大 地 在 蓝 蓝 的 光 影 中 沉 睡 ……
	我 为 何 如 此 忧 伤 难 受 ？
	我 期 待 着 什 么 ？为 什 么 而 伤 悲 ？
	我 对 于 生 活 无 所 期 待 ，
	对 过 去 的 岁 月 毫 不 后 悔 。
	我 在 寻 求 自 由 和 宁 静 ！
	我 愿 忘 怀 一 切 而 入 睡 ！
	但 不 是 在 阴 冷 的 坟 墓 中 长 眠 ……
	我 希 望 永 远 是 这 样 的 睡 眠 ：
	要 胸 中 保 持 着 生 命 的 活 力 ，
	要 呼 吸 均 匀 ，气 息 和 缓 ；
	要 整 日 整 夜 能 够 听 到 
	悦 耳 的 声 音 歌 唱 爱 情 ，
	要 使 我 头 顶 上 茂 盛 的 栎 树 
	随 风 摇 动 ，终 岁 长 青 。
	`}

	for n, t := range text {
		fmt.Printf("***** TEXT %d *****\n", n+1)
		for _, v := range top10(sortPairs(freqAnalyze(t))) {
			fmt.Printf("%s = %d\n", v.Key, v.Value)
		}

	}

}
