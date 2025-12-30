SECURITY
========

Bu belge Project of Thenos projesi için güvenlik politikası, tehdit modeli ve güvenlik açığı bildirim sürecini açıklar. Proje güvenlik-odaklı olduğundan üretime almadan önce bağımsız bir denetim (audit) şiddetle önerilir.

Tehdit Modeli
-------------
- Hedef kullanıcılar: ekip içi gizli iletişim kuran küçük-orta ölçekli takımlar.
- Tehdit vektörleri: aktörlerin ağ üzerinde dinleme (passive eavesdropping), aktif man-in-the-middle, sunucu veya istemci tarafında anahtar sızması, yanlış yapılandırma, kötü amaçlı güncelleme.
- Güvenilir varsayımlar: dağıtımı yapan operatörler yazılımı yönetir; şüpheli altyapı (ör. paylaşılan hosting) için ek risk değerlendirmesi gerekir.

Kritik Güvenlik İlkeleri
------------------------
- Uçtan uca şifreleme ana hedef: mesaj içeriği sunucuda okunamaz olmalıdır.
- Uzun süreli gizli anahtarlar disk üzerinde açık metin olarak saklanmamalıdır.
- Hata ve logging seviyeleri, gizli verileri ifşa etmeyecek şekilde sınırlandırılmalıdır.
- Güvenli varsayılanlar: debug ve verbose loglar kapalı olmalı; TLS/AEAD kullanımı zorunlu kılınmalıdır.

Anahtar Yönetimi
----------------
- Öneri: özel anahtarları sistem anahtar deposunda (OS keystore) veya KMS (HashiCorp Vault, AWS KMS vb.) kullanarak saklayın.
- Anahtar döndürme (rotation) planı oluşturun ve otomatikleştirin.
- Yedekleme işlemlerinde anahtarların şifrelenmiş olarak saklandığından emin olun.

Konfigürasyon ve Dağıtım
------------------------
- Üretim için yapılandırma dosyalarını örnek (`.example`) ile birlikte tutun ve gizli bilgileri ortam değişkenleri veya bir secret store ile sağlayın.
- CI/CD pipeline'da gizli anahtarların çıktıya yazılmadığını doğrulayın.

Logging ve Telemetri
--------------------
- Loglar kullanıcı içeriği (plaintext mesajlar) içermemelidir.
- Hata mesajları saldırganlara bilgi sızdırmayacak şekilde genel tutulmalıdır.

Güvenlik Açığı Bildirimi (Responsible Disclosure)
-----------------------------------------------
- Hassas bir güvenlik açığı keşfederseniz lütfen şu adımları izleyin:
  1. Öncelikle doğrudan projedeki repository sahipleriyle özel iletişime geçin (repo üzerinden private issue veya özel e-posta). Eğer özel iletişim yoksa açık olmayan bir issue başlatın ve "private"/"security" etiketini isteyin.
  2. Mümkünse raporu PGP/PGP-compatible şekilde şifreleyin; projenin PGP anahtarı repository'de belirtilmişse onu kullanın.
  3. Açığın istismar detaylarını ve etki alanını (CVE önerisi için gerekli bilgiler) paylaşın; üretimde hızlı düzeltme mümkünse önerinizi iletin.
  4. Biz raporu aldıktan sonra 72 saat içinde alındığını onaylayacağız; kritik açıklar için acil aksiyon ve koordinasyon sağlayacağız.

İletişim
--------
- Güvenlikle ilgili bildirimler için repository issue'ları veya proje sahibinin belirttiği özel iletişim yolunu kullanın.

Güvenlik Kontrol Listesi (Hızlı)
--------------------------------
- [ ] `core/crypto` ve `core/auth` implementasyonlarını bağımsız bir denetimden geçirin.
- [ ] Diskte anahtar tutmuyorsanız bunun doğrulanması için test ekleyin.
- [ ] Logging seviyelerini üretim için sınırlandırın.
- [ ] CI'da `go vet`, `golangci-lint` ve `go test` çalıştırın.

Notlar
-----
- Bu belge zaman içinde güncellenecektir; lütfen güvenlik süreçleri için en güncel metni repo ana dizininden kontrol edin.
