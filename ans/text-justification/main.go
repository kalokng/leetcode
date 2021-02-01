package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	var lines []string
	n := maxWidth + 1
	var line []string
	for _, w := range words {
		l := len(w) + 1
		if l > n {
			lines = append(lines, fill(line, n, maxWidth))
			n = maxWidth + 1
			line = nil
		}
		n -= l
		line = append(line, w)
	}
	lines = append(lines, strings.Join(line, " ")+strings.Repeat(" ", n))

	return lines
}

func fill(words []string, n, maxWidth int) string {
	var b strings.Builder
	b.WriteString(words[0])
	words = words[1:]
	if len(words) == 0 {
		b.WriteString(strings.Repeat(" ", n))
		return b.String()
	}
	d, r := n/len(words)+1, n%len(words)
	s := strings.Repeat(" ", d)
	for i, w := range words {
		b.WriteString(s)
		if i < r {
			b.WriteRune(' ')
		}
		b.WriteString(w)
	}
	return b.String()
}

func main() {
	tc := []struct {
		words []string
		w     int
	}{
		{
			words: []string{"This", "is", "an", "example", "of", "text", "justification."},
			w:     16,
		},
		{
			words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			w:     10,
		},
		{
			words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			w:     11,
		},
		{
			words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			w:     12,
		},
		{
			words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			w:     13,
		},
		{
			words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			w:     14,
		},
		{
			words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			w:     15,
		},
		{
			words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			w:     16,
		},
		{
			words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			w:     17,
		},
		{
			words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			w:     18,
		},
		{
			words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			w:     19,
		},
		{
			words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			w:     20,
		},
		{
			words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we"},
			w:     20,
		},
		{
			words: []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			w:     16,
		},
		{
			words: []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			w:     14,
		},
	}
	for _, t := range tc {
		for _, l := range fullJustify(t.words, t.w) {
			fmt.Printf("%d: %q\n", len(l), l)
		}
		fmt.Println()
	}
}
