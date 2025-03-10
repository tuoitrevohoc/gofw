export function stringToArrayBuffer(binaryString: string) {
  return Uint8Array.from(binaryString, (c) => c.charCodeAt(0)).buffer;
}

export function base64Encode(binary: ArrayBuffer) {
  const binaryString = String.fromCharCode(...new Uint8Array(binary));
  return btoa(binaryString)
    .replace(/\+/g, "-")
    .replace(/\//g, "_")
    .replace(/=/g, "");
}
