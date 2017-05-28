package acsengine

import "fmt"

const (
	// AzurePublicProdFQDNFormat specifies the format for a prod dns name
	AzurePublicProdFQDNFormat = "%s.%s.cloudapp.azure.com"
	//AzureChinaProdFQDNFormat specify the endpoint of Azure China Cloud
	AzureChinaProdFQDNFormat = "%s.%s.cloudapp.chinacloudapi.cn"
)

// AzureLocations provides all azure regions in prod.
// Related powershell to refresh this list:
//   Get-AzureRmLocation | Select-Object -Property Location
var AzureLocations = []string{
	"australiaeast",
	"australiasoutheast",
	"brazilsouth",
	"canadacentral",
	"canadaeast",
	"centralindia",
	"centralus",
	"centraluseuap",
	"chinaeast",
	"chinanorth",
	"eastasia",
	"eastus",
	"eastus2",
	"eastus2euap",
	"japaneast",
	"japanwest",
	"koreacentral",
	"koreasouth",
	"northcentralus",
	"northeurope",
	"southcentralus",
	"southeastasia",
	"southindia",
	"uksouth",
	"ukwest",
	"westcentralus",
	"westeurope",
	"westindia",
	"westus",
	"westus2",
}

// FormatAzureProdFQDNs constructs all possible Azure prod fqdn
func FormatAzureProdFQDNs(fqdnPrefix string) []string {
	var fqdns []string
	for _, location := range AzureLocations {
		fqdns = append(fqdns, FormatAzureProdFQDN(fqdnPrefix, location))
	}
	return fqdns
}

// FormatAzureProdFQDN constructs an Azure prod fqdn
func FormatAzureProdFQDN(fqdnPrefix string, location string) string {
	FQDNFormat := AzurePublicProdFQDNFormat
	if location == "chinaeast" || location == "chinanorth" {
		FQDNFormat = AzureChinaProdFQDNFormat
	}
	return fmt.Sprintf(FQDNFormat, fqdnPrefix, location)
}

// GetSizeMap returns the size / storage map
func GetSizeMap() string {
	return `    "vmSizesMap": {
    "Standard_A0": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A1": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A10": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A11": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A1_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A2_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A2m_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A4": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A4_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A4m_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A5": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A6": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A7": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A8": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A8_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A8m_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_A9": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D1": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D11": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D11_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D11_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D12": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D12_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D12_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D13": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D13_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D13_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D14": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D14_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D14_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D15_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D1_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D2_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D2_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D3_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D3_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D4": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D4_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D4_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D5_v2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_D5_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_DS1": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS11": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS11_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS11_v2_Promo": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS12": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS12_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS12_v2_Promo": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS13": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS13_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS13_v2_Promo": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS14": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS14_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS14_v2_Promo": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS15_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS1_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS2_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS2_v2_Promo": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS3_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS3_v2_Promo": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS4": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS4_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS4_v2_Promo": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS5_v2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_DS5_v2_Promo": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_F1": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_F16": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_F16s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_F1s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_F2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_F2s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_F4": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_F4s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_F8": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_F8s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_G1": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_G2": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_G3": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_G4": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_G5": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_GS1": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_GS2": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_GS3": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_GS4": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_GS5": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_H16": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_H16m": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_H16mr": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_H16r": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_H8": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_H8m": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_L16s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_L32s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_L4s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_L8s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M128ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M128s": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_M64ms": {
      "storageAccountType": "Premium_LRS"
    },
    "Standard_NC12": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_NC24": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_NC24r": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_NC6": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_NV12": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_NV24": {
      "storageAccountType": "Standard_LRS"
    },
    "Standard_NV6": {
      "storageAccountType": "Standard_LRS"
    }
   }
`
}

// GetClassicSizeMap returns the size / storage map
func GetClassicSizeMap() string {
	return `    "vmSizesMap": {
        "Standard_A0": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_A1": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_A10": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_A11": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_A1_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_A2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_A2_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_A2m_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_A3": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_A4": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_A4_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_A4m_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_A5": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_A6": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_A7": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_A8": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_A8_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_A8m_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_A9": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D1": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D11": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D11_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D11_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D12": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D12_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D12_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D13": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D13_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D13_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D14": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D14_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D14_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D15_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D1_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D2_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D2_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D3": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D3_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D3_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D4": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D4_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D4_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D5_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_D5_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS1": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS11": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS11_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS11_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS12": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS12_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS12_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS13": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS13_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS13_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS14": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS14_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS14_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS15_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS1_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS2_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS2_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS3": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS3_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS3_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS4": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS4_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS4_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS5_v2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_DS5_v2_Promo": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_F1": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_F16": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_F16s": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_F1s": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_F2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_F2s": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_F4": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_F4s": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_F8": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_F8s": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_G1": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_G2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_G3": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_G4": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_G5": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_GS1": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_GS2": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_GS3": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_GS4": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_GS5": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_H16": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_H16m": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_H16mr": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_H16r": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_H8": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_H8m": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_L16s": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_L32s": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_L4s": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_L8s": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_M128ms": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_M128s": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_M64ms": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_NC12": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_NC24": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_NC24r": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_NC6": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_NV12": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_NV24": {
      "storageAccountType": "Standard_LRS"
    },
        "Standard_NV6": {
      "storageAccountType": "Standard_LRS"
    }
   }
`
}
