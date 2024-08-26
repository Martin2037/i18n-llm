// main.go
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms/openai"
)

var (
	i18nDir     string
	sourceDir   string
	targetLangs []string
	allLangs    = []string{
		"en", "es", "fr", "de", "it", "pt", "ru", "zh-CN", "zh-TW", "ja", "ko",
		"ar", "hi", "nl", "pl", "tr", "sv", "da", "fi", "no", "id", "th", "vi",
	}
	llmInstance *openai.LLM
)

var rootCmd = &cobra.Command{
	Use:   "i18n-translator",
	Short: "A CLI tool for translating i18n JSON files",
	Long:  `i18n-translator is a command-line tool that helps you translate JSON files for internationalization.`,
}

var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "Translate i18n JSON files",
	Run:   runTranslate,
}

func init() {
	translateCmd.Flags().StringVarP(&i18nDir, "dir", "d", ".", "i18n directory path")
	translateCmd.Flags().StringVarP(&sourceDir, "source", "s", "", "Source language directory")
	translateCmd.Flags().StringSliceVarP(&targetLangs, "target", "t", []string{}, "Target languages (default: all supported languages except source)")

	translateCmd.MarkFlagRequired("source")

	rootCmd.AddCommand(translateCmd)
}

func main() {
	var err error
	llmInstance, err = openai.New()
	if err != nil {
		fmt.Printf("Error initializing OpenAI: %v\n", err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runTranslate(cmd *cobra.Command, args []string) {
	if len(targetLangs) == 0 {
		targetLangs = getTargetLangs(sourceDir)
	}

	fmt.Printf("Translating from %s to %v\n", sourceDir, targetLangs)
	fmt.Printf("i18n directory: %s\n", i18nDir)

	if err := translateFiles(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func getTargetLangs(sourceDir string) []string {
	targets := make([]string, 0, len(allLangs)-1)
	for _, lang := range allLangs {
		if lang != sourceDir && !(sourceDir == "zh-CN" && lang == "zh-TW") && !(sourceDir == "zh-TW" && lang == "zh-CN") {
			targets = append(targets, lang)
		}
	}
	return targets
}
