package sdk

import (
	"encoding/xml"
	"fmt"

	"github.com/namecheap/go-namecheap-sdk/v2/namecheap"
)

type domainsCheckResult struct {
	Domain                   string  `xml:"Domain,attr"`
	Available                bool    `xml:"Available,attr"`
	Description              string  `xml:"Description,attr"`
	IsPremiumName            bool    `xml:"IsPremiumName,attr"`
	PremiumRegistrationPrice float32 `xml:"PremiumRegistrationPrice,attr"`
	PremiumRenewalPrice      float32 `xml:"PremiumRenewalPrice,attr"`
	PremiumRestorePrice      float32 `xml:"PremiumRestorePrice,attr"`
	PremiumTransferPrice     float32 `xml:"PremiumTransferPrice,attr"`
	IcannFee                 float32 `xml:"IcannFee,attr"`
	EapFee                   float32 `xml:"EapFee,attr"`
}

type domainsCheckResponse struct {
	XMLName *xml.Name `xml:"ApiResponse"`
	Errors  *[]struct {
		Message string `xml:",chardata"`
		Number  string `xml:"Number,attr"`
	} `xml:"Errors>Error"`
	CommandResponse *domainsCheckCommandResponse `xml:"CommandResponse"`
}

type domainsCheckCommandResponse struct {
	Result *domainsCheckResult `xml:"DomainCheckResult"`
}

func DomainsAvailable(client *namecheap.Client, domains string) (*domainsCheckCommandResponse, error) {
	var resp domainsCheckResponse

	params := map[string]string{
		"Command":    "namecheap.domains.check",
		"DomainList": domains,
	}
	if _, err := doXmlWithBackoff(client, params, &resp); err != nil {
		return nil, err
	}

	if resp.Errors != nil && len(*resp.Errors) > 0 {
		apiErr := (*resp.Errors)[0]
		return nil, fmt.Errorf("%s (%s)", apiErr.Message, apiErr.Number)
	}

	return resp.CommandResponse, nil
}
