package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Properties struct {
	IsExpanded                   bool   `json:"IsExpanded, omitempty"`
	BackColorOSX                 string `json:BackColorOSX,omitempty"`
	CursorColor                  string `json:CursorColor,omitempty"`
	BoldColorOSX                 string `json:BoldColorOSX,omitempty"`
	CursorTextColor              string `json:CursorTextColor,omitempty"`
	CustomPalette0NormalBlack    string `json:CustomPalette0NormalBlack,omitempty"`
	CustomPalette10BrightGreen   string `json:CustomPalette10BrightGreen,omitempty"`
	CustomPalette11BrightYellow  string `json:CustomPalette11BrightYellow,omitempty"`
	CustomPalette12BrightBlue    string `json:CustomPalette12BrightBlue,omitempty"`
	CustomPalette13BrightMagenta string `json:CustomPalette13BrightMagenta,omitempty"`
	CustomPalette14BrightCyan    string `json:CustomPalette14BrightCyan,omitempty"`
	CustomPalette15BrightWhite   string `json:CustomPalette15BrightWhite,omitempty"`
	CustomPalette1NormalRed      string `json:CustomPalette1NormalRed,omitempty"`
	CustomPalette2NormalGreen    string `json:CustomPalette2NormalGreen,omitempty"`
	CustomPalette3NormalYellow   string `json:CustomPalette3NormalYellow,omitempty"`
	CustomPalette4NormalBlue     string `json:CustomPalette4NormalBlue,omitempty"`
	CustomPalette5NormalMagenta  string `json:CustomPalette5NormalMagenta,omitempty"`
	CustomPalette6NormalCyan     string `json:CustomPalette6NormalCyan,omitempty"`
	CustomPalette7NormalWhite    string `json:CustomPalette7NormalWhite,omitempty"`
	CustomPalette8BrightBlack    string `json:CustomPalette8BrightBlack,omitempty"`
	CustomPalette9BrightRed      string `json:CustomPalette9BrightRed,omitempty"`
	FontSizeOSX                  string `json:FontSizeOSX,omitempty"`
	ForeColorOSX                 string `json:ForeColorOSX,omitempty"`
	SelectionColor               string `json:SelectionColor,omitempty"`
	SelectionTextColor           string `json:SelectionTextColor,omitempty"`
}

type Object struct {
	Type                    string     `json:"Type"`
	Name                    string     `json:"Name"`
	Username                string     `json:"Username,omitempty"`
	Password                string     `json:"Password,omitempty"`
	ID                      string     `json:"ID,omitempty"`
	CredentialID            string     `json:"CredentialID,omitempty"`
	CredentialName          string     `json:"CredentialName,omitempty"`
	TerminalConnectionType  string     `json:"TerminalConnectionType,omitempty"`
	ComputerName            string     `json:"ComputerName,omitempty"`
	CredentialsFromParent   bool       `json:CredentialsFromParent,omitempty"`
	SecureGatewayFromParent bool       `json:SecureGatewayFromParent,omitempty"`
	Color                   string     `json:Color,omitempty"`
	Properties              Properties `json:"Properties,omitempty"`
	Objects                 []*Object  `json:"Objects,omitempty"`
}

func (obj *Object) AddItem(item *Object) {
	obj.Objects = append(obj.Objects, item)
}

func sortSlice(object *Object) {
	sort.SliceStable(object.Objects, func(i, j int) bool {
		oi, oj := object.Objects[i], object.Objects[j]
		return oi.Name < oj.Name
	})
}

var (
	sshCredentialName string
	rdpCredentialName string
)

func init() {
	flag.StringVar(&sshCredentialName, "ssh_credential_name", "", "RoyalTS Credential Name for SSH")
	flag.StringVar(&rdpCredentialName, "rdp_credential_name", "", "RoyalTS Credential Name for RDP")
	flag.Parse()
}

func main() {

	environments := map[string]*Object{}

	svc := ec2.New(session.New())
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String("running"), aws.String("pending")},
			},
		},
	}
	resp, err := svc.DescribeInstances(params)
	if err != nil {
		fmt.Println("there was an error listing instances in", err.Error())
		log.Fatal(err.Error())
	}
	for idx, _ := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			obj := Object{
				ID:   *inst.InstanceId,
				Name: *inst.InstanceId,
				CredentialsFromParent:   true,
				SecureGatewayFromParent: true,
				Properties: Properties{
					BackColorOSX:                 "#001e26",
					CursorColor:                  "#708284",
					BoldColorOSX:                 "#819090",
					CursorTextColor:              "#002731",
					CustomPalette0NormalBlack:    "#002731",
					CustomPalette10BrightGreen:   "#465b62",
					CustomPalette11BrightYellow:  "#536870",
					CustomPalette12BrightBlue:    "#708284",
					CustomPalette13BrightMagenta: "#5956ba",
					CustomPalette14BrightCyan:    "#819090",
					CustomPalette15BrightWhite:   "#fdf5dd",
					CustomPalette1NormalRed:      "#d11b24",
					CustomPalette2NormalGreen:    "#738a05",
					CustomPalette3NormalYellow:   "#a57706",
					CustomPalette4NormalBlue:     "#2076c8",
					CustomPalette5NormalMagenta:  "#c71b6f",
					CustomPalette6NormalCyan:     "#259286",
					CustomPalette7NormalWhite:    "#eae3cc",
					CustomPalette8BrightBlack:    "#001e26",
					CustomPalette9BrightRed:      "#bd3612",
					FontSizeOSX:                  "14",
					ForeColorOSX:                 "#708284",
					SelectionColor:               "#002731",
					SelectionTextColor:           "#819090",
				},
			}

			obj.SecureGatewayFromParent = true

			if inst.Platform != nil {
				obj.Type = "RemoteDesktopConnection"
				if rdpCredentialName != "" {
					obj.CredentialsFromParent = false
					obj.CredentialName = rdpCredentialName
				}
			} else {
				obj.Type = "TerminalConnection"
				obj.TerminalConnectionType = "SSH"
				if sshCredentialName != "" {
					obj.CredentialsFromParent = false
					obj.CredentialName = sshCredentialName
				}
			}

			t := time.Now().UTC()
			elapsed := t.Sub(*inst.LaunchTime)
			if int(elapsed.Hours()) < 12 {
				obj.Color = "#94ba82"
			}

			// if inst.PublicDnsName != nil && obj.ComputerName == "" {
			// 	obj.ComputerName = *inst.PublicDnsName
			// }
			if inst.PublicIpAddress != nil && obj.ComputerName == "" {
				obj.ComputerName = *inst.PublicIpAddress
			}
			// if inst.PrivateDnsName != nil && obj.ComputerName == "" {
			// 	obj.ComputerName = *inst.PrivateDnsName
			// }
			if inst.PrivateIpAddress != nil && obj.ComputerName == "" {
				obj.ComputerName = *inst.PrivateIpAddress
			}

			var environment string = "unknown"
			for _, tag := range inst.Tags {
				if *tag.Key == "Name" {
					obj.Name = *tag.Value
				}
				if *tag.Key == "Environment" {
					environment = *tag.Value
				}
			}

			if _, ok := environments[environment]; !ok {
				environments[environment] = &Object{
					Type: "Folder",
					Name: environment,
					SecureGatewayFromParent: true,
					CredentialsFromParent:   true,
					Properties: Properties{
						IsExpanded: true,
					},
				}
			}

			data := environments[environment]
			data.AddItem(&obj)

		}
	}
	var region string = "us-west-2"
	connections := Object{
		Type: "Folder",
		Name: region,
		CredentialsFromParent:   true,
		SecureGatewayFromParent: true,
	}

	// spew.Dump(connections)

	for k, _ := range environments {
		// spew.Dump(len(environments[k].Objects))
		sortSlice(environments[k])
		connections.AddItem(environments[k])
	}

	sortSlice(&connections)

	json, err := json.MarshalIndent(connections, "", "  ")
	fmt.Print(string(json))
}
