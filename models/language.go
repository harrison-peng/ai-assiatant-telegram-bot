package models

type Language string

const (
	LanguageEnglish  Language = "en"
	LanguageMandarin Language = "zh"
)

// Validate validates the language.
func (l Language) Validate() bool {
	switch l {
	case LanguageEnglish, LanguageMandarin:
		return true
	default:
		return false
	}
}
