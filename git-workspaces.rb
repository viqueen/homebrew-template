require "language/node"

class GitWorkspaces < Formula
	desc "Collection of useful git extensions that enhance one's local dev environment"
	homepage "https://github.com/viqueen/git-workspaces"
	url "https://registry.npmjs.org/@labset/git-workspaces/-/git-workspaces-4.2.2.tgz"
	sha256 "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	license "Apache-2.0"

	depends_on "node"

	def install
    	system "npm", "install", *Language::Node.std_npm_install_args(libexec)
    	bin.install_symlink Dir["#{libexec}/bin/*"]
  	end

  	test do
    	system "false"
  	end
end
