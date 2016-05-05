require "language/go"

class Ptar < Formula
  desc "ptar: parallel tar (SMP)"
  homepage "https://github.com/moul/ptar"
  url "https://github.com/moul/ptar/archive/v1.0.0.tar.gz"
  sha256 "a23355f496c2940a1b33564264b007ff74d21c7721e954882748e22577878252"

  head "https://github.com/moul/ptar.git"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    ENV["GOBIN"] = buildpath
    ENV["GO15VENDOREXPERIMENT"] = "1"
    (buildpath/"src/github.com/moul/ptar").install Dir["*"]

    system "go", "build", "-o", "#{bin}/ptar", "-v", "github.com/moul/ptar/cmd/ptar/"
  end
end
