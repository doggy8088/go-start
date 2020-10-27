# Go 開發環境驗證專案

## 驗證步驟

1. 下載專案

    ```sh
    git clone https://github.com/doggy8088/go-start.git && cd go-start
    ```

2. 執行專案

    ```sh
    go run .
    ```

3. 開啟 Visual Studio Code 編輯器

    ```sh
    code .
    ```

4. 按下 `Ctrl-Shift-B` 建置專案

    檢查 `dist/` 資料夾是否有建置後的執行檔！

5. 執行建置成品

    macOS/Linux

    ```sh
    cd dist && ./go-start
    ```

    Windows

    ```sh
    cd dist && go-start.exe
    ```

6. 確認 `dist/gorm.db` 資料庫已被建立，並從 VSCode 開啟 `gorm.db` 資料庫 (使用 [SQLite](https://marketplace.visualstudio.com/items?itemName=alexcvzz.vscode-sqlite) 擴充套件)

    ![1-open-sqlite](https://user-images.githubusercontent.com/88981/97464501-f199ea80-197b-11eb-9a02-ee26408d41ea.png)

    ![2-show-table](https://user-images.githubusercontent.com/88981/97464506-f2cb1780-197b-11eb-8363-e549fc9d5d3d.png)
