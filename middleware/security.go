package middleware

import "net/http"

// SecurityHeaders adds security-related HTTP headers to protect against common web vulnerabilities.
func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Content Security Policy - allows external CDNs and inline styles
		w.Header().Set("Content-Security-Policy",
			"default-src 'self'; "+
				"script-src 'self' https://cdn.tailwindcss.com https://unpkg.com https://cdnjs.cloudflare.com 'unsafe-inline'; "+
				"style-src 'self' 'unsafe-inline' https://cdn.tailwindcss.com https://cdnjs.cloudflare.com; "+
				"img-src 'self' data: blob:; "+
				"connect-src 'self'; "+
				"font-src 'self' data:; "+
				"object-src 'none'; "+
				"base-uri 'self'; "+
				"form-action 'self'; "+
				"frame-ancestors 'none';")

		// Prevent browsers from MIME-sniffing
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// Prevent page from being displayed in iframe (clickjacking protection)
		w.Header().Set("X-Frame-Options", "DENY")

		// Control how much referrer information is sent
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

		// Restrict browser features
		w.Header().Set("Permissions-Policy", "geolocation=(), microphone=(), camera=()")

		next.ServeHTTP(w, r)
	})
}
