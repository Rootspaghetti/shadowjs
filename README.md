<h1 align="center">
  <img src="shadowjs.jpg" alt="ShadowJS" width="600">
  <br>ShadowJS
</h1>
<p align="center">
  <strong>JavaScript File Discovery Tool</strong><br>
  Created by <a href="https://github.com/Rootspaghetti">Rootspaghetti</a>
</p>
<p align="center">
  ShadowJS scans domains or subdomains to identify JavaScript files.<br>
  Built for bug bounty hunters and penetration testers.
</p>

<h2>Features</h2>
<ul>
  <li>Detects JavaScript files on domains or subdomains.</li>
  <li>Supports <code>-d</code> (single domain), <code>-subs</code> (subdomain list), and <code>-o</code> (output file).</li>
  <li>Avoids duplicate entries for files with the same name or URL.</li>
</ul>

<h2>Installation</h2>
<p>Install ShadowJS using the following command:</p>
<pre>
<code>go install github.com/Rootspaghetti/shadowjs@latest</code>
</pre>

<h2>Usage</h2>
<p>ShadowJS supports three main parameters:</p>
<ul>
  <li><code>-d</code>: Scan a single domain for JavaScript files.</li>
  <li><code>-subs</code>: Provide a subdomain list file for scanning.</li>
  <li><code>-o</code>: Save results to a specified output file.</li>
</ul>
<h2>Examples</h2>
<div style="background-color: #282c34; color: #abb2bf; padding: 15px; border-radius: 8px; font-family: monospace; font-size: 14px;">
  <p># scan a single domain<br>shadowjs -d example.com</p>
  <p># scan using a subdomain list<br>shadowjs -subs subdomains.txt</p>
  <p># save results to an output file<br>shadowjs -d example.com -o output.txt</p>
</div>

<h2>Contributing</h2>
<p>Feel free to contribute to ShadowJS by creating pull requests or reporting issues in the <a href="https://github.com/Rootspaghetti/shadowjs/issues">issues section</a>.</p>
