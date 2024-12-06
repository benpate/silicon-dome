package dome

// KnownPaths lists out common paths that are scanned by bots
// for vunerabilities.
var KnownPaths = []string{
	"/wp-admin",                // WordPress admin pages
	"/wp-content",              // WordPress content directory
	"/wp-includes",             // WordPress includes directory
	"/.cgi-bin",                // CGI directory
	".php",                     // WordPress XML-RPC interface
	"/.git/",                   // Git repository
	"/phpinfo",                 // PHP
	"/.aws/",                   // AWS directory
	"/.aws.yml",                // AWS configuration file
	"/.env",                    // System Environment file
	"/media/system/js/core.js", // Joomla core.js
}
