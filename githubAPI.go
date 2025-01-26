package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"time"
)

type Config struct {
	DownloadURL string `json:"downloadURL"`
}

func DownloadAndReplace() {
	// メインのExeのパス(名前が変わった際は変更する)
	mainExe := "../VRC-Avatar-Library.exe"

	// 既存のアプリを削除
	fmt.Println("既存のアプリを削除しています...")
	err := os.Remove(mainExe)
	if err != nil {
		fmt.Println("既存のアプリの削除に失敗しました:", err)
	}

	// 1秒待機
	time.Sleep(1 * time.Second)

	// アプリのダウンロード
	fmt.Println("ダウンロードを開始します...")
	downloadURL := readJSON()
	resp, err := http.Get(downloadURL)
	if err != nil {
		fmt.Println("ダウンロードに失敗しました:", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(mainExe)
	if err != nil {
		fmt.Println("ファイルの作成に失敗しました:", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("ファイルの書き込みに失敗しました:", err)
	}

	fmt.Println("ダウンロードが完了しました。")

	// アプリを起動
	fmt.Println("アプリを起動しています...")
	err = startApp(mainExe)
	if err != nil {
		fmt.Println("アプリの起動に失敗しました:", err)
	}
	fmt.Println("アプリの起動が完了しました。")

	time.Sleep(1 * time.Second)
}

func startApp(exePath string) error {
	cmd := exec.Command(exePath)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd.Start()
}

func readJSON() string {
	file, err := os.Open("./Config/config.json")
	if err != nil {
		fmt.Println("設定ファイルの読み込みに失敗しました:", err)
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		fmt.Println("設定ファイルのパースに失敗しました:", err)
	}
	return config.DownloadURL
}