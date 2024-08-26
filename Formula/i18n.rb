class I18n < Formula
  desc "AI-Powered Internationalization Translation Tool"
  homepage "https://github.com/Martin2037/i18n-llm"
  url "https://github.com/Martin2037/i18n-llm/archive/refs/tags/v0.0.1.tar.gz"
  sha256 "119ee836e5c722e0e1fe8ca6b4be6f8cf923adc6b97edd6c8334b8d795b3895f"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", bin/"i18n"
  end

  test do
    assert_match "i18n version 0.0.1", shell_output("#{bin}/i18n --version")
  end
end