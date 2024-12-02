package dome

// KnownPaths lists out common paths that are scanned by bots
// for vunerabilities.
var KnownPaths = []string{
	"/wp-admin",    // WordPress admin pages
	"/wp-content",  // WordPress content directory
	"/wp-includes", // WordPress includes directory
	"/.cgi-bin",    // CGI directory
	".php",         // WordPress XML-RPC interface
}
