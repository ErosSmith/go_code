package beautify_brackets

import (
	"errors"
	// "fmt"
	"strings"
)

func getMatchingBracket(openBracket string) (string, error) {
	switch openBracket {
	case "{":
		return "}", nil
	case "[":
		return "]", nil
	case "(":
		return ")", nil
	default:
		return "", errors.New("Not an open bracket")

	}
}

func getBracketedText(text string) (int, error) {
	//checking required intial conditions
	// fmt.Println("inside getBrackedText")
	if len(text) < 1 {
		return -1, errors.New("Empty string")
	}

	first_char := text[0:1]

	open_brackets := "{[("

	if !strings.Contains(open_brackets, first_char) {
		return -1, errors.New("First character isn't open bracket")
	}
	//intial conditions correct
	matching_bracket, _ := getMatchingBracket(first_char)
	inbetween_bracket := true
	i_index := -1
	m_index := -1
	for inbetween_bracket {
		// fmt.Println(text)
		m_index = strings.Index(text, matching_bracket)
		if m_index == -1 {
			return -1, errors.New("No matching bracket in text")
		}
		i_index = strings.Index(text[1:], first_char)
		i_index = i_index + 1
		//no in between bracket in text
		if i_index == 0 {
			inbetween_bracket = false
			break
		}
		//in between bracket in text remove both brackets
		if i_index < m_index {
			text = text[:i_index] + "*" + text[i_index+1:m_index] + "*" + text[m_index+1:]
		} else {
			break
		}

	}
	// fmt.Println(text, m_index)
	return m_index, nil
}

func indentLevel(text string) string {
	// fmt.Println("inside indent")
	n_split := strings.Split(text, "\n")

	for i, s := range n_split {
		n_split[i] = "\t" + s
	}
	return strings.Join(n_split, "\n")
}

func beautifyBrackededText(text string) string {
	// fmt.Println(text, "in beautify")

	if len(text) < 3 {
		return text
	}

	open_brackets := "{[("
	first_bracket_index := strings.IndexAny(text, open_brackets)
	if !(first_bracket_index > -1) {
		return text
	}

	closing_bracket_index, err := getBracketedText(text[first_bracket_index:])

	if err != nil {
		return ""
	}

	pre_bracket_text := text[:first_bracket_index]
	// fmt.Println("pre_bracket_text", pre_bracket_text)
	bracketed_text := text[first_bracket_index : closing_bracket_index+1+len(pre_bracket_text)]
	// fmt.Println("bracketed_text", bracketed_text)
	remainder := text[closing_bracket_index+1+len(pre_bracket_text) : len(text)]
	// fmt.Println("remainder", remainder)

	inside := beautifyBrackededText(bracketed_text[1 : len(bracketed_text)-1])
	fixed_text := bracketed_text[0:1] + "\n" + indentLevel(inside) + "\n" + bracketed_text[len(bracketed_text)-1:len(bracketed_text)]

	remainder = beautifyBrackededText(remainder)
	output := pre_bracket_text + fixed_text + remainder

	// fmt.Println("output", output)

	return output
}
