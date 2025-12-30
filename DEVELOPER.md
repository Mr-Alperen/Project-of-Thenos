DEVELOPER GUIDE
===============

Lead Developer
--------------
- Name: Alperen ERKAN
- Email: alperenerkan00@gmail.com
- Web: mr-alperen.github.com
- Team: Stux6 Offensive Security Technology Team

Developer 2
-----------
- Name: Eymen Akyüz
- Email: eymenakyuz06@gmail.com
- Web: mr-alperen.github.com
- Team: Stux6 Offensive Security Technology Team

Proje Mimari (Kısa)
--------------------
- `core/` — sunucu uygulaması, protokol, kimlik doğrulama ve kriptografi çekirdeği.
- `client/cli/` — basit komut satırı istemcisi.
- `router/` — mesaj yönlendirme ve dispatcher.
- `server/` — listener, session yönetimi ve konfigürasyon.
- `docs/` — tasarım belgeleri ve protokol açıklamaları.

Geliştirme Ortamı
-----------------
- Gereksinimler:
  - Go 1.20 veya daha yeni
  - `gofmt`, `go vet`
  - Opsiyonel: `golangci-lint`

Kurulum ve Derleme
-------------------
Proje kökünde:

```bash
# tüm paketleri derle
cd /path/to/Project-of-Thenos
go build ./...

# veya sadece core server
cd core
go build ./server
```

Çalıştırma (Geliştirme)
-----------------------
Sunucu:

```bash
cd core
go run ./server
```

İstemci CLI (örnek):

```bash
cd client/cli
go run .
```

Kod Standartları ve İş Akışı
----------------------------
- Kod formatı: `gofmt -w .`
- Statik analiz: `go vet ./...` ve `golangci-lint run` (CI'da çalıştırılmalıdır).
- Branch: her özellik için `feature/` veya `bugfix/` branch'i açın.
- PR: kısa açıklama, testler, ve ilgili issue referansı ekleyin.

Testler
------
- Birim testleri çalıştırma:

```bash
go test ./...
```

- Test eklerken, kriptografi kodu için deterministik olmayan davranışları izole edin ve mock/fixture kullanın.

Güvenlik ve Kriptografi Notları
------------------------------
- `core/crypto` ve `core/auth` içindeki implementasyonlar güvenlik açısından kritik; değişiklik yaparken kripto uzmanı incelemesi gereklidir.
- Özel anahtarları kodda veya repo içinde saklamayın.

Hata Raporlama ve Destek
------------------------
- Yeni hata veya güvenlik bildirimi için lütfen repository üzerinden issue açın.
- Acil güvenlik sorunları için doğrudan e-posta gönderin:
  - Alperen: alperenerkan00@gmail.com
  - Eymen: eymenakyuz06@gmail.com

Ek Kaynaklar
------------
- `README.md` — proje genel bakışı ve hızlı başlangıç.
- `SECURITY.md` — tehdit modeli ve güvenlik prosedürleri.
- `docs/` — detaylı protokol ve kripto dokümanları.
