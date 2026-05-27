const nameInput = document.getElementById("name");
const resultEl = document.getElementById("result");
const runtimeEl = document.getElementById("runtime");

document.getElementById("btn").addEventListener("click", async () => {
  const name = (nameInput.value || "friend").trim();
  try {
    const msg = await window.greet(name);
    resultEl.textContent = msg;
  } catch (err) {
    resultEl.textContent = "Failed calling Go: " + err;
  }
});

(async () => {
  try {
    runtimeEl.textContent = await window.getRuntime();
  } catch {
    runtimeEl.textContent = "Runtime unavailable";
  }
})();
