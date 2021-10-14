package test

import (
	"fmt"
	"strings"
	"testing"

	stringUtils "github.com/nahuelmarianolosada/test_project/stringutils"
)

var helloWorld = " Hello, world "

func showSuccessTest() {
	fmt.Println("\t\t[FINALIZADO]")
}

// go test -cover -v
func TestStringTrimmer(t *testing.T) {
	fmt.Print("Probando Trimmer: \t\t\t")
	if len(stringUtils.Trimmer(helloWorld)) > 12 {
		t.Errorf("No se quitaron los espacios : [%s]", stringUtils.Trimmer(helloWorld))
	}
	showSuccessTest()
}

func TestReverse(t *testing.T) {
	fmt.Print("Probando Inverso de String: ")
	if strings.Compare(stringUtils.Reverse(helloWorld), " dlrow ,olleH ") != 0 {
		t.Errorf("Reverse no coincidente : [%s] != [%s]", stringUtils.Reverse(helloWorld), "olleH")
	}
	showSuccessTest()
}

func TestSubStr(t *testing.T) {
	fmt.Print("Probando Substring: \t\t")
	expected := "Hello"
	if strings.Compare(stringUtils.SubStr(helloWorld, 1, 6), expected) != 0 {
		t.Errorf("SubString no coincidente : [%s] != [%s]", expected, stringUtils.SubStr(helloWorld, 1, 6))
	}

	expected = "world"
	if strings.Compare(expected, stringUtils.Trimmer(stringUtils.SubStrWithoutEnd(helloWorld, 8))) != 0 {
		t.Errorf("SubString sin final no coincidente : [%s] != [%s]", expected, stringUtils.SubStrWithoutEnd(helloWorld, 8))
	}
	showSuccessTest()
}

func TestReplace(t *testing.T) {
	fmt.Print("Probando Replace: \t\t\t")
	expected := "Hello, mars"
	if strings.Compare(stringUtils.Trimmer(stringUtils.Replace(helloWorld, "world", "mars", 1)), expected) != 0 {
		t.Errorf("Replace no coincidente : [%s] != [%s]", expected, stringUtils.Trimmer(stringUtils.Replace(helloWorld, "world", "mars", 1)))
	}
	showSuccessTest()
}

func TestCapitalizer(t *testing.T) {
	fmt.Print("Probando Capitalizador: \t")
	expected := "Hello, World"
	if strings.Compare(stringUtils.Capitalize(stringUtils.Trimmer(helloWorld)), expected) != 0 {
		t.Errorf("Error al capitalizar frase : [%s] != [%s]", expected, stringUtils.Capitalize(stringUtils.Trimmer(helloWorld)))
	}

	if strings.Compare(stringUtils.CapitalizeAll(stringUtils.Trimmer(helloWorld)), "HELLO, WORLD") != 0 {
		t.Errorf("Error al capitalizar frase : [%s] != [%s]", expected, stringUtils.Capitalize(stringUtils.Trimmer(helloWorld)))
	}

	showSuccessTest()
}

func TestIndent(t *testing.T) {
	/*stringUtils.IndentRule(`((Modevent.ElementType == "QUE" && ((Question.Status == "ANSWERED" || Question.Status == "UNANSWERED") || ((((Question.Status == "CLOSED_UNANSWERED" || Question.Status == "DELETED") || (Question.Status == "DISABLED" || Question.Status == "BANNED")) || Question.Status == "UNDER_REVIEW") && Question.Answer.Status == "ACTIVE"))) && Item.HasSalesTermAttribute == true)`)
	fmt.Println("········································································")
	stringUtils.IndentRule(`((Denouncer.Positive > 4 && Denouncer.PositiveTrueAvg < 6) && (Denouncer.Random > 79 && Modevent.ElementType == "QUE"))`)
	fmt.Println("········································································")
	stringUtils.IndentRule(`((Modevent.ElementType == \"QUE\" && ((Question.Status == \"ANSWERED\" || Question.Status == \"UNANSWERED\") || ((((Question.Status == \"CLOSED_UNANSWERED\" || Question.Status == \"DELETED\") || (Question.Status == \"DISABLED\" || Question.Status == \"BANNED\")) || Question.Status == \"UNDER_REVIEW\" ) && Question.Answer.Status == \"ACTIVE\" ))) && ((((User.IsBrand == false && User.IsCBT == false) && Item.IsDeal == false) && User.IsMLeader == false ) && User.Points < 501))`)*/
	/*fmt.Println("································· 34073 ·······································")
	stringUtils.IndentRule(`((Modevent.ElementType == "QUE" && ((Question.Status == "ANSWERED" || Question.Status == "UNANSWERED") || ((((Question.Status == "CLOSED_UNANSWERED" || Question.Status == "DELETED") || (Question.Status == "DISABLED" || Question.Status == "BANNED")) || Question.Status == "UNDER_REVIEW") && Question.Answer.Status == "ACTIVE"))) && (User.IsBrand == true || User.Points > 10000) && (Question.InfractionSection == "QUE"))`)
	fmt.Println("···································· 34074 ····································")
	stringUtils.IndentRule(`((Modevent.ElementType == "QUE" && ((Question.Status == "ANSWERED" || Question.Status == "UNANSWERED") || ((((Question.Status == "CLOSED_UNANSWERED" || Question.Status == "DELETED") || (Question.Status == "DISABLED" || Question.Status == "BANNED")) || Question.Status == "UNDER_REVIEW") && Question.Answer.Status == "ACTIVE"))) && (User.IsBrand == true || User.Points > 10000) && (Question.InfractionSection == "ANS" || Question.InfractionSection == "BOTH"))`)*/
	fmt.Println("···································· 14290 ····································")
	stringUtils.IndentRule(`((((Modevent.ElementType == "QUE" && ((Question.Status == "ANSWERED" || Question.Status == "UNANSWERED") || ((((Question.Status == "CLOSED_UNANSWERED" || Question.Status == "DELETED") || (Question.Status == "DISABLED" || Question.Status == "BANNED")) || Question.Status == "UNDER_REVIEW") && Question.Answer.Status == "ACTIVE"))) && (((User.IsBrand == true || User.IsCBT == true) || Item.IsDeal == true) || User.Points > 10000)) && (Question.InfractionSection == "ANS" || Question.InfractionSection == "BOTH")) && ((Efficiency.Base < 85 || Efficiency.Total < 5) && (Efficiency.Base >= 5 || Efficiency.Total < 10)))`)
}
