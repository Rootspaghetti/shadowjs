<!DOCTYPE html>
<html lang="tr">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body>
  <div style="text-align: center;">
    <img src="docs/shadowjs-logo.png" alt="ShadowJS Logo" style="max-width: 200px; margin-bottom: 20px;">
  </div>

  <h1 style="text-align: center;">ShadowJS</h1>
  
  <p>
    <strong>ShadowJS</strong>, bir domain veya subdomain listesinde bulunan JavaScript dosyalarını (JS) bulmak için geliştirilmiş hafif ve güçlü bir araçtır. Araç, tarama sırasında bulunan tüm <code>.js</code> dosyalarını toplar ve aynı isimli veya aynı URL'ye sahip olan dosyaları tekrarlamadan kaydeder. ShadowJS, özellikle bug bounty avcıları, penetrasyon test uzmanları ve geliştiriciler için hızlı ve kolay bir çözüm sunar.
  </p>

  <hr>

  <h2>Özellikler</h2>
  <ul>
    <li><strong>Domain Taraması:</strong> Bir domain üzerinde tarama yaparak <code>.js</code> dosyalarını toplar (bağlantılardaki tüm JS dosyalarını içerir).</li>
    <li><strong>Subdomain Desteği:</strong> Dosya içindeki subdomain listesini tarar.</li>
    <li><strong>Tekrarsız Sonuçlar:</strong> Aynı isimli veya aynı URL'ye sahip olan dosyaları sadece bir kez kaydeder.</li>
    <li><strong>Kullanım Kolaylığı:</strong> Basit bir komut satırı arayüzü ile kolayca kullanılır.</li>
    <li><strong>Özelleştirilebilir Çıktı:</strong> Sonuçları istediğiniz dosyaya kaydedebilir.</li>
  </ul>

  <hr>

  <h2>Kurulum</h2>
  <p>ShadowJS'yi kullanmak için bilgisayarınızda <a href="https://go.dev/" target="_blank">Go</a> kurulu olmalıdır.</p>
  <ol>
    <li>
      Depoyu klonlayın:
      <pre><code>git clone https://github.com/yourusername/shadowjs.git
cd shadowjs</code></pre>
    </li>
    <li>
      Aracı derleyin:
      <pre><code>go build -o shadowjs</code></pre>
    </li>
    <li>
      Aracı global olarak çalıştırmak için <code>/usr/local/bin</code> dizinine taşıyın:
      <pre><code>sudo cp shadowjs /usr/local/bin</code></pre>
    </li>
    <li>
      Artık aracı terminalden <code>shadowjs</code> komutu ile çalıştırabilirsiniz.
    </li>
  </ol>

  <hr>

  <h2>Kullanım</h2>
  <h3>Tek Bir Domain Taraması</h3>
  <p>Belirli bir domaini tarayıp <code>.js</code> dosyalarını bulmak için:</p>
  <pre><code>shadowjs -d https://example.com -o js_sonuclari.txt</code></pre>

  <h3>Subdomain Listesi Taraması</h3>
  <p>Bir dosyadaki subdomain listesini tarayıp <code>.js</code> dosyalarını bulmak için:</p>
  <pre><code>shadowjs -subs subdomains.txt -o js_sonuclari.txt</code></pre>

  <h3>Parametreler</h3>
  <table border="1">
    <thead>
      <tr>
        <th>Parametre</th>
        <th>Açıklama</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td><code>-d</code></td>
        <td>Tarama yapılacak hedef domain.</td>
      </tr>
      <tr>
        <td><code>-subs</code></td>
        <td>Subdomain listesini içeren dosya.</td>
      </tr>
      <tr>
        <td><code>-o</code></td>
        <td>Sonuçların kaydedileceği dosya (varsayılan: <code>results.txt</code>).</td>
      </tr>
    </tbody>
  </table>

  <hr>

  <h2>Çıktı</h2>
  <p>Çıktı dosyası, tarama sırasında bulunan tüm benzersiz <code>.js</code> dosyalarının URL'lerini içerir. Örnek:</p>
  <pre><code>https://example.com/static/app.js
https://sub.example.com/scripts/main.js</code></pre>

  <hr>

  <h2>Lisans</h2>
  <p>
    ShadowJS, <strong>Apache License 2.0</strong> ile lisanslanmıştır. Daha fazla bilgi için <a href="LICENSE">LICENSE</a> dosyasına göz atabilirsiniz.
  </p>

  <hr>

  <h2>Katkıda Bulunun</h2>
  <p>Yeni özellikler eklemek, hataları düzeltmek veya belgeleri geliştirmek için katkılarınızı bekliyoruz! Pull request veya issue açabilirsiniz.</p>

  <hr>

  <h2>Uyarı</h2>
  <p>
    Bu araç sadece etik kullanım için tasarlanmıştır. Araç ile yapılan yasa dışı veya kötü niyetli faaliyetlerden yazarlar sorumlu değildir. Taramadan önce hedef domain veya subdomain üzerinde yetkiniz olduğundan emin olun.
  </p>

  <hr>

  <h2>Yazar</h2>
  <p>
    ShadowJS, <strong>Root@Spaghetti</strong> tarafından geliştirilmiş ve bakım yapılmaktadır. Sorularınız veya önerileriniz varsa bir issue açabilir veya bizimle iletişime geçebilirsiniz.
  </p>
</body>
</html>
