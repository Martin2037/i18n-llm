class I18n < Formula
  desc "AI-Powered Internationalization Translation Tool"
  homepage "https://github.com/Martin2037/i18n-llm"
  url "https://github.com/Martin2037/i18n-llm/archive/refs/tags/v0.0.1-brew.tar.gz"
  sha256 "e616635ddbfa6263472915979954526100a491bfd0a1c30e180df4790898cf36"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", bin/"i18n"
  end

  test do
    assert_match "i18n version 0.0.1", shell_output("#{bin}/i18n --version")
  end
end