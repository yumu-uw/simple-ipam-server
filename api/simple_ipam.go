package api

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	echo "github.com/labstack/echo/v4"
	"github.com/yumu-uw/simple-ipam-server/lib"
)

type SimpleIpam struct {
	Lock sync.Mutex
}

func sendAssetStoreError(ctx echo.Context, code int, message string) Error {
	AssetErr := Error{
		Code:    int32(code),
		Message: message,
	}
	return AssetErr
}

// AddNewSubnet implements ServerInterface.
func (*SimpleIpam) AddNewSubnet(ctx echo.Context) error {
	var newSubnet NewSubnet
	err := ctx.Bind(&newSubnet)
	if err != nil {
		ase := sendAssetStoreError(ctx, http.StatusBadRequest, "Invalid format for NewSubnet")
		return ctx.JSON(int(ase.Code), ase.Message)
	}
	var subnet Subnet
	subnet.Id = 1
	subnet.Nwaddr = newSubnet.Nwaddr
	subnet.Netmask = newSubnet.Netmask
	return ctx.JSON(http.StatusCreated, subnet)
}

// AddRecord implements ServerInterface.
func (s *SimpleIpam) AddRecord(ctx echo.Context) error {
	panic("unimplemented")
}

// FindAllSubnet implements ServerInterface.
func (s *SimpleIpam) FindAllSubnet(ctx echo.Context) error {
	s.Lock.Lock()
	// var result []netip.Addr
	ips, err := lib.Hosts("192.168.1.240/24")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(ips)
	return ctx.JSON(http.StatusOK, ips)
}

// UpdateRecord implements ServerInterface.
func (s *SimpleIpam) UpdateRecord(ctx echo.Context) error {
	panic("unimplemented")
}

func NewSimpleIpam() *SimpleIpam {
	return &SimpleIpam{}
}
