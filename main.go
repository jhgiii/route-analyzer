/*
TODO: Parsing Output
TODO: Interfaces rather than concrete types
*/
package main

import (
	"fmt"
	"github.com/scrapli/scrapligo/driver/network"
	"github.com/scrapli/scrapligo/driver/options"
	"github.com/scrapli/scrapligo/platform"
	"log"
)

type JuniperRouter struct {
	Hostname string
	MgmtIp   string
	Scrapli  *platform.Platform
	Username string
	Password string
}

func (j *JuniperRouter) GetRoute(prefix, vrf string) (string, error) {
	return "", nil
}

func (j *JuniperRouter) parseRouteTable(routeTable string) (string, error) {
	return "", nil
}

type AristaRouter struct {
	Hostname string
	MgmtIp   string
	Scrapli  *platform.Platform
	Driver   *network.Driver
	Username string
	Password string
}

func NewAristaRouter(hostname, mgmtip, username, password string) *AristaRouter {
	var a AristaRouter
	a.Hostname = hostname
	a.MgmtIp = mgmtip
	a.Username = username
	a.Password = password
	return &a
}
func (a *AristaRouter) SetupDriver() error {
	var err error
	if len(a.MgmtIp) < 1 {
		return fmt.Errorf("password not set. please set MgmtIp")
	}
	if len(a.Username) < 1 {
		return fmt.Errorf("username not set. please set Username")
	}

	if len(a.Password) < 1 {
		return fmt.Errorf("password not set. please set Password")
	}
	a.Scrapli, err = platform.NewPlatform(
		"arista_eos",
		a.MgmtIp,
		options.WithAuthNoStrictKey(),
		options.WithAuthUsername(a.Username),
		options.WithAuthPassword(a.Password))
	if err != nil {
		return err
	}
	a.Driver, err = a.Scrapli.GetNetworkDriver()
	if err != nil {
		return err
	}
	return nil
}

func (a *AristaRouter) Connect() error {
	return a.Driver.Open()
}

func (a *AristaRouter) Close() error {
	return a.Driver.Close()
}

func (a *AristaRouter) GetRoute() (string, error) {
	if a.Driver == nil {
		return "", fmt.Errorf("driver not built, please use AristaRouter.SetupDriver()")
	}
	if a.Scrapli == nil {
		return "", fmt.Errorf("scrapli not setup, please use AristaRouter.SetupDriver()")
	}

	r, err := a.Driver.SendCommand("show ip route vrf all")
	if err != nil {
		return "", err
	}
	return r.Result, nil
}

func (a *AristaRouter) parseRouteTable(routeTable string) (string, error) {
	return "", nil
}

type CiscoRouter struct {
	Hostname string
	MgmtIp   string
	Scrapli  *platform.Platform
}

func (c *CiscoRouter) GetRoute(prefix, vrf string) (string, error) {
	return "", nil
}

func (c *CiscoRouter) parseRouteTable(routeTable string) (string, error) {
	return "", nil
}

type CumulusRouter struct {
	Hostname string
	MgmtIp   string
	Scrapli  *platform.Platform
}

func (c *CumulusRouter) GetRoute(prefix, vrf string) (string, error) {
	return "", nil
}

func (c *CumulusRouter) parseRouteTable(routeTable string) (string, error) {
	return "", nil
}

type Router interface {
	GetRoute() (string, error)
	parseRouteTable(string) (string, error)
}

func main() {
	a := NewAristaRouter("test", "172.20.20.2", "admin", "admin")
	err := a.SetupDriver()
	if err != nil {
		log.Fatal(err)
	}
	err = a.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer a.Close()

	r, err := a.GetRoute()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(r)
}
