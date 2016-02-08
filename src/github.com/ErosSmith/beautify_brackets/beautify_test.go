package beautify_brackets

import (
	"testing"
)

func TestGetMatchingBracket(t *testing.T) {
	s, err := getMatchingBracket("")
	correct := ""
	if err == nil {
		t.Errorf("error not working")
	}

	s, err = getMatchingBracket("{")
	if err != nil {
		t.Errorf("error not working")
	}
	correct = "}"
	if s != correct {
		t.Errorf("Closing bracket not returned. expected %s but got %s", correct, s)
	}

	s, err = getMatchingBracket("[")
	if err != nil {
		t.Errorf("error not working")
	}
	correct = "]"
	if s != correct {
		t.Errorf("Closing bracket not returned. expected %s but got %s", correct, s)
	}

	s, err = getMatchingBracket("(")
	if err != nil {
		t.Errorf("error not working")
	}
	correct = ")"
	if s != correct {
		t.Errorf("Closing bracket not returned. expected %s but got %s", correct, s)
	}

}

func TestGetBracketedText(t *testing.T) {
	s, err := getBracketedText("")

	if err != nil {
		t.Log(err.Error())
	} else {
		t.Error("Should have had an error")
	}
	if s != -1 {
		t.Errorf("weird")
	}

	text := "(     ( )  )   "
	s, err = getBracketedText(text)
	t.Log(s)
	t.Log(text[:s+1])
	if err != nil {
		t.Error(err.Error())
	}
	if s != 11 {
		t.Error("Wrong index")
	}

	text = "(  ( ( ))  ( )  )   "
	s, err = getBracketedText(text)
	t.Log(s)
	t.Log(text[:s+1])
	if err != nil {
		t.Error(err.Error())
	}
	if s != 16 {
		t.Error("Wrong index")
	}

}

func TestIndentLevel(t *testing.T) {
	text := "((\n ))"
	t.Log(text)
	i_text := indentLevel(text)
	t.Log(i_text)

}

func TestBeatifyBracketedText(t *testing.T) {
	text := "((()))"
	t.Log(text)
	i_text := beautifyBrackededText(text)
	t.Log(i_text)

	text = "[((()))(())]"
	t.Log(text)
	i_text = beautifyBrackededText(text)
	t.Log(i_text)

	text = "[((()))sup(hello(hi how are you))]"
	t.Log(text)
	i_text = beautifyBrackededText(text)
	t.Log(i_text)

}
