# WordbookGenerater-Go
単語帳アプリ[WordHolic](https://www.langholic.com/wordholic)用のcsvファイル生成、及び単語テスト用のxlsxファイル生成Webアプリ。


# 対応単語帳
- ターゲット1000
- ターゲット1900
- 古文単語330

# 依存関係
## フロントエンド
[BootStrap](https://getbootstrap.jp)

## バックエンド
github.com/xuri/excelize/v2  
github.com/gin-gonic/gin  
github.com/mattn/go-runewidth  

# メモ
このアプリを立ち上げると以下のAPIが立ち上がります。
~~~
POST /api/wordbook

Parameters
baseWordbookPath string(required)
rng              string(required)

Response
{
   "status":   200,
   "filepath": path,
}
~~~
~~~
POST /api/wordtest

Parameters
baseWordbookPath string(required)
rng              string(required)
isReverse        string
isRandom         string

Response
{
   "status":   200,
   "filepath": path,
}
~~~


