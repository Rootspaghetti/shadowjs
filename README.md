<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body>
  <div style="text-align: center;">
    <img src="shadowjs.jpg" alt="ShadowJS Logo" style="max-width: 200px; margin-bottom: 20px;">
  </div>

  <h1 style="text-align: center;">ShadowJS</h1>
  
  <p style="text-align: center;">
    <a href="README-TR.md" style="text-decoration: none; color: blue; font-weight: bold;">Türkçe Kullanım</a>
  </p>

  <hr>

  <p>
    <strong>ShadowJS</strong> is a lightweight and powerful tool for finding JavaScript files within a domain or its subdomains. It crawls a target domain or a list of subdomains to extract <code>.js</code> files and ensures that only unique files are saved in the output. 
    ShadowJS is designed for bug bounty hunters, penetration testers, and developers looking for a quick and efficient way to gather JavaScript files for further analysis.
  </p>

  <hr>

  <h2>Features</h2>
  <ul>
    <li><strong>Domain Crawling:</strong> Crawls the entire domain to find <code>.js</code> files, including nested links.</li>
    <li><strong>Subdomain Support:</strong> Supports scanning subdomains from a file.</li>
    <li><strong>Unique Results:</strong> Ensures no duplicate JavaScript file URLs are included in the output.</li>
    <li><strong>Easy to Use:</strong> Simple command-line interface for quick integration into your workflow.</li>
    <li><strong>Customizable Output:</strong> Saves results to a file of your choice.</li>
  </ul>

  <hr>

  <h2>Installation</h2>
  <p>To install and use ShadowJS, you need to have <a href="https://go.dev/" target="_blank">Go</a> installed on your machine.</p>
  <ol>
    <li>
      Clone the repository:
      <pre><code>git clone https://github.com/yourusername/shadowjs.git
cd shadowjs</code></pre>
    </li>
    <li>
      Build the tool:
      <pre><code>go build -o shadowjs</code></pre>
    </li>
    <li>
      Move the tool to your <code>/usr/local/bin</code> directory for global access:
      <pre><code>sudo cp shadowjs /usr/local/bin</code></pre>
    </li>
    <li>
      Run the tool:
      <pre><code>shadowjs</code></pre>
    </li>
  </ol>

  <hr>

  <h2>Usage</h2>
  <h3>Scan a Single Domain</h3>
  <p>To crawl a specific domain and find <code>.js</code> files:</p>
  <pre><code>shadowjs -d https://example.com -o js_results.txt</code></pre>

  <h3>Scan Subdomains from a File</h3>
  <p>To scan a list of subdomains for <code>.js</code> files:</p>
  <pre><code>shadowjs -subs subdomains.txt -o js_results.txt</code></pre>

  <h3>Parameters</h3>
  <table border="1">
    <thead>
      <tr>
        <th>Parameter</th>
        <th>Description</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td><code>-d</code></td>
        <td>Target domain to scan.</td>
      </tr>
      <tr>
        <td><code>-subs</code></td>
        <td>File containing a list of subdomains.</td>
      </tr>
      <tr>
        <td><code>-o</code></td>
        <td>Output file to save the results (default: <code>results.txt</code>).</td>
      </tr>
    </tbody>
  </table>

  <hr>

  <h2>Output</h2>
  <p>The output file will contain the URLs of all unique <code>.js</code> files found during the scan. Example:</p>
  <pre><code>https://example.com/static/app.js
https://sub.example.com/scripts/main.js</code></pre>

  <hr>

  <h2>License</h2>
  <p>
    ShadowJS is licensed under the <strong>Apache License 2.0</strong>. See the <a href="LICENSE">LICENSE</a> file for details.
  </p>

  <hr>

  <h2>Contribution</h2>
  <p>Contributions are welcome! If you'd like to add new features, fix bugs, or improve the documentation, feel free to open a pull request or issue.</p>

  <hr>

  <h2>Disclaimer</h2>
  <p>
    This tool is intended for ethical use only. The authors are not responsible for any misuse or illegal activity conducted with this tool. Ensure you have proper authorization before scanning any domain or subdomain.
  </p>

  <hr>

  <h2>Author</h2>
  <p>
    Developed and maintained by <strong>Your Name or Alias</strong>. If you have any questions or suggestions, feel free to open an issue or reach out.
  </p>
</body>
</html>
