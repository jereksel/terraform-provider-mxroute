package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractingDcim(t *testing.T) {
	//var re = regexp.MustCompile(`(?m)".*"`)

	//Given
	str := `



$TTL 14400
@       IN      SOA     ns1.mxrouting.net.      hostmaster.my.domain. (
                                                2020050103
                                                14400
                                                3600
                                                1209600
                                                86400 )

my.domain.	14400	IN	NS	ns1.mxrouting.net.
my.domain.	14400	IN	NS	ns2.mxrouting.net.

my.domain.	14400	IN	A	0.0.0.0
ftp	14400	IN	A	0.0.0.0
mail	14400	IN	A	0.0.0.0
pop	14400	IN	A	0.0.0.0
smtp	14400	IN	A	0.0.0.0
www	14400	IN	A	0.0.0.0

my.domain.	14400	IN	MX	10 mail



my.domain.	14400	IN	TXT	"v=spf1 a mx ip4:0.0.0.0 ~all"
x._domainkey	14400	IN	TXT	( "v=DKIM1; k=rsa; p=LINE1"
					"LINE2"
					"LINE3" )
`

	//When
	dcim, err := ExtractDcim(str)

	//Then
	assert.Empty(t, err)
	assert.Equal(t, "v=DKIM1; k=rsa; p=LINE1LINE2LINE3", dcim)

}
