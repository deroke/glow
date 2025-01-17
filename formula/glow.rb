# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Glow < Formula
  desc ""
  homepage ""
  version "4.3.13"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/meinto/glow/releases/download/v4.3.13/glow_4.3.13_darwin_x86_64.tar.gz"
      sha256 "25b73c361cc9d9db8579669018585ad397afd107e491344e631daf3ab96b1590"

      def install
        bin.install "glow"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/meinto/glow/releases/download/v4.3.13/glow_4.3.13_linux_arm64.tar.gz"
      sha256 "f3b96284cf6f01c928be20ba1b8e0de074c94521bdccd996050412986dbcac4f"

      def install
        bin.install "glow"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/meinto/glow/releases/download/v4.3.13/glow_4.3.13_linux_x86_64.tar.gz"
      sha256 "af7de919c79af1445c47d09c6d55483cde9b59fdd5a2be894eb857b198ef00c0"

      def install
        bin.install "glow"
      end
    end
  end

  depends_on "git"
end
