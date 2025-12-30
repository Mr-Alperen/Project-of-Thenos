PROTOCOL
========

Bu doküman proje içinde kullanılan kabuk (wire) protokolünün temel ilkelerini, frame formatlarını ve örnek mesaj akışlarını açıklar. Amaç: uyumluluk, güvenlik ve hata yönetimi için referans sağlamak.

1. Genel İlkeler
-----------------
- İletişim TCP üzerinden ikili (binary) frameler ile yapılır.
- Tüm multi-byte sayılar ağ bayt sıralaması (big-endian) ile kodlanır.
- Her frame başında sabit uzunluklu bir header bulunur; header frame tipini ve yük uzunluğunu içerir.
- Protokol sürümleriyle geriye dönük uyumluluk sağlamak için her el sıkışma başında versiyon numarası doğrulanır.

2. Frame Header
---------------
- Header (sabir 12 bayt) yapılandırması:
  - 0-1: Magic / Version (uint16)
  - 2: Frame Type (uint8)
  - 3: Flags (uint8)
  - 4-7: Payload Length (uint32)
  - 8-11: Reserved / Check (uint32) — ileride kullanılmak üzere

- Maksimum payload boyutu: 16 MiB (ayarlanabilir); eğer daha büyük veri gönderilecekse chunking uygulanmalıdır.

3. Frame Tipleri (örnek)
------------------------
- 0x01: Handshake Init — el sıkışma başlangıç isteği, payload = client hello (pubkey, nonce, supported algs)
- 0x02: Handshake Reply — sunucu hello (pubkey, nonce, server cert/identity proof)
- 0x03: Handshake Finish — client proof / signature
- 0x10: App Data — AEAD ile şifrelenmiş uygulama verisi (payload = nonce || ciphertext || tag)
- 0x11: Control — kontrol mesajları (ping/pong, keepalive)
- 0xF0: Error — hata bildirimi (hata kodu + kısa mesaj)

4. Handshake Örneği (basitleştirilmiş)
------------------------------------
1) Client -> Server: Handshake Init
   - payload: {version, client_pub_x25519, client_nonce, supported_algorithms}
2) Server -> Client: Handshake Reply
   - payload: {version, server_pub_x25519, server_nonce, server_identity_sig}
3) Client -> Server: Handshake Finish
   - payload: {client_identity_sig}

El sıkışma sonrası her iki taraf da ortak oturum anahtarını HKDF ile türeterek AEAD için anahtar ve nonce salt'ı oluşturur.

5. App Data Frame İçeriği
-------------------------
- AEAD kullanılarak şifrelenmiş gövde gönderilir. Örnek düzen:
  - 0-23: per-message nonce (XChaCha20-Poly1305 kullanılıyorsa 24 bayt)
  - geri kalan: ciphertext || tag

6. Hata Yönetimi ve Kodlar
-------------------------
- Hata frame'ı (`0xF0`) payload olarak {error_code(uint16), message_length(uint16), message(bytes)} içerir.
- Örnek hata kodları:
  - 0x0001: Protocol version mismatch
  - 0x0002: Handshake failed / auth failed
  - 0x0003: Invalid frame format / parse error
  - 0x0004: Internal server error

7. Replay ve Anti-Replay
------------------------
- Handshake sırasında kullanılan noncelar tek kullanımlık olmalıdır.
- Oturum açıldıktan sonra frame sequence number veya per-message nonce ile replay koruması sağlanmalıdır (örn. 64-bit counter, AEAD nonce içinde).

8. Parsers ve Güvenlik
----------------------
- Gelen veriyi parse ederken şu kurallara uyun:
  - Header'ı tamamen okuyun ve `Payload Length` sınırlarını doğrulayın.
  - Beklenmeyen frame tiplerinde bağlantıyı kapatmadan önce uygun hata dönün.
  - Parser'lar için fuzz testleri yazın; `core/protocol/parser.go` ve `core/protocol/frame.go` kritik yüzeylerdir.

9. Backpressure ve Flow Control
-------------------------------
- Sunucu, artan bellek kullanımı veya yavaş istemci durumlarında bağlantıyı kapatmadan önce `Control`/`Pause` benzeri mesajlarla akışı yavaşlatmalıdır.

10. Örnek Binary Frame (el sıkışma init)
---------------------------------------
- Header (12B):
  - Version: 0x0001
  - Type: 0x01
  - Flags: 0x00
  - Payload Length: 0x00000120 (288 byte)
  - Reserved: 0x00000000
- Payload: client hello yapılandırmasına göre değişir (pubkey, nonce, alg listesi)...

11. Geriye Dönük Uyum ve Sürümleme
----------------------------------
- Protokol değişiklikleri version bump ile yapılmalı; yeni alanlar optional olmalı ve eski taraflarda yoksayılmalıdır.

12. Testler
----------
- Frame parsing unit testleri, handshake akışının entegrasyon testleri ve AEAD encrypt/decrypt round-trip testleri yazılmalıdır.
- Parser'ları fuzz ile zorlayarak bellek taşması, panik ve parse hatalarını tespit edin.

13. Uzantılar
------------
- İleride eklenecek özellikler için `Flags` ve `Reserved` alanları kullanılabilir. Eklenecek yeni frame tipleri için registry (dokümantasyon) güncellenecektir.

14. Son Notlar
-------------
- Bu doküman protokolün bir referansıdır; uygulama detayları `core/protocol/frame.go` ve `core/protocol/parser.go` dosyalarındaki implementasyonla uyumlu olmalıdır.
