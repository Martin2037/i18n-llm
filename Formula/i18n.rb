class I18n < Formula
  desc "AI-Powered Internationalization Translation Tool"
  homepage "https://github.com/Martin2037/i18n-llm"
  url "https://github.com/Martin2037/i18n-llm/archive/refs/tags/v0.0.3.tar.gz"
  sha256 "9514513fe42233eb5fde276cd50c6979498edfd86e61b8967c37b822c19d3d37"
  license "MIT"
  version "0.0.3"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", bin/"i18n"
  end

  test do
    assert_match "i18n version 0.0.3", shell_output("#{bin}/i18n --version")
  end
end