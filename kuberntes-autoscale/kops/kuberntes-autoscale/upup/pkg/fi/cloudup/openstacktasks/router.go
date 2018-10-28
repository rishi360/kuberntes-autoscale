/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package openstacktasks

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/upup/pkg/fi/cloudup/openstack"
)

//go:generate fitask -type=Router
type Router struct {
	ID        *string
	Name      *string
	Lifecycle *fi.Lifecycle
}

var _ fi.CompareWithID = &Router{}

func (n *Router) CompareWithID() *string {
	return n.ID
}

func (n *Router) Find(context *fi.Context) (*Router, error) {
	cloud := context.Cloud.(openstack.OpenstackCloud)
	opt := routers.ListOpts{
		Name: fi.StringValue(n.Name),
		ID:   fi.StringValue(n.ID),
	}
	rs, err := cloud.ListRouters(opt)
	if err != nil {
		return nil, err
	}
	if rs == nil {
		return nil, nil
	} else if len(rs) != 1 {
		return nil, fmt.Errorf("found multiple routers with name: %s", fi.StringValue(n.Name))
	}
	v := rs[0]
	actual := &Router{
		ID:        fi.String(v.ID),
		Name:      fi.String(v.Name),
		Lifecycle: n.Lifecycle,
	}
	return actual, nil
}

func (c *Router) Run(context *fi.Context) error {
	return fi.DefaultDeltaRunMethod(c, context)
}

func (_ *Router) CheckChanges(a, e, changes *Router) error {
	if a == nil {
		if e.Name == nil {
			return fi.RequiredField("Name")
		}
	} else {
		if changes.Name != nil {
			return fi.CannotChangeField("Name")
		}
	}
	return nil
}

func (_ *Router) RenderOpenstack(t *openstack.OpenstackAPITarget, a, e, changes *Router) error {
	if a == nil {
		glog.V(2).Infof("Creating Router with name:%q", fi.StringValue(e.Name))

		opt := routers.CreateOpts{
			Name:         fi.StringValue(e.Name),
			AdminStateUp: fi.Bool(true),
		}

		v, err := t.Cloud.CreateRouter(opt)
		if err != nil {
			return fmt.Errorf("Error creating router: %v", err)
		}

		e.ID = fi.String(v.ID)
		glog.V(2).Infof("Creating a new Openstack router, id=%s", v.ID)
		return nil
	}
	e.ID = a.ID
	glog.V(2).Infof("Using an existing Openstack router, id=%s", fi.StringValue(e.ID))
	return nil
}
