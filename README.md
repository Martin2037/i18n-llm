# 🌐 i18n-llm

[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/i18n-llm)](https://goreportcard.com/report/github.com/yourusername/i18n-llm)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

> AI-Powered Intelligent Internationalization Translation Tool

i18n-llm is a revolutionary command-line tool that leverages advanced artificial intelligence to streamline your internationalization (i18n) workflow. Translate your JSON localization files into multiple languages with just a few simple commands, maintaining high quality and consistency.

[//]: # (![i18n-llm Demo]&#40;link_to_your_demo_gif_or_image&#41;)

## ✨ Key Features

- 🚀 **Fast and Efficient**: Translate multiple files to multiple languages with a single command
- 🧠 **AI-Driven**: Utilizes advanced language models to ensure translation quality
- 🔄 **Structure Preservation**: Perfectly maintains the original JSON structure, translating only the values
- 🌍 **Multi-Language Support**: Supports 20+ common languages
- 💡 **Intelligent Processing**: Automatically skips the source language to avoid unnecessary translations
- 🛠 **Flexible Configuration**: Customize source language, target languages, and directories

## 🚀 Quick Start

### Installation

```bash
go install github.com/yourusername/i18n-llm@latest
```

### Usage

1. Set your OpenAI API key:

```bash
export OPENAI_API_KEY=your_api_key_here
```

2. Run the translation command:

```bash
i18n-llm translate -s zh-CN
```

This will translate all JSON files in the `zh-CN` folder of your current directory into all supported languages.

### More Examples

Translate to specific languages:
```bash
i18n-llm translate -s en -t fr,de,es
```

Specify the i18n directory:
```bash
i18n-llm translate -d /path/to/i18n -s ja
```

[//]: # (## 📚 Detailed Documentation)

[//]: # ()
[//]: # (Check out our [full documentation]&#40;link_to_your_documentation&#41; for more advanced usage and configuration options.)

## 🤝 Contributing

We welcome contributions of all forms! Whether it's new features, bug fixes, or documentation improvements. Check out our [contribution guidelines](CONTRIBUTING.md) to get started.

## 📄 License

i18n-llm is released under the MIT License. See the [LICENSE](LICENSE) file for more details.

## 🙏 Acknowledgements

- Thanks to [OpenAI](https://openai.com/) for providing powerful language model support
- Thanks to [langchaingo](https://github.com/tmc/langchaingo) for the convenient Go language AI interaction interface

---

If you find i18n-llm useful, please consider giving it a star ⭐️!

[Back to Top](#i18n-llm)