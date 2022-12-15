# xk6-plist
This is a [k6](https://github.com/grafana/k6) extension using the
[xk6](https://github.com/grafana/xk6) system.

## Build

To build a `k6` binary with this plugin, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Install `xk6`:
  ```shell
  go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```shell
  xk6 build --with github.com/mcosta74/xk6-plist
  ```

## Example

```javascript
// script.js
import plist, { XMLFormat } from "k6/x/plist";

const data = `<?xml version="1.0" encoding="UTF-8"?>
  <!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
  <plist version="1.0">
    <dict>
      <key>TheString</key>
      <string>Hello World</string>
      <key>TheInt</key>
      <integer>1</integer>
      <key>TheReal</key>
      <real>1.7</real>
      <key>TheBool</key>
      <true />
      <key>TheArray</key>
      <array>
        <integer>1</integer>
        <integer>2</integer>
        <integer>3</integer>
      </array>
    </dict>
  </plist>`;

export default function() {
  const [xxx, format] = plist.parse(data);
  console.log(xxx, plist.getFormatName(format))

  let foo = {
    A: 1,
    B: 2,
    C: ['a', 'b'],
    D: {
      d1: 1,
      d2: 2.3,
    },
  };

  const str1 = plist.serialize(foo, XMLFormat);
  console.log(str1)

  const str2 = plist.serializeIndent(foo, XMLFormat, "\t");
  console.log(str2);
}```