Project of Thenos
=================

Kısa Açıklama
-------------
Project of Thenos, ekip içi gizli iletişim için tasarlanmış uçtan-uça şifrelemeye odaklı bir chat sunucusu ve istemci setidir. Performans, gizlilik ve kapalı devre kullanım (private deployments) ön planda tutulur.

Hızlı Başlangıç
----------------
Gereksinimler:
- Go 1.20+
- `make` (opsiyonel)

Derleme:

```bash
cd /path/to/Project-of-Thenos/core
go build ./...
```

Sunucu çalıştırma (örnek):

```bash
cd core
go run ./server
```

İstemci CLI (örnek):

```bash
cd client/cli
go run .
```

Proje Yapısı (kısa):
- `core/` — sunucu, protokol, kimlik doğrulama, kripto çekirdeği
- `client/cli/` — basit CLI istemcisi
- `crypto/`, `auth/` — kriptografik fonksiyonlar ve el sıkışma mekanizmaları
- `router/`, `server/` — mesaj yönlendirme ve TCP listener
- `docs/` — tasarım ve protokol belgeleri

Güvenlik Notları (özet):
- Proje hassas kriptografik kurallara dayandığından `core/crypto` ve `core/auth` içindeki eksik/boş dosyalar kritik; üretime almadan önce bir kripto denetimi (audit) önerilir.
- Gizli anahtar yönetimi, disk üzerinde açık metin tutulmamalıdır.
- Varsayılan olarak debugging ve logging seviyeleri gizli verileri açığa çıkarabilir; üretimde minimize edilmelidir.

Nasıl katkıda bulunulur?
- Kodu biçimlendir: `gofmt -w .`
- Test ve lint çalıştır: `go test ./...` ve `golangci-lint run` (konfigüre edilmemişse önce ekleyin)
- Yeni özellikler için issue açın ve PR gönderin.

Daha fazla doküman
- `docs/crypto.md`, `docs/protocol.md`, `PROTOCOL.md` dosyalarını inceleyin.

Sorumluluk reddi
- Bu proje güvenlik kritik yazılımdır; dağıtıma almadan önce bağımsız bir güvenlik denetimi yaptırın.

İrtibat
- Proje deposunda issue'lar üzerinden iletişim kurun.
# Project-of-Thenos