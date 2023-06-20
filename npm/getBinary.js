const { Binary } = require("binary-install");
const os = require("os");

function getSystemInfo() {
  const type = os.type();
  const arch = os.arch();

  if (type === "Windows_NT" && arch === "x64") {
    return {
      platform: "win32",
      arch: "x64",
    };
  }
  if (type === "Windows_NT") {
    return {
      platform: "win32",
      arch: "x86",
    };
  }
  if (type === "Linux" && arch === "x64") {
    return {
      platform: "linux",
      arch: "x64",
    };
  }
  if (type === "Darwin" && arch === "x64") {
    return {
      platform: "Darwin",
      arch: "x64",
    };
  }
  if (type === "Darwin" && arch === "arm64") {
    return {
      platform: "Darwin",
      arch: "arm64",
    };
  }

  throw new Error(`Unsupported platform: ${type} ${arch}`);
}

function getBinary() {
  const { platform, arch } = getSystemInfo();
  const version = require("../package.json").version;
  //   https://github.com/Onboardbase/onboardbasekit/releases/download/v0.0.7/onboardbasekit_Darwin_x86_64.tar.gz
  const url = `https://github.com/Onboardbase/onboardbasekit/releases/download/v${version}/onboardbasekit_${platform}_${arch}.tar.gz`;
  //   url: "https://github.com/Onboardbase/onboardbasekit/releases/download/v{{version}}/onboardbasekit_{{platform}}_{{arch}}.tar.gz",
  const name = "onboardbasekit";
  return new Binary(name, url);
}

module.exports = getBinary;
