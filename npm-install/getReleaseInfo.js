import * as fs from "fs/promises";
import os from "os";

function getSystemInfo() {
  const type = os.type();
  const arch = os.arch();
  if (type === "Windows_NT" && arch === "x64") {
    return {
      platform: "Windows",
      arch: "x86_64",
    };
  }
  if (type === "Windows_NT") {
    return {
      platform: "Windows",
      arch: "i386",
    };
  }
  if (type === "Linux" && arch === "x64") {
    return {
      platform: "Linux",
      arch: "arm64",
    };
  }
  if (type === "Linux" && arch === "x32") {
    return {
      platform: "Linux",
      arch: "i386",
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

async function getReleaseInfo() {
  const { platform, arch } = getSystemInfo();
  const packageJson = await fs.readFile("package.json").then(JSON.parse);
  let version = packageJson.version;
  const url = `https://github.com/Onboardbase/onboardbasekit/releases/download/v${version}/onboardbasekit_${platform}_${arch}.tar.gz`;
  const name = "onboardbasekit";
  return { url, name };
}

export default getReleaseInfo;
