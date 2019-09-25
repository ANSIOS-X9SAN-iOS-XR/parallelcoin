package save

import (
   "io/ioutil"
   
   "github.com/pelletier/go-toml"
   
   "github.com/p9c/pod/app/apputil"
   "github.com/p9c/pod/pkg/pod"
)

func Pod(c *pod.Config) (success bool) {
	if yp, e := toml.Marshal(c); e == nil {
		apputil.EnsureDir(*c.ConfigFile)
		if e := ioutil.WriteFile(*c.ConfigFile, yp, 0600); e != nil {
			return
		}
		return true
	}
	return
}
