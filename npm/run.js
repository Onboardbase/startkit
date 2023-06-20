#!/usr/bin/env node

const { Binary } = require("binary-install");
let binary = new Binary(
  "onboardbasekit",
  "https://github.com/Onboardbase/onboardbasekit/releases/download/v1.0.1/onboardbasekit_Darwin_arm64.tar.gz"
);
binary.run();