INSTALLATION & BUILD
=====================

Bu doküman, Project of Thenos projesini geliştirme/üretim ortamında derleme ve çalıştırma adımlarını açıklar.

1) Gereksinimler
-----------------
- Go 1.20 veya daha yeni
- `git`
- (opsiyonel) `make`, `golangci-lint`

macOS (Homebrew önerilir):

```bash
brew install go git
```

Ubuntu/Debian:

```bash
sudo apt update
sudo apt install -y golang-go git make
```

2) Depoyu klonlama
-------------------

```bash
git clone <repo-url> Project-of-Thenos
cd Project-of-Thenos
```

3) Derleme (tüm paketler)
-------------------------

```bash
go build ./...
```

Core sunucuyu derleyip çalıştırma (geliştirme amaçlı):

```bash
cd core
go run ./server
```

CLI istemcisini çalıştırma:

```bash
cd client/cli
go run .
```

4) Konfigürasyon
----------------
- `core/config.go` içinde varsayılan ayarlar ve `server/config.go` yapılandırması bulunur; üretim için bir konfigürasyon dosyası veya ortam değişkenleri ile yapılandırma sağlamanız önerilir.
- Gizli anahtarları çevresel değişkenler veya bir secret store (Vault, OS keystore) ile yönetin; konfig dosyalarına düz metin anahtar koymayın.

5) Systemd servis (örnek)
-------------------------
Projede `deployment/systemd/thenos.service` benzeri bir dosya bulunuyorsa onu kullanabilirsiniz. Örnek:

```ini
[Unit]
Description=Thenos Server
After=network.target

[Service]
Type=simple
WorkingDirectory=/opt/Project-of-Thenos/core
ExecStart=/usr/bin/env go run ./server
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

Üretim için: derlenmiş ikiliyi (`go build`) uygun bir yere kopyalayın ve `ExecStart`'ı o ikiliye işaret edecek şekilde güncelleyin.

6) Logging ve Debugging
------------------------
- Geliştirme sırasında verbose loglar faydalıdır; üretimde log seviyesini düşürün ve mesaj içeriği (plaintext) yazdırılmadığından emin olun.

7) CI / Lint / Test Önerileri
---------------------------
- Otomatik kontroller için pipeline önerileri:

```bash
gofmt -w .
go vet ./...
golangci-lint run
go test ./...
```

8) Troubleshooting
------------------
- "bind: address already in use" hatası: port çakışması var — `lsof -i :<port>` ile kontrol edin.
- Bağlantı/parsing hatalarında, `core/protocol/parser.go` ve `core/protocol/frame.go` loglarını gözden geçirin.

9) Güvenlik Hatırlatmaları
-------------------------
- `core/crypto` ve `core/auth` içindeki anahtar/implementasyon eksiklerini üretime almadan önce tamamlayın ve bağımsız bir kripto denetimi yaptırın.
- Anahtarları kod veya konfig içinde saklamayın.

10) Daha Fazla Bilgi
--------------------
- Geliştirme için bakılacak dosyalar: `core/server`, `core/router`, `core/protocol`, `core/auth`, `core/crypto`.
- Dokümanlar: `README.md`, `DEVELOPER.md`, `SECURITY.md`, `CRYPTO.md`, `PROTOCOL.md`.
