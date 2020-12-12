package controllers

import (
    "context"
    "testing"

    cachev1alpha1 "github.com/AndrienkoAleksandr/memcached-operator/pkg/apis/cache/v1alpha1"
    "k8s.io/apimachinery/pkg/runtime"
    "sigs.k8s.io/controller-runtime/pkg/client"
    "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestMemcachedController(t *testing.T) {
    // A Memcached object with metadata and spec.
    memcached := &cachev1alpha1.Memcached{
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
