package stringUtils

import (
	"fmt"
	"regexp"
	"strings"
)

func IndentRule(rule string) {
	//rule := `((Modevent.ElementType == "QUE" && ((Question.Status == "ANSWERED" || Question.Status == "UNANSWERED") || ((((Question.Status == "CLOSED_UNANSWERED" || Question.Status == "DELETED") || (Question.Status == "DISABLED" || Question.Status == "BANNED")) || Question.Status == "UNDER_REVIEW") && Question.Answer.Status == "ACTIVE"))) && Item.HasSalesTermAttribute == true)`
	//rule := `((Denouncer.Positive > 4 && Denouncer.PositiveTrueAvg < 6) && (Denouncer.Random > 79 && Modevent.ElementType == "QUE"))`

	re0 := regexp.MustCompile(` `)
	ruleTest0 := re0.ReplaceAllString(rule, "")

	re1 := regexp.MustCompile(`\(`)
	ruleTest := re1.ReplaceAllString(ruleTest0, "\n\t(")

	re2 := regexp.MustCompile(`\) \[|\] `)
	ruleTest2 := re2.ReplaceAllString(ruleTest, "\n\r)")

	re3 := regexp.MustCompile(`\)`)
	ruleTest3 := re3.ReplaceAllString(ruleTest2, ")\n\t") //Hasta aquí todo piola

	//Esta parte quiere indentar
	fields := strings.Split(ruleTest3, "\n")

	for i := range(fields) {
		if isTheSameLevel(fields[i]) {
			if i != 0 {
				tabCount := strings.Count(fields[i - 1], `	`)
				if isToAddTab(fields[i-1]) {
					tabCount++
				}
				re4 := regexp.MustCompile(`	`)
				field := re4.ReplaceAllString(fields[i], next(`	`, tabCount))
				fields[i] = field
			}
		}

		if isToAddTab(fields[i]) {
			tabCount := strings.Count(fields[i - 1], `	`)
			re4 := regexp.MustCompile(`	`)
			field := re4.ReplaceAllString(fields[i], next(`	`, tabCount + 1))
			fields[i] = field
		} else if isToRemoveTab(fields[i]) {
			tabCount := strings.Count(fields[i - 1], `	`)
			re4 := regexp.MustCompile(`	`)
			field := re4.ReplaceAllString(fields[i], next(`	`, tabCount - 1))
			fields[i] = field
		}

		fmt.Println(fields[i])
    }
}

func isTheSameLevel(field string) bool {
	if (strings.Contains(field, "(") && strings.Contains(field, ")")) || (
		!strings.Contains(field, "(") && !strings.Contains(field, ")")) { return true }
	return false
}

func isToAddTab(field string) bool {
	if strings.Contains(field, "(") && !strings.Contains(field, ")") { return true }
	return false
}

func isToRemoveTab(field string) bool {
	if !strings.Contains(field, "(") && strings.Contains(field, ")") { return true }
	return false
}

func next(word string , repetitions int) (final string) {
	for i := 0 ; i<repetitions; i++ {
		final += word
	}
	return final
}

