package trie

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func New() *Trie {
	return &Trie{}
}

func (t *Trie) Insert(word string) {
	cur := t
	for _, ch := range word {
		ch -= 'a'
		if cur.children[ch] == nil {
			cur.children[ch] = &Trie{}
		}
		cur = cur.children[ch]
	}
	cur.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)

	return node != nil && node.isEnd
}

func (t *Trie) StartWith(prefix string) bool {
	node := t.searchPrefix(prefix)

	return node != nil
}

func (t *Trie) searchPrefix(prefix string) *Trie {
	cur := t
	for _, ch := range prefix {
		ch -= 'a'
		if cur.children[ch] == nil {
			return nil
		}
		cur = cur.children[ch]
	}

	return cur
}

func (t *Trie) Empty() bool {
	for i := 0; i < 26; i++ {
		if t.children[i] != nil {
			return false
		}
	}

	return true
}