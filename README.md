# freee-redirector
freeeの今月の勤怠ページへとリダイレクトします。

`https://p.secure.freee.co.jp/#work_records/#{年}/#{給与支払い月}/employees/#{従業員番号？}`
のようなURI(毎月変わる)が常に開けるURIが欲しかった。

## 設定項目

従業員番号: `template.yaml` の `employees` を自分のものにすればよい。
