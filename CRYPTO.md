CRYPTO DESIGN & GUIDELINES
==========================

Bu belge `core/crypto` ve `core/auth` için tasarım kararları, tercih edilen kriptografik yapı taşları, uygulama uyarıları ve denetim notlarını içerir. Proje güvenlik kritik olduğundan burada yazılanlar yol gösterici olup, kritik değişiklikler için kripto uzmanı onayı gereklidir.

1. Temel Prensipler
-------------------
- Güvenilir, iyi incelenmiş ve modern primitifleri kullanın — "home-grown" kriptografi yazmaktan kaçının.
- AEAD (Authenticated Encryption with Associated Data) kullanımı zorunlu olsun.
- Anahtar ve nonce yönetimi hataları en yaygın ve tehlikeli hatalardır; bunlara öncelik verin.
- Tüm kriptografik işlemler mümkünse sabit zamanlı (constant-time) yapılmalıdır.

2. Tavsiye Edilen Primitifler ve Kütüphaneler
---------------------------------------------
- ECDH: X25519 (Curve25519) — performans ve güvenlik dengesi için önerilir.
- İmza: Ed25519 veya Schnorr over Ristretto25519 (Schnorr tercih edilecekse Ristretto/ristretto255 tercih edin).
- AEAD: XChaCha20-Poly1305 (uygun nonce/IV kullanım kolaylığı) veya AES-GCM (donanım hızlandırma mevcutsa).
- KDF: HKDF-SHA256 (RFC 5869).
- Hash: SHA-256 (veya SHA-512 gerektiğinde).

Go ekosistemi için önerilen paketler:
- `golang.org/x/crypto/curve25519` (X25519)
- `golang.org/x/crypto/chacha20poly1305` (XChaCha20-Poly1305)
- `crypto/ed25519` veya `github.com/bwesterb/go-ristretto` / `filippo.io/edwards25519` + Ristretto dönüşümleri (Schnorr için)

3. El Sıkışma (Handshake)
-------------------------
- El sıkışma, anahtar değiş tokuşu + birbirini doğrulama (authentication) + oturum anahtarlarının türetilmesini içermelidir.
- Önerilen akış (örnek, basitleştirilmiş):
  1. X25519 ECDH temelinde ephemeral key exchange.
  2. Karşılıklı kimlik doğrulama için uzun süreli anahtarlarla (Ed25519/Schnorr) imza doğrulama.
  3. Ortak gizli anahtardan HKDF ile AEAD için session keys türetme.

- Not: Replay ve downgrade saldırılarına karşı protokolde versiyon ve nonce/anti-replay mekanizmaları ekleyin.

4. AEAD ve Nonce Yönetimi
-------------------------
- AEAD kullanımında nonce tekrarından (nonce reuse) kaçının — tekrar aynı nonce ile aynı anahtar kullanmak gizliliği tamamen yok eder.
- XChaCha20-Poly1305 tercih edilirse 24 bayt uzunluğunda nonce vardır; ephemeral/non-repeating nonce stratejisi kullanın (örn. counter + per-session random salt).

5. Anahtar Yönetimi
-------------------
- Uzun süreli anahtarlar (identity keys) disk üzerinde düz metin olarak saklanmamalıdır.
- Öneri: OS keystore (macOS Keychain, Linux Secret Service), Vault veya HSM kullanımı.
- Anahtar rotasyonu (rotation) planı hazırlayın ve protokole uyumlu şekilde tasarlayın (eski oturumların güvenli sonlandırılması).

6. Kriptografik Hatalara Karşı Korunma
------------------------------------
- Tüm hata mesajları saldırganlara bilgi vermeyecek şekilde genel tutulmalı (ör. "authentication failed" yerine iç detay vermeyin).
- İmzaların ve MAC doğrulamalarının zamanlamasını saldırıya açık bırakmayın — constant-time karşılaştırma kullanın.

7. Testler ve Denetimler
-----------------------
- Birim testleri: deterministik olmayan kripto işlevlerini izole eden test harness'ları yazın (mock RNG veya fixed seeds test-only ortamında kullanılabilir).
- Fuzzing: protokol parser'ları ve frame işlemeyi fuzz ile test edin (`go-fuzz` / `AFL` gibi araçlar).
- Farklı implementasyonlarla uyumluluk testleri oluşturun (ör. bir JS/Python istemci ile el sıkışma senaryoları).
- Değişikliklerde kripto uzmanı tarafından kod incelemesi (code review) ve gerekirse bağımsız audit isteyin.

8. Performans ve Yan Kanal (Side-channel) Dikkatleri
--------------------------------------------------
- Sabit zamanlı implementasyonlar önemlidir; Go kütüphanelerinin çoğu buna dikkat eder ama dikkatli olun.
- CPU önbellek, branch-prediction gibi yan kanallara karşı kritik kodlar için ek önlemler düşünün.

9. Uygulamada Riskli veya Boş Dosyalar
--------------------------------------
- Repository'de `core/auth/schnorr.go` ve `core/crypto/params.go` gibi boş veya eksik dosyalar tespit edildi. Bunlar kritik işlevleri temsil ediyor olabilir; prod ortamına almadan önce:
  - Bu dosyaların gerçek, güvenli implementasyonları eklenmeli.
  - Eğer özel Schnorr implementasyonu kullanılacaksa, Ristretto/Curve25519 tabanlı, testleri ve referans vektörleri sağlanmış bir yaklaşım tercih edin.

10. Sürümleme ve Güvenlik Bildirimleri
-------------------------------------
- Kriptografide yapılan değişiklikler hem protokol sürümünü hem de wire-format uyumluluğunu etkileyebilir — sürümlendirme politikasını netleştirin.
- Güvenlik düzeltmeleri için koordineli açıklama (coordinated disclosure) süreçlerini `SECURITY.md`'deki yönergelere göre yürütün.

11. Hızlı Uygulama Kontrolleri (Checklist)
-----------------------------------------
- [ ] AEAD kullanılıyor ve nonce yönetimi doğru.
- [ ] KDF = HKDF-SHA256 veya eşdeğeri kullanılıyor.
- [ ] ECDH = X25519 tercih ediliyor.
- [ ] İmzalar Ed25519 veya Schnorr+Ristretto tabanlı.
- [ ] Tüm anahtarlar güvenli bir yerde saklanıyor (Keystore/Vault/HSM).
- [ ] Parser ve frame işleme kodları için fuzz testleri mevcut.
- [ ] `core/auth/schnorr.go` ve `core/crypto/params.go` implementasyonları planlandı/eklendi.

12. Kaynaklar ve Referanslar
---------------------------
- RFCs: RFC 5869 (HKDF), RFC 8439 (ChaCha20-Poly1305)
- libs/implementations: libsodium, `golang.org/x/crypto`, `filippo.io/edwards25519`.

Not: Bu belge teknik rehberlik sağlar ama üretime geçmeden önce bağımsız bir kripto denetimi (audit) zorunludur.
