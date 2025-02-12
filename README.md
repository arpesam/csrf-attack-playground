# CSRF Attack Simulation

This repository demonstrates a **Cross-Site Request Forgery (CSRF)** attack and provides a simulated environment for understanding how CSRF works. The setup includes:

1. A vulnerable **web server** written in Go that simulates a scenario where sensitive resources are accessible without proper CSRF protection.
2. An **attacker's page** designed to exploit this vulnerability by forging unauthorized requests to the server.

This project is meant for **educational purposes only** and should not be used for malicious activities.

---

## **How It Works**

### **1. The Vulnerable Web Server**
- The web server sets a `session_token` cookie for authenticated users.
- It includes two endpoints:
  - `/set`: Sets a cookie with the value `secure123` to simulate a login mechanism.
  - `/protected-resource`: Serves a protected resource, but the server only checks for the presence of the `session_token` cookie without validating the origin of the request.

#### **How it's vulnerable**
- The server does not validate whether the request comes from a trusted source (`Origin` or `Referer` headers are logged but not enforced).
- This makes it susceptible to CSRF attacks, where a malicious page can forge requests to `/protected-resource` using the victim's session.

---

### **2. The Attacker Page (csrf.html)**
- The attacker page is a simple HTML file that automatically submits a forged request to the vulnerable server's `/protected-resource` endpoint.
- When a victim visits the attacker page while logged into the vulnerable server, the attacker exploits the victim's session token stored in the browser cookies.

#### **How the attack works**
- The victim visits the attacker page.
- The page sends a request to `http://localhost:8080/protected-resource` using the victim's cookies, which the server accepts as legitimate due to the lack of CSRF protection.

---

## **Setup and Usage**

### **Prerequisites**
- Go installed on your system.
- A modern web browser to test the simulation.
- ngrok to run https locally in the webserver

### **Steps to Run**

#### **1. Clone the repository**
```bash
git clone <your-repo-url>
cd <your-repo-name>
```

#### **2. Run the Vulnerable Web Server**
```bash
go run main.go
```
- The web server will start on `http://localhost:8080`.

#### **3. Test the Vulnerable Web Server**
1. Visit `http://localhost:8080/set` to set the `session_token` cookie.
2. Visit `http://localhost:8080/protected-resource` to access the protected resource.

#### **4. Simulate the CSRF Attack**
1. Open the `csrf.html` file in a browser (e.g., by double-clicking it).
2. Observe that the attacker page automatically submits a request to `https://<NGROK_URL>:8080/protected-resource`.
3. The server processes the request as if it were made by the victim, demonstrating the CSRF vulnerability.

---

## **Example Walkthrough**

### **Legitimate Use Case**
1. Open the web server at `https://<NGROK_URL>::8080`.
2. Visit `/set` to set the cookie.
3. Visit `/protected-resource` to verify access to the protected resource.

### **Simulated Attack**
1. Open the `csrf.html` page in a browser.
2. The attacker page forges a request to `/protected-resource` without any interaction from the victim.
3. The server serves the protected resource, as it assumes the request is legitimate.

---

## **How to Fix the Vulnerability**
To prevent CSRF attacks, implement the following measures on the server:

1. **CSRF Tokens**
   - Generate a unique token for each user session.
   - Require the token to be submitted with sensitive requests.
   - Validate the token on the server.

2. **SameSite Cookies**
   - Set the `SameSite` attribute for session cookies to `Strict` or `Lax` to prevent them from being sent with cross-origin requests.

3. **Validate the Origin or Referer Header**
   - Ensure that sensitive requests come from trusted origins.

4. **Content Security Policy (CSP)**
   - Use CSP to restrict the loading and execution of resources from untrusted sources.

---

## **Disclaimer**
This project is for **educational purposes only**. Do not use this knowledge for malicious purposes. Always ensure your applications are secured against CSRF and other vulnerabilities.

---

## **Contributing**
Feel free to submit pull requests or report issues to improve the simulation or documentation.

---

## **License**
This project is licensed under the MIT License. See the LICENSE file for details.

