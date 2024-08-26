class I18n < Formula
  desc "AI-Powered Internationalization Translation Tool"
  homepage "https://github.com/Martin2037/i18n-llm"
  url "https://github.com/Martin2037/i18n-llm/archive/refs/tags/v0.0.2.tar.gz"
  sha256 "d18b274d07781dadc8d6a76867a09017d7343df980e191068d69b9591f75811d"
  license "MIT"
  version "0.0.2"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", bin/"i18n"
  end

  test do
    assert_match "i18n version 0.0.2", shell_output("#{bin}/i18n --version")
  end
end