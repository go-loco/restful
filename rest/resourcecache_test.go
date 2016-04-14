package rest_test

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/go-loco/restful/rest"
)

func TestCacheGetLowCacheMaxSize(t *testing.T) {

	mcs := rest.MaxCacheSize
	defer func() { rest.MaxCacheSize = mcs }()

	rest.MaxCacheSize = 500

	var f [1000]*rest.Response

	for i := range f {
		f[i] = rb.Get("/cache/user")

		if f[i].Response.StatusCode != http.StatusOK {
			t.Fatal("f Status != OK (200)")
		}

	}

}

func TestCacheGet(t *testing.T) {

	var f [1000]*rest.Response

	for i := range f {
		f[i] = rb.Get("/cache/user")

		if f[i].Response.StatusCode != http.StatusOK {
			t.Fatal("f Status != OK (200)")
		}

	}

}

func TestCacheGetEtag(t *testing.T) {

	var f [100]*rest.Response

	for i := range f {
		f[i] = rb.Get("/cache/etag/user")

		if f[i].Response.StatusCode != http.StatusOK {
			t.Fatal("f Status != OK (200)")
		}

	}

}

func TestCacheGetLastModified(t *testing.T) {

	var f [100]*rest.Response

	for i := range f {
		f[i] = rb.Get("/cache/lastmodified/user")

		if f[i].Response.StatusCode != http.StatusOK {
			t.Fatal("f Status != OK (200)")
		}

	}

}

func TestCacheGetExpires(t *testing.T) {

	var f [100]*rest.Response

	for i := range f {
		f[i] = rb.Get("/cache/expires/user")

		if f[i].Response.StatusCode != http.StatusOK {
			t.Fatal("f Status != OK (200)")
		}

	}

}

func TestCacheForkJoinGet(t *testing.T) {

	var f [100]*rest.FutureResponse

	for x := 0; x < 1000; x++ {

		rb.ForkJoin(func(cr *rest.Concurrent) {
			for i := range f {
				f[i] = cr.Get("/cache/user")
			}
		})

		for i := range f {
			if f[i].Response().StatusCode != http.StatusOK {
				t.Fatal("f[" + strconv.Itoa(i) + "] Status != OK (200)")
			}
		}

	}

}

func TestCacheSlowGet(t *testing.T) {

	var f [1000]*rest.Response

	for i := range f {
		f[i] = rb.Get("/cache/user")

		if f[i].Response.StatusCode != http.StatusOK {
			t.Fatal("f Status != OK (200)")
		}

		//Wait for so we get cache eviction
		time.Sleep(3 * time.Millisecond)

	}

}

func TestCacheSlowForkJoinGet(t *testing.T) {

	var f [100]*rest.FutureResponse

	for x := 0; x < 10; x++ {

		rb.ForkJoin(func(cr *rest.Concurrent) {
			for i := range f {
				f[i] = cr.Get("/slow/cache/user")
			}
		})

		for i := range f {
			if f[i].Response().StatusCode != http.StatusOK {
				t.Fatal("f[" + strconv.Itoa(i) + "] Status != OK (200)")
			}
		}

		//Wait for so we get cache eviction
		time.Sleep(300 * time.Millisecond)

	}

}
