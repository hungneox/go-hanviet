class Hanviet < Formula
    desc "A CLI to vietnamtudien.org/hanviet"
    homepage "https://github.com/hungneox/go-hanviet"
    version "1.0.1"
    sha256 "045a4cc734e8d621f00e7848d262660744b75074fc165a87a24407e141fbbd71"
    url "https://github.com/hungneox/go-hanviet/releases/download/1.0.1/darwin-amd64-1.0.1.zip"
  
    def install
      bin.install "hanviet"
    end
  end