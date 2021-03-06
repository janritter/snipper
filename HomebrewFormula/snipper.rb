# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Snipper < Formula
  desc "Tool to get various snippets directly from your CLI"
  homepage "https://github.com/janritter/snipper"
  version "0.1.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/janritter/snipper/releases/download/0.1.0/snipper_0.1.0_darwin_arm64.tar.gz"
      sha256 "12a24b7c47cfe94c17b85d147957ae3f263a56eae037488d7a034613e9024998"

      def install
        bin.install "snipper"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/janritter/snipper/releases/download/0.1.0/snipper_0.1.0_darwin_amd64.tar.gz"
      sha256 "2e3c29f895fecc4981363a3b7651e8713d3a9c0fc2caeb0e5b684ac7951297b2"

      def install
        bin.install "snipper"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/janritter/snipper/releases/download/0.1.0/snipper_0.1.0_linux_arm64.tar.gz"
      sha256 "87baf6bdaab14e07e8dd69925c1a78db5eadcb364e51b3ba7d8c6fcd20a49006"

      def install
        bin.install "snipper"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/janritter/snipper/releases/download/0.1.0/snipper_0.1.0_linux_amd64.tar.gz"
      sha256 "d9221638b05a2759d7d0f1857da76ca7002f9f24faa15456b13b53351e6ee4d7"

      def install
        bin.install "snipper"
      end
    end
  end
end
