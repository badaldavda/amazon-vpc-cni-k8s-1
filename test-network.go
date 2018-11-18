package main

import (
	"net"
	"os"
	//	"time"

	"github.com/aws/amazon-vpc-cni-k8s/pkg/networkutils"
	log "github.com/cihub/seelog"
)

func main() {

	testcase := os.Args[1]

	switch testcase {
	case "1":
		log.Infof("testing %v, adding nat rules", testcase)
		runTestCase1()
	case "2":
		log.Infof("testing %v, updating nat rules", testcase)
		runTestCase2()
	case "3":
		log.Infof("testing %v, updating non-nat rules", testcase)
		runTestCase3()
	case "del":
		log.Infof("testing %v", testcase)
		runTestCaseDel()
	default:
		log.Info("testing default")
	}

	//	time.Sleep(5 * time.Second)
}

func runTestCaseTotal() {
	networkClient := networkutils.New()

	rules, err := networkClient.GetRuleList()

	log.Infof(" rules: %v, err %v\n", rules, err)

	src := "192.158.1.10"

	srcIPNet := net.IPNet{IP: net.ParseIP(src), Mask: net.IPv4Mask(255, 255, 255, 255)}

	log.Infof("srcIPNet %v \n", srcIPNet)
	srcRules, err := networkClient.GetRuleListBySrc(rules, srcIPNet)

	log.Infof(" source rules: %v, err %v\n", srcRules, err)

	toCIDRs := []string{"192.168.0.0/24", "192.168.1.0/24", "192.168.2.0/24"}

	networkClient.UpdateRuleListBySrc(rules, srcIPNet, toCIDRs, true)
	curRules, err := networkClient.GetRuleList()
	log.Infof(" use TNAT, rules: %v, err %v\n", curRules, err)

	networkClient.UpdateRuleListBySrc(curRules, srcIPNet, toCIDRs, false)
	curRules, err = networkClient.GetRuleList()
	log.Infof("not use TONAT rules: %v, err %v\n", curRules, err)

	toCIDRs = []string{"192.168.0.0/24", "192.168.1.0/24", "192.168.2.0/24"}

	networkClient.UpdateRuleListBySrc(rules, srcIPNet, toCIDRs, true)
	curRules, err = networkClient.GetRuleList()
	log.Infof(" use TNAT, rules: %v, err %v\n", curRules, err)

	// testing delete
	networkClient.DelRuleListBySrc(srcIPNet)

}

func runTestCase1() {
	networkClient := networkutils.New()

	rules, err := networkClient.GetRuleList()

	log.Infof(" rules: %v, err %v\n", rules, err)

	src := "192.158.1.10"

	srcIPNet := net.IPNet{IP: net.ParseIP(src), Mask: net.IPv4Mask(255, 255, 255, 255)}

	log.Infof("srcIPNet %v \n", srcIPNet)
	srcRules, err := networkClient.GetRuleListBySrc(rules, srcIPNet)

	log.Infof(" source rules: %v, err %v\n", srcRules, err)

	toCIDRs := []string{"192.168.0.0/24", "192.168.1.0/24", "192.168.2.0/24"}

	networkClient.UpdateRuleListBySrc(rules, srcIPNet, toCIDRs, true)
	curRules, err := networkClient.GetRuleList()
	log.Infof(" use TNAT, rules: %v, err %v\n", curRules, err)

}

func runTestCase2() {
	networkClient := networkutils.New()

	rules, err := networkClient.GetRuleList()

	log.Infof(" rules: %v, err %v\n", rules, err)

	src := "192.158.1.10"

	srcIPNet := net.IPNet{IP: net.ParseIP(src), Mask: net.IPv4Mask(255, 255, 255, 255)}

	log.Infof("srcIPNet %v \n", srcIPNet)
	srcRules, err := networkClient.GetRuleListBySrc(rules, srcIPNet)

	log.Infof(" source rules: %v, err %v\n", srcRules, err)

	toCIDRs := []string{"192.168.0.0/24", "192.168.3.0/24", "192.168.2.0/24"}

	networkClient.UpdateRuleListBySrc(rules, srcIPNet, toCIDRs, true)
	curRules, err := networkClient.GetRuleList()
	log.Infof(" use TNAT, rules: %v, err %v\n", curRules, err)

}
func runTestCase3() {
	networkClient := networkutils.New()

	rules, err := networkClient.GetRuleList()

	log.Infof(" rules: %v, err %v\n", rules, err)

	src := "192.158.1.10"

	srcIPNet := net.IPNet{IP: net.ParseIP(src), Mask: net.IPv4Mask(255, 255, 255, 255)}

	log.Infof("srcIPNet %v \n", srcIPNet)
	srcRules, err := networkClient.GetRuleListBySrc(rules, srcIPNet)

	log.Infof(" source rules: %v, err %v\n", srcRules, err)

	toCIDRs := []string{"192.168.0.0/24", "192.168.3.0/24", "192.168.2.0/24"}

	networkClient.UpdateRuleListBySrc(rules, srcIPNet, toCIDRs, false)
	curRules, err := networkClient.GetRuleList()
	log.Infof(" use TNAT, rules: %v, err %v\n", curRules, err)

}

func runTestCaseDel() {
	networkClient := networkutils.New()

	src := "192.158.1.10"

	srcIPNet := net.IPNet{IP: net.ParseIP(src), Mask: net.IPv4Mask(255, 255, 255, 255)}

	log.Infof("srcIPNet %v \n", srcIPNet)

	networkClient.DelRuleListBySrc(srcIPNet)
}
