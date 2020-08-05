package fixtures

var gomod = `
bitbucket.org/bertimus9/systemstat v0.0.0-20180207000608-0eeff89b0690
github.com/Azure/azure-sdk-for-go v43.0.0+incompatible
github.com/Azure/go-autorest/autorest v0.9.6
github.com/Azure/go-autorest/autorest/adal v0.8.2
github.com/Azure/go-autorest/autorest/to v0.2.0
github.com/GoogleCloudPlatform/k8s-cloud-provider v0.0.0-20200415212048-7901bc822317
github.com/JeffAshton/win_pdh v0.0.0-20161109143554-76bb4ee9f0ab
github.com/Microsoft/go-winio v0.4.15-0.20190919025122-fc70bd9a86b5
github.com/Microsoft/hcsshim v0.8.10-0.20200715222032-5eafd1556990
github.com/PuerkitoBio/purell v1.1.1
github.com/armon/circbuf v0.0.0-20150827004946-bbbad097214e
github.com/auth0/go-jwt-middleware v0.0.0-20170425171159-5493cabe49f7 // indirect
github.com/aws/aws-sdk-go v1.28.2
github.com/blang/semver v3.5.0+incompatible
github.com/boltdb/bolt v1.3.1 // indirect
github.com/caddyserver/caddy v1.0.3
github.com/clusterhq/flocker-go v0.0.0-20160920122132-2b8b7259d313
github.com/codegangsta/negroni v1.0.0 // indirect
github.com/container-storage-interface/spec v1.2.0
github.com/containernetworking/cni v0.8.0
github.com/coredns/corefile-migration v1.0.10
github.com/coreos/go-oidc v2.1.0+incompatible
github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e
github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f
github.com/cpuguy83/go-md2man/v2 v2.0.0
github.com/davecgh/go-spew v1.1.1
github.com/docker/distribution v2.7.1+incompatible
github.com/docker/docker v1.4.2-0.20200309214505-aa6a9891b09c
github.com/docker/go-connections v0.4.0
github.com/docker/go-units v0.4.0
github.com/elazarl/goproxy v0.0.0-20180725130230-947c36da3153
github.com/emicklei/go-restful v2.9.5+incompatible
github.com/evanphx/json-patch v0.0.0-20190815234213-e83c0a1c26c8
github.com/fsnotify/fsnotify v1.4.9
github.com/go-bindata/go-bindata v3.1.1+incompatible
github.com/go-openapi/analysis v0.19.5
github.com/go-openapi/loads v0.19.4
github.com/go-openapi/spec v0.19.3
github.com/go-openapi/strfmt v0.19.3
github.com/go-openapi/validate v0.19.5
github.com/go-ozzo/ozzo-validation v3.5.0+incompatible // indirect
github.com/gogo/protobuf v1.3.1
github.com/golang/groupcache v0.0.0-20191227052852-215e87163ea7
github.com/golang/mock v1.3.1
github.com/google/cadvisor v0.37.0
github.com/google/go-cmp v0.4.0
github.com/google/gofuzz v1.1.0
github.com/google/uuid v1.1.1
github.com/googleapis/gnostic v0.4.1
github.com/gorilla/context v1.1.1 // indirect
github.com/hashicorp/golang-lru v0.5.1
github.com/heketi/heketi v9.0.1-0.20190917153846-c2e2a4ab7ab9+incompatible
github.com/heketi/tests v0.0.0-20151005000721-f3775cbcefd6 // indirect
github.com/ishidawataru/sctp v0.0.0-20190723014705-7c296d48a2b5
github.com/json-iterator/go v1.1.10
github.com/libopenstorage/openstorage v1.0.0
github.com/lithammer/dedent v1.1.0
github.com/lpabon/godbc v0.1.1 // indirect
github.com/magiconair/properties v1.8.1 // indirect
github.com/miekg/dns v1.1.4
github.com/moby/ipvs v1.0.1
github.com/mohae/deepcopy v0.0.0-20170603005431-491d3605edfb // indirect
github.com/mrunalp/fileutils v0.0.0-20200520151820-abd8a0e76976
github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822
github.com/mvdan/xurls v1.1.0
github.com/onsi/ginkgo v1.11.0
github.com/onsi/gomega v1.7.0
github.com/opencontainers/go-digest v1.0.0-rc1
github.com/opencontainers/runc v1.0.0-rc91.0.20200707015106-819fcc687efb
github.com/opencontainers/selinux v1.5.2
github.com/pkg/errors v0.9.1
github.com/pmezard/go-difflib v1.0.0
github.com/prometheus/client_golang v1.7.1
github.com/prometheus/client_model v0.2.0
github.com/prometheus/common v0.10.0
github.com/quobyte/api v0.1.2
github.com/robfig/cron v1.1.0
github.com/spf13/afero v1.2.2
github.com/spf13/cobra v1.0.0
github.com/spf13/jwalterweatherman v1.1.0 // indirect
github.com/spf13/pflag v1.0.5
github.com/spf13/viper v1.4.0
github.com/storageos/go-api v0.0.0-20180912212459-343b3eff91fc
github.com/stretchr/testify v1.4.0
github.com/thecodeteam/goscaleio v0.1.0
github.com/urfave/negroni v1.0.0 // indirect
github.com/vishvananda/netlink v1.1.0
github.com/vmware/govmomi v0.20.3
go.etcd.io/etcd v0.5.0-alpha.5.0.20200716221620-18dfb9cca345
golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
golang.org/x/net v0.0.0-20200707034311-ab3426394381
golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6
golang.org/x/sys v0.0.0-20200622214017-ed371f2e16b4
golang.org/x/time v0.0.0-20191024005414-555d28b269f0
golang.org/x/tools v0.0.0-20200616133436-c1934b75d054
gonum.org/v1/gonum v0.6.2
gonum.org/v1/netlib v0.0.0-20190331212654-76723241ea4e // indirect
google.golang.org/api v0.15.1
google.golang.org/grpc v1.27.0
gopkg.in/gcfg.v1 v1.2.0
gopkg.in/square/go-jose.v2 v2.2.2
gopkg.in/yaml.v2 v2.2.8
k8s.io/api v0.0.0
k8s.io/apiextensions-apiserver v0.0.0
k8s.io/apimachinery v0.0.0
k8s.io/apiserver v0.0.0
k8s.io/cli-runtime v0.0.0
k8s.io/client-go v0.0.0
k8s.io/cloud-provider v0.0.0
k8s.io/cluster-bootstrap v0.0.0
k8s.io/code-generator v0.0.0
k8s.io/component-base v0.0.0
k8s.io/cri-api v0.0.0
k8s.io/csi-translation-lib v0.0.0
k8s.io/gengo v0.0.0-20200428234225-8167cfdcfc14
k8s.io/heapster v1.2.0-beta.1
k8s.io/klog/v2 v2.2.0
k8s.io/kube-aggregator v0.0.0
k8s.io/kube-controller-manager v0.0.0
k8s.io/kube-openapi v0.0.0-20200427153329-656914f816f9
k8s.io/kube-proxy v0.0.0
k8s.io/kube-scheduler v0.0.0
k8s.io/kubectl v0.0.0
k8s.io/kubelet v0.0.0
k8s.io/legacy-cloud-providers v0.0.0
k8s.io/metrics v0.0.0
k8s.io/sample-apiserver v0.0.0
k8s.io/system-validators v1.1.2
k8s.io/utils v0.0.0-20200729134348-d5654de09c73
sigs.k8s.io/kustomize v2.0.3+incompatible
sigs.k8s.io/yaml v1.2.0
`

var Gomod = readLines(gomod)
