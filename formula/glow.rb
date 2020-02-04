# This file was generated by GoReleaser. DO NOT EDIT.
class Glow < Formula
  desc ""
  homepage ""
  version "4.1.0"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/meinto/glow/releases/download/v4.1.0/glow_4.1.0_darwin_x86_64.tar.gz"
    sha256 "3fa6fffb0a506563a400d8c9d16cd035f4bf577d1b9c9ee4bab63dc4371ad8d6"
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://github.com/meinto/glow/releases/download/v4.1.0/glow_4.1.0_linux_x86_64.tar.gz"
      sha256 "a42356ce19feec39f3254fb76a49a0160abce63e312e854fd42db357408ef465"
    end
  end
  
  depends_on "git"

  def install
    bin.install "glow"
  end
end
