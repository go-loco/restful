package rest_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/go-loco/restful/rest"
)

func TestGet(t *testing.T) {
	resp := rest.Get(server.URL + "/user")

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

}

func TestSlowGet(t *testing.T) {

	var f [100]*rest.Response

	for i := range f {
		f[i] = rb.Get("/slow/user")

		if f[i].Response.StatusCode != http.StatusOK {
			t.Fatal("f Status != OK (200)")
		}

	}

}

func TestHead(t *testing.T) {
	resp := rest.Head(server.URL + "/user")

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

}

func TestPost(t *testing.T) {
	resp := rest.Post(server.URL+"/user", &User{Name: "Matilda"})

	if resp.StatusCode != http.StatusCreated {
		t.Fatal("Status != OK (201)")
	}
}

func TestPostXML(t *testing.T) {

	rbXML := rest.RequestBuilder{
		BaseURL:     server.URL,
		ContentType: rest.XML,
	}

	resp := rbXML.Post("/xml/user", &User{Name: "Matilda"})

	if resp.StatusCode != http.StatusCreated {
		t.Fatal("Status != OK (201)")
	}
}

func TestPut(t *testing.T) {
	resp := rest.Put(server.URL+"/user/3", &User{Name: "Pichucha"})

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200")
	}
}

func TestPatch(t *testing.T) {
	resp := rest.Patch(server.URL+"/user/3", &User{Name: "Pichucha"})

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200")
	}
}

func TestDelete(t *testing.T) {
	resp := rest.Delete(server.URL + "/user/4")

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200")
	}
}

func TestOptions(t *testing.T) {
	resp := rest.Options(server.URL + "/user")

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200")
	}
}

func TestAsyncGet(t *testing.T) {

	rest.AsyncGet(server.URL+"/user", func(r *rest.Response) {
		if r.StatusCode != http.StatusOK {
			t.Fatal("Status != OK (200)")
		}
	})

	time.Sleep(50 * time.Millisecond)
}

func TestAsyncHead(t *testing.T) {

	rest.AsyncHead(server.URL+"/user", func(r *rest.Response) {
		if r.StatusCode != http.StatusOK {
			t.Fatal("Status != OK (200)")
		}
	})

	time.Sleep(50 * time.Millisecond)
}

func TestAsyncPost(t *testing.T) {

	rest.AsyncPost(server.URL+"/user", &User{Name: "Matilda"}, func(r *rest.Response) {
		if r.StatusCode != http.StatusCreated {
			t.Fatal("Status != OK (201)")
		}
	})

	time.Sleep(50 * time.Millisecond)
}

func TestAsyncPut(t *testing.T) {

	rest.AsyncPut(server.URL+"/user/3", &User{Name: "Pichucha"}, func(r *rest.Response) {
		if r.StatusCode != http.StatusOK {
			t.Fatal("Status != OK (200)")
		}
	})

	time.Sleep(50 * time.Millisecond)
}

func TestAsyncPatch(t *testing.T) {

	rest.AsyncPatch(server.URL+"/user/3", &User{Name: "Pichucha"}, func(r *rest.Response) {
		if r.StatusCode != http.StatusOK {
			t.Fatal("Status != OK (200)")
		}
	})

	time.Sleep(50 * time.Millisecond)
}

func TestAsyncDelete(t *testing.T) {

	rest.AsyncDelete(server.URL+"/user/4", func(r *rest.Response) {
		if r.StatusCode != http.StatusOK {
			t.Fatal("Status != OK (200)")
		}
	})

	time.Sleep(50 * time.Millisecond)
}

func TestAsyncOptions(t *testing.T) {

	rest.AsyncOptions(server.URL+"/user", func(r *rest.Response) {
		if r.StatusCode != http.StatusOK {
			t.Fatal("Status != OK (200)")
		}
	})

	time.Sleep(50 * time.Millisecond)
}

func TestHeaders(t *testing.T) {

	h := make(http.Header)
	h.Add("X-Test", "test")

	builder := rest.RequestBuilder{
		BaseURL: server.URL,
		Headers: h,
	}

	r := builder.Get("/header")

	if r.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

}

func TestWrongURL(t *testing.T) {
	r := rest.Get("foo")
	if r.Err == nil {
		t.Fatal("Wrong URL should get an error")
	}
}
