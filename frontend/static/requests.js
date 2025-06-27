// 共通のAPIリクエスト関数
async function postJSON(url, data) {
  try {
    const res = await fetch(url, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data)
    });

    if (!res.ok) {
      throw new Error(`サーバーエラー: ${res.status}`);
    }

    return await res.json();
  } catch (error) {
    console.error("API通信エラー:", error);
    throw error;
  }
}

async function executeDownload(path) {
  const fullPath = path;
  const fileName = fullPath.split("/").pop();
  const a = document.createElement("a");
  document.body.appendChild(a);
  a.download = fileName;
  a.href = fullPath;
  a.click();
  a.remove();
}


async function requestGenerateWordBook() {
  const wordbookValue = document.getElementById("wordbookSelect").value;
  const rngValue = document.getElementById("rngInput").value.replace(/\s/g, '');

  if (!wordbookValue || !rngValue) {
    alert("入力されていない項目があります。");
    return;
  }

  if (/[^0-9,\-~]/.test(rngValue)) {
    alert("範囲の入力が不正です。");
    return;
  }

  try {
    const result = await postJSON("/api/wordbook", {
      baseWordbookPath: wordbookValue,
      rng: rngValue
    });

    if (result["status"] !== 200) {
      alert("生成の途中でエラーが発生しました");
      return;
    }

    if (result["filepath"]) {
        executeDownload(result["filepath"]);
    }
  } catch (error) {
    alert("通信中にエラーが発生しました。");
  }
}


async function requestGenerateWordTest() {
  const wordbookValue = document.getElementById("wordbookSelect").value;
  const rngValue = document.getElementById("rngInput").value.replace(/\s/g, '');
  const isReverse = document.getElementById("reverseCheckSwitch").checked;
  const isRandom = document.getElementById("randomCheckSwitch").checked;

  console.log("isReverse:", isReverse);

  if (!wordbookValue || !rngValue) {
    alert("入力されていない項目があります。");
    return;
  }

  if (/[^0-9,\-~]/.test(rngValue)) {
    alert("範囲の入力が不正です。");
    return;
  }

  try {
    const result = await postJSON("/api/wordtest", {
      baseWordbookPath: wordbookValue,
      rng: rngValue,
      isReverse: isReverse,
      isRandom: isRandom
    });

    if (result["status"] !== 200) {
      alert("生成の途中でエラーが発生しました");
      return;
    }

    if (result["filepath"]) {
        executeDownload(result["filepath"]);
    }
  } catch (error) {
    alert("通信中にエラーが発生しました。");
  }
}
