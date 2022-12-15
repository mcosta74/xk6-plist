import plist from "k6/x/plist";

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
  console.log(xxx, format)

  let foo = {
    A: 1,
    B: 2,
    C: ['a', 'b'],
    D: {
      d1: 1,
      d2: 2.3,
    },
  };

  const str1 = plist.serialize(foo, 1);
  console.log(str1)

  const str2 = plist.serializeIndent(foo, 1, "\t");
  console.log(str2);
}