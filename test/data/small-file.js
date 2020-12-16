// this is my best code ever!
function test() {
  throw new Error("Oops...");
}
if (window.test) {
  test();
}