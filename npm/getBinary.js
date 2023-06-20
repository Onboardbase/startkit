const { Binary } = require("binary-install");
const os = require("os");

function getPlatform() {
  const type = os.type();
  const arch = os.arch();

  if (type === "Windows_NT" && arch === "x64") return "win64";
  if (type === "Windows_NT") return "win32";
  if (type === "Linux" && arch === "x64") return "linux";
  if (type === "Darwin" && arch === "x64") return "macos";

  throw new Error(`Unsupported platform: ${type} ${arch}`);
}

function getBinary() {
  const platform = getPlatform();
  const version = require("../package.json").version;
  //   https://github.com/Onboardbase/onboardbasekit/releases/download/v0.0.7/onboardbasekit_Darwin_x86_64.tar.gz
  const url = `https://github.com/Onboardbase/onboardbasekit/releases/download/v${version}/onboardbasekit_${platform}.tar.gz`;
  //   url: "https://github.com/Onboardbase/onboardbasekit/releases/download/v{{version}}/onboardbasekit_{{platform}}_{{arch}}.tar.gz",
  const name = "onboardbasekit";
  return new Binary(url, { name });
}

module.exports = getBinary;
