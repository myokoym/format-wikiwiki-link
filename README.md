# README

YouTube動画リンク取得ツール for https://wikiwiki.jp/

## Install

```
git clone https://github.com/myokoym/format-wikiwiki-link.git
cd ormat-wikiwiki-link
go mod init format-json
```

## Usage

```
curl -s "https://www.googleapis.com/youtube/v3/search?channelId=UCDrG8pzkq_cCCLcN5wgrzNw&key=<YouTube APIキー>&part=snippet&maxResults=3&fields=nextPageToken,items(id(videoId),snippet(publishedAt,title))&order=date&type=video" | jq . | tee -a result-makoh.json
go run format-json.go result-makoh.json >> result-makoh.txt
```
