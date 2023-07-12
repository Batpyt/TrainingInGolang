package tree

/*
请你实现 Trie 类：

Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。


示例：

输入
inputs = ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
inputs = [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出
[null, null, true, false, true, null, true]

解释
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/QC3q1f
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
树的节点结构，children数组存该节点下一个节点的字符在a-z中的下标（a：0，b：1。。。）
isEnd表示是否该单词结束
*/

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

/** Initialize your data structure here. */

func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */

func (t *Trie) Insert(word string) {
	root := t
	//遍历该词
	for _, ch := range word {
		ch -= 'a'
		//若没有对应字符的子节点，就创建一个
		if root.children[ch] == nil {
			root.children[ch] = &Trie{}
		}
		root = root.children[ch]
	}
	root.isEnd = true
}

/** Returns if the word is in the trie. */

func (t *Trie) Search(word string) bool {
	root := t.SearchLetterEnds(word)
	return root != nil && root.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchLetterEnds(prefix) != nil
}

//该方法用于搜索给定词的最后一个字符所在的节点

func (t *Trie) SearchLetterEnds(letter string) *Trie {
	root := t
	for _, ch := range letter {
		ch -= 'a'
		//如果不存在改词会返回nil
		if root.children[ch] == nil {
			return nil
		}
		root = root.children[ch]
	}
	return root
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
