package testinteg

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

var baseURL = flag.String("url", "http://localhost:8080", "base URL of the site to test")

func waitForSite(timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		resp, err := http.Get(*baseURL + "/")
		if err == nil && resp.StatusCode == http.StatusOK {
			resp.Body.Close()
			return nil
		}
		if resp != nil {
			resp.Body.Close()
		}
		time.Sleep(2 * time.Second)
	}
	return fmt.Errorf("site did not return 200 within %s", timeout)
}

func TestMain(m *testing.M) {
	flag.Parse()
	if err := waitForSite(30 * time.Second); err != nil {
		fmt.Println("FAIL:", err)
		return
	}
	m.Run()
}

func get(t *testing.T, path string) *http.Response {
	t.Helper()
	resp, err := http.Get(*baseURL + path)
	if err != nil {
		t.Fatalf("GET %s failed: %v", path, err)
	}
	return resp
}

func body(t *testing.T, resp *http.Response) string {
	t.Helper()
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("could not read response body: %v", err)
	}
	return string(b)
}

func TestSiteResponds(t *testing.T) {
	resp := get(t, "/")
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestIndexContainsSiteName(t *testing.T) {
	resp := get(t, "/")
	if !strings.Contains(body(t, resp), "Bodily Integrity Commons") {
		t.Error("expected page to contain 'Bodily Integrity Commons'")
	}
}

func TestNoUnreplacedPlaceholders(t *testing.T) {
	resp := get(t, "/")
	if strings.Contains(body(t, resp), "__COMMIT_SHA__") {
		t.Error("__COMMIT_SHA__ placeholder was not replaced")
	}
}

func TestCSSServed(t *testing.T) {
	resp := get(t, "/css/styles.css")
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 for CSS, got %d", resp.StatusCode)
	}
}

func TestOrganizationsPresent(t *testing.T) {
	resp := get(t, "/organizations.html")
	b := body(t, resp)
	orgs := []string{"Intact Global", "Intaction", "Intact America"}
	for _, org := range orgs {
		if !strings.Contains(b, org) {
			t.Errorf("expected organizations page to contain organization %q", org)
		}
	}
}

func TestJSServed(t *testing.T) {
	resp := get(t, "/js/nav.js")
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 for nav.js, got %d", resp.StatusCode)
	}
}

func TestAskButtonPresent(t *testing.T) {
	askButtonName := "prompt-submit"
	resp := get(t, "/")
	if !strings.Contains(body(t, resp), askButtonName) {
		t.Errorf("expected page to contain the ask prompt button '%s'", askButtonName)
	}
}

func TestAskReturnsResponse(t *testing.T) {
	resp := get(t, "/ask?q=circumcision")
	b := body(t, resp)
	if !strings.Contains(b, "circumcision") {
		t.Error("expected /ask response to contain the query term")
	}
}

func TestUsage(t *testing.T) {
	fmt.Printf("Running integration tests against: %s\n", *baseURL)
}
