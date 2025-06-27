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
      const res = await fetch("/api/wordbook", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          baseWordbookPath: wordbookValue,
          rng: rngValue
        })
      });
  
      const result = await res.json();
  
      if (result["status"] !== 200) {
        alert("生成の途中でエラーが発生しました");
        return;
      }
  
      if (result["filepath"]){
        console.log("ファイルパス:", result["filepath"]);
        const fullPath = result["filepath"];
        const fileName = fullPath.split("/").pop();
        
        const a = document.createElement("a");
        document.body.appendChild(a);
        a.download = fileName;
        a.href = fullPath;
        a.click();
        a.remove();
        return;
      }

    } catch (error) {
      console.error("通信エラー:", error);
      alert("通信中にエラーが発生しました。");
      return;
    }
}
  