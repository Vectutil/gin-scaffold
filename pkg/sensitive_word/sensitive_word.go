package sensitive_word

import (
	"errors"
	"unicode"
)

// TrieNode Define the nodes of DFA
type TrieNode struct {
	Children map[rune]*TrieNode
	IsEnd    bool
}

// NewTrieNode Create a new Trie node
func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make(map[rune]*TrieNode),
		IsEnd:    false,
	}
}

// DFA The structure contains the root node of the DFA
type DFA struct {
	Root *TrieNode
}

// NewDFA Initialize a new DFA
func NewDFA() *DFA {
	return &DFA{
		Root: NewTrieNode(),
	}
}

// AddWord Add sensitive words to DFA
func (d *DFA) AddWord(word string) {
	node := d.Root
	for _, char := range word {
		if _, exists := node.Children[char]; !exists {
			node.Children[char] = NewTrieNode()
		}
		node = node.Children[char]
	}
	node.IsEnd = true
}

// UpdateOldWord update old word
func (d *DFA) UpdateOldWord(oldWord, newWord string) {
	d.DeleteWord(oldWord)
	d.AddWord(newWord)
}

// DeleteWord delete word
func (d *DFA) DeleteWord(word string) bool {
	result := []rune(word)
	// 辅助函数用于递归删除节点
	var deleteNode func(node *TrieNode, index int) bool
	deleteNode = func(node *TrieNode, index int) bool {
		if index == len(result) {
			// 如果该词不存在，直接返回
			if !node.IsEnd {
				return false
			}
			// 清除该词的结束标记
			node.IsEnd = false
			// 如果该节点没有子节点，可以删除
			return len(node.Children) == 0
		}

		char := result[index]
		child, exists := node.Children[char]
		if !exists {
			return false // 如果路径不存在，则不做任何操作
		}

		// 递归删除子节点
		shouldDeleteChild := deleteNode(child, index+1)
		if shouldDeleteChild {
			// 删除当前节点的子节点
			delete(node.Children, char)
			// 如果当前节点没有其他子节点且不是词尾节点，返回 true
			return len(node.Children) == 0 && !node.IsEnd
		}
		return false
	}

	// 调用递归函数删除指定的词
	return deleteNode(d.Root, 0)
}

// Filter the input text and replace sensitive words
func (d *DFA) Filter(text string, isPreprocessText bool) string {
	if isPreprocessText {
		text = d.PreprocessText(text)
	}
	result := []rune(text)
	for i := 0; i < len(result); i++ {
		node := d.Root
		j := i
		for j < len(result) {
			if nextNode, exists := node.Children[result[j]]; exists {
				node = nextNode
				if node.IsEnd {
					for k := i; k <= j; k++ {
						result[k] = '*'
					}
				}
				j++
			} else {
				break
			}
		}
	}
	return string(result)
}

// Check  if the input text contains sensitive words
func (d *DFA) Check(text string, isPreprocessText bool) error {
	if isPreprocessText {
		text = d.PreprocessText(text)
	}
	result := []rune(text)
	for i := 0; i < len(result); {
		node := d.Root
		start := i
		matched := false
		for j := i; j < len(result); j++ {
			char := result[j]
			if nextNode, exists := node.Children[char]; exists {
				node = nextNode
				if node.IsEnd {
					matched = true
					return errors.New("包含敏感词: " + string(result[start:j+1]))
				}
			} else {
				break
			}
		}
		if !matched {
			i++
		}
	}
	return nil
}

// PreprocessText Preprocess text and remove invalid characters
func (d *DFA) PreprocessText(text string) string {
	var result []rune
	for _, r := range text {
		if d.isValidChar(r) {
			result = append(result, r)
		}
	}
	return string(result)
}

// IsValidChar Check if the characters are valid
func (d *DFA) isValidChar(r rune) bool {
	// 允许字母、数字
	if unicode.IsLetter(r) || unicode.IsNumber(r) {
		return true
	}
	// 允许汉字
	if unicode.In(r, unicode.Han) {
		return true
	}
	// 可选：允许某些标点符号或其他特殊字符
	// 这里我们使用 unicode.IsPunct 来允许所有的标点符号
	//if unicode.IsPunct(r) {
	//	return true
	//}
	// 默认情况下不允许其他字符
	return false
}
