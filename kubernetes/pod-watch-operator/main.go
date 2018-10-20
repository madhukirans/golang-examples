package main
import (
	"fmt"
	"log"
	"net/http"
	"time"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/cache"
	"k8s.io/kubernetes/pkg/client/restclient"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/util/wait"
	"k8s.io/client-go/tools/clientcmd"
	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
)

func podCreated(obj interface{}) {
	pod := obj.(*api.Pod)
	fmt.Println("Pod created: " + pod.ObjectMeta.Name)
}
func podDeleted(obj interface{}) {
	pod := obj.(*api.Pod)
	fmt.Println("Pod deleted: " + pod.ObjectMeta.Name)
}
func watchPods(client *client.Client, store cache.Store) cache.Store {
	//Define what we want to look for (Pods)
	watchlist := cache.NewListWatchFromClient(client, "pods", api.NamespaceAll, fields.Everything())
	resyncPeriod := 30 * time.Minute
	//Setup an informer to call functions when the watchlist changes
	eStore, eController := framework.NewInformer(
		watchlist,
		&api.Pod{},
		resyncPeriod,
		framework.ResourceEventHandlerFuncs{
			AddFunc:    podCreated,
			DeleteFunc: podDeleted,
		},
	)
	//Run the controller as a goroutine
	go eController.Run(wait.NeverStop)
	return eStore
}
func main() {

	cfg, err := clientcmd.BuildConfigFromFlags("", "/Users/madhuseelam/sre/terraform-kubernetes-installer/generated/kubeconfig")
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	//Create a new client to interact with cluster and freak if it doesn't work
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalln("Client not created sucessfully:", err)
	}
	//Create a cache to store Pods
	var podsStore cache.Store
	//Watch for Pods
	podsStore = watchPods(kubeClient, podsStore)
	//Keep alive
	log.Fatal(http.ListenAndServe(":8080", nil))
}
