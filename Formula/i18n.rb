class I18n < Formula
  desc "AI-Powered Internationalization Translation Tool"
  homepage "https://github.com/Martin2037/i18n-llm"
  url "https://github.com/Martin2037/i18n-llm/archive/refs/tags/v0.0.1.tar.gz"
  sha256 "5fa9c08cccd78ae51bc4d71a75b947844ca2653bbd262801c628c3d52c01fba3"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", bin/"i18n"
  end

  test do
    assert_match "i18n version 0.0.1", shell_output("#{bin}/i18n --version")
  end
end