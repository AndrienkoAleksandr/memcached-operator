package controllers

import (
	"context"
	"testing"

	cachev1alpha1 "github.com/AndrienkoAleksandr/memcached-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestMemcachedController(t *testing.T) {
    // A Memcached object with metadata and spec.

    scheme := scheme.Scheme
	scheme.AddKnownTypes(cachev1alpha1.GroupVersion, &cachev1alpha1.Memcached{}, &cachev1alpha1.MemcachedList{})

    memcached := &cachev1alpha1.Memcached{
        TypeMeta: metav1.TypeMeta {
            APIVersion: "cache.example.com/v1alpha1",
            Kind: "Memcached",
        },
        ObjectMeta: metav1.ObjectMeta{
            Name:      "memcached",
            Namespace: "memcached-operator",
            Labels: map[string]string{
                "label-key": "label-value",
            },
        },
    }

    // Objects to track in the fake client.
    objs := []runtime.Object{memcached}

    // Create a fake client to mock API calls.
    cl := fake.NewFakeClient(objs...)

    // List Memcached objects filtering by labels
    opt := client.MatchingLabels(map[string]string{"label-key": "label-value"})
    memcachedList := &cachev1alpha1.MemcachedList{}
    err := cl.List(context.TODO(), memcachedList, opt)
    if err != nil {
        t.Fatalf("list memcached: (%v)", err)
    }
}
