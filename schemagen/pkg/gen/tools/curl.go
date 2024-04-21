package tools

import (
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/props"
	"github.com/UnstoppableMango/pulumi-kubernetes-the-hard-way/schemagen/pkg/gen/types"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

// Oh boy this is gonna take a while

func generateCurl() tool {
	inputs := map[string]schema.PropertySpec{
		"abstractUnixSocket": props.String("(HTTP) Connect through an abstract Unix domain socket, instead of using the network."),
		"altSvc":             props.String("(HTTPS)  This  option enables the alt-svc parser in curl."),
		"anyAuth":            props.Boolean("(HTTP) Tells curl to figure out authentication method by itself, and use the most secure one the remote site claims to support."),
		"append":             props.Boolean("(FTP SFTP) When used in an upload, this makes curl append to the target file instead of overwriting it."),
		"awsSigv4":           props.String("Use AWS V4 signature authentication in the transfer."),
		"basic":              props.Boolean("(HTTP) Tells curl to use HTTP Basic authentication with the remote host."),
		"cacert":             props.String("(TLS)  Tells curl to use the specified certificate file to verify the peer."),
		"capath":             props.String("(TLS) Tells curl to use the specified certificate directory to verify the peer."),
		"cert":               props.String("(TLS)  Tells  curl to use the specified client certificate file when getting a file with HTTPS, FTPS or another SSL-based protocol."),
		"certStatus":         props.Boolean("(TLS) Tells curl to verify the status of the server certificate by using the Certificate Status Request (aka. OCSP stapling) TLS extension."),
		"certType": {
			Description: "(TLS) Tells curl what type the provided client certificate is using.",
			TypeSpec:    types.LocalType("CurlCertType", "tools"),
		},
		"ciphers":        props.String("(TLS) Specifies which ciphers to use in the connection."),
		"compressed":     props.Boolean("(HTTP) Request a compressed response using one of the algorithms curl supports, and automatically decompress the content."),
		"compressedSsh":  props.Boolean("(SCP SFTP) Enables built-in SSH compression."),
		"config":         props.String("Specify a text file to read curl arguments from."),
		"connectTimeout": props.Integer("Maximum  time  in seconds that you allow curl's connection to take."),
		"connectTo":      props.String("For a request to the given HOST1:PORT1 pair, connect to HOST2:PORT2 instead."),
		"continueAt":     props.String("Continue/Resume a previous file transfer at the given offset."),
		"cookieJar":      props.String("(HTTP)  Specify  to which file you want curl to write all cookies after a completed operation."),
		"cookie":         props.String("(HTTP) Pass the data to the HTTP server in the Cookie header."),
		"createDirs":     props.Boolean("When used in conjunction with the -o, --output option, curl will create the necessary local directory hierarchy as needed."),
		"createFileMode": props.String("(SFTP SCP FILE) When curl is used to create files remotely using one of the supported protocols, this option allows the user to set which 'mode' to set on the file at creation time, instead of the default 0644."),
		"crlf":           props.Boolean("(FTP SMTP) Convert LF to CRLF in upload. Useful for MVS (OS/390)."),
		"crlfFile":       props.String("(TLS) Provide a file using PEM format with a Certificate Revocation List that may specify peer certificates that are to be considered revoked."),
		"curves":         props.String("(TLS)  Tells curl to request specific curves to use during SSL session establishment according to RFC 8422, 5.1."),
		"data":           props.String("(HTTP  MQTT) Sends the specified data in a POST request to the HTTP server, in the same way that a browser does when a user has filled in an HTML form and presses the submit button."),
		"dataAscii":      props.String("(HTTP) This is just an alias for -d, --data."),
		"dataBinary":     props.String("(HTTP) This posts data exactly as specified with no extra processing whatsoever."),
		"dataRaw":        props.String("(HTTP) This posts data similarly to -d, --data but without the special interpretation of the @ character."),
		"dataUrlEncode":  props.String("(HTTP) This posts data, similar to the other -d, --data options with the exception that this performs URL-encoding."),
		"delegation": {
			Description: "(GSS/kerberos) Set LEVEL to tell the server what it is allowed to delegate when it comes to user credentials.",
			TypeSpec:    types.LocalType("CurlDelegationLevel", "tools"),
		},
		"digest":                props.Boolean("(HTTP) Enables HTTP Digest authentication."),
		"disable":               props.Boolean("If used as the first parameter on the command line, the curlrc config file will not be read and used."),
		"disableEprt":           props.Boolean("(FTP) Tell curl to disable the use of the EPRT and LPRT commands when doing active FTP transfers"),
		"disableEpsv":           props.Boolean("(FTP)  Tell  curl to disable the use of the EPSV command when doing passive FTP transfers"),
		"disallowUsernameInUrl": props.Boolean("(HTTP) This tells curl to exit if passed a url containing a username."),
		"dnsInterface":          props.String("(DNS)  Tell  curl  to send outgoing DNS requests through <interface>."),
		"dnsIpv4Addr":           props.String("(DNS) Tell curl to bind to <ip-address> when making IPv4 DNS requests, so that the DNS requests originate from this address."),
		"dnsIpv6Addr":           props.String("(DNS) Tell curl to bind to <ip-address> when making IPv6 DNS requests, so that the DNS requests originate from this address."),
		"dnsServers":            props.String("Set the list of DNS servers to be used instead of the system default."),
		"dohCertStatus":         props.Boolean("Same as --cert-status but used for DoH (DNS-over-HTTPS)."),
		"dohInsecure":           props.Boolean("Same as -k, --insecure but used for DoH (DNS-over-HTTPS)."),
		"dohUrl":                props.String("Specifies which DNS-over-HTTPS (DoH) server to use to resolve hostnames, instead of using the default name resolver mechanism."),
		"dumpHeader":            props.String("(HTTP FTP) Write the received protocol headers to the specified file."),
		"egdFile":               props.String("(TLS) Specify the path name to the Entropy Gathering Daemon socket."),
		"engine":                props.String("(TLS) Select the OpenSSL crypto engine to use for cipher operations."),
		"etagCompare":           props.String("(HTTP) This option makes a conditional HTTP request for the specific ETag read from the given file by sending a custom If-None-Match header using the stored ETag."),
		"etagSave":              props.String("(HTTP) This option saves an HTTP ETag to the specified file."),
		"expect100Timeout":      props.Integer("(HTTP)  Maximum  time in seconds that you allow curl to wait for a 100-continue response when curl emits an Expects: 100-continue header in its request."),
		"failEarly":             props.Boolean("Fail and exit on the first detected transfer error."),
		"failWithBody":          props.Boolean("(HTTP)  Return an error on server errors where the HTTP response code is 400 or greater)."),
		"fail":                  props.Boolean("(HTTP) Fail silently (no output at all) on server errors."),
		"falseStart":            props.Boolean("(TLS) Tells curl to use false start during the TLS handshake."),
		"form":                  props.String("(HTTP SMTP IMAP) For HTTP protocol family, this lets curl emulate a filled-in form in which a user has pressed the submit button."),
		"formEscape":            props.Boolean("(HTTP) Tells curl to pass on names of multipart form fields and files using backslash-escaping instead of percent-encoding."),
		"formName":              props.String("(HTTP SMTP IMAP) Similar to -F, --form except that the value string for the named parameter is used literally."),
		"ftpAccount":            props.String("(FTP) When an FTP server asks for 'account data' after user name and password has been provided, this data is sent off using the ACCT command."),
		"ftpAlternativeUser":    props.String("(FTP) If authenticating with the USER and PASS commands fails, send this command."),
		"ftpCreateDirs":         props.Boolean("(FTP  SFTP) When an FTP or SFTP URL/operation uses a path that does not currently exist on the server, the standard behavior of curl is to fail."),
		"ftpMethod":             props.String("(FTP) Control what method curl should use to reach a file on an FTP(S) server."),
		"ftpPasv":               props.Boolean("(FTP) Use passive mode for the data connection."),
		"ftpPort":               props.String("(FTP)  Reverses  the  default  initiator/listener  roles  when connecting with FTP."),
		"ftpPret":               props.Boolean("(FTP) Tell curl to send a PRET command before PASV (and EPSV)."),
		"ftpSkipPasvIp":         props.Boolean("(FTP)  Tell  curl  to  not use the IP address the server suggests in its response to curl's PASV command when curl connects the data connection."),
		"ftpSslCccMode":         props.String("(FTP) Sets the CCC mode. The passive mode will not initiate the shutdown, but instead wait for the server to do it, and will not reply to the shutdown from the  server."),
		"urls":                  props.ArrayOf("string", "Corresponds to the URLs argument."),
	}

	required := []string{"urls"}

	typ := schema.ComplexTypeSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "Abstraction over the `curl` utility on a remote system. Transfer a URL.",
			Type:        "object",
			Properties:  inputs,
			Required:    required,
		},
	}

	return tool{optsType: typ, types: map[string]schema.ComplexTypeSpec{}}
}

// If we ever get a way to add the "required outputs" logic around a complexType
// resource := schema.ResourceSpec{
// 	ObjectTypeSpec: schema.ObjectTypeSpec{
// 		Description: "Abstraction over the `curl` utility on a remote system. Transfer a URL.",
// 		Properties:  implicitOutputs(inputs, map[string]schema.PropertySpec{}),
// 		Required:    required,
// 	},
// 	InputProperties: inputs,
// 	RequiredInputs:  required,
// }
