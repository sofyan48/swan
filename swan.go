package swan

import (
	"net"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sofyan48/swan/libs"
)

// ChekcIPRange Check ACL Addr range for ip and subneting
func ChekcIPRange(wip []string, addr string) bool {

	for _, ip := range wip {
		ipfix := strings.ReplaceAll(ip, " ", "")
		_, ipv4Net, err := net.ParseCIDR(ipfix)
		if err != nil {
			logrus.Infof("Ip Not Supported : %s", err)
		}
		if ipv4Net.Contains(net.ParseIP(addr)) {
			return true
		}
	}
	return false
}

// CheckValidIP Checking Valid IP
func CheckValidIP(c *gin.Context, ip string) {
	whitelistaddr := libs.GetEnvVariabel("ACL_ADDR", os.Getenv("ACL_ADDR"))
	cidr := strings.Split(whitelistaddr, ",")
	logrus.Infof("IP : %s", ip)
	if ip == "::1" {
		c.Next()
	} else {
		check := ChekcIPRange(cidr, ip)
		if check != true {
			c.JSON(401, gin.H{
				"message": "Not Authorize",
			})
			logrus.Infof("Ip Not Supported : %s", "Not Authorize")
			c.Abort()
		} else {
			c.Next()
		}
	}
}

// SwanACL Initial
func SwanACL() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		CheckValidIP(c, ip)
	}
}
