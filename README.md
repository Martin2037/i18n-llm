# 🚀 i18n-llm: Supercharge Your i18n Workflow!

[![Go Report Card](https://goreportcard.com/badge/github.com/Martin2037/i18n-llm)](https://goreportcard.com/report/github.com/Martin2037/i18n-llm)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

> Blazing Fast, AI-Powered Internationalization for Modern Web Projects

Tired of slow, inaccurate translations holding back your global reach? Say hello to i18n-llm, the game-changing CLI tool that's revolutionizing how developers handle internationalization.

## 🌟 Why i18n-llm Will Make You Love i18n Again

- ⚡️ **Lightning Fast**: Translate entire projects in minutes, not hours
- 🧠 **AI-Driven Accuracy**: Harness the power of advanced language models for context-aware translations
- 🔀 **Universal Compatibility**: Perfect for React, Next.js, Vue, or any project with JSON-based translations
- 🎯 **Pinpoint Precision**: Maintains your JSON structure flawlessly, only touching what needs translation
- 🌍 **Polyglot Power**: Support for 20+ languages at your fingertips
- 🛠 **Flexible Integration**: Works with your existing i18n setup, whether it's part of your app or standalone files

## 🚀 From Zero to Global in 60 Seconds

```bash
export OPENAI_API_KEY=your_api_key_here
i18n-llm translate -s en -t zh,es,fr -d ./src/locales
```

Boom! Your entire app just learned to speak Chinese, Spanish, and French. It's that simple.

## 🎭 Seamless Integration for All Projects

Whether you're working on:
- A cutting-edge React or Next.js application
- A Vue.js or Angular project
- Or managing standalone localization files

i18n-llm has got you covered. It adapts to your project structure, not the other way around.

## 📁 Flexible Directory Structure

i18n-llm works with various project layouts:

```
your_project/
│
├── src/
│   ├── locales/  (for integrated app translations)
│   │   ├── en/
│   │   │   ├── common.json
│   │   │   └── home.json
│   │   ├── zh/
│   │   └── ... (other language folders)
│
└── i18n/  (or any name for standalone translation files)
    ├── en/
    │   ├── messages.json
    │   └── errors.json
    ├── es/
    └── ... (other language folders)
```

As long as you have language-specific folders containing JSON files, i18n-llm will work its magic!

## 🚀 Installation

Get up and running in no time:

### Using Go

```bash
go install github.com/Martin2037/i18n-llm@latest
```

### Using Homebrew (macOS)

```bash
brew tap Martin2037/i18n-llm https://github.com/Martin2037/i18n-llm.git
brew update
brew install i18n
```

## 🔧 Usage That Just Makes Sense

1. Set your OpenAI API key:

```bash
export OPENAI_API_KEY=your_api_key_here
```

2. Navigate to your project directory.

3. Let the magic happen:

```bash
i18n-llm translate -s en -d ./path/to/your/locales
```

This command translates all JSON files from English to all supported languages. It's that easy!

### More Wizardry

Translate to specific languages:
```bash
i18n-llm translate -s en -t fr,de,es -d ./src/i18n
```

## 🚀 Real-World Example: Going Global in Minutes

1. Ensure your project's localization files are structured (as shown above).

2. Set your OpenAI API key:
   ```bash
   export OPENAI_API_KEY=your_api_key_here
   ```

3. Cast the translation spell:
   ```bash
   i18n-llm translate -s en -t zh,ja,ko,es,fr -d ./src/locales
   ```

   Watch as i18n-llm:
   - Uses English as the source
   - Conjures translations in Chinese, Japanese, Korean, Spanish, and French
   - Finds your source files in `./src/locales/en/`
   - Creates or updates files in the respective language folders

4. Marvel at your newly internationalized project!

## 🤝 Join the i18n Revolution

Excited? We are too! Check out our [contribution guidelines](CONTRIBUTING.md) and help make i18n-llm even more awesome.

## 📄 License

i18n-llm is released under the MIT License. See the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgements

- Powered by [OpenAI](https://openai.com/)'s cutting-edge language models
- Built with [langchaingo](https://github.com/tmc/langchaingo) for seamless AI interactions

---

Found i18n-llm helpful? Give us a star ⭐️ and spread the word!

[Beam me up!](#i18n-llm-supercharge-your-i18n-workflow)