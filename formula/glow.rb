# This file was generated by GoReleaser. DO NOT EDIT.
class Glow < Formula
  desc ""
  homepage ""
  url "https://github.com/meinto/glow/releases/download/v1.5.3/glow_1.5.3_darwin_x86_64.tar.gz"
  version "1.5.3"
  sha256 "cef3eec0c356f7c648c18d70378412909ab76b42c8c6d1fa1e3d79d3f69405e7"
  
  depends_on "git"

  def install
    bin.install "glow"
  end
end