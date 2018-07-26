package tls

import (
	"crypto/tls"
	"fmt"
	"sync"
	"time"
)

var (
	norm = make([]domTime, 0)
	warn = make([]domTime, 0)
	wg   sync.WaitGroup

	now, gap time.Time
)

type domTime struct {
	domain string
	exp    time.Time
	beg    time.Time
}

// Check for tls check
func Check(days int32, doms []string) {
	if days <= 0 || len(doms) == 0 {
		fmt.Println("invalid parameters")
		return
	}
	fmt.Printf("check if domain is expired within [%d] days later:\n", days)
	now = time.Now()
	gap = now.Add(time.Hour * 24 * time.Duration(days))

	getExpireTime(doms)
	printExpire()
}

func getExpireTime(doms []string) {
	for _, dom := range doms {
		wg.Add(1)
		go getEachTime(dom)
	}
	wg.Wait()
}

// get tls time info
func getEachTime(dom string) {
	defer wg.Done()

	conn, err := tls.Dial("tcp", dom+":443", nil)
	if err != nil {
		fmt.Printf("  %s error:%+v\n", dom, err)
		return
	}
	state := conn.ConnectionState()
	for _, p := range state.PeerCertificates {
		if !p.IsCA {
			// p is *x509.Certificate;
			// NotBefore, NotAfter time.Time for Validity bounds.
			classify(dom, p.NotBefore, p.NotAfter)
		}
	}
}

// classify to save dom in warn or norm slice
func classify(dom string, before, after time.Time) {
	diff := after.Sub(gap)
	dt := domTime{dom, before, after}
	if int(diff.Hours()) > 0 {
		norm = append(norm, dt)
	} else {
		warn = append(warn, dt)
	}
}

func printExpire() {
	const FMT = "Jan 02，2006"

	fmt.Print("\nstatus\tbegin-with\texpired-at\tdomain-name\n")
	for _, d := range warn {
		fmt.Printf("warning\t%s\t%s\t%s\n", d.exp.Format(FMT), d.beg.Format(FMT), d.domain)
	}
	if len(warn) != 0 {
		fmt.Println()
	}
	for _, d := range norm {
		fmt.Printf("normal\t%s\t%s\t%s\n", d.exp.Format(FMT), d.beg.Format(FMT), d.domain)
	}
}
