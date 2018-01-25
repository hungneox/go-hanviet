class Hanviet < Formula
    desc "A CLI to vietnamtudien.org/hanviet"
    homepage "https://github.com/hungneox/go-hanviet"
    version "1.0.0"
    sha256 "a759bc358a513c93963d22b044c9542651059b9e1a4000d5d08db20e9ace1697"
    url "https://github.com/hungneox/go-hanviet/releases/download/1.0.0/darwin-amd64-1.0.0.zip"
  
    def install
      bin.install "hanviet"
    end
  end