# This file was generated by GoReleaser. DO NOT EDIT.
class Glow < Formula
  desc ""
  homepage ""
  url "https://github.com/meinto/glow/releases/download/v1.12.3/glow_1.12.3_darwin_x86_64.tar.gz"
  version "1.12.3"
  sha256 "fd1987e05cb6c22fa58f8eacbe408a00cfd34942309d022275b160301ea6fc90"
  
  depends_on "git"

  def install
    bin.install "glow"
  end
end
