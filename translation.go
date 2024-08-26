// translator.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/schollz/progressbar/v3"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

var completedFiles int64
var totalFiles int
var mutex sync.Mutex

func translateFiles() error {
	sourceFiles, err := filepath.Glob(filepath.Join(i18nDir, sourceDir, "*.json"))
	if err != nil {
		return fmt.Errorf("failed to read source files: %v", err)
	}

	totalFiles = len(sourceFiles)
	bar := progressbar.Default(int64(totalFiles * 100)) // 100 steps per file for smoother progress

	for _, sourceFile := range sourceFiles {
		sourceContent, err := readJSONFile(sourceFile)
		if err != nil {
			return fmt.Errorf("failed to read source file %s: %v", sourceFile, err)
		}

		if err := translateContent(llmInstance, sourceContent, targetLangs, sourceFile, bar); err != nil {
			return fmt.Errorf("failed to translate: %v", err)
		}

		atomic.AddInt64(&completedFiles, 1)
	}

	// Ensure the progress bar reaches 100% only after all files are processed
	bar.Set(totalFiles * 100)

	fmt.Println("\nTranslation completed successfully!")
	return nil
}

func readJSONFile(filePath string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	return result, err
}

func writeJSONFile(filePath string, content map[string]interface{}) error {
	data, err := json.MarshalIndent(content, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}

func translateContent(llm *openai.LLM, content map[string]interface{}, targetLangs []string, sourceFile string, bar *progressbar.ProgressBar) error {
	inputJSON, err := json.Marshal(content)
	if err != nil {
		return fmt.Errorf("failed to marshal input content: %v", err)
	}

	prompt := fmt.Sprintf(`
Translate the following JSON from %s to %s. Please follow these guidelines:

1. The translated JSON must maintain the exact same structure, format, and key order as the original JSON.
2. Translate only the string values; ignore numbers, booleans, and null values.
3. Preserve the formatting within the strings (e.g., line breaks, special characters).
4. Keep placeholders (e.g., {name}, {{variable}}) unchanged and untranslated.
5. Translate each element in arrays individually, maintaining the array structure.
6. Apply these rules recursively to nested objects and arrays.
7. Ensure translations are simple, straightforward, and suitable for web display, with correct grammar.
8. Retain proper nouns and technical terms in their original form.
Input JSON:
%s

Output the translations in the following format:
{
"en": {translated JSON for English},
"ja": {translated JSON for Japanese},
...
}

Important:

1. The output must be a valid JSON object.
2. The structure, formatting, and key order of the translated JSON must exactly match the original.
3. Include translations for all specified target languages.
4. Ensure proper use of braces, brackets, quotes, and commas.
5. Do not add any text or comments outside the JSON structure.

If you encounter any issues or ambiguities, maintain the original content and structure.
`, sourceDir, strings.Join(targetLangs, ", "), string(inputJSON))

	ctx := context.Background()

	var streamedDataSize int64
	expectedSize := int64(len(inputJSON) * len(targetLangs) * 2) // Multiply by 2 to account for potential expansion during translation
	var lastProgress int64

	result, err := llm.Call(ctx, prompt,
		llms.WithJSONMode(),
		llms.WithModel("gpt-4o-mini"),
		llms.WithTemperature(0.1),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			mutex.Lock()
			streamedDataSize += int64(len(chunk))
			progress := int64(float64(streamedDataSize) / float64(expectedSize) * 90) // Use 90 out of 100 steps for streaming progress
			if progress > 90 {
				progress = 90
			}
			if progress > lastProgress {
				bar.Add(int(progress - lastProgress))
				lastProgress = progress
			}
			mutex.Unlock()

			return nil
		}))
	if err != nil {
		return fmt.Errorf("failed to call LLM: %v", err)
	}

	var translations map[string]map[string]interface{}
	err = json.Unmarshal([]byte(result), &translations)
	if err != nil {
		return fmt.Errorf("failed to unmarshal translated content: %v", err)
	}

	// Write translated content to files
	for lang, translatedContent := range translations {
		targetDir := filepath.Join(i18nDir, lang)
		if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", targetDir, err)
		}

		targetFile := filepath.Join(targetDir, filepath.Base(sourceFile))
		if err := writeJSONFile(targetFile, translatedContent); err != nil {
			return fmt.Errorf("failed to write translated file %s: %v", targetFile, err)
		}
	}

	// Update progress for file completion
	bar.Add(int(100 - lastProgress)) // Use remaining steps for file completion

	return nil
}
